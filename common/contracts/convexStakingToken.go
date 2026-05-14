package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ConvexStakingTokenABI is the ABI used to query convexPoolId() on a Convex staking token.
const ConvexStakingTokenABI = "[{\"inputs\":[],\"name\":\"convexPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ConvexStakingToken is a minimal binding for convexPoolId() on a Convex staking token.
type ConvexStakingToken struct {
	contract *bind.BoundContract
}

// NewConvexStakingToken creates a minimal caller for stakingToken.convexPoolId().
func NewConvexStakingToken(address common.Address, backend bind.ContractBackend) (*ConvexStakingToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ConvexStakingTokenABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, parsed, backend, backend, backend)
	return &ConvexStakingToken{contract: contract}, nil
}

// ConvexPoolId calls convexPoolId() on a staking token.
func (c *ConvexStakingToken) ConvexPoolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "convexPoolId")
	if err != nil {
		return *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}
