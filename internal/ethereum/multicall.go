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
	"github.com/majorfi/ydaemon/internal/contracts"
	"github.com/majorfi/ydaemon/internal/logs"
)

type Call struct {
	Name     string         `json:"name"`
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
	logs.Pretty(rpcURI)
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

// Execute will take a group of calls and execute them in a single transaction via
// the multicall3 contract.
func (caller *TEthMultiCaller) Execute(calls []Call) map[string]CallResponse {
	var responses []CallResponse
	// Create mapping for results. Be aware that we sometimes get two empty results initially, unsure why
	results := make(map[string]CallResponse)

	var multiCalls = make([]contracts.Multicall2Call, 0, len(calls))

	// Add calls to multicall structure for the contract
	for _, call := range calls {
		multiCalls = append(multiCalls, call.GetMultiCall())
	}

	// Prepare calldata for multicall
	abi, _ := contracts.Multicall3MetaData.GetAbi()
	callData, err := abi.Pack("tryAggregate", true, multiCalls)
	if err != nil {
		logs.Error("Failed to pack tryAggregate")
		return results
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
		return results
	}

	// Unpack results
	unpackedResp, _ := caller.Abi.Unpack("tryAggregate", resp)
	a, err := json.Marshal(unpackedResp[0])
	if err != nil {
		logs.Error("Failed to unmarshal response: " + err.Error())
		return results
	}

	// Unpack results
	if err := json.Unmarshal(a, &responses); err != nil {
		logs.Error("Failed to unmarshal response: " + err.Error())
		return results
	}

	for i, response := range responses {
		results[calls[i].Name] = response
	}

	return results
}

// ExecuteByBatch will take a group of calls, split them in fixed-size group to
// avoid the gasLimit error, and execute as many transactions as required to get
// the results.
// Note: should be used on Fantom for example.
func (caller *TEthMultiCaller) ExecuteByBatch(calls []Call, batchSize int) map[string]CallResponse {
	var responses []CallResponse
	// Create mapping for results. Be aware that we sometimes get two empty results initially, unsure why
	results := make(map[string]CallResponse)

	var multiCalls = make([]contracts.Multicall2Call, 0, len(calls))

	// Add calls to multicall structure for the contract
	for _, call := range calls {
		multiCalls = append(multiCalls, call.GetMultiCall())
	}

	var resp []byte
	for i := 0; i < len(multiCalls); i += batchSize {
		var group []contracts.Multicall2Call

		if i+batchSize >= len(multiCalls) {
			group = multiCalls[i:]
		} else {
			group = multiCalls[i : i+batchSize]
		}

		// Prepare calldata for multicall
		abi, _ := contracts.Multicall3MetaData.GetAbi()
		callData, err := abi.Pack("tryAggregate", true, group)
		if err != nil {
			logs.Error("Failed to pack tryAggregate")
			return results
		}

		// Perform multicall
		tempresp, err := caller.Client.CallContract(
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
			return results
		}
		// append tempresp to resp
		resp = append(resp, tempresp...)
	}

	// Unpack results
	unpackedResp, _ := caller.Abi.Unpack("tryAggregate", resp)
	a, err := json.Marshal(unpackedResp[0])
	if err != nil {
		logs.Error("Failed to unmarshal response: " + err.Error())
		return results
	}

	// Unpack results
	if err := json.Unmarshal(a, &responses); err != nil {
		logs.Error("Failed to unmarshal response: " + err.Error())
		return results
	}

	for i, response := range responses {
		results[calls[i].Name] = response
	}

	return results
}
