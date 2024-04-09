package ethereum

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
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

const SHOULD_LOG_WARNINGS = false

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

func (call Call) GetMultiCall() contracts.Multicall3Call {
	return contracts.Multicall3Call{
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
		logs.Error(err)
		time.Sleep(time.Second)
		return NewMulticall(rpcURI, multicallAddress)
	}

	// Load Multicall abi for later use
	mcAbi, err := contracts.Multicall3MetaData.GetAbi()
	if err != nil {
		logs.Error(err)
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
	multiCallGroup []contracts.Multicall3Call,
	blockNumber *big.Int,
) ([]byte, error) {
	abi, _ := contracts.Multicall3MetaData.GetAbi()
	callData, err := abi.Pack("tryAggregate", false, multiCallGroup)
	if err != nil {
		return []byte{}, err
	}
	// Perform multicall
	resp, err := caller.Client.CallContract(
		context.Background(),
		ethereum.CallMsg{
			To:   &caller.ContractAddress,
			Data: callData,
		},
		blockNumber,
	)
	if err != nil {
		chainID, _ := caller.Client.ChainID(context.Background())
		return []byte{}, errors.New("Failed to perform multicall for:" + chainID.String() + " | " + err.Error())
	}
	return resp, nil
}

// ExecuteByBatch will take a group of calls, split them in fixed-size group to
// avoid the gasLimit error, and execute as many transactions as required to get
// the results
func (caller *TEthMultiCaller) ExecuteByBatch(
	calls []Call,
	batchSize uint64,
	blockNumber *big.Int,
) map[string][]interface{} {
	if caller.Client == nil {
		logs.Error("No client provided.")
		return nil
	}
	var initialBatchSize = batchSize
	var responses []CallResponse
	// Create mapping for results. Be aware that we sometimes get two empty results initially, unsure why
	results := make(map[string][]interface{})

	// Add calls to multicall structure for the contract
	var multiCalls = make([]contracts.Multicall3Call, 0, len(calls))
	var rawCalls = make([]Call, 0, len(calls))
	for _, call := range calls {
		multiCalls = append(multiCalls, call.GetMultiCall())
		rawCalls = append(rawCalls, call)
	}

	//Store the chainID for logging
	chainID, _ := caller.Client.ChainID(context.Background())
	chainIDStr := "unknown"
	if chainID != nil {
		chainIDStr = strconv.Itoa(int(chainID.Int64()))
	}

	for i := uint64(0); i < uint64(len(multiCalls)); {
		var group []contracts.Multicall3Call
		var rawCallsGroup []Call
		if i >= uint64(len(multiCalls)) {
			break
		} else if (i + batchSize) >= uint64(len(multiCalls)) {
			group = multiCalls[i:]
			rawCallsGroup = rawCalls[i:]
		} else {
			group = multiCalls[i : i+batchSize]
			rawCallsGroup = rawCalls[i : i+batchSize]
		}

		_ = rawCallsGroup

		tempPackedResp, err := caller.execute(group, blockNumber)
		if err != nil {
			LIMIT_ERROR := strings.Contains(strings.ToLower(err.Error()), "call retuned result on length") && strings.Contains(strings.ToLower(err.Error()), "exceeding limit")
			SIZE_ERROR := strings.Contains(strings.ToLower(err.Error()), "request entity too large")
			OUT_OF_GAS_ERROR := strings.Contains(strings.ToLower(err.Error()), "out of gas")
			isAssumingOutOfGas := false

			if LIMIT_ERROR {
				if SHOULD_LOG_WARNINGS {
					logs.Warning("Multicall gas limit error, retrying with smaller batch size: " + strconv.FormatUint(batchSize, 10) + " on chain " + chainID.Text(10))
				}
			} else if SIZE_ERROR {
				if SHOULD_LOG_WARNINGS {
					logs.Warning("Multicall size error, retrying with smaller batch size: " + strconv.FormatUint(batchSize, 10) + " on chain " + chainID.Text(10))
				}
			} else if OUT_OF_GAS_ERROR {
				if SHOULD_LOG_WARNINGS {
					logs.Warning("Multicall out of gas error, retrying with smaller batch size: " + strconv.FormatUint(batchSize, 10) + " on chain " + chainID.Text(10))
				}
			} else {
				//assume it's out of gas for a few tries
				isAssumingOutOfGas = true
			}

			//check if error is a request entity too large
			if LIMIT_ERROR || SIZE_ERROR || OUT_OF_GAS_ERROR || isAssumingOutOfGas {
				if batchSize == math.MaxInt64 {
					batchSize = 10_000
					continue
				}
				if batchSize <= 1 {
					if SHOULD_LOG_WARNINGS {
						logs.Error(`Multicall failed on chain ` + chainIDStr + `! See error: ` + err.Error())
					}
					return nil
				}
				if isAssumingOutOfGas && SHOULD_LOG_WARNINGS {
					logs.Error(`Multicall failed on chain ` + chainIDStr + `! See error: ` + err.Error())
				}
				batchSize = batchSize / 2
				continue
			} else {
				logs.Error(err)
				//sleep a few ms and retry
				time.Sleep(2000 * time.Millisecond)
				if SHOULD_LOG_WARNINGS {
					logs.Warning(`Retrying with initial batch size of ` + strconv.FormatUint(initialBatchSize, 10))
				}
				continue
			}
		}

		if len(tempPackedResp) == 0 {
			logs.Error("Empty response from multicall for " + strconv.FormatInt(int64(len(group)), 10) + " calls")
			if batchSize == math.MaxInt64 {
				batchSize = 10_000
				continue
			}
			if batchSize <= 1 {
				if SHOULD_LOG_WARNINGS {
					logs.Error(`Multicall failed on chain ` + chainIDStr + `! See error: ` + err.Error())
				}
				return nil
			}
			batchSize = batchSize / 2
			continue
		}

		// Unpack results
		unpackedResp, err := caller.Abi.Unpack("tryAggregate", tempPackedResp)
		if err != nil {
			logs.Error(err)
			continue
		}

		a, err := json.Marshal(unpackedResp[0])
		if err != nil {
			logs.Error(err)
			continue
		}

		// Unpack results
		var tempResp []CallResponse
		if err := json.Unmarshal(a, &tempResp); err != nil {
			logs.Error(err)
			continue
		}

		responses = append(responses, tempResp...)
		i += batchSize
		if batchSize != initialBatchSize {
			batchSize = initialBatchSize
		}
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
