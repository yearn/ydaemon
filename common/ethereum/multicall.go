package ethereum

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/logs"
)

type Call struct {
	Name     string         `json:"name"`
	Method   string         `json:"method"`
	Version  string         `json:"version"`
	Abi      *abi.ABI       `json:"abi"`
	Target   common.Address `json:"target"`
	CallData []byte         `json:"call_data"`
}

type CallResponse struct {
	Success    bool   `json:"success"`
	ReturnData []byte `json:"returnData"`
}

type TEthMultiCaller struct {
	Signer          *bind.TransactOpts
	Client          *ethclient.Client
	Abi             *abi.ABI
	ContractAddress common.Address
}

func (call Call) GetMultiCall() contracts.Multicall2Call {
	return contracts.Multicall2Call{
		Target:   call.Target,
		CallData: call.CallData,
	}
}

// NewMulticall creates a new instance of a TEthMultiCaller. This is the instance we
// will later use to perform multiple ethereum calls batched in the same transaction.
// For performance reason, this should be initialized once and then reused.
func NewMulticall(rpcURI string, multicallAddress common.Address) TEthMultiCaller {
	if rpcURI == "" {
		logs.Error("No rpcURI provided.")
		return TEthMultiCaller{}
	}
	client, err := ethclient.Dial(rpcURI)
	if err != nil {
		logs.
			Capture(`error`, `failed to connect to node`).
			SetEntity(`multicall`).
			SetExtra(`error`, err.Error()).
			SetTag(`rpcURI`, rpcURI).
			SetTag(`multicallAddress`, multicallAddress.Hex()).
			Send()
		time.Sleep(time.Second)
		return NewMulticall(rpcURI, multicallAddress)
	}

	// Load Multicall abi for later use
	mcAbi, err := contracts.Multicall2MetaData.GetAbi()
	if err != nil {
		currentChainID, _ := client.ChainID(context.Background())
		logs.
			Capture(`error`, `failed to decode Multicall ABI`).
			SetEntity(`multicall`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, currentChainID.String()).
			SetTag(`rpcURI`, rpcURI).
			SetTag(`multicallAddress`, multicallAddress.Hex()).
			Send()
		//TODO: FIX HERE TO PREVENT LOOP
		time.Sleep(time.Second)
		return NewMulticall(rpcURI, multicallAddress)
	}

	return TEthMultiCaller{
		Signer:          randomSigner(),
		Client:          client,
		Abi:             mcAbi,
		ContractAddress: multicallAddress,
	}
}

func (caller *TEthMultiCaller) execute(
	multiCallGroup []contracts.Multicall2Call,
	blockNumber *big.Int,
) ([]byte, error) {
	abi, _ := contracts.Multicall2MetaData.GetAbi()
	callData, err := abi.Pack("tryAggregate", false, multiCallGroup)

	if err != nil {
		logs.Error(`failed to pack tryAggregate: ` + err.Error())
		return []byte{}, err
	}

	// Perform multicall
	resp, err := caller.Client.CallContract(
		context.Background(),
		ethereum.CallMsg{
			To:   &caller.ContractAddress,
			Data: callData,
			Gas:  0,
			From: caller.Signer.From,
		},
		blockNumber,
	)
	if err != nil {
		chainID, _ := caller.Client.ChainID(context.Background())
		return []byte{}, errors.New("Failed to perform multicall for: " + chainID.String() + " | " + err.Error())
	}
	return resp, nil
}

