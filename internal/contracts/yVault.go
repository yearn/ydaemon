// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// YearnVaultMetaData contains all meta data concerning the YearnVault contract.
var YearnVaultMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"debtRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minDebtPerHarvest\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maxDebtPerHarvest\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"performanceFee\",\"type\":\"uint256\"}],\"name\":\"StrategyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"gain\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"loss\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"debtPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalGain\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalLoss\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"debtAdded\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"debtRatio\",\"type\":\"uint256\"}],\"name\":\"StrategyReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"governance\",\"type\":\"address\"}],\"name\":\"UpdateGovernance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"management\",\"type\":\"address\"}],\"name\":\"UpdateManagement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guestList\",\"type\":\"address\"}],\"name\":\"UpdateGuestList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"rewards\",\"type\":\"address\"}],\"name\":\"UpdateRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateDepositLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"performanceFee\",\"type\":\"uint256\"}],\"name\":\"UpdatePerformanceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"managementFee\",\"type\":\"uint256\"}],\"name\":\"UpdateManagementFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"UpdateGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"EmergencyShutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"queue\",\"type\":\"address[20]\"}],\"name\":\"UpdateWithdrawalQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"debtRatio\",\"type\":\"uint256\"}],\"name\":\"StrategyUpdateDebtRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"minDebtPerHarvest\",\"type\":\"uint256\"}],\"name\":\"StrategyUpdateMinDebtPerHarvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"maxDebtPerHarvest\",\"type\":\"uint256\"}],\"name\":\"StrategyUpdateMaxDebtPerHarvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"performanceFee\",\"type\":\"uint256\"}],\"name\":\"StrategyUpdatePerformanceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldVersion\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"StrategyMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyRemovedFromQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyAddedToQueue\",\"type\":\"event\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"nameOverride\",\"type\":\"string\"},{\"name\":\"symbolOverride\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"nameOverride\",\"type\":\"string\"},{\"name\":\"symbolOverride\",\"type\":\"string\"},{\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":4546,\"inputs\":[],\"name\":\"apiVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"gas\":107044,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":71894,\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"setSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":36365,\"inputs\":[{\"name\":\"governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37637,\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37775,\"inputs\":[{\"name\":\"management\",\"type\":\"address\"}],\"name\":\"setManagement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37805,\"inputs\":[{\"name\":\"guestList\",\"type\":\"address\"}],\"name\":\"setGuestList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37835,\"inputs\":[{\"name\":\"rewards\",\"type\":\"address\"}],\"name\":\"setRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":36519,\"inputs\":[{\"name\":\"degration\",\"type\":\"uint256\"}],\"name\":\"setLockedProfitDegration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37795,\"inputs\":[{\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setDepositLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37929,\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":37959,\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setManagementFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":39203,\"inputs\":[{\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"setGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":39274,\"inputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"setEmergencyShutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":763950,\"inputs\":[{\"name\":\"queue\",\"type\":\"address[20]\"}],\"name\":\"setWithdrawalQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":76768,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":116531,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":38271,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":40312,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":40336,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":81264,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"permit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":4098,\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":383839,\"inputs\":[],\"name\":\"maxAvailableShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"maxShares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"maxShares\",\"type\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"maxShares\",\"type\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":18195,\"inputs\":[],\"name\":\"pricePerShare\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1485796,\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"debtRatio\",\"type\":\"uint256\"},{\"name\":\"minDebtPerHarvest\",\"type\":\"uint256\"},{\"name\":\"maxDebtPerHarvest\",\"type\":\"uint256\"},{\"name\":\"performanceFee\",\"type\":\"uint256\"}],\"name\":\"addStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":115193,\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"debtRatio\",\"type\":\"uint256\"}],\"name\":\"updateStrategyDebtRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":42441,\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"minDebtPerHarvest\",\"type\":\"uint256\"}],\"name\":\"updateStrategyMinDebtPerHarvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":42471,\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"maxDebtPerHarvest\",\"type\":\"uint256\"}],\"name\":\"updateStrategyMaxDebtPerHarvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":41251,\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"performanceFee\",\"type\":\"uint256\"}],\"name\":\"updateStrategyPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":1141468,\"inputs\":[{\"name\":\"oldVersion\",\"type\":\"address\"},{\"name\":\"newVersion\",\"type\":\"address\"}],\"name\":\"migrateStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"revokeStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":1199804,\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"addStrategyToQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":23088703,\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"removeStrategyFromQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debtOutstanding\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"debtOutstanding\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"creditAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":9551,\"inputs\":[],\"name\":\"availableDepositLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"expectedReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":1015170,\"inputs\":[{\"name\":\"gain\",\"type\":\"uint256\"},{\"name\":\"loss\",\"type\":\"uint256\"},{\"name\":\"_debtPayment\",\"type\":\"uint256\"}],\"name\":\"report\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"gas\":8750,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":7803,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2408,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2438,\"inputs\":[],\"name\":\"precisionFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2683,\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2928,\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2528,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2558,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2588,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2618,\"inputs\":[],\"name\":\"management\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2648,\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2678,\"inputs\":[],\"name\":\"guestList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":11031,\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"strategies\",\"outputs\":[{\"name\":\"performanceFee\",\"type\":\"uint256\"},{\"name\":\"activation\",\"type\":\"uint256\"},{\"name\":\"debtRatio\",\"type\":\"uint256\"},{\"name\":\"minDebtPerHarvest\",\"type\":\"uint256\"},{\"name\":\"maxDebtPerHarvest\",\"type\":\"uint256\"},{\"name\":\"lastReport\",\"type\":\"uint256\"},{\"name\":\"totalDebt\",\"type\":\"uint256\"},{\"name\":\"totalGain\",\"type\":\"uint256\"},{\"name\":\"totalLoss\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2847,\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"withdrawalQueue\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2768,\"inputs\":[],\"name\":\"emergencyShutdown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2798,\"inputs\":[],\"name\":\"depositLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2828,\"inputs\":[],\"name\":\"debtRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2858,\"inputs\":[],\"name\":\"totalDebt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2888,\"inputs\":[],\"name\":\"lastReport\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2918,\"inputs\":[],\"name\":\"activation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2948,\"inputs\":[],\"name\":\"lockedProfit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":2978,\"inputs\":[],\"name\":\"lockedProfitDegration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":3008,\"inputs\":[],\"name\":\"rewards\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":3038,\"inputs\":[],\"name\":\"managementFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":3068,\"inputs\":[],\"name\":\"performanceFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":3313,\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"gas\":3128,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61542f56600436101561000d57613d3b565b600035601c52600051341561002157600080fd5b6383b43589811415610037573361022052610068565b63a5b81fdf8114156100635760a43560a01c1561005357600080fd5b602060a461022037600050610068565b610705565b60043560a01c1561007857600080fd5b60243560a01c1561008857600080fd5b60443560a01c1561009857600080fd5b60606064356004016101403760406064356004013511156100b857600080fd5b60406084356004016101c03760206084356004013511156100d857600080fd5b601454156100e557600080fd5b600435600755600061028052610280805160208201209050610140805160208201209050141561022c576000606061032060046395d89b416102c0526102dc6004355afa61013257600080fd5b603f3d1161013f57600080fd5b60156103206103205101511061015457600080fd5b6000506103406014806020846103e001018260208501600060045af15050805182019150506007610380527f20795661756c74000000000000000000000000000000000000000000000000006103a0526103806007806020846103e001018260208501600060045af1505080518201915050806103e0526103e0905080600060c052602060c020602082510161012060006002818352015b826101205160200211156101ff57610221565b61012051602002850151610120518501555b81516001018083528114156101ec575b505050505050610287565b61014080600060c052602060c020602082510161012060006003818352015b8261012051602002111561025e57610280565b61012051602002850151610120518501555b815160010180835281141561024b575b5050505050505b6000610280526102808051602082012090506101c080516020820120905014156103c857600060026102c0527f79760000000000000000000000000000000000000000000000000000000000006102e0526102c06002806020846103e001018260208501600060045af1505080518201915050606061038060046395d89b416103205261033c6004355afa61031b57600080fd5b603f3d1161032857600080fd5b60156103806103805101511061033d57600080fd5b6000506103a06014806020846103e001018260208501600060045af1505080518201915050806103e0526103e0905080600160c052602060c020602082510161012060006002818352015b8261012051602002111561039b576103bd565b61012051602002850151610120518501555b8151600101808352811415610388575b505050505050610423565b6101c080600160c052602060c020602082510161012060006002818352015b826101205160200211156103fa5761041c565b61012051602002850151610120518501555b81516001018083528114156103e7575b5050505050505b60206102a0600463313ce5676102405261025c6004355afa61044457600080fd5b601f3d1161045157600080fd5b6000506102a051600255601260025410156104af57604e60126002548082101561047a57600080fd5b808203905090501061048b57600080fd5b60126002548082101561049d57600080fd5b80820390509050600a0a6003556104b5565b60016003555b602435600855602435610240527f8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b876020610240a1602435600955602435610240527fff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef4386020610240a1604435601755604435610240527fdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d6020610240a161022051600a5561022051610240527f837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c6020610240a16103e86019556103e86102405261024051610260527f0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb6020610260a160c860185560c86102405261024051610260527f7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf6020610260a142601355426014556529d635a8e00060165560007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f602082610620010152602081019050600b610500527f596561726e205661756c740000000000000000000000000000000000000000006105205261050080516020820120905060208261062001015260208101905060056105c0527f302e332e350000000000000000000000000000000000000000000000000000006105e0526105c0805160208201209050602082610620010152602081019050466020826106200101526020810190503060208261062001015260208101905080610620526106209050805160208201209050601b55005b632582941081141561079a576005610140527f302e332e35000000000000000000000000000000000000000000000000000000610160526101408051602001806101e08284600060045af161075957600080fd5b50506101e0518061020001818260206001820306601f820103905003368237505060206101c05260406101e0510160206001820306601f82010390506101c0f35b63c47f002781141561083057604a60043560040161014037602a6004356004013511156107c657600080fd5b60085433146107d457600080fd5b61014080600060c052602060c020602082510161012060006003818352015b8261012051602002111561080657610828565b61012051602002850151610120518501555b81516001018083528114156107f3575b505050505050005b63b84c82468114156108c657603460043560040161014037601460043560040135111561085c57600080fd5b600854331461086a57600080fd5b61014080600160c052602060c020602082510161012060006002818352015b8261012051602002111561089c576108be565b61012051602002850151610120518501555b8151600101808352811415610889575b505050505050005b63ab033ea98114156108f85760043560a01c156108e257600080fd5b60085433146108f057600080fd5b600435600b55005b63238efcbc81141561094457600b54331461091257600080fd5b3360085533610140527f8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b876020610140a1005b63d4a22bde8114156109a45760043560a01c1561096057600080fd5b600854331461096e57600080fd5b600435600955600435610140527fff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef4386020610140a1005b630b5b78eb811415610a045760043560a01c156109c057600080fd5b60085433146109ce57600080fd5b600435600c55600435610140527f6d674c311329fb38bbc96dc33d2aad03b9bf9fcfdd8f5e5054fda291a5b3c1f86020610140a1005b63ec38a862811415610a645760043560a01c15610a2057600080fd5b6008543314610a2e57600080fd5b600435601755600435610140527fdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d6020610140a1005b638402a84f811415610a9d576008543314610a7e57600080fd5b670de0b6b3a76400006004351115610a9557600080fd5b600435601655005b63bdc8144b811415610aed576008543314610ab757600080fd5b600435601055600435610140527fae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd626020610140a1005b6370897b23811415610b4e576008543314610b0757600080fd5b6127106004351115610b1857600080fd5b600435601955600435610140527f0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb6020610140a1005b63fe56e232811415610baf576008543314610b6857600080fd5b6127106004351115610b7957600080fd5b600435601855600435610140527f7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf6020610140a1005b638a0dac4a811415610c605760043560a01c15610bcb57600080fd5b600a54610160526008546101805260006101405261014061012060006002818352015b610120516020026101600151331415610c0a5760018352610c1a565b8151600101808352811415610bee575b50505061014051610c2a57600080fd5b600435600a55600435610140527f837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c6020610140a1005b6314c64402811415610d2c5760043560011c15610c7c57600080fd5b60043515610ce857600a54610160526008546101805260006101405261014061012060006002818352015b610120516020026101600151331415610cc35760018352610cd3565b8151600101808352811415610ca7575b50505061014051610ce357600080fd5b610cf6565b6008543314610cf657600080fd5b600435600f55600435610140527fba40372a3a724dca3c57156128ef1e896724b65b37a17f190b1ad5de68f3a4f36020610140a1005b6394148415811415610f71576000610120525b610120516004013560a01c15610d5457600080fd5b6020610120510161012052610280610120511015610d7157610d3f565b600954610160526008546101805260006101405261014061012060006002818352015b610120516020026101600151331415610db05760018352610dc0565b8151600101808352811415610d94575b50505061014051610dd057600080fd5b61014060006014818352015b60046101405160148110610def57600080fd5b60200201351515610e1d576101405160148110610e0b57600080fd5b600e60c052602060c020015415610e20565b60005b15610e2a57610ead565b60006001600d60046101405160148110610e4357600080fd5b602002013560e05260c052604060c02060c052602060c020015411610e6757600080fd5b60046101405160148110610e7a57600080fd5b60200201356101405160148110610e9057600080fd5b600e60c052602060c02001555b8151600101808352811415610ddc575b50506004356101405260243561016052604435610180526064356101a0526084356101c05260a4356101e05260c4356102005260e43561022052610104356102405261012435610260526101443561028052610164356102a052610184356102c0526101a4356102e0526101c435610300526101e43561032052610204356103405261022435610360526102443561038052610264356103a0527f695ac3ac73f08f2002284ffe563cefe798ee2878a5e04219522e2e99eb89d168610280610140a1005b63a9059cbb811415610fc35760043560a01c15610f8d57600080fd5b336101405260043561016052602435610180526101805161016051610140516006580161408c565b600050600160005260206000f35b6323b872dd8114156110f85760043560a01c15610fdf57600080fd5b60243560a01c15610fef57600080fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600560043560e05260c052604060c0203360e05260c052604060c0205410156110c057600560043560e05260c052604060c0203360e05260c052604060c020546044358082101561106057600080fd5b808203905090506101405261014051600560043560e05260c052604060c0203360e05260c052604060c020556101405161016052336004357f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9256020610160a35b6004356101405260243561016052604435610180526101805161016051610140516006580161408c565b600050600160005260206000f35b63095ea7b38114156111715760043560a01c1561111457600080fd5b60243560053360e05260c052604060c02060043560e05260c052604060c0205560243561014052600435337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9256020610140a3600160005260206000f35b633950935181141561121e5760043560a01c1561118d57600080fd5b60053360e05260c052604060c02060043560e05260c052604060c02080546024358181830110156111bd57600080fd5b8082019050905081555060053360e05260c052604060c02060043560e05260c052604060c0205461014052600435337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9256020610140a3600160005260206000f35b63a457c2d78114156112c95760043560a01c1561123a57600080fd5b60053360e05260c052604060c02060043560e05260c052604060c02080546024358082101561126857600080fd5b8082039050905081555060053360e05260c052604060c02060043560e05260c052604060c0205461014052600435337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9256020610140a3600160005260206000f35b639fd5a6cf8114156116c65760043560a01c156112e557600080fd5b60243560a01c156112f557600080fd5b606160843560040161014037604160843560040135111561131557600080fd5b60006004351861132457600080fd5b606435151561133457600161133b565b4260643510155b61134457600080fd5b601a60043560e05260c052604060c020546101e05260006002610520527f19010000000000000000000000000000000000000000000000000000000000006105405261052060028060208461078001018260208501600060045af1505080518201915050601b5460208261078001015260208101905060007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c96020826106800101526020810190506004356020826106800101526020810190506024356020826106800101526020810190506044356020826106800101526020810190506101e05160208261068001015260208101905060643560208261068001015260208101905080610680526106809050805160208201209050602082610780010152602081019050806107805261078090508051602082012090506102005260006020602082066103000161014051828401111561149e57600080fd5b6041806103208260206020880688030161014001600060045af15050818152809050905090508060200151600082518060209013156114dc57600080fd5b80919012156114ea57600080fd5b806020036101000a82049050905090506102205260206020602082066103200161014051828401111561151c57600080fd5b6041806103408260206020880688030161014001600060045af150508181528090509050905080602001516000825180602090131561155a57600080fd5b809190121561156857600080fd5b806020036101000a82049050905090506102405260406001602082066103400161014051828401111561159a57600080fd5b6041806103608260206020880688030161014001600060045af15050818152809050905090508060200151600082518060209013156115d857600080fd5b80919012156115e657600080fd5b806020036101000a8204905090509050610260526004356102005161028052610260516102a052610220516102c052610240516102e052602060c0608061028060015afa5060c0511461163857600080fd5b604435600560043560e05260c052604060c02060243560e05260c052604060c020556101e051600181818301101561166f57600080fd5b80820190509050601a60043560e05260c052604060c02055604435610280526024356004357f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9256020610280a3600160005260206000f35b6301e1d1148114156116ec5760065801614197565b610140526101405160005260206000f35b63d0e30db0811415611727577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610140523361016052611781565b63b6b55f258114156117485733610160526020600461014037600050611781565b636e553f6581141561177c57602060046101403760243560a01c1561176c57600080fd5b6020602461016037600050611781565b6119d5565b601c541561178e57600080fd5b6001601c55600f54156117a057600080fd5b61014051610180527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61018051141561186a5760105461014051610160516101805160065801614197565b6101a0526101805261016052610140526101a0518082101561180c57600080fd5b80820390509050602061024060246370a082316101c052336101e0526101dc6007545afa61183957600080fd5b601f3d1161184657600080fd5b600050610240518082111561185b578061185d565b815b90509050610180526118bb565b60105461014051610160516101805160065801614197565b6101a0526101805261016052610140526101a051610180518181830110156118a957600080fd5b8082019050905011156118bb57600080fd5b600061018051116118cb57600080fd5b6000600c5418156119215760206102406044635ed7660e6101a052336101c052610180516101e0526101bc600c545afa61190457600080fd5b601f3d1161191157600080fd5b6000506102405161192157600080fd5b6101405161016051610180516101a051610160516101c052610180516101e0526101e0516101c051600658016141fa565b610240526101a052610180526101605261014052610240516101a0526101405161016051610180516101a0516007546101c052336101e0523061020052610180516102205261022051610200516101e0516101c05160065801613edb565b6101a0526101805261016052610140526000506101a0516000526000601c5560206000f35b6375de2902811415611b335760206101e060246370a0823161016052306101805261017c6007545afa611a0757600080fd5b601f3d11611a1457600080fd5b6000506101e051610200526101405161016051610180516101a0516101c0516101e0516102005161020051610220526102205160065801614532565b61028052610200526101e0526101c0526101a052610180526101605261014052610280516101405261018060006014818352015b61018051600e60c052602060c020015461016052610160511515611aa757611b24565b61014080516101405161016051610180516006600d6101605160e05260c052604060c02060c052602060c02001546101a0526101a05160065801614532565b6102005261018052610160526101405261020051818183011015611b0957600080fd5b808201905090508152505b8151600101808352811415611a84575b50506101405160005260206000f35b633ccfd60b811415611b74577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610140523361016052600161018052611c15565b632e1a7d4d811415611b9b5733610160526001610180526020600461014037600050611c15565b62f714ce811415611bd457600161018052602060046101403760243560a01c15611bc457600080fd5b6020602461016037600050611c15565b63e63697c8811415611c1057602060046101403760243560a01c15611bf857600080fd5b60206024610160376020604461018037600050611c15565b61224e565b601c5415611c2257600080fd5b6001601c55610140516101a052612710610180511115611c4157600080fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101a0511415611c805760043360e05260c052604060c020546101a0525b60043360e05260c052604060c020546101a0511115611c9e57600080fd5b60006101a05111611cae57600080fd5b6101405161016051610180516101a0516101c0516101a0516101e0526101e05160065801614377565b610240526101c0526101a052610180526101605261014052610240516101c05260006101e052602061028060246370a0823161020052306102205261021c6007545afa611d2357600080fd5b601f3d11611d3057600080fd5b600050610280516101c0511115611ff8576102c060006014818352015b6102c051600e60c052602060c02001546102a0526102a0511515611d7057611ff5565b602061038060246370a0823161030052306103205261031c6007545afa611d9657600080fd5b601f3d11611da357600080fd5b600050610380516102e0526102e0516101c051111515611dc257611ff5565b6101c0516102e05180821015611dd757600080fd5b8082039050905061030052610300516006600d6102a05160e05260c052604060c02060c052602060c020015480821115611e115780611e13565b815b9050905061030052610300511515611e2a57611fe5565b60206103c06024632e1a7d4d61034052610300516103605261035c60006102a0515af1611e5657600080fd5b601f3d11611e6357600080fd5b6000506103c0516103205260206103e060246370a0823161036052306103805261037c6007545afa611e9457600080fd5b601f3d11611ea157600080fd5b6000506103e0516102e05180821015611eb957600080fd5b80820390509050610340526000610320511115611f8c576101c080516103205180821015611ee657600080fd5b808203905090508152506101e0805161032051818183011015611f0857600080fd5b80820190509050815250610140610360525b61036051516020610360510161036052610360610360511015611f3c57611f1a565b6102a05161038052610320516103a0526103a0516103805160065801614622565b610340610360525b6103605152602061036051036103605261014061036051101515611f8857611f65565b6000505b6006600d6102a05160e05260c052604060c02060c052602060c0200180546103405180821015611fbb57600080fd5b80820390509050815550601280546103405180821015611fda57600080fd5b808203905090508155505b8151600101808352811415611d4d575b50505b60206102a060246370a0823161022052306102405261023c6007545afa61201e57600080fd5b601f3d1161202b57600080fd5b6000506102a05161020052610200516101c05111156120c057610200516101c0526101405161016051610180516101a0516101c0516101e051610200516101c0516101e05181818301101561207f57600080fd5b80820190509050610220526102205160065801614532565b61028052610200526101e0526101c0526101a052610180526101605261014052610280516101a0525b60035461022052610220516101805180820282158284830414176120e357600080fd5b809050905090506101c0516101e05181818301101561210157600080fd5b80820190509050808202821582848304141761211c57600080fd5b809050905090506127108082049050905061022051808061213c57600080fd5b8204905090506101e051111561215157600080fd5b600680546101a0518082101561216657600080fd5b8082039050905081555060043360e05260c052604060c02080546101a0518082101561219157600080fd5b808203905090508155506101a051610240526000337fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020610240a36101405161016051610180516101a0516101c0516101e05161020051610220516007546102405261016051610260526101c0516102805261028051610260516102405160065801613d41565b61022052610200526101e0526101c0526101a0526101805261016052610140526000506101c0516000526000601c5560206000f35b6399530b0681141561229157604e6002541061226957600080fd5b600254600a0a610140526101405160065801614377565b6101a0526101a05160005260206000f35b6314b4e26e8114156124d65760043560a01c156122ad57600080fd5b6013600e60c052602060c0200154156122c557600080fd5b600f54156122d257600080fd5b60085433146122e057600080fd5b6000600435186122ef57600080fd5b6001600d60043560e05260c052604060c02060c052602060c02001541561231557600080fd5b60206101a0600463fbfa77cf6101405261015c6004355afa61233657600080fd5b601f3d1161234357600080fd5b6000506101a051301461235557600080fd5b60206101a06004631f1fcd516101405261015c6004355afa61237657600080fd5b601f3d1161238357600080fd5b6000506101a0516007541461239757600080fd5b6127106011546024358181830110156123af57600080fd5b8082019050905011156123c157600080fd5b60643560443511156123d257600080fd5b612710601954808210156123e557600080fd5b8082039050905060843511156123fa57600080fd5b600d60043560e05260c052604060c02060c052602060c0206084358155426001820155602435600282015560443560038201556064356004820155426005820155600060068201556000600782015560006008820155506024356101405260443561016052606435610180526084356101a0526004357f5a6abd2af9fe6c0554fa08649e2d86e4393ff19dc304d072d38d295c9291d4dc6080610140a2601180546024358181830110156124ad57600080fd5b808201905090508155506004356013600e60c052602060c020015560065801614847565b600050005b637c6a4f248114156126335760043560a01c156124f257600080fd5b600954610160526008546101805260006101405261014061012060006002818352015b6101205160200261016001513314156125315760018352612541565b8151600101808352811415612515575b5050506101405161255157600080fd5b60006001600d60043560e05260c052604060c02060c052602060c02001541161257957600080fd5b601180546002600d60043560e05260c052604060c02060c052602060c0200154808210156125a657600080fd5b808203905090508155506024356002600d60043560e05260c052604060c02060c052602060c0200155601180546024358181830110156125e557600080fd5b80820190509050815550612710601154111561260057600080fd5b602435610140526004357fbda9398315c83ccef012bcaa318a2ff7b680f36429d36597bd4bc25ac11ead596020610140a2005b63e722befe8114156127525760043560a01c1561264f57600080fd5b600954610160526008546101805260006101405261014061012060006002818352015b61012051602002610160015133141561268e576001835261269e565b8151600101808352811415612672575b505050610140516126ae57600080fd5b60006001600d60043560e05260c052604060c02060c052602060c0200154116126d657600080fd5b6024356004600d60043560e05260c052604060c02060c052602060c0200154101561270057600080fd5b6024356003600d60043560e05260c052604060c02060c052602060c0200155602435610140526004357f0b728ad785976532c4aaadde09b1cba5f262a7090e83c62d2377bc405678b29c6020610140a2005b634757a1568114156128715760043560a01c1561276e57600080fd5b600954610160526008546101805260006101405261014061012060006002818352015b6101205160200261016001513314156127ad57600183526127bd565b8151600101808352811415612791575b505050610140516127cd57600080fd5b60006001600d60043560e05260c052604060c02060c052602060c0200154116127f557600080fd5b6024356003600d60043560e05260c052604060c02060c052602060c0200154111561281f57600080fd5b6024356004600d60043560e05260c052604060c02060c052602060c0200155602435610140526004357f1796a8e0760e2de5b72e7bf64fccb7666c48ceab94cb6cae7cb7eff4b6f641ab6020610140a2005b63d0194ed681141561293a5760043560a01c1561288d57600080fd5b600854331461289b57600080fd5b612710601954808210156128ae57600080fd5b8082039050905060243511156128c357600080fd5b60006001600d60043560e05260c052604060c02060c052602060c0200154116128eb57600080fd5b602435600d60043560e05260c052604060c02060c052602060c02055602435610140526004357fe57488a65fa53066d4c25bac90db47dda4e5de3025ac12bf76ff07211cf7f39e6020610140a2005b636cb56d19811415612c525760043560a01c1561295657600080fd5b60243560a01c1561296657600080fd5b600854331461297457600080fd5b60006024351861298357600080fd5b60006001600d60043560e05260c052604060c02060c052602060c0200154116129ab57600080fd5b6001600d60243560e05260c052604060c02060c052602060c0200154156129d157600080fd5b610140600d60043560e05260c052604060c0208060c052602060c02054825260018160c052602060c0200154826020015260028160c052602060c0200154826040015260038160c052602060c0200154826060015260048160c052602060c0200154826080015260058160c052602060c02001548260a0015260068160c052602060c02001548260c0015260078160c052602060c02001548260e0015260088160c052602060c020015482610100015250506101405161016051610180516101a0516101c0516101e05161020051610220516102405160043561026052610260516006580161492a565b6102405261022052610200526101e0526101c0526101a0526101805261016052610140526000506011805461018051818183011015612af957600080fd5b8082019050905081555060006006600d60043560e05260c052604060c02060c052602060c0200155600d60243560e05260c052604060c02060c052602060c0206101405181556101e05160018201556101805160028201556101a05160038201556101c05160048201556101e05160058201556102005160068201556000600782015560006008820155506004353b612b9157600080fd5b60006000602463ce5494bb610260526024356102805261027c60006004355af1612bba57600080fd5b6024356004357f100b69bb6b504e1252e36b375233158edee64d071b399e2f81473a695fd1b02160006000a361026060006014818352015b6004356102605160148110612c0657600080fd5b600e60c052602060c02001541415612c3e576024356102605160148110612c2c57600080fd5b600e60c052602060c020015560006000f35b8151600101808352811415612bf2575b5050005b63a0e4af9a811415612c68573361014052612c99565b63bb994d48811415612c945760043560a01c15612c8457600080fd5b6020600461014037600050612c99565b612d22565b61014051610180526008546101a052600a546101c05260006101605261016061012060006003818352015b610120516020026101800151331415612ce05760018352612cf0565b8151600101808352811415612cc4575b50505061016051612d0057600080fd5b610140516101405161016052610160516006580161492a565b61014052600050005b63f76e4caa811415612e9e5760043560a01c15612d3e57600080fd5b600954610160526008546101805260006101405261014061012060006002818352015b610120516020026101600151331415612d7d5760018352612d8d565b8151600101808352811415612d61575b50505061014051612d9d57600080fd5b60006001600d60043560e05260c052604060c02060c052602060c020015411612dc557600080fd5b60006101405261018060006014818352015b61018051600e60c052602060c020015461016052610160511515612dfa57612e3c565b6004356101605118612e0b57600080fd5b61014080516001818183011015612e2157600080fd5b808201905090508152505b8151600101808352811415612dd7575b505060146101405110612e4e57600080fd5b6004356013600e60c052602060c02001556101405160065801614847565b610140526000506004357fa8727d412c6fa1e2497d6d6f275e2d9fe4d9318d5b793632e60ad9d38ee8f1fa60006000a2005b63b22439f5811415612fc55760043560a01c15612eba57600080fd5b600954610160526008546101805260006101405261014061012060006002818352015b610120516020026101600151331415612ef95760018352612f09565b8151600101808352811415612edd575b50505061014051612f1957600080fd5b61014060006014818352015b6004356101405160148110612f3957600080fd5b600e60c052602060c02001541415612fad5760006101405160148110612f5e57600080fd5b600e60c052602060c02001556101405160065801614847565b610140526000506004357f8e1ec3c16d6a67ea8effe2ac7adef9c2de0bc0dc47c49cdf18f6a8b0048085be60006000a260006000f35b8151600101808352811415612f25575b505060006000fd5b63bf3759b5811415612fdb57336101405261300c565b63bdcf36bb8114156130075760043560a01c15612ff757600080fd5b602060046101403760005061300c565b61303a565b61014051610140516101605261016051600658016149b9565b6101c052610140526101c05160005260206000f35b63112c1f9b811415613050573361014052613081565b63d764801381141561307c5760043560a01c1561306c57600080fd5b6020600461014037600050613081565b6130af565b6101405161014051610160526101605160065801614afb565b6101c052610140526101c05160005260206000f35b63153c27c481141561311f5760065801614197565b61014052610140516010541115613112576010546101405160065801614197565b610160526101405261016051808210156130fe57600080fd5b8082039050905060005260206000f361311d565b600060005260206000f35b005b63d3406abd811415613135573361014052613166565b6333586b678114156131615760043560a01c1561315157600080fd5b6020600461014037600050613166565b613194565b6101405161014051610160526101605160065801614dcd565b6101c052610140526101c05160005260206000f35b63a1d9bafc8114156136cd5760006001600d3360e05260c052604060c02060c052602060c0200154116131c657600080fd5b6004356044358181830110156131db57600080fd5b8082019050905060206101c060246370a0823161014052336101605261015c6007545afa61320857600080fd5b601f3d1161321557600080fd5b6000506101c051101561322757600080fd5b6000602435111561325357336101405260243561016052610160516101405160065801614622565b6000505b61014051336101605260043561018052610180516101605160065801614f5b565b6101e052610140526101e051610140526007600d3360e05260c052604060c02060c052602060c0200180546004358181830110156132b157600080fd5b808201905090508155506101405161016051336101805261018051600658016149b9565b6101e05261016052610140526101e0516101605260443561016051808211156132fe5780613300565b815b9050905061018052600061018051111561338a576006600d3360e05260c052604060c02060c052602060c020018054610180518082101561334057600080fd5b8082039050905081555060128054610180518082101561335f57600080fd5b808203905090508155506101608051610180518082101561337f57600080fd5b808203905090508152505b6101405161016051610180516101a051336101c0526101c05160065801614afb565b610220526101a052610180526101605261014052610220516101a05260006101a051111561342e576006600d3360e05260c052604060c02060c052602060c0200180546101a05181818301101561340257600080fd5b80820190509050815550601280546101a05181818301101561342357600080fd5b808201905090508155505b6004356101805181818301101561344457600080fd5b808201905090506101c0526101a0516101c05110156134ce576101405161016051610180516101a0516101c0516007546101e05233610200526101a0516101c0518082101561349257600080fd5b808203905090506102205261022051610200516101e05160065801613d41565b6101c0526101a052610180526101605261014052600050613552565b6101a0516101c0511115613552576101405161016051610180516101a0516101c0516007546101e052336102005230610220526101c0516101a0518082101561351657600080fd5b80820390509050610240526102405161022051610200516101e05160065801613edb565b6101c0526101a0526101805261016052610140526000505b426005600d3360e05260c052604060c02060c052602060c020015542601355600435610140518082101561358557600080fd5b808203905090506015556004356101e0526024356102005261018051610220526007600d3360e05260c052604060c02060c052602060c0200154610240526008600d3360e05260c052604060c02060c052602060c0200154610260526006600d3360e05260c052604060c02060c052602060c0200154610280526101a0516102a0526002600d3360e05260c052604060c02060c052602060c02001546102c052337f67f96d2854a335a4cadb49f84fd3ca6f990744ddb3feceeb4b349d2d53d32ad36101006101e0a26002600d3360e05260c052604060c02060c052602060c02001541515613675576001613679565b600f545b156136be576020610240600463efbb5cb06101e0526101fc335afa61369d57600080fd5b601f3d116136aa57600080fd5b6000506102405160005260206000f36136cb565b6101605160005260206000f35b005b6301681a62811415613703577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61014052613724565b636ea056a981141561371f576020602461014037600050613724565b613804565b60043560a01c1561373457600080fd5b600854331461374257600080fd5b6007546004351861375257600080fd5b61014051610160527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101605114156137c457602061020060246370a0823161018052306101a05261019c6004355afa6137ab57600080fd5b601f3d116137b857600080fd5b60005061020051610160525b6101405161016051600435610180526008546101a052610160516101c0526101c0516101a0516101805160065801613d41565b6101605261014052600050005b6306fdde038114156138a95760008060c052602060c020610180602082540161012060006003818352015b8261012051602002111561384257613864565b61012051850154610120516020028501525b815160010180835281141561382f575b50505050505061018051806101a001818260206001820306601f82010390500336823750506020610160526040610180510160206001820306601f8201039050610160f35b6395d89b4181141561394e5760018060c052602060c020610180602082540161012060006002818352015b826101205160200211156138e757613909565b61012051850154610120516020028501525b81516001018083528114156138d4575b50505050505061018051806101a001818260206001820306601f82010390500336823750506020610160526040610180510160206001820306601f8201039050610160f35b63313ce5678114156139665760025460005260206000f35b639d902fc081141561397e5760035460005260206000f35b6370a082318114156139b45760043560a01c1561399a57600080fd5b600460043560e05260c052604060c0205460005260206000f35b63dd62ed3e811415613a085760043560a01c156139d057600080fd5b60243560a01c156139e057600080fd5b600560043560e05260c052604060c02060243560e05260c052604060c0205460005260206000f35b6318160ddd811415613a205760065460005260206000f35b63fc0c546a811415613a385760075460005260206000f35b635aa6e675811415613a505760085460005260206000f35b6388a8d602811415613a685760095460005260206000f35b63452a9320811415613a8057600a5460005260206000f35b6346d55875811415613a9857600c5460005260206000f35b6339ebf823811415613bb25760043560a01c15613ab457600080fd5b600d60043560e05260c052604060c0206101408080808460c052602060c0205481525050602081019050808060018560c052602060c020015481525050602081019050808060028560c052602060c020015481525050602081019050808060038560c052602060c020015481525050602081019050808060048560c052602060c020015481525050602081019050808060058560c052602060c020015481525050602081019050808060068560c052602060c020015481525050602081019050808060078560c052602060c020015481525050602081019050808060088560c052602060c0200154815250506101209050905060c05260c051610140f35b63c822adda811415613be35760043560148110613bce57600080fd5b600e60c052602060c020015460005260206000f35b633403c2fc811415613bfb57600f5460005260206000f35b63ecf70858811415613c135760105460005260206000f35b63cea55f57811415613c2b5760115460005260206000f35b63fc7b9c18811415613c435760125460005260206000f35b63c3535b52811415613c5b5760135460005260206000f35b633629c8de811415613c735760145460005260206000f35b6344b81396811415613c8b5760155460005260206000f35b632140254d811415613ca35760165460005260206000f35b639ec5a894811415613cbb5760175460005260206000f35b63a6f7f5d6811415613cd35760185460005260206000f35b6387788782811415613ceb5760195460005260206000f35b637ecebe00811415613d215760043560a01c15613d0757600080fd5b601a60043560e05260c052604060c0205460005260206000f35b633644e515811415613d3957601b5460005260206000f35b505b60006000fd5b6101a05261014052610160526101805260006004610220527fa9059cbb000000000000000000000000000000000000000000000000000000006102405261022060048060208461028001018260208501600060045af15050805182019150506101605160208261028001015260208101905061018051602082610280010152602081019050806102805261028090508051602001806103208284600060045af1613dea57600080fd5b505060206103e0610320516103406000610140515af1613e0957600080fd5b60203d80821115613e1a5780613e1c565b815b905090506103c0526103c08051602001806101c08284600060045af1613e4157600080fd5b505060006101c0511115613ed5576101c0806020015160008251806020901315613e6a57600080fd5b8091901215613e7857600080fd5b806020036101000a820490509050905015151515613ed5576308c379a0610220526020610240526010610260527f5472616e73666572206661696c656421000000000000000000000000000000006102805261026050606461023cfd5b6101a051565b6101c0526101405261016052610180526101a05260006004610240527f23b872dd00000000000000000000000000000000000000000000000000000000610260526102406004806020846102a001018260208501600060045af1505080518201915050610160516020826102a0010152602081019050610180516020826102a00101526020810190506101a0516020826102a0010152602081019050806102a0526102a090508051602001806103608284600060045af1613f9b57600080fd5b50506020610440610360516103806000610140515af1613fba57600080fd5b60203d80821115613fcb5780613fcd565b815b90509050610420526104208051602001806101e08284600060045af1613ff257600080fd5b505060006101e0511115614086576101e080602001516000825180602090131561401b57600080fd5b809190121561402957600080fd5b806020036101000a820490509050905015151515614086576308c379a0610240526020610260526010610280527f5472616e73666572206661696c656421000000000000000000000000000000006102a05261028050606461025cfd5b6101c051565b6101a052610140526101605261018052306101e05260006102005260006101c0526101c061012060006002818352015b610120516020026101e001516101605114156140db57600183526140eb565b81516001018083528114156140bc575b5050506101c051156140fc57600080fd5b60046101405160e05260c052604060c0208054610180518082101561412057600080fd5b8082039050905081555060046101605160e05260c052604060c02080546101805181818301101561415057600080fd5b80820190509050815550610180516101c05261016051610140517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60206101c0a36101a051565b6101405260206101e060246370a0823161016052306101805261017c6007545afa6141c157600080fd5b601f3d116141ce57600080fd5b6000506101e0516012548181830110156141e757600080fd5b8082019050905060005260005161014051565b61018052610140526101605260006101a0526006546101c05260006101c05111156142d8576003546101e0526101e05161016051808202821582848304141761424257600080fd5b809050905090506101c051808202821582848304141761426157600080fd5b809050905090506101405161016051610180516101a0516101c0516101e05160065801614197565b610200526101e0526101c0526101a0526101805261016052610140526102005180806142b457600080fd5b8204905090506101e05180806142c957600080fd5b8204905090506101a0526142e1565b610160516101a0525b6101c0516101a0518181830110156142f857600080fd5b8082019050905060065560046101405160e05260c052604060c02080546101a05181818301101561432857600080fd5b808201905090508155506101a0516101e0526101405160007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60206101e0a36101a05160005260005161018051565b61016052610140526006541515614398576101405160005260005161016051565b42601354808210156143a957600080fd5b8082039050905060165480820282158284830414176143c757600080fd5b80905090509050610180526101405161016051610180516101a05160065801614197565b6101c0526101a0526101805261016052610140526101c0516101a0526003546101c052670de0b6b3a76400006101805110156144bb576101a080516015546101c05161018051808202821582848304141761444557600080fd5b80905090509050601554808202821582848304141761446357600080fd5b80905090509050670de0b6b3a7640000808204905090506101c051808061448957600080fd5b8204905090508082101561449c57600080fd5b80820390509050808210156144b057600080fd5b808203905090508152505b6101c0516101405180820282158284830414176144d757600080fd5b809050905090506101a05180820282158284830414176144f657600080fd5b80905090509050600654808061450b57600080fd5b8204905090506101c051808061452057600080fd5b82049050905060005260005161016051565b61016052610140526000610140516101605160065801614197565b610180526101605261014052610180511115614612576003546101a0526101a05161014051808202821582848304141761458657600080fd5b8090509050905060065480820282158284830414176145a457600080fd5b809050905090506101405161016051610180516101a05160065801614197565b6101c0526101a0526101805261016052610140526101c05180806145e757600080fd5b8204905090506101a05180806145fc57600080fd5b8204905090506000526000516101605156614620565b600060005260005161016051565b005b6101805261014052610160526006600d6101405160e05260c052604060c02060c052602060c02001546101a052610160516101a051101561466257600080fd5b6008600d6101405160e05260c052604060c02060c052602060c0200180546101605181818301101561469357600080fd5b808201905090508155506101a05161016051808210156146b257600080fd5b808203905090506006600d6101405160e05260c052604060c02060c052602060c02001556012805461016051808210156146eb57600080fd5b808203905090508155506002600d6101405160e05260c052604060c02060c052602060c02001546101c0526003546101e0526101e05161016051808202821582848304141761473957600080fd5b80905090509050612710808202821582848304141761475757600080fd5b809050905090506101405161016051610180516101a0516101c0516101e0516102005160065801614197565b61022052610200526101e0526101c0526101a0526101805261016052610140526102205180806147b257600080fd5b8204905090506101e05180806147c757600080fd5b8204905090506101c051808211156147df57806147e1565b815b90509050610200526002600d6101405160e05260c052604060c02060c052602060c020018054610200518082101561481857600080fd5b8082039050905081555060118054610200518082101561483757600080fd5b8082039050905081555061018051565b6101405260006101605261018060006014818352015b610180516014811061486e57600080fd5b600e60c052602060c02001546101a0526101a05115156148ad576101608051600181818301101561489e57600080fd5b80820190509050815250614912565b6000610160511115614912576101a0516101805161016051808210156148d257600080fd5b80820390509050601481106148e657600080fd5b600e60c052602060c02001556000610180516014811061490557600080fd5b600e60c052602060c02001555b815160010180835281141561485d575b505061014051565b6101605261014052601180546002600d6101405160e05260c052604060c02060c052602060c02001548082101561496057600080fd5b8082039050905081555060006002600d6101405160e05260c052604060c02060c052602060c0200155610140517f4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a60006000a261016051565b610160526101405260035461018052610180516002600d6101405160e05260c052604060c02060c052602060c020015480820282158284830414176149fd57600080fd5b809050905090506101405161016051610180516101a05160065801614197565b6101c0526101a0526101805261016052610140526101c0518082028215828483041417614a4957600080fd5b8090509050905061271080820490509050610180518080614a6957600080fd5b8204905090506101a0526006600d6101405160e05260c052604060c02060c052602060c02001546101c052600f5415614ab0576101c0516000526000516101605156614af9565b6101a0516101c051111515614ad15760006000526000516101605156614af9565b6101c0516101a05180821015614ae657600080fd5b8082039050905060005260005161016051565b005b6101605261014052600f5415614b1957600060005260005161016051565b600354610180526101405161016051610180516101a05160065801614197565b6101c0526101a0526101805261016052610140526101c0516101a052610180516011548082028215828483041417614b7057600080fd5b809050905090506101a0518082028215828483041417614b8f57600080fd5b8090509050905061271080820490509050610180518080614baf57600080fd5b8204905090506101c0526012546101e052610180516002600d6101405160e05260c052604060c02060c052602060c02001548082028215828483041417614bf557600080fd5b809050905090506101a0518082028215828483041417614c1457600080fd5b8090509050905061271080820490509050610180518080614c3457600080fd5b820490509050610200526006600d6101405160e05260c052604060c02060c052602060c0200154610220526003600d6101405160e05260c052604060c02060c052602060c0200154610240526004600d6101405160e05260c052604060c02060c052602060c0200154610260526102205161020051111515614cb7576001614cc2565b6101e0516101c05111155b15614cd557600060005260005161016051565b610200516102205180821015614cea57600080fd5b8082039050905061028052610280516101c0516101e05180821015614d0e57600080fd5b8082039050905080821115614d235780614d25565b815b905090506102805261028051602061032060246370a082316102a052306102c0526102bc6007545afa614d5757600080fd5b601f3d11614d6457600080fd5b6000506103205180821115614d795780614d7b565b815b905090506102805261024051610280511015614da35760006000526000516101605156614dcb565b610280516102605180821115614db95780614dbb565b815b9050905060005260005161016051565b005b61016052610140526005600d6101405160e05260c052604060c02060c052602060c020015461018052426101805180821015614e0857600080fd5b808203905090506101a052610180516001600d6101405160e05260c052604060c02060c052602060c020015480821015614e4157600080fd5b808203905090506101c05260006101a0511115614ea75760006101c0511115614e9f5760206102c060046322f3e2d46102605261027c610140515afa614e8657600080fd5b601f3d11614e9357600080fd5b6000506102c051614ea2565b60005b614eaa565b60005b15614f4b576003546102e0526102e0516007600d6101405160e05260c052604060c02060c052602060c02001548082028215828483041417614eeb57600080fd5b809050905090506101a0518082028215828483041417614f0a57600080fd5b809050905090506101c0518080614f2057600080fd5b8204905090506102e0518080614f3557600080fd5b8204905090506000526000516101605156614f59565b600060005260005161016051565b005b6101805261014052610160526003546101a0526101a0516006600d6101405160e05260c052604060c02060c052602060c020015460206102406004638e6350e26101e0526101fc610140515afa614fb157600080fd5b601f3d11614fbe57600080fd5b6000506102405180821015614fd257600080fd5b80820390509050426005600d6101405160e05260c052604060c02060c052602060c02001548082101561500457600080fd5b80820390509050808202821582848304141761501f57600080fd5b80905090509050601854808202821582848304141761503d57600080fd5b80905090509050808202821582848304141761505857600080fd5b80905090509050612710808204905090506301e18558808204905090506101a051808061508457600080fd5b8204905090506101c0526040366101e0376000610160511115615181576101a0516101605180820282158284830414176150bd57600080fd5b80905090509050600d6101405160e05260c052604060c02060c052602060c0205480820282158284830414176150f257600080fd5b80905090509050612710808204905090506101a051808061511257600080fd5b8204905090506101e0526101a05161016051808202821582848304141761513857600080fd5b80905090509050601954808202821582848304141761515657600080fd5b80905090509050612710808204905090506101a051808061517657600080fd5b820490509050610200525b610200516101e05181818301101561519857600080fd5b808201905090506101c0518181830110156151b257600080fd5b80820190509050610220526101605161022051111561520c5761016051610220526101605161020051808210156151e857600080fd5b808203905090506101e0518082101561520057600080fd5b808203905090506101c0525b600061022051111561541b576101405161016051610180516101a0516101c0516101e051610200516102205161024051306102605261022051610280526102805161026051600658016141fa565b6102e0526102405261022052610200526101e0526101c0526101a0526101805261016052610140526102e0516102405260006101e0511115615384576101a0516101e05180820282158284830414176152b257600080fd5b809050905090506102405180820282158284830414176152d157600080fd5b809050905090506102205180806152e757600080fd5b8204905090506101a05180806152fc57600080fd5b820490509050610260526101405161016051610180516101a0516101c0516101e051610200516102205161024051610260513061028052610140516102a052610260516102c0526102c0516102a051610280516006580161408c565b610260526102405261022052610200526101e0526101c0526101a0526101805261016052610140526000505b600060043060e05260c052604060c02054111561541b576101405161016051610180516101a0516101c0516101e05161020051610220516102405130610260526017546102805260043060e05260c052604060c020546102a0526102a05161028051610260516006580161408c565b6102405261022052610200526101e0526101c0526101a0526101805261016052610140526000505b6102205160005260005161018051565b61000461542f0361000460003961000461542f036000f3",
}

// YearnVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use YearnVaultMetaData.ABI instead.
var YearnVaultABI = YearnVaultMetaData.ABI

// YearnVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use YearnVaultMetaData.Bin instead.
var YearnVaultBin = YearnVaultMetaData.Bin

// DeployYearnVault deploys a new Ethereum contract, binding an instance of YearnVault to it.
func DeployYearnVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *YearnVault, error) {
	parsed, err := YearnVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(YearnVaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &YearnVault{YearnVaultCaller: YearnVaultCaller{contract: contract}, YearnVaultTransactor: YearnVaultTransactor{contract: contract}, YearnVaultFilterer: YearnVaultFilterer{contract: contract}}, nil
}

// YearnVault is an auto generated Go binding around an Ethereum contract.
type YearnVault struct {
	YearnVaultCaller     // Read-only binding to the contract
	YearnVaultTransactor // Write-only binding to the contract
	YearnVaultFilterer   // Log filterer for contract events
}

// YearnVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type YearnVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YearnVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YearnVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YearnVaultSession struct {
	Contract     *YearnVault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YearnVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YearnVaultCallerSession struct {
	Contract *YearnVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// YearnVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YearnVaultTransactorSession struct {
	Contract     *YearnVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// YearnVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type YearnVaultRaw struct {
	Contract *YearnVault // Generic contract binding to access the raw methods on
}

// YearnVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YearnVaultCallerRaw struct {
	Contract *YearnVaultCaller // Generic read-only contract binding to access the raw methods on
}

// YearnVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YearnVaultTransactorRaw struct {
	Contract *YearnVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearnVault creates a new instance of YearnVault, bound to a specific deployed contract.
func NewYearnVault(address common.Address, backend bind.ContractBackend) (*YearnVault, error) {
	contract, err := bindYearnVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YearnVault{YearnVaultCaller: YearnVaultCaller{contract: contract}, YearnVaultTransactor: YearnVaultTransactor{contract: contract}, YearnVaultFilterer: YearnVaultFilterer{contract: contract}}, nil
}

// NewYearnVaultCaller creates a new read-only instance of YearnVault, bound to a specific deployed contract.
func NewYearnVaultCaller(address common.Address, caller bind.ContractCaller) (*YearnVaultCaller, error) {
	contract, err := bindYearnVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YearnVaultCaller{contract: contract}, nil
}

// NewYearnVaultTransactor creates a new write-only instance of YearnVault, bound to a specific deployed contract.
func NewYearnVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*YearnVaultTransactor, error) {
	contract, err := bindYearnVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YearnVaultTransactor{contract: contract}, nil
}

// NewYearnVaultFilterer creates a new log filterer instance of YearnVault, bound to a specific deployed contract.
func NewYearnVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*YearnVaultFilterer, error) {
	contract, err := bindYearnVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YearnVaultFilterer{contract: contract}, nil
}

