package multicalls

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
)

var VeloSugarABI, _ = contracts.VeloSugarOracleMetaData.GetAbi()

func GetManyRatesWithConnectors(name string, contractAddress common.Address, len uint8, connectors []common.Address) ethereum.Call {
	parsedData, err := VeloSugarABI.Pack("getManyRatesWithConnectors", len, connectors)
	if err != nil {
		logs.Error("Error packing VeloSugarABI getManyRatesWithConnectors", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      VeloSugarABI,
		Method:   `getManyRatesWithConnectors`,
		CallData: parsedData,
		Name:     name,
	}
}
