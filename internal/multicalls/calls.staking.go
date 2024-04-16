package multicalls

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
)

var StakingABI, _ = contracts.YOptimismStakingRewardMetaData.GetAbi()
var JuicedStakingABI, _ = contracts.JuicedStakingRewardsMetaData.GetAbi()

func GetPeriodFinish(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := StakingABI.Pack("periodFinish")
	if err != nil {
		logs.Error("Error packing StakingABI periodFinish", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      StakingABI,
		Method:   `periodFinish`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetRewardRate(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := StakingABI.Pack("rewardRate")
	if err != nil {
		logs.Error("Error packing StakingABI rewardRate", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      StakingABI,
		Method:   `rewardRate`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetRewardsToken(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := StakingABI.Pack("rewardsToken")
	if err != nil {
		logs.Error("Error packing StakingABI rewardsToken", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      StakingABI,
		Method:   `rewardsToken`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetRewardToken(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := StakingABI.Pack("rewardToken")
	if err != nil {
		logs.Error("Error packing StakingABI rewardToken", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      StakingABI,
		Method:   `rewardToken`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetRewardTokens(name string, contractAddress common.Address, index int64) ethereum.Call {
	parsedData, err := JuicedStakingABI.Pack("rewardTokens", big.NewInt(index))
	if err != nil {
		logs.Error("Error packing JuicedStakingABI rewardTokens", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      JuicedStakingABI,
		Method:   `rewardTokens`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetRewardData(name string, contractAddress common.Address, rewardTokenAddress common.Address) ethereum.Call {
	parsedData, err := JuicedStakingABI.Pack("rewardData", rewardTokenAddress)
	if err != nil {
		logs.Error("Error packing JuicedStakingABI rewardData", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      JuicedStakingABI,
		Method:   `rewardData`,
		CallData: parsedData,
		Name:     name,
	}
}
