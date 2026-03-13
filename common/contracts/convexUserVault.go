package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ConvexUserVaultABI is the ABI used to query stakingToken() on a Convex user vault.
const ConvexUserVaultABI = "[{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ConvexUserVault is a minimal binding for stakingToken() on a Convex user vault.
type ConvexUserVault struct {
	contract *bind.BoundContract
}

// NewConvexUserVault creates a minimal caller for userVault.stakingToken().
func NewConvexUserVault(address common.Address, backend bind.ContractBackend) (*ConvexUserVault, error) {
	parsed, err := abi.JSON(strings.NewReader(ConvexUserVaultABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, parsed, backend, backend, backend)
	return &ConvexUserVault{contract: contract}, nil
}

// StakingToken calls stakingToken() on a user vault.
func (c *ConvexUserVault) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "stakingToken")
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