// ExecuteByBatch will take a group of calls, split them in fixed-size group to
// avoid the gasLimit error, and execute as many transactions as required to get
// the results
func (caller *TEthMultiCaller) ExecuteByBatch(
	calls []Call,
	batchSize int,
	blockNumber *big.Int,
) map[string][]interface{} {
	var responses []CallResponse
	// Create mapping for results. Be aware that we sometimes get two empty results initially, unsure why
	results := make(map[string][]interface{})

	// Add calls to multicall structure for the contract
	var multiCalls = make([]contracts.Multicall2Call, 0, len(calls))
	var rawCalls = make([]Call, 0, len(calls))
	for _, call := range calls {
		multiCalls = append(multiCalls, call.GetMultiCall())
		rawCalls = append(rawCalls, call)
	}

	for i := 0; i < len(multiCalls); i += batchSize {
		var group []contracts.Multicall2Call
		var rawCallsGroup []Call

		if i+batchSize >= len(multiCalls) {
			group = multiCalls[i:]
			rawCallsGroup = rawCalls[i:]
		} else {
			group = multiCalls[i : i+batchSize]
			rawCallsGroup = rawCalls[i : i+batchSize]
		}
		_ = rawCallsGroup

		tempPackedResp, err := caller.execute(group, blockNumber)
		if err != nil {
			//print associated rawCalls
			// logs.Error(rawCallsGroup)
			LIMIT_ERROR := strings.Contains(strings.ToLower(err.Error()), "call retuned result on length") && strings.Contains(strings.ToLower(err.Error()), "exceeding limit")
			SIZE_ERROR := strings.Contains(strings.ToLower(err.Error()), "request entity too large")
			OUT_OF_GAS_ERROR := strings.Contains(strings.ToLower(err.Error()), "out of gas")
			isAssumingOutOfGas := false

			if LIMIT_ERROR {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Warning("Multicall gas limit error, retrying with smaller batch size", "batchSize", batchSize)
				}
			} else if SIZE_ERROR {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Warning("Multicall size error, retrying with smaller batch size", "batchSize", batchSize)
				}
			} else if OUT_OF_GAS_ERROR {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Warning("Multicall out of gas error, retrying with smaller batch size", "batchSize", batchSize)
				}
			} else {
				//assume it's out of gas for a few tries
				isAssumingOutOfGas = true
			}

			//check if error is a request entity too large
			if LIMIT_ERROR || SIZE_ERROR || OUT_OF_GAS_ERROR || isAssumingOutOfGas {
				if batchSize == math.MaxInt64 {
					return caller.ExecuteByBatch(calls, 10000, blockNumber)
				}
				if isAssumingOutOfGas && batchSize <= 1 {
					for i := 0; i < 100; i++ {
						pc, file, line, ok := runtime.Caller(i)
						if ok {
							if strings.Contains(file, "/multicall.go") || strings.Contains(file, "/multicalls/perform.go") {
								continue
							}
							logs.Error(`Multicall failed at [` + runtime.FuncForPC(pc).Name() + `:` + strconv.Itoa(line) + `]: ` + err.Error())
							break
						}
					}
					return nil
				}
				if batchSize <= 1 {
					logs.Error(`Multicall failed: ` + err.Error())
					for i := 0; i < 100; i++ {
						pc, file, line, ok := runtime.Caller(i)
						if ok {
							if strings.Contains(file, "/multicall.go") || strings.Contains(file, "/multicalls/perform.go") {
								continue
							}
							logs.Error(`Multicall failed at ` + runtime.FuncForPC(pc).Name() + `:` + strconv.Itoa(line) + `: ` + err.Error())
							break
						}
					}
					return nil
				}
				return caller.ExecuteByBatch(calls, batchSize/2, blockNumber)
			} else {
				logs.Error(err)
				//sleep a few ms and retry
				time.Sleep(2000 * time.Millisecond)
				return caller.ExecuteByBatch(calls, batchSize, blockNumber)
			}
		}

		// Unpack results
		unpackedResp, err := caller.Abi.Unpack("tryAggregate", tempPackedResp)
		if err != nil {
			logs.
				Capture(`error`, `failed to unpack response`).
				SetEntity(`multicall`).
				SetExtra(`error`, err.Error()).
				SetExtra(`tempPackedResp`, string(tempPackedResp)).
				SetTag(`blockNumber`, blockNumber.String()).
				Send()
			continue
		}

		a, err := json.Marshal(unpackedResp[0])
		if err != nil {
			logs.
				Capture(`error`, `failed to marshal response`).
				SetEntity(`multicall`).
				SetExtra(`error`, err.Error()).
				SetTag(`blockNumber`, blockNumber.String()).
				Send()
			continue
		}

		// Unpack results
		var tempResp []CallResponse
		if err := json.Unmarshal(a, &tempResp); err != nil {
			logs.
				Capture(`error`, `failed to unmarshal response`).
				SetEntity(`multicall`).
				SetExtra(`error`, err.Error()).
				SetTag(`blockNumber`, blockNumber.String()).
				Send()
			continue
		}

		responses = append(responses, tempResp...)
	}

	for i, response := range responses {
		if response.ReturnData != nil {
			unpacked, err := calls[i].Abi.Unpack(calls[i].Method, response.ReturnData)
			if err != nil {
				// logs.Warning("Failed to unpack method " + calls[i].Method + " for " + calls[i].Name + " : " + err.Error())
				results[calls[i].Name+calls[i].Method] = nil
			} else {
				results[calls[i].Name+calls[i].Method] = unpacked
			}
		} else {
			results[calls[i].Name+calls[i].Method] = nil
		}
	}

	return results
}