// bindYearnVault binds a generic wrapper to an already deployed contract.
func bindYearnVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YearnVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnVault *YearnVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnVault.Contract.YearnVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnVault *YearnVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnVault.Contract.YearnVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnVault *YearnVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnVault.Contract.YearnVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnVault *YearnVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnVault *YearnVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnVault *YearnVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnVault.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YearnVault *YearnVaultCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YearnVault *YearnVaultSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _YearnVault.Contract.DOMAINSEPARATOR(&_YearnVault.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YearnVault *YearnVaultCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _YearnVault.Contract.DOMAINSEPARATOR(&_YearnVault.CallOpts)
}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_YearnVault *YearnVaultCaller) Activation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "activation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_YearnVault *YearnVaultSession) Activation() (*big.Int, error) {
	return _YearnVault.Contract.Activation(&_YearnVault.CallOpts)
}

// Activation is a free data retrieval call binding the contract method 0x3629c8de.
//
// Solidity: function activation() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) Activation() (*big.Int, error) {
	return _YearnVault.Contract.Activation(&_YearnVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_YearnVault *YearnVaultCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_YearnVault *YearnVaultSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YearnVault.Contract.Allowance(&_YearnVault.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YearnVault.Contract.Allowance(&_YearnVault.CallOpts, arg0, arg1)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YearnVault *YearnVaultCaller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YearnVault *YearnVaultSession) ApiVersion() (string, error) {
	return _YearnVault.Contract.ApiVersion(&_YearnVault.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_YearnVault *YearnVaultCallerSession) ApiVersion() (string, error) {
	return _YearnVault.Contract.ApiVersion(&_YearnVault.CallOpts)
}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_YearnVault *YearnVaultCaller) AvailableDepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "availableDepositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_YearnVault *YearnVaultSession) AvailableDepositLimit() (*big.Int, error) {
	return _YearnVault.Contract.AvailableDepositLimit(&_YearnVault.CallOpts)
}

// AvailableDepositLimit is a free data retrieval call binding the contract method 0x153c27c4.
//
// Solidity: function availableDepositLimit() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) AvailableDepositLimit() (*big.Int, error) {
	return _YearnVault.Contract.AvailableDepositLimit(&_YearnVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_YearnVault *YearnVaultCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_YearnVault *YearnVaultSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _YearnVault.Contract.BalanceOf(&_YearnVault.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _YearnVault.Contract.BalanceOf(&_YearnVault.CallOpts, arg0)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_YearnVault *YearnVaultCaller) CreditAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "creditAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_YearnVault *YearnVaultSession) CreditAvailable() (*big.Int, error) {
	return _YearnVault.Contract.CreditAvailable(&_YearnVault.CallOpts)
}

// CreditAvailable is a free data retrieval call binding the contract method 0x112c1f9b.
//
// Solidity: function creditAvailable() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) CreditAvailable() (*big.Int, error) {
	return _YearnVault.Contract.CreditAvailable(&_YearnVault.CallOpts)
}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_YearnVault *YearnVaultCaller) CreditAvailable0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "creditAvailable0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_YearnVault *YearnVaultSession) CreditAvailable0(strategy common.Address) (*big.Int, error) {
	return _YearnVault.Contract.CreditAvailable0(&_YearnVault.CallOpts, strategy)
}

// CreditAvailable0 is a free data retrieval call binding the contract method 0xd7648013.
//
// Solidity: function creditAvailable(address strategy) view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) CreditAvailable0(strategy common.Address) (*big.Int, error) {
	return _YearnVault.Contract.CreditAvailable0(&_YearnVault.CallOpts, strategy)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_YearnVault *YearnVaultCaller) DebtOutstanding(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "debtOutstanding")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_YearnVault *YearnVaultSession) DebtOutstanding() (*big.Int, error) {
	return _YearnVault.Contract.DebtOutstanding(&_YearnVault.CallOpts)
}

// DebtOutstanding is a free data retrieval call binding the contract method 0xbf3759b5.
//
// Solidity: function debtOutstanding() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) DebtOutstanding() (*big.Int, error) {
	return _YearnVault.Contract.DebtOutstanding(&_YearnVault.CallOpts)
}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_YearnVault *YearnVaultCaller) DebtOutstanding0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "debtOutstanding0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_YearnVault *YearnVaultSession) DebtOutstanding0(strategy common.Address) (*big.Int, error) {
	return _YearnVault.Contract.DebtOutstanding0(&_YearnVault.CallOpts, strategy)
}

// DebtOutstanding0 is a free data retrieval call binding the contract method 0xbdcf36bb.
//
// Solidity: function debtOutstanding(address strategy) view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) DebtOutstanding0(strategy common.Address) (*big.Int, error) {
	return _YearnVault.Contract.DebtOutstanding0(&_YearnVault.CallOpts, strategy)
}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_YearnVault *YearnVaultCaller) DebtRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "debtRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_YearnVault *YearnVaultSession) DebtRatio() (*big.Int, error) {
	return _YearnVault.Contract.DebtRatio(&_YearnVault.CallOpts)
}

// DebtRatio is a free data retrieval call binding the contract method 0xcea55f57.
//
// Solidity: function debtRatio() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) DebtRatio() (*big.Int, error) {
	return _YearnVault.Contract.DebtRatio(&_YearnVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_YearnVault *YearnVaultCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_YearnVault *YearnVaultSession) Decimals() (*big.Int, error) {
	return _YearnVault.Contract.Decimals(&_YearnVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) Decimals() (*big.Int, error) {
	return _YearnVault.Contract.Decimals(&_YearnVault.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_YearnVault *YearnVaultCaller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "depositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_YearnVault *YearnVaultSession) DepositLimit() (*big.Int, error) {
	return _YearnVault.Contract.DepositLimit(&_YearnVault.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) DepositLimit() (*big.Int, error) {
	return _YearnVault.Contract.DepositLimit(&_YearnVault.CallOpts)
}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_YearnVault *YearnVaultCaller) EmergencyShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "emergencyShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_YearnVault *YearnVaultSession) EmergencyShutdown() (bool, error) {
	return _YearnVault.Contract.EmergencyShutdown(&_YearnVault.CallOpts)
}

// EmergencyShutdown is a free data retrieval call binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() view returns(bool)
func (_YearnVault *YearnVaultCallerSession) EmergencyShutdown() (bool, error) {
	return _YearnVault.Contract.EmergencyShutdown(&_YearnVault.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_YearnVault *YearnVaultCaller) ExpectedReturn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "expectedReturn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_YearnVault *YearnVaultSession) ExpectedReturn() (*big.Int, error) {
	return _YearnVault.Contract.ExpectedReturn(&_YearnVault.CallOpts)
}

// ExpectedReturn is a free data retrieval call binding the contract method 0xd3406abd.
//
// Solidity: function expectedReturn() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) ExpectedReturn() (*big.Int, error) {
	return _YearnVault.Contract.ExpectedReturn(&_YearnVault.CallOpts)
}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_YearnVault *YearnVaultCaller) ExpectedReturn0(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "expectedReturn0", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_YearnVault *YearnVaultSession) ExpectedReturn0(strategy common.Address) (*big.Int, error) {
	return _YearnVault.Contract.ExpectedReturn0(&_YearnVault.CallOpts, strategy)
}

// ExpectedReturn0 is a free data retrieval call binding the contract method 0x33586b67.
//
// Solidity: function expectedReturn(address strategy) view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) ExpectedReturn0(strategy common.Address) (*big.Int, error) {
	return _YearnVault.Contract.ExpectedReturn0(&_YearnVault.CallOpts, strategy)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YearnVault *YearnVaultCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YearnVault *YearnVaultSession) Governance() (common.Address, error) {
	return _YearnVault.Contract.Governance(&_YearnVault.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YearnVault *YearnVaultCallerSession) Governance() (common.Address, error) {
	return _YearnVault.Contract.Governance(&_YearnVault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_YearnVault *YearnVaultCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_YearnVault *YearnVaultSession) Guardian() (common.Address, error) {
	return _YearnVault.Contract.Guardian(&_YearnVault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_YearnVault *YearnVaultCallerSession) Guardian() (common.Address, error) {
	return _YearnVault.Contract.Guardian(&_YearnVault.CallOpts)
}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_YearnVault *YearnVaultCaller) GuestList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "guestList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_YearnVault *YearnVaultSession) GuestList() (common.Address, error) {
	return _YearnVault.Contract.GuestList(&_YearnVault.CallOpts)
}

// GuestList is a free data retrieval call binding the contract method 0x46d55875.
//
// Solidity: function guestList() view returns(address)
func (_YearnVault *YearnVaultCallerSession) GuestList() (common.Address, error) {
	return _YearnVault.Contract.GuestList(&_YearnVault.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_YearnVault *YearnVaultCaller) LastReport(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "lastReport")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_YearnVault *YearnVaultSession) LastReport() (*big.Int, error) {
	return _YearnVault.Contract.LastReport(&_YearnVault.CallOpts)
}

// LastReport is a free data retrieval call binding the contract method 0xc3535b52.
//
// Solidity: function lastReport() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) LastReport() (*big.Int, error) {
	return _YearnVault.Contract.LastReport(&_YearnVault.CallOpts)
}

// LockedProfit is a free data retrieval call binding the contract method 0x44b81396.
//
// Solidity: function lockedProfit() view returns(uint256)
func (_YearnVault *YearnVaultCaller) LockedProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "lockedProfit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedProfit is a free data retrieval call binding the contract method 0x44b81396.
//
// Solidity: function lockedProfit() view returns(uint256)
func (_YearnVault *YearnVaultSession) LockedProfit() (*big.Int, error) {
	return _YearnVault.Contract.LockedProfit(&_YearnVault.CallOpts)
}

// LockedProfit is a free data retrieval call binding the contract method 0x44b81396.
//
// Solidity: function lockedProfit() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) LockedProfit() (*big.Int, error) {
	return _YearnVault.Contract.LockedProfit(&_YearnVault.CallOpts)
}

// LockedProfitDegration is a free data retrieval call binding the contract method 0x2140254d.
//
// Solidity: function lockedProfitDegration() view returns(uint256)
func (_YearnVault *YearnVaultCaller) LockedProfitDegration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "lockedProfitDegration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedProfitDegration is a free data retrieval call binding the contract method 0x2140254d.
//
// Solidity: function lockedProfitDegration() view returns(uint256)
func (_YearnVault *YearnVaultSession) LockedProfitDegration() (*big.Int, error) {
	return _YearnVault.Contract.LockedProfitDegration(&_YearnVault.CallOpts)
}

// LockedProfitDegration is a free data retrieval call binding the contract method 0x2140254d.
//
// Solidity: function lockedProfitDegration() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) LockedProfitDegration() (*big.Int, error) {
	return _YearnVault.Contract.LockedProfitDegration(&_YearnVault.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YearnVault *YearnVaultCaller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YearnVault *YearnVaultSession) Management() (common.Address, error) {
	return _YearnVault.Contract.Management(&_YearnVault.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_YearnVault *YearnVaultCallerSession) Management() (common.Address, error) {
	return _YearnVault.Contract.Management(&_YearnVault.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_YearnVault *YearnVaultCaller) ManagementFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "managementFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_YearnVault *YearnVaultSession) ManagementFee() (*big.Int, error) {
	return _YearnVault.Contract.ManagementFee(&_YearnVault.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) ManagementFee() (*big.Int, error) {
	return _YearnVault.Contract.ManagementFee(&_YearnVault.CallOpts)
}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_YearnVault *YearnVaultCaller) MaxAvailableShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "maxAvailableShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_YearnVault *YearnVaultSession) MaxAvailableShares() (*big.Int, error) {
	return _YearnVault.Contract.MaxAvailableShares(&_YearnVault.CallOpts)
}

