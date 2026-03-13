package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// FraxBaseStrategyABI is the ABI used to query userVault() on Frax base strategies.
const FraxBaseStrategyABI = "[{\"inputs\":[],\"name\":\"userVault\",\"outputs\":[{\"internalType\":\"contract IConvexFrax\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FraxBaseStrategy is a minimal binding for userVault() on Frax base strategies.
type FraxBaseStrategy struct {
	contract *bind.BoundContract
}

// NewFraxBaseStrategy creates a minimal caller for fraxBaseStrategy.userVault().
func NewFraxBaseStrategy(address common.Address, backend bind.ContractBackend) (*FraxBaseStrategy, error) {
	parsed, err := abi.JSON(strings.NewReader(FraxBaseStrategyABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, parsed, backend, backend, backend)
	return &FraxBaseStrategy{contract: contract}, nil
}

// UserVault calls userVault() on a Frax base strategy.
func (f *FraxBaseStrategy) UserVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "userVault")
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
