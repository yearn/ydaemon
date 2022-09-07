package ethereum

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yearn/ydaemon/internal/contracts"
	"github.com/yearn/ydaemon/internal/logs"
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
	return contracts.Multicall2Call{Target: call.Target, CallData: call.CallData}
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
		logs.Error("Failed to connect to Ethereum node")
		time.Sleep(time.Second)
		return NewMulticall(rpcURI, multicallAddress)
	}

	// Load Multicall abi for later use
	mcAbi, err := contracts.Multicall2MetaData.GetAbi()
	if err != nil {
		logs.Error("Failed to decode Multicall ABI")
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

func (caller *TEthMultiCaller) execute(multiCallGroup []contracts.Multicall2Call) ([]byte, error) {
	// Prepare calldata for multicall
	abi, _ := contracts.Multicall2MetaData.GetAbi()
	callData, err := abi.Pack("tryAggregate", false, multiCallGroup)
	if err != nil {
		logs.Error("Failed to pack tryAggregate")
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
		nil,
	)
	if err != nil {
		logs.Error("Failed to perform multicall: " + err.Error())
		return []byte{}, err
	}
	return resp, nil
}

// ExecuteByBatch will take a group of calls, split them in fixed-size group to
// avoid the gasLimit error, and execute as many transactions as required to get
// the results.
func (caller *TEthMultiCaller) ExecuteByBatch(calls []Call, batchSize int) map[string][]interface{} {
	var responses []CallResponse
	// Create mapping for results. Be aware that we sometimes get two empty results initially, unsure why
	results := make(map[string][]interface{})

	// Add calls to multicall structure for the contract
	var multiCalls = make([]contracts.Multicall2Call, 0, len(calls))
	for _, call := range calls {
		multiCalls = append(multiCalls, call.GetMultiCall())
	}

	for i := 0; i < len(multiCalls); i += batchSize {
		var group []contracts.Multicall2Call

		if i+batchSize >= len(multiCalls) {
			group = multiCalls[i:]
		} else {
			group = multiCalls[i : i+batchSize]
		}

		tempPackedResp, err := caller.execute(group)
		if err != nil {
			logs.Error(err)
			continue
		}

		// Unpack results
		unpackedResp, _ := caller.Abi.Unpack("tryAggregate", tempPackedResp)
		a, err := json.Marshal(unpackedResp[0])
		if err != nil {
			logs.Error("Failed to unmarshal response: " + err.Error())
			continue
		}

		// Unpack results
		var tempResp []CallResponse
		if err := json.Unmarshal(a, &tempResp); err != nil {
			logs.Error("Failed to unmarshal response: " + err.Error())
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