// MaxAvailableShares is a free data retrieval call binding the contract method 0x75de2902.
//
// Solidity: function maxAvailableShares() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) MaxAvailableShares() (*big.Int, error) {
	return _YearnVault.Contract.MaxAvailableShares(&_YearnVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnVault *YearnVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnVault *YearnVaultSession) Name() (string, error) {
	return _YearnVault.Contract.Name(&_YearnVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnVault *YearnVaultCallerSession) Name() (string, error) {
	return _YearnVault.Contract.Name(&_YearnVault.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_YearnVault *YearnVaultCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_YearnVault *YearnVaultSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _YearnVault.Contract.Nonces(&_YearnVault.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _YearnVault.Contract.Nonces(&_YearnVault.CallOpts, arg0)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_YearnVault *YearnVaultCaller) PerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "performanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_YearnVault *YearnVaultSession) PerformanceFee() (*big.Int, error) {
	return _YearnVault.Contract.PerformanceFee(&_YearnVault.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) PerformanceFee() (*big.Int, error) {
	return _YearnVault.Contract.PerformanceFee(&_YearnVault.CallOpts)
}

// PrecisionFactor is a free data retrieval call binding the contract method 0x9d902fc0.
//
// Solidity: function precisionFactor() view returns(uint256)
func (_YearnVault *YearnVaultCaller) PrecisionFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "precisionFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrecisionFactor is a free data retrieval call binding the contract method 0x9d902fc0.
//
// Solidity: function precisionFactor() view returns(uint256)
func (_YearnVault *YearnVaultSession) PrecisionFactor() (*big.Int, error) {
	return _YearnVault.Contract.PrecisionFactor(&_YearnVault.CallOpts)
}

// PrecisionFactor is a free data retrieval call binding the contract method 0x9d902fc0.
//
// Solidity: function precisionFactor() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) PrecisionFactor() (*big.Int, error) {
	return _YearnVault.Contract.PrecisionFactor(&_YearnVault.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnVault *YearnVaultCaller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnVault *YearnVaultSession) PricePerShare() (*big.Int, error) {
	return _YearnVault.Contract.PricePerShare(&_YearnVault.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) PricePerShare() (*big.Int, error) {
	return _YearnVault.Contract.PricePerShare(&_YearnVault.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_YearnVault *YearnVaultCaller) Rewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_YearnVault *YearnVaultSession) Rewards() (common.Address, error) {
	return _YearnVault.Contract.Rewards(&_YearnVault.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(address)
func (_YearnVault *YearnVaultCallerSession) Rewards() (common.Address, error) {
	return _YearnVault.Contract.Rewards(&_YearnVault.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_YearnVault *YearnVaultCaller) Strategies(opts *bind.CallOpts, arg0 common.Address) (struct {
	PerformanceFee    *big.Int
	Activation        *big.Int
	DebtRatio         *big.Int
	MinDebtPerHarvest *big.Int
	MaxDebtPerHarvest *big.Int
	LastReport        *big.Int
	TotalDebt         *big.Int
	TotalGain         *big.Int
	TotalLoss         *big.Int
}, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "strategies", arg0)

	outstruct := new(struct {
		PerformanceFee    *big.Int
		Activation        *big.Int
		DebtRatio         *big.Int
		MinDebtPerHarvest *big.Int
		MaxDebtPerHarvest *big.Int
		LastReport        *big.Int
		TotalDebt         *big.Int
		TotalGain         *big.Int
		TotalLoss         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PerformanceFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Activation = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DebtRatio = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MinDebtPerHarvest = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxDebtPerHarvest = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LastReport = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalDebt = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TotalGain = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TotalLoss = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_YearnVault *YearnVaultSession) Strategies(arg0 common.Address) (struct {
	PerformanceFee    *big.Int
	Activation        *big.Int
	DebtRatio         *big.Int
	MinDebtPerHarvest *big.Int
	MaxDebtPerHarvest *big.Int
	LastReport        *big.Int
	TotalDebt         *big.Int
	TotalGain         *big.Int
	TotalLoss         *big.Int
}, error) {
	return _YearnVault.Contract.Strategies(&_YearnVault.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns(uint256 performanceFee, uint256 activation, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 lastReport, uint256 totalDebt, uint256 totalGain, uint256 totalLoss)
func (_YearnVault *YearnVaultCallerSession) Strategies(arg0 common.Address) (struct {
	PerformanceFee    *big.Int
	Activation        *big.Int
	DebtRatio         *big.Int
	MinDebtPerHarvest *big.Int
	MaxDebtPerHarvest *big.Int
	LastReport        *big.Int
	TotalDebt         *big.Int
	TotalGain         *big.Int
	TotalLoss         *big.Int
}, error) {
	return _YearnVault.Contract.Strategies(&_YearnVault.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnVault *YearnVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnVault *YearnVaultSession) Symbol() (string, error) {
	return _YearnVault.Contract.Symbol(&_YearnVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnVault *YearnVaultCallerSession) Symbol() (string, error) {
	return _YearnVault.Contract.Symbol(&_YearnVault.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_YearnVault *YearnVaultCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_YearnVault *YearnVaultSession) Token() (common.Address, error) {
	return _YearnVault.Contract.Token(&_YearnVault.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_YearnVault *YearnVaultCallerSession) Token() (common.Address, error) {
	return _YearnVault.Contract.Token(&_YearnVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnVault *YearnVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnVault *YearnVaultSession) TotalAssets() (*big.Int, error) {
	return _YearnVault.Contract.TotalAssets(&_YearnVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _YearnVault.Contract.TotalAssets(&_YearnVault.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YearnVault *YearnVaultCaller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YearnVault *YearnVaultSession) TotalDebt() (*big.Int, error) {
	return _YearnVault.Contract.TotalDebt(&_YearnVault.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) TotalDebt() (*big.Int, error) {
	return _YearnVault.Contract.TotalDebt(&_YearnVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnVault *YearnVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnVault *YearnVaultSession) TotalSupply() (*big.Int, error) {
	return _YearnVault.Contract.TotalSupply(&_YearnVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _YearnVault.Contract.TotalSupply(&_YearnVault.CallOpts)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_YearnVault *YearnVaultCaller) WithdrawalQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "withdrawalQueue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_YearnVault *YearnVaultSession) WithdrawalQueue(arg0 *big.Int) (common.Address, error) {
	return _YearnVault.Contract.WithdrawalQueue(&_YearnVault.CallOpts, arg0)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0xc822adda.
//
// Solidity: function withdrawalQueue(uint256 arg0) view returns(address)
func (_YearnVault *YearnVaultCallerSession) WithdrawalQueue(arg0 *big.Int) (common.Address, error) {
	return _YearnVault.Contract.WithdrawalQueue(&_YearnVault.CallOpts, arg0)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_YearnVault *YearnVaultTransactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_YearnVault *YearnVaultSession) AcceptGovernance() (*types.Transaction, error) {
	return _YearnVault.Contract.AcceptGovernance(&_YearnVault.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_YearnVault *YearnVaultTransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _YearnVault.Contract.AcceptGovernance(&_YearnVault.TransactOpts)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x14b4e26e.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee) returns()
func (_YearnVault *YearnVaultTransactor) AddStrategy(opts *bind.TransactOpts, strategy common.Address, debtRatio *big.Int, minDebtPerHarvest *big.Int, maxDebtPerHarvest *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "addStrategy", strategy, debtRatio, minDebtPerHarvest, maxDebtPerHarvest, performanceFee)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x14b4e26e.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee) returns()
func (_YearnVault *YearnVaultSession) AddStrategy(strategy common.Address, debtRatio *big.Int, minDebtPerHarvest *big.Int, maxDebtPerHarvest *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.AddStrategy(&_YearnVault.TransactOpts, strategy, debtRatio, minDebtPerHarvest, maxDebtPerHarvest, performanceFee)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x14b4e26e.
//
// Solidity: function addStrategy(address strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee) returns()
func (_YearnVault *YearnVaultTransactorSession) AddStrategy(strategy common.Address, debtRatio *big.Int, minDebtPerHarvest *big.Int, maxDebtPerHarvest *big.Int, performanceFee *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.AddStrategy(&_YearnVault.TransactOpts, strategy, debtRatio, minDebtPerHarvest, maxDebtPerHarvest, performanceFee)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_YearnVault *YearnVaultTransactor) AddStrategyToQueue(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "addStrategyToQueue", strategy)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_YearnVault *YearnVaultSession) AddStrategyToQueue(strategy common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.AddStrategyToQueue(&_YearnVault.TransactOpts, strategy)
}

// AddStrategyToQueue is a paid mutator transaction binding the contract method 0xf76e4caa.
//
// Solidity: function addStrategyToQueue(address strategy) returns()
func (_YearnVault *YearnVaultTransactorSession) AddStrategyToQueue(strategy common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.AddStrategyToQueue(&_YearnVault.TransactOpts, strategy)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Approve(&_YearnVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Approve(&_YearnVault.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.DecreaseAllowance(&_YearnVault.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultTransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.DecreaseAllowance(&_YearnVault.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YearnVault *YearnVaultTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YearnVault *YearnVaultSession) Deposit() (*types.Transaction, error) {
	return _YearnVault.Contract.Deposit(&_YearnVault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YearnVault *YearnVaultTransactorSession) Deposit() (*types.Transaction, error) {
	return _YearnVault.Contract.Deposit(&_YearnVault.TransactOpts)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_YearnVault *YearnVaultTransactor) Deposit0(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "deposit0", _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_YearnVault *YearnVaultSession) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Deposit0(&_YearnVault.TransactOpts, _amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(uint256)
func (_YearnVault *YearnVaultTransactorSession) Deposit0(_amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Deposit0(&_YearnVault.TransactOpts, _amount)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_YearnVault *YearnVaultTransactor) Deposit1(opts *bind.TransactOpts, _amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "deposit1", _amount, recipient)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_YearnVault *YearnVaultSession) Deposit1(_amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.Deposit1(&_YearnVault.TransactOpts, _amount, recipient)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address recipient) returns(uint256)
func (_YearnVault *YearnVaultTransactorSession) Deposit1(_amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.Deposit1(&_YearnVault.TransactOpts, _amount, recipient)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "increaseAllowance", spender, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultSession) IncreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.IncreaseAllowance(&_YearnVault.TransactOpts, spender, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultTransactorSession) IncreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.IncreaseAllowance(&_YearnVault.TransactOpts, spender, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_YearnVault *YearnVaultTransactor) Initialize(opts *bind.TransactOpts, token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "initialize", token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_YearnVault *YearnVaultSession) Initialize(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _YearnVault.Contract.Initialize(&_YearnVault.TransactOpts, token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride) returns()
func (_YearnVault *YearnVaultTransactorSession) Initialize(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string) (*types.Transaction, error) {
	return _YearnVault.Contract.Initialize(&_YearnVault.TransactOpts, token, governance, rewards, nameOverride, symbolOverride)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_YearnVault *YearnVaultTransactor) Initialize0(opts *bind.TransactOpts, token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "initialize0", token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_YearnVault *YearnVaultSession) Initialize0(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.Initialize0(&_YearnVault.TransactOpts, token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa5b81fdf.
//
// Solidity: function initialize(address token, address governance, address rewards, string nameOverride, string symbolOverride, address guardian) returns()
func (_YearnVault *YearnVaultTransactorSession) Initialize0(token common.Address, governance common.Address, rewards common.Address, nameOverride string, symbolOverride string, guardian common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.Initialize0(&_YearnVault.TransactOpts, token, governance, rewards, nameOverride, symbolOverride, guardian)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_YearnVault *YearnVaultTransactor) MigrateStrategy(opts *bind.TransactOpts, oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "migrateStrategy", oldVersion, newVersion)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_YearnVault *YearnVaultSession) MigrateStrategy(oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.MigrateStrategy(&_YearnVault.TransactOpts, oldVersion, newVersion)
}

// MigrateStrategy is a paid mutator transaction binding the contract method 0x6cb56d19.
//
// Solidity: function migrateStrategy(address oldVersion, address newVersion) returns()
func (_YearnVault *YearnVaultTransactorSession) MigrateStrategy(oldVersion common.Address, newVersion common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.MigrateStrategy(&_YearnVault.TransactOpts, oldVersion, newVersion)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_YearnVault *YearnVaultTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "permit", owner, spender, amount, expiry, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_YearnVault *YearnVaultSession) Permit(owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _YearnVault.Contract.Permit(&_YearnVault.TransactOpts, owner, spender, amount, expiry, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 expiry, bytes signature) returns(bool)
func (_YearnVault *YearnVaultTransactorSession) Permit(owner common.Address, spender common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _YearnVault.Contract.Permit(&_YearnVault.TransactOpts, owner, spender, amount, expiry, signature)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_YearnVault *YearnVaultTransactor) RemoveStrategyFromQueue(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "removeStrategyFromQueue", strategy)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_YearnVault *YearnVaultSession) RemoveStrategyFromQueue(strategy common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.RemoveStrategyFromQueue(&_YearnVault.TransactOpts, strategy)
}

// RemoveStrategyFromQueue is a paid mutator transaction binding the contract method 0xb22439f5.
//
// Solidity: function removeStrategyFromQueue(address strategy) returns()
func (_YearnVault *YearnVaultTransactorSession) RemoveStrategyFromQueue(strategy common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.RemoveStrategyFromQueue(&_YearnVault.TransactOpts, strategy)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_YearnVault *YearnVaultTransactor) Report(opts *bind.TransactOpts, gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "report", gain, loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_YearnVault *YearnVaultSession) Report(gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Report(&_YearnVault.TransactOpts, gain, loss, _debtPayment)
}

// Report is a paid mutator transaction binding the contract method 0xa1d9bafc.
//
// Solidity: function report(uint256 gain, uint256 loss, uint256 _debtPayment) returns(uint256)
func (_YearnVault *YearnVaultTransactorSession) Report(gain *big.Int, loss *big.Int, _debtPayment *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Report(&_YearnVault.TransactOpts, gain, loss, _debtPayment)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_YearnVault *YearnVaultTransactor) RevokeStrategy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "revokeStrategy")
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_YearnVault *YearnVaultSession) RevokeStrategy() (*types.Transaction, error) {
	return _YearnVault.Contract.RevokeStrategy(&_YearnVault.TransactOpts)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xa0e4af9a.
//
// Solidity: function revokeStrategy() returns()
func (_YearnVault *YearnVaultTransactorSession) RevokeStrategy() (*types.Transaction, error) {
	return _YearnVault.Contract.RevokeStrategy(&_YearnVault.TransactOpts)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_YearnVault *YearnVaultTransactor) RevokeStrategy0(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "revokeStrategy0", strategy)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_YearnVault *YearnVaultSession) RevokeStrategy0(strategy common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.RevokeStrategy0(&_YearnVault.TransactOpts, strategy)
}

// RevokeStrategy0 is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_YearnVault *YearnVaultTransactorSession) RevokeStrategy0(strategy common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.RevokeStrategy0(&_YearnVault.TransactOpts, strategy)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_YearnVault *YearnVaultTransactor) SetDepositLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setDepositLimit", limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_YearnVault *YearnVaultSession) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.SetDepositLimit(&_YearnVault.TransactOpts, limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_YearnVault *YearnVaultTransactorSession) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.SetDepositLimit(&_YearnVault.TransactOpts, limit)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_YearnVault *YearnVaultTransactor) SetEmergencyShutdown(opts *bind.TransactOpts, active bool) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setEmergencyShutdown", active)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_YearnVault *YearnVaultSession) SetEmergencyShutdown(active bool) (*types.Transaction, error) {
	return _YearnVault.Contract.SetEmergencyShutdown(&_YearnVault.TransactOpts, active)
}

// SetEmergencyShutdown is a paid mutator transaction binding the contract method 0x14c64402.
//
// Solidity: function setEmergencyShutdown(bool active) returns()
func (_YearnVault *YearnVaultTransactorSession) SetEmergencyShutdown(active bool) (*types.Transaction, error) {
	return _YearnVault.Contract.SetEmergencyShutdown(&_YearnVault.TransactOpts, active)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_YearnVault *YearnVaultTransactor) SetGovernance(opts *bind.TransactOpts, governance common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setGovernance", governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_YearnVault *YearnVaultSession) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetGovernance(&_YearnVault.TransactOpts, governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_YearnVault *YearnVaultTransactorSession) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetGovernance(&_YearnVault.TransactOpts, governance)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_YearnVault *YearnVaultTransactor) SetGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setGuardian", guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_YearnVault *YearnVaultSession) SetGuardian(guardian common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetGuardian(&_YearnVault.TransactOpts, guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address guardian) returns()
func (_YearnVault *YearnVaultTransactorSession) SetGuardian(guardian common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetGuardian(&_YearnVault.TransactOpts, guardian)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_YearnVault *YearnVaultTransactor) SetGuestList(opts *bind.TransactOpts, guestList common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setGuestList", guestList)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_YearnVault *YearnVaultSession) SetGuestList(guestList common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetGuestList(&_YearnVault.TransactOpts, guestList)
}

// SetGuestList is a paid mutator transaction binding the contract method 0x0b5b78eb.
//
// Solidity: function setGuestList(address guestList) returns()
func (_YearnVault *YearnVaultTransactorSession) SetGuestList(guestList common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetGuestList(&_YearnVault.TransactOpts, guestList)
}

// SetLockedProfitDegration is a paid mutator transaction binding the contract method 0x8402a84f.
//
// Solidity: function setLockedProfitDegration(uint256 degration) returns()
func (_YearnVault *YearnVaultTransactor) SetLockedProfitDegration(opts *bind.TransactOpts, degration *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setLockedProfitDegration", degration)
}

// SetLockedProfitDegration is a paid mutator transaction binding the contract method 0x8402a84f.
//
// Solidity: function setLockedProfitDegration(uint256 degration) returns()
func (_YearnVault *YearnVaultSession) SetLockedProfitDegration(degration *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.SetLockedProfitDegration(&_YearnVault.TransactOpts, degration)
}

// SetLockedProfitDegration is a paid mutator transaction binding the contract method 0x8402a84f.
//
// Solidity: function setLockedProfitDegration(uint256 degration) returns()
func (_YearnVault *YearnVaultTransactorSession) SetLockedProfitDegration(degration *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.SetLockedProfitDegration(&_YearnVault.TransactOpts, degration)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_YearnVault *YearnVaultTransactor) SetManagement(opts *bind.TransactOpts, management common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setManagement", management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_YearnVault *YearnVaultSession) SetManagement(management common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetManagement(&_YearnVault.TransactOpts, management)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address management) returns()
func (_YearnVault *YearnVaultTransactorSession) SetManagement(management common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetManagement(&_YearnVault.TransactOpts, management)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_YearnVault *YearnVaultTransactor) SetManagementFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setManagementFee", fee)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_YearnVault *YearnVaultSession) SetManagementFee(fee *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.SetManagementFee(&_YearnVault.TransactOpts, fee)
}

// SetManagementFee is a paid mutator transaction binding the contract method 0xfe56e232.
//
// Solidity: function setManagementFee(uint256 fee) returns()
func (_YearnVault *YearnVaultTransactorSession) SetManagementFee(fee *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.SetManagementFee(&_YearnVault.TransactOpts, fee)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_YearnVault *YearnVaultTransactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_YearnVault *YearnVaultSession) SetName(name string) (*types.Transaction, error) {
	return _YearnVault.Contract.SetName(&_YearnVault.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_YearnVault *YearnVaultTransactorSession) SetName(name string) (*types.Transaction, error) {
	return _YearnVault.Contract.SetName(&_YearnVault.TransactOpts, name)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_YearnVault *YearnVaultTransactor) SetPerformanceFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setPerformanceFee", fee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_YearnVault *YearnVaultSession) SetPerformanceFee(fee *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.SetPerformanceFee(&_YearnVault.TransactOpts, fee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 fee) returns()
func (_YearnVault *YearnVaultTransactorSession) SetPerformanceFee(fee *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.SetPerformanceFee(&_YearnVault.TransactOpts, fee)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_YearnVault *YearnVaultTransactor) SetRewards(opts *bind.TransactOpts, rewards common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setRewards", rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_YearnVault *YearnVaultSession) SetRewards(rewards common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetRewards(&_YearnVault.TransactOpts, rewards)
}

// SetRewards is a paid mutator transaction binding the contract method 0xec38a862.
//
// Solidity: function setRewards(address rewards) returns()
func (_YearnVault *YearnVaultTransactorSession) SetRewards(rewards common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetRewards(&_YearnVault.TransactOpts, rewards)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_YearnVault *YearnVaultTransactor) SetSymbol(opts *bind.TransactOpts, symbol string) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setSymbol", symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_YearnVault *YearnVaultSession) SetSymbol(symbol string) (*types.Transaction, error) {
	return _YearnVault.Contract.SetSymbol(&_YearnVault.TransactOpts, symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_YearnVault *YearnVaultTransactorSession) SetSymbol(symbol string) (*types.Transaction, error) {
	return _YearnVault.Contract.SetSymbol(&_YearnVault.TransactOpts, symbol)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_YearnVault *YearnVaultTransactor) SetWithdrawalQueue(opts *bind.TransactOpts, queue [20]common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "setWithdrawalQueue", queue)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_YearnVault *YearnVaultSession) SetWithdrawalQueue(queue [20]common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetWithdrawalQueue(&_YearnVault.TransactOpts, queue)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x94148415.
//
// Solidity: function setWithdrawalQueue(address[20] queue) returns()
func (_YearnVault *YearnVaultTransactorSession) SetWithdrawalQueue(queue [20]common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.SetWithdrawalQueue(&_YearnVault.TransactOpts, queue)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_YearnVault *YearnVaultTransactor) Sweep(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "sweep", token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_YearnVault *YearnVaultSession) Sweep(token common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.Sweep(&_YearnVault.TransactOpts, token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_YearnVault *YearnVaultTransactorSession) Sweep(token common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.Sweep(&_YearnVault.TransactOpts, token)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_YearnVault *YearnVaultTransactor) Sweep0(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "sweep0", token, amount)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_YearnVault *YearnVaultSession) Sweep0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Sweep0(&_YearnVault.TransactOpts, token, amount)
}

// Sweep0 is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address token, uint256 amount) returns()
func (_YearnVault *YearnVaultTransactorSession) Sweep0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Sweep0(&_YearnVault.TransactOpts, token, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "transfer", receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Transfer(&_YearnVault.TransactOpts, receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultTransactorSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Transfer(&_YearnVault.TransactOpts, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "transferFrom", sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.TransferFrom(&_YearnVault.TransactOpts, sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_YearnVault *YearnVaultTransactorSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.TransferFrom(&_YearnVault.TransactOpts, sender, receiver, amount)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_YearnVault *YearnVaultTransactor) UpdateStrategyDebtRatio(opts *bind.TransactOpts, strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "updateStrategyDebtRatio", strategy, debtRatio)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_YearnVault *YearnVaultSession) UpdateStrategyDebtRatio(strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.UpdateStrategyDebtRatio(&_YearnVault.TransactOpts, strategy, debtRatio)
}

// UpdateStrategyDebtRatio is a paid mutator transaction binding the contract method 0x7c6a4f24.
//
// Solidity: function updateStrategyDebtRatio(address strategy, uint256 debtRatio) returns()
func (_YearnVault *YearnVaultTransactorSession) UpdateStrategyDebtRatio(strategy common.Address, debtRatio *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.UpdateStrategyDebtRatio(&_YearnVault.TransactOpts, strategy, debtRatio)
}

// UpdateStrategyMaxDebtPerHarvest is a paid mutator transaction binding the contract method 0x4757a156.
//
// Solidity: function updateStrategyMaxDebtPerHarvest(address strategy, uint256 maxDebtPerHarvest) returns()
func (_YearnVault *YearnVaultTransactor) UpdateStrategyMaxDebtPerHarvest(opts *bind.TransactOpts, strategy common.Address, maxDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "updateStrategyMaxDebtPerHarvest", strategy, maxDebtPerHarvest)
}

// UpdateStrategyMaxDebtPerHarvest is a paid mutator transaction binding the contract method 0x4757a156.
//
// Solidity: function updateStrategyMaxDebtPerHarvest(address strategy, uint256 maxDebtPerHarvest) returns()
func (_YearnVault *YearnVaultSession) UpdateStrategyMaxDebtPerHarvest(strategy common.Address, maxDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.UpdateStrategyMaxDebtPerHarvest(&_YearnVault.TransactOpts, strategy, maxDebtPerHarvest)
}

// UpdateStrategyMaxDebtPerHarvest is a paid mutator transaction binding the contract method 0x4757a156.
//
// Solidity: function updateStrategyMaxDebtPerHarvest(address strategy, uint256 maxDebtPerHarvest) returns()
func (_YearnVault *YearnVaultTransactorSession) UpdateStrategyMaxDebtPerHarvest(strategy common.Address, maxDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.UpdateStrategyMaxDebtPerHarvest(&_YearnVault.TransactOpts, strategy, maxDebtPerHarvest)
}

// UpdateStrategyMinDebtPerHarvest is a paid mutator transaction binding the contract method 0xe722befe.
//
// Solidity: function updateStrategyMinDebtPerHarvest(address strategy, uint256 minDebtPerHarvest) returns()
func (_YearnVault *YearnVaultTransactor) UpdateStrategyMinDebtPerHarvest(opts *bind.TransactOpts, strategy common.Address, minDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "updateStrategyMinDebtPerHarvest", strategy, minDebtPerHarvest)
}

// UpdateStrategyMinDebtPerHarvest is a paid mutator transaction binding the contract method 0xe722befe.
//
// Solidity: function updateStrategyMinDebtPerHarvest(address strategy, uint256 minDebtPerHarvest) returns()
func (_YearnVault *YearnVaultSession) UpdateStrategyMinDebtPerHarvest(strategy common.Address, minDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.UpdateStrategyMinDebtPerHarvest(&_YearnVault.TransactOpts, strategy, minDebtPerHarvest)
}

// UpdateStrategyMinDebtPerHarvest is a paid mutator transaction binding the contract method 0xe722befe.
//
// Solidity: function updateStrategyMinDebtPerHarvest(address strategy, uint256 minDebtPerHarvest) returns()
func (_YearnVault *YearnVaultTransactorSession) UpdateStrategyMinDebtPerHarvest(strategy common.Address, minDebtPerHarvest *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.UpdateStrategyMinDebtPerHarvest(&_YearnVault.TransactOpts, strategy, minDebtPerHarvest)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_YearnVault *YearnVaultTransactor) UpdateStrategyPerformanceFee(opts *bind.TransactOpts, strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "updateStrategyPerformanceFee", strategy, performanceFee)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_YearnVault *YearnVaultSession) UpdateStrategyPerformanceFee(strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.UpdateStrategyPerformanceFee(&_YearnVault.TransactOpts, strategy, performanceFee)
}

// UpdateStrategyPerformanceFee is a paid mutator transaction binding the contract method 0xd0194ed6.
//
// Solidity: function updateStrategyPerformanceFee(address strategy, uint256 performanceFee) returns()
func (_YearnVault *YearnVaultTransactorSession) UpdateStrategyPerformanceFee(strategy common.Address, performanceFee *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.UpdateStrategyPerformanceFee(&_YearnVault.TransactOpts, strategy, performanceFee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YearnVault *YearnVaultTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YearnVault *YearnVaultSession) Withdraw() (*types.Transaction, error) {
	return _YearnVault.Contract.Withdraw(&_YearnVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YearnVault *YearnVaultTransactorSession) Withdraw() (*types.Transaction, error) {
	return _YearnVault.Contract.Withdraw(&_YearnVault.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_YearnVault *YearnVaultTransactor) Withdraw0(opts *bind.TransactOpts, maxShares *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "withdraw0", maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_YearnVault *YearnVaultSession) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Withdraw0(&_YearnVault.TransactOpts, maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_YearnVault *YearnVaultTransactorSession) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Withdraw0(&_YearnVault.TransactOpts, maxShares)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_YearnVault *YearnVaultTransactor) Withdraw1(opts *bind.TransactOpts, maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "withdraw1", maxShares, recipient)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_YearnVault *YearnVaultSession) Withdraw1(maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.Withdraw1(&_YearnVault.TransactOpts, maxShares, recipient)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address recipient) returns(uint256)
func (_YearnVault *YearnVaultTransactorSession) Withdraw1(maxShares *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.Withdraw1(&_YearnVault.TransactOpts, maxShares, recipient)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_YearnVault *YearnVaultTransactor) Withdraw2(opts *bind.TransactOpts, maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "withdraw2", maxShares, recipient, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_YearnVault *YearnVaultSession) Withdraw2(maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Withdraw2(&_YearnVault.TransactOpts, maxShares, recipient, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address recipient, uint256 maxLoss) returns(uint256)
func (_YearnVault *YearnVaultTransactorSession) Withdraw2(maxShares *big.Int, recipient common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Withdraw2(&_YearnVault.TransactOpts, maxShares, recipient, maxLoss)
}

// YearnVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the YearnVault contract.
type YearnVaultApprovalIterator struct {
	Event *YearnVaultApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultApproval represents a Approval event raised by the YearnVault contract.
type YearnVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnVault *YearnVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YearnVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultApprovalIterator{contract: _YearnVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnVault *YearnVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YearnVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultApproval)
				if err := _YearnVault.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnVault *YearnVaultFilterer) ParseApproval(log types.Log) (*YearnVaultApproval, error) {
	event := new(YearnVaultApproval)
	if err := _YearnVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultEmergencyShutdownIterator is returned from FilterEmergencyShutdown and is used to iterate over the raw logs and unpacked data for EmergencyShutdown events raised by the YearnVault contract.
type YearnVaultEmergencyShutdownIterator struct {
	Event *YearnVaultEmergencyShutdown // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultEmergencyShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultEmergencyShutdown)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultEmergencyShutdown)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultEmergencyShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultEmergencyShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultEmergencyShutdown represents a EmergencyShutdown event raised by the YearnVault contract.
type YearnVaultEmergencyShutdown struct {
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyShutdown is a free log retrieval operation binding the contract event 0xba40372a3a724dca3c57156128ef1e896724b65b37a17f190b1ad5de68f3a4f3.
//
// Solidity: event EmergencyShutdown(bool active)
func (_YearnVault *YearnVaultFilterer) FilterEmergencyShutdown(opts *bind.FilterOpts) (*YearnVaultEmergencyShutdownIterator, error) {

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "EmergencyShutdown")
	if err != nil {
		return nil, err
	}
	return &YearnVaultEmergencyShutdownIterator{contract: _YearnVault.contract, event: "EmergencyShutdown", logs: logs, sub: sub}, nil
}

// WatchEmergencyShutdown is a free log subscription operation binding the contract event 0xba40372a3a724dca3c57156128ef1e896724b65b37a17f190b1ad5de68f3a4f3.
//
// Solidity: event EmergencyShutdown(bool active)
func (_YearnVault *YearnVaultFilterer) WatchEmergencyShutdown(opts *bind.WatchOpts, sink chan<- *YearnVaultEmergencyShutdown) (event.Subscription, error) {

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "EmergencyShutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultEmergencyShutdown)
				if err := _YearnVault.contract.UnpackLog(event, "EmergencyShutdown", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyShutdown is a log parse operation binding the contract event 0xba40372a3a724dca3c57156128ef1e896724b65b37a17f190b1ad5de68f3a4f3.
//
// Solidity: event EmergencyShutdown(bool active)
func (_YearnVault *YearnVaultFilterer) ParseEmergencyShutdown(log types.Log) (*YearnVaultEmergencyShutdown, error) {
	event := new(YearnVaultEmergencyShutdown)
	if err := _YearnVault.contract.UnpackLog(event, "EmergencyShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultStrategyAddedIterator is returned from FilterStrategyAdded and is used to iterate over the raw logs and unpacked data for StrategyAdded events raised by the YearnVault contract.
type YearnVaultStrategyAddedIterator struct {
	Event *YearnVaultStrategyAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultStrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultStrategyAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultStrategyAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultStrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultStrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultStrategyAdded represents a StrategyAdded event raised by the YearnVault contract.
type YearnVaultStrategyAdded struct {
	Strategy          common.Address
	DebtRatio         *big.Int
	MinDebtPerHarvest *big.Int
	MaxDebtPerHarvest *big.Int
	PerformanceFee    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStrategyAdded is a free log retrieval operation binding the contract event 0x5a6abd2af9fe6c0554fa08649e2d86e4393ff19dc304d072d38d295c9291d4dc.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee)
func (_YearnVault *YearnVaultFilterer) FilterStrategyAdded(opts *bind.FilterOpts, strategy []common.Address) (*YearnVaultStrategyAddedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultStrategyAddedIterator{contract: _YearnVault.contract, event: "StrategyAdded", logs: logs, sub: sub}, nil
}

// WatchStrategyAdded is a free log subscription operation binding the contract event 0x5a6abd2af9fe6c0554fa08649e2d86e4393ff19dc304d072d38d295c9291d4dc.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee)
func (_YearnVault *YearnVaultFilterer) WatchStrategyAdded(opts *bind.WatchOpts, sink chan<- *YearnVaultStrategyAdded, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultStrategyAdded)
				if err := _YearnVault.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyAdded is a log parse operation binding the contract event 0x5a6abd2af9fe6c0554fa08649e2d86e4393ff19dc304d072d38d295c9291d4dc.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 debtRatio, uint256 minDebtPerHarvest, uint256 maxDebtPerHarvest, uint256 performanceFee)
func (_YearnVault *YearnVaultFilterer) ParseStrategyAdded(log types.Log) (*YearnVaultStrategyAdded, error) {
	event := new(YearnVaultStrategyAdded)
	if err := _YearnVault.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultStrategyAddedToQueueIterator is returned from FilterStrategyAddedToQueue and is used to iterate over the raw logs and unpacked data for StrategyAddedToQueue events raised by the YearnVault contract.
type YearnVaultStrategyAddedToQueueIterator struct {
	Event *YearnVaultStrategyAddedToQueue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultStrategyAddedToQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultStrategyAddedToQueue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultStrategyAddedToQueue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultStrategyAddedToQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultStrategyAddedToQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultStrategyAddedToQueue represents a StrategyAddedToQueue event raised by the YearnVault contract.
type YearnVaultStrategyAddedToQueue struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToQueue is a free log retrieval operation binding the contract event 0xa8727d412c6fa1e2497d6d6f275e2d9fe4d9318d5b793632e60ad9d38ee8f1fa.
//
// Solidity: event StrategyAddedToQueue(address indexed strategy)
func (_YearnVault *YearnVaultFilterer) FilterStrategyAddedToQueue(opts *bind.FilterOpts, strategy []common.Address) (*YearnVaultStrategyAddedToQueueIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "StrategyAddedToQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultStrategyAddedToQueueIterator{contract: _YearnVault.contract, event: "StrategyAddedToQueue", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToQueue is a free log subscription operation binding the contract event 0xa8727d412c6fa1e2497d6d6f275e2d9fe4d9318d5b793632e60ad9d38ee8f1fa.
//
// Solidity: event StrategyAddedToQueue(address indexed strategy)
func (_YearnVault *YearnVaultFilterer) WatchStrategyAddedToQueue(opts *bind.WatchOpts, sink chan<- *YearnVaultStrategyAddedToQueue, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "StrategyAddedToQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultStrategyAddedToQueue)
				if err := _YearnVault.contract.UnpackLog(event, "StrategyAddedToQueue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyAddedToQueue is a log parse operation binding the contract event 0xa8727d412c6fa1e2497d6d6f275e2d9fe4d9318d5b793632e60ad9d38ee8f1fa.
//
// Solidity: event StrategyAddedToQueue(address indexed strategy)
func (_YearnVault *YearnVaultFilterer) ParseStrategyAddedToQueue(log types.Log) (*YearnVaultStrategyAddedToQueue, error) {
	event := new(YearnVaultStrategyAddedToQueue)
	if err := _YearnVault.contract.UnpackLog(event, "StrategyAddedToQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultStrategyMigratedIterator is returned from FilterStrategyMigrated and is used to iterate over the raw logs and unpacked data for StrategyMigrated events raised by the YearnVault contract.
type YearnVaultStrategyMigratedIterator struct {
	Event *YearnVaultStrategyMigrated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultStrategyMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultStrategyMigrated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultStrategyMigrated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultStrategyMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultStrategyMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultStrategyMigrated represents a StrategyMigrated event raised by the YearnVault contract.
type YearnVaultStrategyMigrated struct {
	OldVersion common.Address
	NewVersion common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategyMigrated is a free log retrieval operation binding the contract event 0x100b69bb6b504e1252e36b375233158edee64d071b399e2f81473a695fd1b021.
//
// Solidity: event StrategyMigrated(address indexed oldVersion, address indexed newVersion)
func (_YearnVault *YearnVaultFilterer) FilterStrategyMigrated(opts *bind.FilterOpts, oldVersion []common.Address, newVersion []common.Address) (*YearnVaultStrategyMigratedIterator, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "StrategyMigrated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultStrategyMigratedIterator{contract: _YearnVault.contract, event: "StrategyMigrated", logs: logs, sub: sub}, nil
}

// WatchStrategyMigrated is a free log subscription operation binding the contract event 0x100b69bb6b504e1252e36b375233158edee64d071b399e2f81473a695fd1b021.
//
// Solidity: event StrategyMigrated(address indexed oldVersion, address indexed newVersion)
func (_YearnVault *YearnVaultFilterer) WatchStrategyMigrated(opts *bind.WatchOpts, sink chan<- *YearnVaultStrategyMigrated, oldVersion []common.Address, newVersion []common.Address) (event.Subscription, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "StrategyMigrated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultStrategyMigrated)
				if err := _YearnVault.contract.UnpackLog(event, "StrategyMigrated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyMigrated is a log parse operation binding the contract event 0x100b69bb6b504e1252e36b375233158edee64d071b399e2f81473a695fd1b021.
//
// Solidity: event StrategyMigrated(address indexed oldVersion, address indexed newVersion)
func (_YearnVault *YearnVaultFilterer) ParseStrategyMigrated(log types.Log) (*YearnVaultStrategyMigrated, error) {
	event := new(YearnVaultStrategyMigrated)
	if err := _YearnVault.contract.UnpackLog(event, "StrategyMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultStrategyRemovedFromQueueIterator is returned from FilterStrategyRemovedFromQueue and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromQueue events raised by the YearnVault contract.
type YearnVaultStrategyRemovedFromQueueIterator struct {
	Event *YearnVaultStrategyRemovedFromQueue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultStrategyRemovedFromQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultStrategyRemovedFromQueue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultStrategyRemovedFromQueue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultStrategyRemovedFromQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultStrategyRemovedFromQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultStrategyRemovedFromQueue represents a StrategyRemovedFromQueue event raised by the YearnVault contract.
type YearnVaultStrategyRemovedFromQueue struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromQueue is a free log retrieval operation binding the contract event 0x8e1ec3c16d6a67ea8effe2ac7adef9c2de0bc0dc47c49cdf18f6a8b0048085be.
//
// Solidity: event StrategyRemovedFromQueue(address indexed strategy)
func (_YearnVault *YearnVaultFilterer) FilterStrategyRemovedFromQueue(opts *bind.FilterOpts, strategy []common.Address) (*YearnVaultStrategyRemovedFromQueueIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "StrategyRemovedFromQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultStrategyRemovedFromQueueIterator{contract: _YearnVault.contract, event: "StrategyRemovedFromQueue", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromQueue is a free log subscription operation binding the contract event 0x8e1ec3c16d6a67ea8effe2ac7adef9c2de0bc0dc47c49cdf18f6a8b0048085be.
//
// Solidity: event StrategyRemovedFromQueue(address indexed strategy)
func (_YearnVault *YearnVaultFilterer) WatchStrategyRemovedFromQueue(opts *bind.WatchOpts, sink chan<- *YearnVaultStrategyRemovedFromQueue, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "StrategyRemovedFromQueue", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultStrategyRemovedFromQueue)
				if err := _YearnVault.contract.UnpackLog(event, "StrategyRemovedFromQueue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyRemovedFromQueue is a log parse operation binding the contract event 0x8e1ec3c16d6a67ea8effe2ac7adef9c2de0bc0dc47c49cdf18f6a8b0048085be.
//
// Solidity: event StrategyRemovedFromQueue(address indexed strategy)
func (_YearnVault *YearnVaultFilterer) ParseStrategyRemovedFromQueue(log types.Log) (*YearnVaultStrategyRemovedFromQueue, error) {
	event := new(YearnVaultStrategyRemovedFromQueue)
	if err := _YearnVault.contract.UnpackLog(event, "StrategyRemovedFromQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultStrategyReportedIterator is returned from FilterStrategyReported and is used to iterate over the raw logs and unpacked data for StrategyReported events raised by the YearnVault contract.
type YearnVaultStrategyReportedIterator struct {
	Event *YearnVaultStrategyReported // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultStrategyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultStrategyReported)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultStrategyReported)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultStrategyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultStrategyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultStrategyReported represents a StrategyReported event raised by the YearnVault contract.
type YearnVaultStrategyReported struct {
	Strategy  common.Address
	Gain      *big.Int
	Loss      *big.Int
	DebtPaid  *big.Int
	TotalGain *big.Int
	TotalLoss *big.Int
	TotalDebt *big.Int
	DebtAdded *big.Int
	DebtRatio *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStrategyReported is a free log retrieval operation binding the contract event 0x67f96d2854a335a4cadb49f84fd3ca6f990744ddb3feceeb4b349d2d53d32ad3.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 debtPaid, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtRatio)
func (_YearnVault *YearnVaultFilterer) FilterStrategyReported(opts *bind.FilterOpts, strategy []common.Address) (*YearnVaultStrategyReportedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultStrategyReportedIterator{contract: _YearnVault.contract, event: "StrategyReported", logs: logs, sub: sub}, nil
}

// WatchStrategyReported is a free log subscription operation binding the contract event 0x67f96d2854a335a4cadb49f84fd3ca6f990744ddb3feceeb4b349d2d53d32ad3.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 debtPaid, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtRatio)
func (_YearnVault *YearnVaultFilterer) WatchStrategyReported(opts *bind.WatchOpts, sink chan<- *YearnVaultStrategyReported, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultStrategyReported)
				if err := _YearnVault.contract.UnpackLog(event, "StrategyReported", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyReported is a log parse operation binding the contract event 0x67f96d2854a335a4cadb49f84fd3ca6f990744ddb3feceeb4b349d2d53d32ad3.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 debtPaid, uint256 totalGain, uint256 totalLoss, uint256 totalDebt, uint256 debtAdded, uint256 debtRatio)
func (_YearnVault *YearnVaultFilterer) ParseStrategyReported(log types.Log) (*YearnVaultStrategyReported, error) {
	event := new(YearnVaultStrategyReported)
	if err := _YearnVault.contract.UnpackLog(event, "StrategyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultStrategyRevokedIterator is returned from FilterStrategyRevoked and is used to iterate over the raw logs and unpacked data for StrategyRevoked events raised by the YearnVault contract.
type YearnVaultStrategyRevokedIterator struct {
	Event *YearnVaultStrategyRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultStrategyRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultStrategyRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultStrategyRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultStrategyRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultStrategyRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultStrategyRevoked represents a StrategyRevoked event raised by the YearnVault contract.
type YearnVaultStrategyRevoked struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRevoked is a free log retrieval operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_YearnVault *YearnVaultFilterer) FilterStrategyRevoked(opts *bind.FilterOpts, strategy []common.Address) (*YearnVaultStrategyRevokedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "StrategyRevoked", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultStrategyRevokedIterator{contract: _YearnVault.contract, event: "StrategyRevoked", logs: logs, sub: sub}, nil
}

// WatchStrategyRevoked is a free log subscription operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_YearnVault *YearnVaultFilterer) WatchStrategyRevoked(opts *bind.WatchOpts, sink chan<- *YearnVaultStrategyRevoked, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "StrategyRevoked", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultStrategyRevoked)
				if err := _YearnVault.contract.UnpackLog(event, "StrategyRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyRevoked is a log parse operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_YearnVault *YearnVaultFilterer) ParseStrategyRevoked(log types.Log) (*YearnVaultStrategyRevoked, error) {
	event := new(YearnVaultStrategyRevoked)
	if err := _YearnVault.contract.UnpackLog(event, "StrategyRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultStrategyUpdateDebtRatioIterator is returned from FilterStrategyUpdateDebtRatio and is used to iterate over the raw logs and unpacked data for StrategyUpdateDebtRatio events raised by the YearnVault contract.
type YearnVaultStrategyUpdateDebtRatioIterator struct {
	Event *YearnVaultStrategyUpdateDebtRatio // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultStrategyUpdateDebtRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultStrategyUpdateDebtRatio)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultStrategyUpdateDebtRatio)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultStrategyUpdateDebtRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultStrategyUpdateDebtRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultStrategyUpdateDebtRatio represents a StrategyUpdateDebtRatio event raised by the YearnVault contract.
type YearnVaultStrategyUpdateDebtRatio struct {
	Strategy  common.Address
	DebtRatio *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdateDebtRatio is a free log retrieval operation binding the contract event 0xbda9398315c83ccef012bcaa318a2ff7b680f36429d36597bd4bc25ac11ead59.
//
// Solidity: event StrategyUpdateDebtRatio(address indexed strategy, uint256 debtRatio)
func (_YearnVault *YearnVaultFilterer) FilterStrategyUpdateDebtRatio(opts *bind.FilterOpts, strategy []common.Address) (*YearnVaultStrategyUpdateDebtRatioIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "StrategyUpdateDebtRatio", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultStrategyUpdateDebtRatioIterator{contract: _YearnVault.contract, event: "StrategyUpdateDebtRatio", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdateDebtRatio is a free log subscription operation binding the contract event 0xbda9398315c83ccef012bcaa318a2ff7b680f36429d36597bd4bc25ac11ead59.
//
// Solidity: event StrategyUpdateDebtRatio(address indexed strategy, uint256 debtRatio)
func (_YearnVault *YearnVaultFilterer) WatchStrategyUpdateDebtRatio(opts *bind.WatchOpts, sink chan<- *YearnVaultStrategyUpdateDebtRatio, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "StrategyUpdateDebtRatio", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultStrategyUpdateDebtRatio)
				if err := _YearnVault.contract.UnpackLog(event, "StrategyUpdateDebtRatio", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyUpdateDebtRatio is a log parse operation binding the contract event 0xbda9398315c83ccef012bcaa318a2ff7b680f36429d36597bd4bc25ac11ead59.
//
// Solidity: event StrategyUpdateDebtRatio(address indexed strategy, uint256 debtRatio)
func (_YearnVault *YearnVaultFilterer) ParseStrategyUpdateDebtRatio(log types.Log) (*YearnVaultStrategyUpdateDebtRatio, error) {
	event := new(YearnVaultStrategyUpdateDebtRatio)
	if err := _YearnVault.contract.UnpackLog(event, "StrategyUpdateDebtRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultStrategyUpdateMaxDebtPerHarvestIterator is returned from FilterStrategyUpdateMaxDebtPerHarvest and is used to iterate over the raw logs and unpacked data for StrategyUpdateMaxDebtPerHarvest events raised by the YearnVault contract.
type YearnVaultStrategyUpdateMaxDebtPerHarvestIterator struct {
	Event *YearnVaultStrategyUpdateMaxDebtPerHarvest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultStrategyUpdateMaxDebtPerHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultStrategyUpdateMaxDebtPerHarvest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultStrategyUpdateMaxDebtPerHarvest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultStrategyUpdateMaxDebtPerHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultStrategyUpdateMaxDebtPerHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultStrategyUpdateMaxDebtPerHarvest represents a StrategyUpdateMaxDebtPerHarvest event raised by the YearnVault contract.
type YearnVaultStrategyUpdateMaxDebtPerHarvest struct {
	Strategy          common.Address
	MaxDebtPerHarvest *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdateMaxDebtPerHarvest is a free log retrieval operation binding the contract event 0x1796a8e0760e2de5b72e7bf64fccb7666c48ceab94cb6cae7cb7eff4b6f641ab.
//
// Solidity: event StrategyUpdateMaxDebtPerHarvest(address indexed strategy, uint256 maxDebtPerHarvest)
func (_YearnVault *YearnVaultFilterer) FilterStrategyUpdateMaxDebtPerHarvest(opts *bind.FilterOpts, strategy []common.Address) (*YearnVaultStrategyUpdateMaxDebtPerHarvestIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "StrategyUpdateMaxDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultStrategyUpdateMaxDebtPerHarvestIterator{contract: _YearnVault.contract, event: "StrategyUpdateMaxDebtPerHarvest", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdateMaxDebtPerHarvest is a free log subscription operation binding the contract event 0x1796a8e0760e2de5b72e7bf64fccb7666c48ceab94cb6cae7cb7eff4b6f641ab.
//
// Solidity: event StrategyUpdateMaxDebtPerHarvest(address indexed strategy, uint256 maxDebtPerHarvest)
func (_YearnVault *YearnVaultFilterer) WatchStrategyUpdateMaxDebtPerHarvest(opts *bind.WatchOpts, sink chan<- *YearnVaultStrategyUpdateMaxDebtPerHarvest, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "StrategyUpdateMaxDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultStrategyUpdateMaxDebtPerHarvest)
				if err := _YearnVault.contract.UnpackLog(event, "StrategyUpdateMaxDebtPerHarvest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyUpdateMaxDebtPerHarvest is a log parse operation binding the contract event 0x1796a8e0760e2de5b72e7bf64fccb7666c48ceab94cb6cae7cb7eff4b6f641ab.
//
// Solidity: event StrategyUpdateMaxDebtPerHarvest(address indexed strategy, uint256 maxDebtPerHarvest)
func (_YearnVault *YearnVaultFilterer) ParseStrategyUpdateMaxDebtPerHarvest(log types.Log) (*YearnVaultStrategyUpdateMaxDebtPerHarvest, error) {
	event := new(YearnVaultStrategyUpdateMaxDebtPerHarvest)
	if err := _YearnVault.contract.UnpackLog(event, "StrategyUpdateMaxDebtPerHarvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultStrategyUpdateMinDebtPerHarvestIterator is returned from FilterStrategyUpdateMinDebtPerHarvest and is used to iterate over the raw logs and unpacked data for StrategyUpdateMinDebtPerHarvest events raised by the YearnVault contract.
type YearnVaultStrategyUpdateMinDebtPerHarvestIterator struct {
	Event *YearnVaultStrategyUpdateMinDebtPerHarvest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultStrategyUpdateMinDebtPerHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultStrategyUpdateMinDebtPerHarvest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultStrategyUpdateMinDebtPerHarvest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultStrategyUpdateMinDebtPerHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultStrategyUpdateMinDebtPerHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultStrategyUpdateMinDebtPerHarvest represents a StrategyUpdateMinDebtPerHarvest event raised by the YearnVault contract.
type YearnVaultStrategyUpdateMinDebtPerHarvest struct {
	Strategy          common.Address
	MinDebtPerHarvest *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdateMinDebtPerHarvest is a free log retrieval operation binding the contract event 0x0b728ad785976532c4aaadde09b1cba5f262a7090e83c62d2377bc405678b29c.
//
// Solidity: event StrategyUpdateMinDebtPerHarvest(address indexed strategy, uint256 minDebtPerHarvest)
func (_YearnVault *YearnVaultFilterer) FilterStrategyUpdateMinDebtPerHarvest(opts *bind.FilterOpts, strategy []common.Address) (*YearnVaultStrategyUpdateMinDebtPerHarvestIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "StrategyUpdateMinDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultStrategyUpdateMinDebtPerHarvestIterator{contract: _YearnVault.contract, event: "StrategyUpdateMinDebtPerHarvest", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdateMinDebtPerHarvest is a free log subscription operation binding the contract event 0x0b728ad785976532c4aaadde09b1cba5f262a7090e83c62d2377bc405678b29c.
//
// Solidity: event StrategyUpdateMinDebtPerHarvest(address indexed strategy, uint256 minDebtPerHarvest)
func (_YearnVault *YearnVaultFilterer) WatchStrategyUpdateMinDebtPerHarvest(opts *bind.WatchOpts, sink chan<- *YearnVaultStrategyUpdateMinDebtPerHarvest, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "StrategyUpdateMinDebtPerHarvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultStrategyUpdateMinDebtPerHarvest)
				if err := _YearnVault.contract.UnpackLog(event, "StrategyUpdateMinDebtPerHarvest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyUpdateMinDebtPerHarvest is a log parse operation binding the contract event 0x0b728ad785976532c4aaadde09b1cba5f262a7090e83c62d2377bc405678b29c.
//
// Solidity: event StrategyUpdateMinDebtPerHarvest(address indexed strategy, uint256 minDebtPerHarvest)
func (_YearnVault *YearnVaultFilterer) ParseStrategyUpdateMinDebtPerHarvest(log types.Log) (*YearnVaultStrategyUpdateMinDebtPerHarvest, error) {
	event := new(YearnVaultStrategyUpdateMinDebtPerHarvest)
	if err := _YearnVault.contract.UnpackLog(event, "StrategyUpdateMinDebtPerHarvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultStrategyUpdatePerformanceFeeIterator is returned from FilterStrategyUpdatePerformanceFee and is used to iterate over the raw logs and unpacked data for StrategyUpdatePerformanceFee events raised by the YearnVault contract.
type YearnVaultStrategyUpdatePerformanceFeeIterator struct {
	Event *YearnVaultStrategyUpdatePerformanceFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultStrategyUpdatePerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultStrategyUpdatePerformanceFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultStrategyUpdatePerformanceFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultStrategyUpdatePerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultStrategyUpdatePerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultStrategyUpdatePerformanceFee represents a StrategyUpdatePerformanceFee event raised by the YearnVault contract.
type YearnVaultStrategyUpdatePerformanceFee struct {
	Strategy       common.Address
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdatePerformanceFee is a free log retrieval operation binding the contract event 0xe57488a65fa53066d4c25bac90db47dda4e5de3025ac12bf76ff07211cf7f39e.
//
// Solidity: event StrategyUpdatePerformanceFee(address indexed strategy, uint256 performanceFee)
func (_YearnVault *YearnVaultFilterer) FilterStrategyUpdatePerformanceFee(opts *bind.FilterOpts, strategy []common.Address) (*YearnVaultStrategyUpdatePerformanceFeeIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "StrategyUpdatePerformanceFee", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultStrategyUpdatePerformanceFeeIterator{contract: _YearnVault.contract, event: "StrategyUpdatePerformanceFee", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdatePerformanceFee is a free log subscription operation binding the contract event 0xe57488a65fa53066d4c25bac90db47dda4e5de3025ac12bf76ff07211cf7f39e.
//
// Solidity: event StrategyUpdatePerformanceFee(address indexed strategy, uint256 performanceFee)
func (_YearnVault *YearnVaultFilterer) WatchStrategyUpdatePerformanceFee(opts *bind.WatchOpts, sink chan<- *YearnVaultStrategyUpdatePerformanceFee, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "StrategyUpdatePerformanceFee", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultStrategyUpdatePerformanceFee)
				if err := _YearnVault.contract.UnpackLog(event, "StrategyUpdatePerformanceFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyUpdatePerformanceFee is a log parse operation binding the contract event 0xe57488a65fa53066d4c25bac90db47dda4e5de3025ac12bf76ff07211cf7f39e.
//
// Solidity: event StrategyUpdatePerformanceFee(address indexed strategy, uint256 performanceFee)
func (_YearnVault *YearnVaultFilterer) ParseStrategyUpdatePerformanceFee(log types.Log) (*YearnVaultStrategyUpdatePerformanceFee, error) {
	event := new(YearnVaultStrategyUpdatePerformanceFee)
	if err := _YearnVault.contract.UnpackLog(event, "StrategyUpdatePerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the YearnVault contract.
type YearnVaultTransferIterator struct {
	Event *YearnVaultTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultTransfer represents a Transfer event raised by the YearnVault contract.
type YearnVaultTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_YearnVault *YearnVaultFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*YearnVaultTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &YearnVaultTransferIterator{contract: _YearnVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_YearnVault *YearnVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YearnVaultTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultTransfer)
				if err := _YearnVault.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_YearnVault *YearnVaultFilterer) ParseTransfer(log types.Log) (*YearnVaultTransfer, error) {
	event := new(YearnVaultTransfer)
	if err := _YearnVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultUpdateDepositLimitIterator is returned from FilterUpdateDepositLimit and is used to iterate over the raw logs and unpacked data for UpdateDepositLimit events raised by the YearnVault contract.
type YearnVaultUpdateDepositLimitIterator struct {
	Event *YearnVaultUpdateDepositLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultUpdateDepositLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultUpdateDepositLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultUpdateDepositLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultUpdateDepositLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultUpdateDepositLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultUpdateDepositLimit represents a UpdateDepositLimit event raised by the YearnVault contract.
type YearnVaultUpdateDepositLimit struct {
	DepositLimit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimit is a free log retrieval operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 depositLimit)
func (_YearnVault *YearnVaultFilterer) FilterUpdateDepositLimit(opts *bind.FilterOpts) (*YearnVaultUpdateDepositLimitIterator, error) {

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return &YearnVaultUpdateDepositLimitIterator{contract: _YearnVault.contract, event: "UpdateDepositLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimit is a free log subscription operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 depositLimit)
func (_YearnVault *YearnVaultFilterer) WatchUpdateDepositLimit(opts *bind.WatchOpts, sink chan<- *YearnVaultUpdateDepositLimit) (event.Subscription, error) {

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultUpdateDepositLimit)
				if err := _YearnVault.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateDepositLimit is a log parse operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 depositLimit)
func (_YearnVault *YearnVaultFilterer) ParseUpdateDepositLimit(log types.Log) (*YearnVaultUpdateDepositLimit, error) {
	event := new(YearnVaultUpdateDepositLimit)
	if err := _YearnVault.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultUpdateGovernanceIterator is returned from FilterUpdateGovernance and is used to iterate over the raw logs and unpacked data for UpdateGovernance events raised by the YearnVault contract.
type YearnVaultUpdateGovernanceIterator struct {
	Event *YearnVaultUpdateGovernance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultUpdateGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultUpdateGovernance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultUpdateGovernance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultUpdateGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultUpdateGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultUpdateGovernance represents a UpdateGovernance event raised by the YearnVault contract.
type YearnVaultUpdateGovernance struct {
	Governance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateGovernance is a free log retrieval operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address governance)
func (_YearnVault *YearnVaultFilterer) FilterUpdateGovernance(opts *bind.FilterOpts) (*YearnVaultUpdateGovernanceIterator, error) {

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "UpdateGovernance")
	if err != nil {
		return nil, err
	}
	return &YearnVaultUpdateGovernanceIterator{contract: _YearnVault.contract, event: "UpdateGovernance", logs: logs, sub: sub}, nil
}

// WatchUpdateGovernance is a free log subscription operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address governance)
func (_YearnVault *YearnVaultFilterer) WatchUpdateGovernance(opts *bind.WatchOpts, sink chan<- *YearnVaultUpdateGovernance) (event.Subscription, error) {

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "UpdateGovernance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultUpdateGovernance)
				if err := _YearnVault.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateGovernance is a log parse operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address governance)
func (_YearnVault *YearnVaultFilterer) ParseUpdateGovernance(log types.Log) (*YearnVaultUpdateGovernance, error) {
	event := new(YearnVaultUpdateGovernance)
	if err := _YearnVault.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultUpdateGuardianIterator is returned from FilterUpdateGuardian and is used to iterate over the raw logs and unpacked data for UpdateGuardian events raised by the YearnVault contract.
type YearnVaultUpdateGuardianIterator struct {
	Event *YearnVaultUpdateGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultUpdateGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultUpdateGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultUpdateGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultUpdateGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultUpdateGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultUpdateGuardian represents a UpdateGuardian event raised by the YearnVault contract.
type YearnVaultUpdateGuardian struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateGuardian is a free log retrieval operation binding the contract event 0x837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c.
//
// Solidity: event UpdateGuardian(address guardian)
func (_YearnVault *YearnVaultFilterer) FilterUpdateGuardian(opts *bind.FilterOpts) (*YearnVaultUpdateGuardianIterator, error) {

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "UpdateGuardian")
	if err != nil {
		return nil, err
	}
	return &YearnVaultUpdateGuardianIterator{contract: _YearnVault.contract, event: "UpdateGuardian", logs: logs, sub: sub}, nil
}

// WatchUpdateGuardian is a free log subscription operation binding the contract event 0x837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c.
//
// Solidity: event UpdateGuardian(address guardian)
func (_YearnVault *YearnVaultFilterer) WatchUpdateGuardian(opts *bind.WatchOpts, sink chan<- *YearnVaultUpdateGuardian) (event.Subscription, error) {

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "UpdateGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultUpdateGuardian)
				if err := _YearnVault.contract.UnpackLog(event, "UpdateGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateGuardian is a log parse operation binding the contract event 0x837b9ad138a0a1839a9637afce5306a5c13e23eb63365686843a5319a243609c.
//
// Solidity: event UpdateGuardian(address guardian)
func (_YearnVault *YearnVaultFilterer) ParseUpdateGuardian(log types.Log) (*YearnVaultUpdateGuardian, error) {
	event := new(YearnVaultUpdateGuardian)
	if err := _YearnVault.contract.UnpackLog(event, "UpdateGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultUpdateGuestListIterator is returned from FilterUpdateGuestList and is used to iterate over the raw logs and unpacked data for UpdateGuestList events raised by the YearnVault contract.
type YearnVaultUpdateGuestListIterator struct {
	Event *YearnVaultUpdateGuestList // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultUpdateGuestListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultUpdateGuestList)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultUpdateGuestList)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultUpdateGuestListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultUpdateGuestListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultUpdateGuestList represents a UpdateGuestList event raised by the YearnVault contract.
type YearnVaultUpdateGuestList struct {
	GuestList common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateGuestList is a free log retrieval operation binding the contract event 0x6d674c311329fb38bbc96dc33d2aad03b9bf9fcfdd8f5e5054fda291a5b3c1f8.
//
// Solidity: event UpdateGuestList(address guestList)
func (_YearnVault *YearnVaultFilterer) FilterUpdateGuestList(opts *bind.FilterOpts) (*YearnVaultUpdateGuestListIterator, error) {

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "UpdateGuestList")
	if err != nil {
		return nil, err
	}
	return &YearnVaultUpdateGuestListIterator{contract: _YearnVault.contract, event: "UpdateGuestList", logs: logs, sub: sub}, nil
}

// WatchUpdateGuestList is a free log subscription operation binding the contract event 0x6d674c311329fb38bbc96dc33d2aad03b9bf9fcfdd8f5e5054fda291a5b3c1f8.
//
// Solidity: event UpdateGuestList(address guestList)
func (_YearnVault *YearnVaultFilterer) WatchUpdateGuestList(opts *bind.WatchOpts, sink chan<- *YearnVaultUpdateGuestList) (event.Subscription, error) {

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "UpdateGuestList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultUpdateGuestList)
				if err := _YearnVault.contract.UnpackLog(event, "UpdateGuestList", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateGuestList is a log parse operation binding the contract event 0x6d674c311329fb38bbc96dc33d2aad03b9bf9fcfdd8f5e5054fda291a5b3c1f8.
//
// Solidity: event UpdateGuestList(address guestList)
func (_YearnVault *YearnVaultFilterer) ParseUpdateGuestList(log types.Log) (*YearnVaultUpdateGuestList, error) {
	event := new(YearnVaultUpdateGuestList)
	if err := _YearnVault.contract.UnpackLog(event, "UpdateGuestList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultUpdateManagementIterator is returned from FilterUpdateManagement and is used to iterate over the raw logs and unpacked data for UpdateManagement events raised by the YearnVault contract.
type YearnVaultUpdateManagementIterator struct {
	Event *YearnVaultUpdateManagement // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultUpdateManagementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultUpdateManagement)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultUpdateManagement)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultUpdateManagementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultUpdateManagementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultUpdateManagement represents a UpdateManagement event raised by the YearnVault contract.
type YearnVaultUpdateManagement struct {
	Management common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagement is a free log retrieval operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address management)
func (_YearnVault *YearnVaultFilterer) FilterUpdateManagement(opts *bind.FilterOpts) (*YearnVaultUpdateManagementIterator, error) {

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "UpdateManagement")
	if err != nil {
		return nil, err
	}
	return &YearnVaultUpdateManagementIterator{contract: _YearnVault.contract, event: "UpdateManagement", logs: logs, sub: sub}, nil
}

// WatchUpdateManagement is a free log subscription operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address management)
func (_YearnVault *YearnVaultFilterer) WatchUpdateManagement(opts *bind.WatchOpts, sink chan<- *YearnVaultUpdateManagement) (event.Subscription, error) {

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "UpdateManagement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultUpdateManagement)
				if err := _YearnVault.contract.UnpackLog(event, "UpdateManagement", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateManagement is a log parse operation binding the contract event 0xff54978127edd34aec0f9061fb3b155fbe0ededdfa881ee3e0d541d3a1eef438.
//
// Solidity: event UpdateManagement(address management)
func (_YearnVault *YearnVaultFilterer) ParseUpdateManagement(log types.Log) (*YearnVaultUpdateManagement, error) {
	event := new(YearnVaultUpdateManagement)
	if err := _YearnVault.contract.UnpackLog(event, "UpdateManagement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultUpdateManagementFeeIterator is returned from FilterUpdateManagementFee and is used to iterate over the raw logs and unpacked data for UpdateManagementFee events raised by the YearnVault contract.
type YearnVaultUpdateManagementFeeIterator struct {
	Event *YearnVaultUpdateManagementFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultUpdateManagementFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultUpdateManagementFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultUpdateManagementFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultUpdateManagementFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultUpdateManagementFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultUpdateManagementFee represents a UpdateManagementFee event raised by the YearnVault contract.
type YearnVaultUpdateManagementFee struct {
	ManagementFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagementFee is a free log retrieval operation binding the contract event 0x7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf.
//
// Solidity: event UpdateManagementFee(uint256 managementFee)
func (_YearnVault *YearnVaultFilterer) FilterUpdateManagementFee(opts *bind.FilterOpts) (*YearnVaultUpdateManagementFeeIterator, error) {

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "UpdateManagementFee")
	if err != nil {
		return nil, err
	}
	return &YearnVaultUpdateManagementFeeIterator{contract: _YearnVault.contract, event: "UpdateManagementFee", logs: logs, sub: sub}, nil
}

// WatchUpdateManagementFee is a free log subscription operation binding the contract event 0x7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf.
//
// Solidity: event UpdateManagementFee(uint256 managementFee)
func (_YearnVault *YearnVaultFilterer) WatchUpdateManagementFee(opts *bind.WatchOpts, sink chan<- *YearnVaultUpdateManagementFee) (event.Subscription, error) {

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "UpdateManagementFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultUpdateManagementFee)
				if err := _YearnVault.contract.UnpackLog(event, "UpdateManagementFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateManagementFee is a log parse operation binding the contract event 0x7a7883b0074f96e2c7fab65eb25abf624c488761a5db889e3bb84855dcc6daaf.
//
// Solidity: event UpdateManagementFee(uint256 managementFee)
func (_YearnVault *YearnVaultFilterer) ParseUpdateManagementFee(log types.Log) (*YearnVaultUpdateManagementFee, error) {
	event := new(YearnVaultUpdateManagementFee)
	if err := _YearnVault.contract.UnpackLog(event, "UpdateManagementFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultUpdatePerformanceFeeIterator is returned from FilterUpdatePerformanceFee and is used to iterate over the raw logs and unpacked data for UpdatePerformanceFee events raised by the YearnVault contract.
type YearnVaultUpdatePerformanceFeeIterator struct {
	Event *YearnVaultUpdatePerformanceFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultUpdatePerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultUpdatePerformanceFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultUpdatePerformanceFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultUpdatePerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultUpdatePerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultUpdatePerformanceFee represents a UpdatePerformanceFee event raised by the YearnVault contract.
type YearnVaultUpdatePerformanceFee struct {
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdatePerformanceFee is a free log retrieval operation binding the contract event 0x0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb.
//
// Solidity: event UpdatePerformanceFee(uint256 performanceFee)
func (_YearnVault *YearnVaultFilterer) FilterUpdatePerformanceFee(opts *bind.FilterOpts) (*YearnVaultUpdatePerformanceFeeIterator, error) {

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "UpdatePerformanceFee")
	if err != nil {
		return nil, err
	}
	return &YearnVaultUpdatePerformanceFeeIterator{contract: _YearnVault.contract, event: "UpdatePerformanceFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePerformanceFee is a free log subscription operation binding the contract event 0x0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb.
//
// Solidity: event UpdatePerformanceFee(uint256 performanceFee)
func (_YearnVault *YearnVaultFilterer) WatchUpdatePerformanceFee(opts *bind.WatchOpts, sink chan<- *YearnVaultUpdatePerformanceFee) (event.Subscription, error) {

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "UpdatePerformanceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultUpdatePerformanceFee)
				if err := _YearnVault.contract.UnpackLog(event, "UpdatePerformanceFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatePerformanceFee is a log parse operation binding the contract event 0x0810a1c261ca2c0cd86a0152c51c43ba9dc329639d2349f98140891b2ea798eb.
//
// Solidity: event UpdatePerformanceFee(uint256 performanceFee)
func (_YearnVault *YearnVaultFilterer) ParseUpdatePerformanceFee(log types.Log) (*YearnVaultUpdatePerformanceFee, error) {
	event := new(YearnVaultUpdatePerformanceFee)
	if err := _YearnVault.contract.UnpackLog(event, "UpdatePerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultUpdateRewardsIterator is returned from FilterUpdateRewards and is used to iterate over the raw logs and unpacked data for UpdateRewards events raised by the YearnVault contract.
type YearnVaultUpdateRewardsIterator struct {
	Event *YearnVaultUpdateRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultUpdateRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultUpdateRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultUpdateRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultUpdateRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultUpdateRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultUpdateRewards represents a UpdateRewards event raised by the YearnVault contract.
type YearnVaultUpdateRewards struct {
	Rewards common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateRewards is a free log retrieval operation binding the contract event 0xdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d.
//
// Solidity: event UpdateRewards(address rewards)
func (_YearnVault *YearnVaultFilterer) FilterUpdateRewards(opts *bind.FilterOpts) (*YearnVaultUpdateRewardsIterator, error) {

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "UpdateRewards")
	if err != nil {
		return nil, err
	}
	return &YearnVaultUpdateRewardsIterator{contract: _YearnVault.contract, event: "UpdateRewards", logs: logs, sub: sub}, nil
}

// WatchUpdateRewards is a free log subscription operation binding the contract event 0xdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d.
//
// Solidity: event UpdateRewards(address rewards)
func (_YearnVault *YearnVaultFilterer) WatchUpdateRewards(opts *bind.WatchOpts, sink chan<- *YearnVaultUpdateRewards) (event.Subscription, error) {

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "UpdateRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultUpdateRewards)
				if err := _YearnVault.contract.UnpackLog(event, "UpdateRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateRewards is a log parse operation binding the contract event 0xdf3c41a916aecbf42361a147f8348c242662c3ce20ecef30e826b80642477a3d.
//
// Solidity: event UpdateRewards(address rewards)
func (_YearnVault *YearnVaultFilterer) ParseUpdateRewards(log types.Log) (*YearnVaultUpdateRewards, error) {
	event := new(YearnVaultUpdateRewards)
	if err := _YearnVault.contract.UnpackLog(event, "UpdateRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnVaultUpdateWithdrawalQueueIterator is returned from FilterUpdateWithdrawalQueue and is used to iterate over the raw logs and unpacked data for UpdateWithdrawalQueue events raised by the YearnVault contract.
type YearnVaultUpdateWithdrawalQueueIterator struct {
	Event *YearnVaultUpdateWithdrawalQueue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YearnVaultUpdateWithdrawalQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnVaultUpdateWithdrawalQueue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YearnVaultUpdateWithdrawalQueue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YearnVaultUpdateWithdrawalQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnVaultUpdateWithdrawalQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnVaultUpdateWithdrawalQueue represents a UpdateWithdrawalQueue event raised by the YearnVault contract.
type YearnVaultUpdateWithdrawalQueue struct {
	Queue [20]common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateWithdrawalQueue is a free log retrieval operation binding the contract event 0x695ac3ac73f08f2002284ffe563cefe798ee2878a5e04219522e2e99eb89d168.
//
// Solidity: event UpdateWithdrawalQueue(address[20] queue)
func (_YearnVault *YearnVaultFilterer) FilterUpdateWithdrawalQueue(opts *bind.FilterOpts) (*YearnVaultUpdateWithdrawalQueueIterator, error) {

	logs, sub, err := _YearnVault.contract.FilterLogs(opts, "UpdateWithdrawalQueue")
	if err != nil {
		return nil, err
	}
	return &YearnVaultUpdateWithdrawalQueueIterator{contract: _YearnVault.contract, event: "UpdateWithdrawalQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateWithdrawalQueue is a free log subscription operation binding the contract event 0x695ac3ac73f08f2002284ffe563cefe798ee2878a5e04219522e2e99eb89d168.
//
// Solidity: event UpdateWithdrawalQueue(address[20] queue)
func (_YearnVault *YearnVaultFilterer) WatchUpdateWithdrawalQueue(opts *bind.WatchOpts, sink chan<- *YearnVaultUpdateWithdrawalQueue) (event.Subscription, error) {

	logs, sub, err := _YearnVault.contract.WatchLogs(opts, "UpdateWithdrawalQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnVaultUpdateWithdrawalQueue)
				if err := _YearnVault.contract.UnpackLog(event, "UpdateWithdrawalQueue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateWithdrawalQueue is a log parse operation binding the contract event 0x695ac3ac73f08f2002284ffe563cefe798ee2878a5e04219522e2e99eb89d168.
//
// Solidity: event UpdateWithdrawalQueue(address[20] queue)
func (_YearnVault *YearnVaultFilterer) ParseUpdateWithdrawalQueue(log types.Log) (*YearnVaultUpdateWithdrawalQueue, error) {
	event := new(YearnVaultUpdateWithdrawalQueue)
	if err := _YearnVault.contract.UnpackLog(event, "UpdateWithdrawalQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

