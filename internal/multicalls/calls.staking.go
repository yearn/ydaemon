package multicalls

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
)

var StakingABI, _ = contracts.YOptimismStakingRewardMetaData.GetAbi()

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
