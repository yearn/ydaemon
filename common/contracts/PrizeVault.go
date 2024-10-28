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
	_ = abi.ConvertType
)

// PrizeHooks is an auto generated low-level Go binding around an user-defined struct.
type PrizeHooks struct {
	UseBeforeClaimPrize bool
	UseAfterClaimPrize  bool
	Implementation      common.Address
}

// PrizeVaultMetaData contains all meta data concerning the PrizeVault contract.
var PrizeVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"contractIERC4626\",\"name\":\"yieldVault_\",\"type\":\"address\"},{\"internalType\":\"contractPrizePool\",\"name\":\"prizePool_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"claimer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"yieldFeeRecipient_\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"yieldFeePercentage_\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"yieldBuffer_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BurnZeroShares\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"CallerNotClaimer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidationPair\",\"type\":\"address\"}],\"name\":\"CallerNotLP\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"yieldFeeRecipient\",\"type\":\"address\"}],\"name\":\"CallerNotYieldFeeRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimRecipientZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimerZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositZeroAssets\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"FailedToGetAssetDecimals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LPZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidationAmountOutZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalToWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableYield\",\"type\":\"uint256\"}],\"name\":\"LiquidationExceedsAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prizeToken\",\"type\":\"address\"}],\"name\":\"LiquidationTokenInNotPrizeToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"LiquidationTokenOutNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAssets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"LossyDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"}],\"name\":\"MaxSharesExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAssets\",\"type\":\"uint256\"}],\"name\":\"MinAssetsNotReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"name\":\"MintLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PermitCallerNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrizePoolZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"returnDataSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookDataLimit\",\"type\":\"uint256\"}],\"name\":\"ReturnDataOverLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"yieldFeeBalance\",\"type\":\"uint256\"}],\"name\":\"SharesExceedsYieldFeeBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TwabControllerZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawZeroAssets\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"yieldFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxYieldFeePercentage\",\"type\":\"uint256\"}],\"name\":\"YieldFeePercentageExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"YieldVaultZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroTotalAssets\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"ClaimYieldFeeShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"ClaimerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidationPair\",\"type\":\"address\"}],\"name\":\"LiquidationPairSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"useBeforeClaimPrize\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useAfterClaimPrize\",\"type\":\"bool\"},{\"internalType\":\"contractIPrizeHooks\",\"name\":\"implementation\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structPrizeHooks\",\"name\":\"hooks\",\"type\":\"tuple\"}],\"name\":\"SetHooks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidationPair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yieldFee\",\"type\":\"uint256\"}],\"name\":\"TransferYieldOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yieldFeePercentage\",\"type\":\"uint256\"}],\"name\":\"YieldFeePercentageSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"yieldFeeRecipient\",\"type\":\"address\"}],\"name\":\"YieldFeeRecipientSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_PRECISION\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOOK_GAS\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOOK_RETURN_DATA_LIMIT\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_YIELD_FEE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableYieldBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_winner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_prizeIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"_reward\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"_rewardRecipient\",\"type\":\"address\"}],\"name\":\"claimPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"claimYieldFeeShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentYieldBuffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"depositWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getHooks\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"useBeforeClaimPrize\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useAfterClaimPrize\",\"type\":\"bool\"},{\"internalType\":\"contractIPrizeHooks\",\"name\":\"implementation\",\"type\":\"address\"}],\"internalType\":\"structPrizeHooks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidationPair\",\"type\":\"address\"}],\"name\":\"isLiquidationPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"}],\"name\":\"liquidatableBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizePool\",\"outputs\":[{\"internalType\":\"contractPrizePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minAssets\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_claimer\",\"type\":\"address\"}],\"name\":\"setClaimer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"useBeforeClaimPrize\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useAfterClaimPrize\",\"type\":\"bool\"},{\"internalType\":\"contractIPrizeHooks\",\"name\":\"implementation\",\"type\":\"address\"}],\"internalType\":\"structPrizeHooks\",\"name\":\"hooks\",\"type\":\"tuple\"}],\"name\":\"setHooks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidationPair\",\"type\":\"address\"}],\"name\":\"setLiquidationPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_yieldFeePercentage\",\"type\":\"uint32\"}],\"name\":\"setYieldFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_yieldFeeRecipient\",\"type\":\"address\"}],\"name\":\"setYieldFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"targetOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPreciseAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalYieldBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"name\":\"transferTokensOut\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twabController\",\"outputs\":[{\"internalType\":\"contractTwabController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"verifyTokensIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxShares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldBuffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldFeePercentage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldVault\",\"outputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PrizeVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use PrizeVaultMetaData.ABI instead.
var PrizeVaultABI = PrizeVaultMetaData.ABI

// PrizeVault is an auto generated Go binding around an Ethereum contract.
type PrizeVault struct {
	PrizeVaultCaller     // Read-only binding to the contract
	PrizeVaultTransactor // Write-only binding to the contract
	PrizeVaultFilterer   // Log filterer for contract events
}

// PrizeVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrizeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizeVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrizeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizeVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrizeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizeVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrizeVaultSession struct {
	Contract     *PrizeVault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrizeVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrizeVaultCallerSession struct {
	Contract *PrizeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PrizeVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrizeVaultTransactorSession struct {
	Contract     *PrizeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PrizeVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrizeVaultRaw struct {
	Contract *PrizeVault // Generic contract binding to access the raw methods on
}

// PrizeVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrizeVaultCallerRaw struct {
	Contract *PrizeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// PrizeVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrizeVaultTransactorRaw struct {
	Contract *PrizeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrizeVault creates a new instance of PrizeVault, bound to a specific deployed contract.
func NewPrizeVault(address common.Address, backend bind.ContractBackend) (*PrizeVault, error) {
	contract, err := bindPrizeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrizeVault{PrizeVaultCaller: PrizeVaultCaller{contract: contract}, PrizeVaultTransactor: PrizeVaultTransactor{contract: contract}, PrizeVaultFilterer: PrizeVaultFilterer{contract: contract}}, nil
}

// NewPrizeVaultCaller creates a new read-only instance of PrizeVault, bound to a specific deployed contract.
func NewPrizeVaultCaller(address common.Address, caller bind.ContractCaller) (*PrizeVaultCaller, error) {
	contract, err := bindPrizeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultCaller{contract: contract}, nil
}

// NewPrizeVaultTransactor creates a new write-only instance of PrizeVault, bound to a specific deployed contract.
func NewPrizeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*PrizeVaultTransactor, error) {
	contract, err := bindPrizeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultTransactor{contract: contract}, nil
}

// NewPrizeVaultFilterer creates a new log filterer instance of PrizeVault, bound to a specific deployed contract.
func NewPrizeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*PrizeVaultFilterer, error) {
	contract, err := bindPrizeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultFilterer{contract: contract}, nil
}

// bindPrizeVault binds a generic wrapper to an already deployed contract.
func bindPrizeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrizeVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrizeVault *PrizeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrizeVault.Contract.PrizeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrizeVault *PrizeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizeVault.Contract.PrizeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrizeVault *PrizeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrizeVault.Contract.PrizeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrizeVault *PrizeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrizeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrizeVault *PrizeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrizeVault *PrizeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrizeVault.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PrizeVault *PrizeVaultCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PrizeVault *PrizeVaultSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _PrizeVault.Contract.DOMAINSEPARATOR(&_PrizeVault.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PrizeVault *PrizeVaultCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _PrizeVault.Contract.DOMAINSEPARATOR(&_PrizeVault.CallOpts)
}

// FEEPRECISION is a free data retrieval call binding the contract method 0xe63a391f.
//
// Solidity: function FEE_PRECISION() view returns(uint32)
func (_PrizeVault *PrizeVaultCaller) FEEPRECISION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "FEE_PRECISION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FEEPRECISION is a free data retrieval call binding the contract method 0xe63a391f.
//
// Solidity: function FEE_PRECISION() view returns(uint32)
func (_PrizeVault *PrizeVaultSession) FEEPRECISION() (uint32, error) {
	return _PrizeVault.Contract.FEEPRECISION(&_PrizeVault.CallOpts)
}

// FEEPRECISION is a free data retrieval call binding the contract method 0xe63a391f.
//
// Solidity: function FEE_PRECISION() view returns(uint32)
func (_PrizeVault *PrizeVaultCallerSession) FEEPRECISION() (uint32, error) {
	return _PrizeVault.Contract.FEEPRECISION(&_PrizeVault.CallOpts)
}

// HOOKGAS is a free data retrieval call binding the contract method 0x4006522b.
//
// Solidity: function HOOK_GAS() view returns(uint24)
func (_PrizeVault *PrizeVaultCaller) HOOKGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "HOOK_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HOOKGAS is a free data retrieval call binding the contract method 0x4006522b.
//
// Solidity: function HOOK_GAS() view returns(uint24)
func (_PrizeVault *PrizeVaultSession) HOOKGAS() (*big.Int, error) {
	return _PrizeVault.Contract.HOOKGAS(&_PrizeVault.CallOpts)
}

// HOOKGAS is a free data retrieval call binding the contract method 0x4006522b.
//
// Solidity: function HOOK_GAS() view returns(uint24)
func (_PrizeVault *PrizeVaultCallerSession) HOOKGAS() (*big.Int, error) {
	return _PrizeVault.Contract.HOOKGAS(&_PrizeVault.CallOpts)
}

// HOOKRETURNDATALIMIT is a free data retrieval call binding the contract method 0xd30ae1b4.
//
// Solidity: function HOOK_RETURN_DATA_LIMIT() view returns(uint16)
func (_PrizeVault *PrizeVaultCaller) HOOKRETURNDATALIMIT(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "HOOK_RETURN_DATA_LIMIT")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// HOOKRETURNDATALIMIT is a free data retrieval call binding the contract method 0xd30ae1b4.
//
// Solidity: function HOOK_RETURN_DATA_LIMIT() view returns(uint16)
func (_PrizeVault *PrizeVaultSession) HOOKRETURNDATALIMIT() (uint16, error) {
	return _PrizeVault.Contract.HOOKRETURNDATALIMIT(&_PrizeVault.CallOpts)
}

// HOOKRETURNDATALIMIT is a free data retrieval call binding the contract method 0xd30ae1b4.
//
// Solidity: function HOOK_RETURN_DATA_LIMIT() view returns(uint16)
func (_PrizeVault *PrizeVaultCallerSession) HOOKRETURNDATALIMIT() (uint16, error) {
	return _PrizeVault.Contract.HOOKRETURNDATALIMIT(&_PrizeVault.CallOpts)
}

// MAXYIELDFEE is a free data retrieval call binding the contract method 0x1caa73cb.
//
// Solidity: function MAX_YIELD_FEE() view returns(uint32)
func (_PrizeVault *PrizeVaultCaller) MAXYIELDFEE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "MAX_YIELD_FEE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXYIELDFEE is a free data retrieval call binding the contract method 0x1caa73cb.
//
// Solidity: function MAX_YIELD_FEE() view returns(uint32)
func (_PrizeVault *PrizeVaultSession) MAXYIELDFEE() (uint32, error) {
	return _PrizeVault.Contract.MAXYIELDFEE(&_PrizeVault.CallOpts)
}

// MAXYIELDFEE is a free data retrieval call binding the contract method 0x1caa73cb.
//
// Solidity: function MAX_YIELD_FEE() view returns(uint32)
func (_PrizeVault *PrizeVaultCallerSession) MAXYIELDFEE() (uint32, error) {
	return _PrizeVault.Contract.MAXYIELDFEE(&_PrizeVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.Allowance(&_PrizeVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.Allowance(&_PrizeVault.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_PrizeVault *PrizeVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_PrizeVault *PrizeVaultSession) Asset() (common.Address, error) {
	return _PrizeVault.Contract.Asset(&_PrizeVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_PrizeVault *PrizeVaultCallerSession) Asset() (common.Address, error) {
	return _PrizeVault.Contract.Asset(&_PrizeVault.CallOpts)
}

// AvailableYieldBalance is a free data retrieval call binding the contract method 0x0d1e5255.
//
// Solidity: function availableYieldBalance() view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) AvailableYieldBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "availableYieldBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableYieldBalance is a free data retrieval call binding the contract method 0x0d1e5255.
//
// Solidity: function availableYieldBalance() view returns(uint256)
func (_PrizeVault *PrizeVaultSession) AvailableYieldBalance() (*big.Int, error) {
	return _PrizeVault.Contract.AvailableYieldBalance(&_PrizeVault.CallOpts)
}

// AvailableYieldBalance is a free data retrieval call binding the contract method 0x0d1e5255.
//
// Solidity: function availableYieldBalance() view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) AvailableYieldBalance() (*big.Int, error) {
	return _PrizeVault.Contract.AvailableYieldBalance(&_PrizeVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.BalanceOf(&_PrizeVault.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.BalanceOf(&_PrizeVault.CallOpts, _account)
}

// Claimer is a free data retrieval call binding the contract method 0xd379be23.
//
// Solidity: function claimer() view returns(address)
func (_PrizeVault *PrizeVaultCaller) Claimer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "claimer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Claimer is a free data retrieval call binding the contract method 0xd379be23.
//
// Solidity: function claimer() view returns(address)
func (_PrizeVault *PrizeVaultSession) Claimer() (common.Address, error) {
	return _PrizeVault.Contract.Claimer(&_PrizeVault.CallOpts)
}

// Claimer is a free data retrieval call binding the contract method 0xd379be23.
//
// Solidity: function claimer() view returns(address)
func (_PrizeVault *PrizeVaultCallerSession) Claimer() (common.Address, error) {
	return _PrizeVault.Contract.Claimer(&_PrizeVault.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) ConvertToAssets(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "convertToAssets", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.ConvertToAssets(&_PrizeVault.CallOpts, _shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.ConvertToAssets(&_PrizeVault.CallOpts, _shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) ConvertToShares(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "convertToShares", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.ConvertToShares(&_PrizeVault.CallOpts, _assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.ConvertToShares(&_PrizeVault.CallOpts, _assets)
}

// CurrentYieldBuffer is a free data retrieval call binding the contract method 0x237fd108.
//
// Solidity: function currentYieldBuffer() view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) CurrentYieldBuffer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "currentYieldBuffer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentYieldBuffer is a free data retrieval call binding the contract method 0x237fd108.
//
// Solidity: function currentYieldBuffer() view returns(uint256)
func (_PrizeVault *PrizeVaultSession) CurrentYieldBuffer() (*big.Int, error) {
	return _PrizeVault.Contract.CurrentYieldBuffer(&_PrizeVault.CallOpts)
}

// CurrentYieldBuffer is a free data retrieval call binding the contract method 0x237fd108.
//
// Solidity: function currentYieldBuffer() view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) CurrentYieldBuffer() (*big.Int, error) {
	return _PrizeVault.Contract.CurrentYieldBuffer(&_PrizeVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PrizeVault *PrizeVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PrizeVault *PrizeVaultSession) Decimals() (uint8, error) {
	return _PrizeVault.Contract.Decimals(&_PrizeVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PrizeVault *PrizeVaultCallerSession) Decimals() (uint8, error) {
	return _PrizeVault.Contract.Decimals(&_PrizeVault.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PrizeVault *PrizeVaultCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PrizeVault *PrizeVaultSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _PrizeVault.Contract.Eip712Domain(&_PrizeVault.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PrizeVault *PrizeVaultCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _PrizeVault.Contract.Eip712Domain(&_PrizeVault.CallOpts)
}

// GetHooks is a free data retrieval call binding the contract method 0xde03f408.
//
// Solidity: function getHooks(address account) view returns((bool,bool,address))
func (_PrizeVault *PrizeVaultCaller) GetHooks(opts *bind.CallOpts, account common.Address) (PrizeHooks, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "getHooks", account)

	if err != nil {
		return *new(PrizeHooks), err
	}

	out0 := *abi.ConvertType(out[0], new(PrizeHooks)).(*PrizeHooks)

	return out0, err

}

// GetHooks is a free data retrieval call binding the contract method 0xde03f408.
//
// Solidity: function getHooks(address account) view returns((bool,bool,address))
func (_PrizeVault *PrizeVaultSession) GetHooks(account common.Address) (PrizeHooks, error) {
	return _PrizeVault.Contract.GetHooks(&_PrizeVault.CallOpts, account)
}

// GetHooks is a free data retrieval call binding the contract method 0xde03f408.
//
// Solidity: function getHooks(address account) view returns((bool,bool,address))
func (_PrizeVault *PrizeVaultCallerSession) GetHooks(account common.Address) (PrizeHooks, error) {
	return _PrizeVault.Contract.GetHooks(&_PrizeVault.CallOpts, account)
}

// IsLiquidationPair is a free data retrieval call binding the contract method 0x1b571924.
//
// Solidity: function isLiquidationPair(address _tokenOut, address _liquidationPair) view returns(bool)
func (_PrizeVault *PrizeVaultCaller) IsLiquidationPair(opts *bind.CallOpts, _tokenOut common.Address, _liquidationPair common.Address) (bool, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "isLiquidationPair", _tokenOut, _liquidationPair)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidationPair is a free data retrieval call binding the contract method 0x1b571924.
//
// Solidity: function isLiquidationPair(address _tokenOut, address _liquidationPair) view returns(bool)
func (_PrizeVault *PrizeVaultSession) IsLiquidationPair(_tokenOut common.Address, _liquidationPair common.Address) (bool, error) {
	return _PrizeVault.Contract.IsLiquidationPair(&_PrizeVault.CallOpts, _tokenOut, _liquidationPair)
}

// IsLiquidationPair is a free data retrieval call binding the contract method 0x1b571924.
//
// Solidity: function isLiquidationPair(address _tokenOut, address _liquidationPair) view returns(bool)
func (_PrizeVault *PrizeVaultCallerSession) IsLiquidationPair(_tokenOut common.Address, _liquidationPair common.Address) (bool, error) {
	return _PrizeVault.Contract.IsLiquidationPair(&_PrizeVault.CallOpts, _tokenOut, _liquidationPair)
}

// LiquidatableBalanceOf is a free data retrieval call binding the contract method 0xb0fcf626.
//
// Solidity: function liquidatableBalanceOf(address _tokenOut) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) LiquidatableBalanceOf(opts *bind.CallOpts, _tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "liquidatableBalanceOf", _tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidatableBalanceOf is a free data retrieval call binding the contract method 0xb0fcf626.
//
// Solidity: function liquidatableBalanceOf(address _tokenOut) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) LiquidatableBalanceOf(_tokenOut common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.LiquidatableBalanceOf(&_PrizeVault.CallOpts, _tokenOut)
}

// LiquidatableBalanceOf is a free data retrieval call binding the contract method 0xb0fcf626.
//
// Solidity: function liquidatableBalanceOf(address _tokenOut) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) LiquidatableBalanceOf(_tokenOut common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.LiquidatableBalanceOf(&_PrizeVault.CallOpts, _tokenOut)
}

// LiquidationPair is a free data retrieval call binding the contract method 0x649f2394.
//
// Solidity: function liquidationPair() view returns(address)
func (_PrizeVault *PrizeVaultCaller) LiquidationPair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "liquidationPair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidationPair is a free data retrieval call binding the contract method 0x649f2394.
//
// Solidity: function liquidationPair() view returns(address)
func (_PrizeVault *PrizeVaultSession) LiquidationPair() (common.Address, error) {
	return _PrizeVault.Contract.LiquidationPair(&_PrizeVault.CallOpts)
}

// LiquidationPair is a free data retrieval call binding the contract method 0x649f2394.
//
// Solidity: function liquidationPair() view returns(address)
func (_PrizeVault *PrizeVaultCallerSession) LiquidationPair() (common.Address, error) {
	return _PrizeVault.Contract.LiquidationPair(&_PrizeVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.MaxDeposit(&_PrizeVault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.MaxDeposit(&_PrizeVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) MaxMint(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "maxMint", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) MaxMint(_owner common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.MaxMint(&_PrizeVault.CallOpts, _owner)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address _owner) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) MaxMint(_owner common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.MaxMint(&_PrizeVault.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) MaxRedeem(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "maxRedeem", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.MaxRedeem(&_PrizeVault.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.MaxRedeem(&_PrizeVault.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) MaxWithdraw(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "maxWithdraw", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.MaxWithdraw(&_PrizeVault.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.MaxWithdraw(&_PrizeVault.CallOpts, _owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PrizeVault *PrizeVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PrizeVault *PrizeVaultSession) Name() (string, error) {
	return _PrizeVault.Contract.Name(&_PrizeVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PrizeVault *PrizeVaultCallerSession) Name() (string, error) {
	return _PrizeVault.Contract.Name(&_PrizeVault.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) Nonces(owner common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.Nonces(&_PrizeVault.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _PrizeVault.Contract.Nonces(&_PrizeVault.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrizeVault *PrizeVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrizeVault *PrizeVaultSession) Owner() (common.Address, error) {
	return _PrizeVault.Contract.Owner(&_PrizeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrizeVault *PrizeVaultCallerSession) Owner() (common.Address, error) {
	return _PrizeVault.Contract.Owner(&_PrizeVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PrizeVault *PrizeVaultCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PrizeVault *PrizeVaultSession) PendingOwner() (common.Address, error) {
	return _PrizeVault.Contract.PendingOwner(&_PrizeVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PrizeVault *PrizeVaultCallerSession) PendingOwner() (common.Address, error) {
	return _PrizeVault.Contract.PendingOwner(&_PrizeVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) pure returns(uint256)
func (_PrizeVault *PrizeVaultCaller) PreviewDeposit(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "previewDeposit", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) pure returns(uint256)
func (_PrizeVault *PrizeVaultSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.PreviewDeposit(&_PrizeVault.CallOpts, _assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) pure returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.PreviewDeposit(&_PrizeVault.CallOpts, _assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) pure returns(uint256)
func (_PrizeVault *PrizeVaultCaller) PreviewMint(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "previewMint", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) pure returns(uint256)
func (_PrizeVault *PrizeVaultSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.PreviewMint(&_PrizeVault.CallOpts, _shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) pure returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.PreviewMint(&_PrizeVault.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) PreviewRedeem(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "previewRedeem", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) PreviewRedeem(_shares *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.PreviewRedeem(&_PrizeVault.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) PreviewRedeem(_shares *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.PreviewRedeem(&_PrizeVault.CallOpts, _shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) PreviewWithdraw(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "previewWithdraw", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_PrizeVault *PrizeVaultSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.PreviewWithdraw(&_PrizeVault.CallOpts, _assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _PrizeVault.Contract.PreviewWithdraw(&_PrizeVault.CallOpts, _assets)
}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(address)
func (_PrizeVault *PrizeVaultCaller) PrizePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "prizePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(address)
func (_PrizeVault *PrizeVaultSession) PrizePool() (common.Address, error) {
	return _PrizeVault.Contract.PrizePool(&_PrizeVault.CallOpts)
}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(address)
func (_PrizeVault *PrizeVaultCallerSession) PrizePool() (common.Address, error) {
	return _PrizeVault.Contract.PrizePool(&_PrizeVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PrizeVault *PrizeVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PrizeVault *PrizeVaultSession) Symbol() (string, error) {
	return _PrizeVault.Contract.Symbol(&_PrizeVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PrizeVault *PrizeVaultCallerSession) Symbol() (string, error) {
	return _PrizeVault.Contract.Symbol(&_PrizeVault.CallOpts)
}

// TargetOf is a free data retrieval call binding the contract method 0x700f04ef.
//
// Solidity: function targetOf(address ) view returns(address)
func (_PrizeVault *PrizeVaultCaller) TargetOf(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "targetOf", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetOf is a free data retrieval call binding the contract method 0x700f04ef.
//
// Solidity: function targetOf(address ) view returns(address)
func (_PrizeVault *PrizeVaultSession) TargetOf(arg0 common.Address) (common.Address, error) {
	return _PrizeVault.Contract.TargetOf(&_PrizeVault.CallOpts, arg0)
}

// TargetOf is a free data retrieval call binding the contract method 0x700f04ef.
//
// Solidity: function targetOf(address ) view returns(address)
func (_PrizeVault *PrizeVaultCallerSession) TargetOf(arg0 common.Address) (common.Address, error) {
	return _PrizeVault.Contract.TargetOf(&_PrizeVault.CallOpts, arg0)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_PrizeVault *PrizeVaultSession) TotalAssets() (*big.Int, error) {
	return _PrizeVault.Contract.TotalAssets(&_PrizeVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _PrizeVault.Contract.TotalAssets(&_PrizeVault.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_PrizeVault *PrizeVaultSession) TotalDebt() (*big.Int, error) {
	return _PrizeVault.Contract.TotalDebt(&_PrizeVault.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) TotalDebt() (*big.Int, error) {
	return _PrizeVault.Contract.TotalDebt(&_PrizeVault.CallOpts)
}

// TotalPreciseAssets is a free data retrieval call binding the contract method 0x4244aa9b.
//
// Solidity: function totalPreciseAssets() view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) TotalPreciseAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "totalPreciseAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPreciseAssets is a free data retrieval call binding the contract method 0x4244aa9b.
//
// Solidity: function totalPreciseAssets() view returns(uint256)
func (_PrizeVault *PrizeVaultSession) TotalPreciseAssets() (*big.Int, error) {
	return _PrizeVault.Contract.TotalPreciseAssets(&_PrizeVault.CallOpts)
}

// TotalPreciseAssets is a free data retrieval call binding the contract method 0x4244aa9b.
//
// Solidity: function totalPreciseAssets() view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) TotalPreciseAssets() (*big.Int, error) {
	return _PrizeVault.Contract.TotalPreciseAssets(&_PrizeVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PrizeVault *PrizeVaultSession) TotalSupply() (*big.Int, error) {
	return _PrizeVault.Contract.TotalSupply(&_PrizeVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _PrizeVault.Contract.TotalSupply(&_PrizeVault.CallOpts)
}

// TotalYieldBalance is a free data retrieval call binding the contract method 0xd4122abf.
//
// Solidity: function totalYieldBalance() view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) TotalYieldBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "totalYieldBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalYieldBalance is a free data retrieval call binding the contract method 0xd4122abf.
//
// Solidity: function totalYieldBalance() view returns(uint256)
func (_PrizeVault *PrizeVaultSession) TotalYieldBalance() (*big.Int, error) {
	return _PrizeVault.Contract.TotalYieldBalance(&_PrizeVault.CallOpts)
}

// TotalYieldBalance is a free data retrieval call binding the contract method 0xd4122abf.
//
// Solidity: function totalYieldBalance() view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) TotalYieldBalance() (*big.Int, error) {
	return _PrizeVault.Contract.TotalYieldBalance(&_PrizeVault.CallOpts)
}

// TwabController is a free data retrieval call binding the contract method 0xb0812d7b.
//
// Solidity: function twabController() view returns(address)
func (_PrizeVault *PrizeVaultCaller) TwabController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "twabController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TwabController is a free data retrieval call binding the contract method 0xb0812d7b.
//
// Solidity: function twabController() view returns(address)
func (_PrizeVault *PrizeVaultSession) TwabController() (common.Address, error) {
	return _PrizeVault.Contract.TwabController(&_PrizeVault.CallOpts)
}

// TwabController is a free data retrieval call binding the contract method 0xb0812d7b.
//
// Solidity: function twabController() view returns(address)
func (_PrizeVault *PrizeVaultCallerSession) TwabController() (common.Address, error) {
	return _PrizeVault.Contract.TwabController(&_PrizeVault.CallOpts)
}

// YieldBuffer is a free data retrieval call binding the contract method 0x13dc6c5d.
//
// Solidity: function yieldBuffer() view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) YieldBuffer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "yieldBuffer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YieldBuffer is a free data retrieval call binding the contract method 0x13dc6c5d.
//
// Solidity: function yieldBuffer() view returns(uint256)
func (_PrizeVault *PrizeVaultSession) YieldBuffer() (*big.Int, error) {
	return _PrizeVault.Contract.YieldBuffer(&_PrizeVault.CallOpts)
}

// YieldBuffer is a free data retrieval call binding the contract method 0x13dc6c5d.
//
// Solidity: function yieldBuffer() view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) YieldBuffer() (*big.Int, error) {
	return _PrizeVault.Contract.YieldBuffer(&_PrizeVault.CallOpts)
}

// YieldFeeBalance is a free data retrieval call binding the contract method 0x9a1adf9a.
//
// Solidity: function yieldFeeBalance() view returns(uint256)
func (_PrizeVault *PrizeVaultCaller) YieldFeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "yieldFeeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YieldFeeBalance is a free data retrieval call binding the contract method 0x9a1adf9a.
//
// Solidity: function yieldFeeBalance() view returns(uint256)
func (_PrizeVault *PrizeVaultSession) YieldFeeBalance() (*big.Int, error) {
	return _PrizeVault.Contract.YieldFeeBalance(&_PrizeVault.CallOpts)
}

// YieldFeeBalance is a free data retrieval call binding the contract method 0x9a1adf9a.
//
// Solidity: function yieldFeeBalance() view returns(uint256)
func (_PrizeVault *PrizeVaultCallerSession) YieldFeeBalance() (*big.Int, error) {
	return _PrizeVault.Contract.YieldFeeBalance(&_PrizeVault.CallOpts)
}

// YieldFeePercentage is a free data retrieval call binding the contract method 0xe16777c1.
//
// Solidity: function yieldFeePercentage() view returns(uint32)
func (_PrizeVault *PrizeVaultCaller) YieldFeePercentage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "yieldFeePercentage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// YieldFeePercentage is a free data retrieval call binding the contract method 0xe16777c1.
//
// Solidity: function yieldFeePercentage() view returns(uint32)
func (_PrizeVault *PrizeVaultSession) YieldFeePercentage() (uint32, error) {
	return _PrizeVault.Contract.YieldFeePercentage(&_PrizeVault.CallOpts)
}

// YieldFeePercentage is a free data retrieval call binding the contract method 0xe16777c1.
//
// Solidity: function yieldFeePercentage() view returns(uint32)
func (_PrizeVault *PrizeVaultCallerSession) YieldFeePercentage() (uint32, error) {
	return _PrizeVault.Contract.YieldFeePercentage(&_PrizeVault.CallOpts)
}

// YieldFeeRecipient is a free data retrieval call binding the contract method 0xedb8eb80.
//
// Solidity: function yieldFeeRecipient() view returns(address)
func (_PrizeVault *PrizeVaultCaller) YieldFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "yieldFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldFeeRecipient is a free data retrieval call binding the contract method 0xedb8eb80.
//
// Solidity: function yieldFeeRecipient() view returns(address)
func (_PrizeVault *PrizeVaultSession) YieldFeeRecipient() (common.Address, error) {
	return _PrizeVault.Contract.YieldFeeRecipient(&_PrizeVault.CallOpts)
}

// YieldFeeRecipient is a free data retrieval call binding the contract method 0xedb8eb80.
//
// Solidity: function yieldFeeRecipient() view returns(address)
func (_PrizeVault *PrizeVaultCallerSession) YieldFeeRecipient() (common.Address, error) {
	return _PrizeVault.Contract.YieldFeeRecipient(&_PrizeVault.CallOpts)
}

// YieldVault is a free data retrieval call binding the contract method 0xa7f8a5e2.
//
// Solidity: function yieldVault() view returns(address)
func (_PrizeVault *PrizeVaultCaller) YieldVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizeVault.contract.Call(opts, &out, "yieldVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldVault is a free data retrieval call binding the contract method 0xa7f8a5e2.
//
// Solidity: function yieldVault() view returns(address)
func (_PrizeVault *PrizeVaultSession) YieldVault() (common.Address, error) {
	return _PrizeVault.Contract.YieldVault(&_PrizeVault.CallOpts)
}

// YieldVault is a free data retrieval call binding the contract method 0xa7f8a5e2.
//
// Solidity: function yieldVault() view returns(address)
func (_PrizeVault *PrizeVaultCallerSession) YieldVault() (common.Address, error) {
	return _PrizeVault.Contract.YieldVault(&_PrizeVault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PrizeVault *PrizeVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PrizeVault *PrizeVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.Approve(&_PrizeVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PrizeVault *PrizeVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.Approve(&_PrizeVault.TransactOpts, spender, amount)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PrizeVault *PrizeVaultTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PrizeVault *PrizeVaultSession) ClaimOwnership() (*types.Transaction, error) {
	return _PrizeVault.Contract.ClaimOwnership(&_PrizeVault.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PrizeVault *PrizeVaultTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _PrizeVault.Contract.ClaimOwnership(&_PrizeVault.TransactOpts)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x2895cace.
//
// Solidity: function claimPrize(address _winner, uint8 _tier, uint32 _prizeIndex, uint96 _reward, address _rewardRecipient) returns(uint256)
func (_PrizeVault *PrizeVaultTransactor) ClaimPrize(opts *bind.TransactOpts, _winner common.Address, _tier uint8, _prizeIndex uint32, _reward *big.Int, _rewardRecipient common.Address) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "claimPrize", _winner, _tier, _prizeIndex, _reward, _rewardRecipient)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x2895cace.
//
// Solidity: function claimPrize(address _winner, uint8 _tier, uint32 _prizeIndex, uint96 _reward, address _rewardRecipient) returns(uint256)
func (_PrizeVault *PrizeVaultSession) ClaimPrize(_winner common.Address, _tier uint8, _prizeIndex uint32, _reward *big.Int, _rewardRecipient common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.ClaimPrize(&_PrizeVault.TransactOpts, _winner, _tier, _prizeIndex, _reward, _rewardRecipient)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x2895cace.
//
// Solidity: function claimPrize(address _winner, uint8 _tier, uint32 _prizeIndex, uint96 _reward, address _rewardRecipient) returns(uint256)
func (_PrizeVault *PrizeVaultTransactorSession) ClaimPrize(_winner common.Address, _tier uint8, _prizeIndex uint32, _reward *big.Int, _rewardRecipient common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.ClaimPrize(&_PrizeVault.TransactOpts, _winner, _tier, _prizeIndex, _reward, _rewardRecipient)
}

// ClaimYieldFeeShares is a paid mutator transaction binding the contract method 0x353d5a18.
//
// Solidity: function claimYieldFeeShares(uint256 _shares) returns()
func (_PrizeVault *PrizeVaultTransactor) ClaimYieldFeeShares(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "claimYieldFeeShares", _shares)
}

// ClaimYieldFeeShares is a paid mutator transaction binding the contract method 0x353d5a18.
//
// Solidity: function claimYieldFeeShares(uint256 _shares) returns()
func (_PrizeVault *PrizeVaultSession) ClaimYieldFeeShares(_shares *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.ClaimYieldFeeShares(&_PrizeVault.TransactOpts, _shares)
}

// ClaimYieldFeeShares is a paid mutator transaction binding the contract method 0x353d5a18.
//
// Solidity: function claimYieldFeeShares(uint256 _shares) returns()
func (_PrizeVault *PrizeVaultTransactorSession) ClaimYieldFeeShares(_shares *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.ClaimYieldFeeShares(&_PrizeVault.TransactOpts, _shares)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PrizeVault *PrizeVaultTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PrizeVault *PrizeVaultSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.DecreaseAllowance(&_PrizeVault.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PrizeVault *PrizeVaultTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.DecreaseAllowance(&_PrizeVault.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_PrizeVault *PrizeVaultTransactor) Deposit(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "deposit", _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_PrizeVault *PrizeVaultSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.Deposit(&_PrizeVault.TransactOpts, _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_PrizeVault *PrizeVaultTransactorSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.Deposit(&_PrizeVault.TransactOpts, _assets, _receiver)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x50921b23.
//
// Solidity: function depositWithPermit(uint256 _assets, address _owner, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_PrizeVault *PrizeVaultTransactor) DepositWithPermit(opts *bind.TransactOpts, _assets *big.Int, _owner common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "depositWithPermit", _assets, _owner, _deadline, _v, _r, _s)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x50921b23.
//
// Solidity: function depositWithPermit(uint256 _assets, address _owner, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_PrizeVault *PrizeVaultSession) DepositWithPermit(_assets *big.Int, _owner common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PrizeVault.Contract.DepositWithPermit(&_PrizeVault.TransactOpts, _assets, _owner, _deadline, _v, _r, _s)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x50921b23.
//
// Solidity: function depositWithPermit(uint256 _assets, address _owner, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_PrizeVault *PrizeVaultTransactorSession) DepositWithPermit(_assets *big.Int, _owner common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PrizeVault.Contract.DepositWithPermit(&_PrizeVault.TransactOpts, _assets, _owner, _deadline, _v, _r, _s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PrizeVault *PrizeVaultTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PrizeVault *PrizeVaultSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.IncreaseAllowance(&_PrizeVault.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PrizeVault *PrizeVaultTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.IncreaseAllowance(&_PrizeVault.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_PrizeVault *PrizeVaultTransactor) Mint(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "mint", _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_PrizeVault *PrizeVaultSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.Mint(&_PrizeVault.TransactOpts, _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_PrizeVault *PrizeVaultTransactorSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.Mint(&_PrizeVault.TransactOpts, _shares, _receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PrizeVault *PrizeVaultTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PrizeVault *PrizeVaultSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PrizeVault.Contract.Permit(&_PrizeVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PrizeVault *PrizeVaultTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PrizeVault.Contract.Permit(&_PrizeVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner, uint256 _minAssets) returns(uint256)
func (_PrizeVault *PrizeVaultTransactor) Redeem(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address, _owner common.Address, _minAssets *big.Int) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "redeem", _shares, _receiver, _owner, _minAssets)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner, uint256 _minAssets) returns(uint256)
func (_PrizeVault *PrizeVaultSession) Redeem(_shares *big.Int, _receiver common.Address, _owner common.Address, _minAssets *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.Redeem(&_PrizeVault.TransactOpts, _shares, _receiver, _owner, _minAssets)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner, uint256 _minAssets) returns(uint256)
func (_PrizeVault *PrizeVaultTransactorSession) Redeem(_shares *big.Int, _receiver common.Address, _owner common.Address, _minAssets *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.Redeem(&_PrizeVault.TransactOpts, _shares, _receiver, _owner, _minAssets)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256)
func (_PrizeVault *PrizeVaultTransactor) Redeem0(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "redeem0", _shares, _receiver, _owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256)
func (_PrizeVault *PrizeVaultSession) Redeem0(_shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.Redeem0(&_PrizeVault.TransactOpts, _shares, _receiver, _owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256)
func (_PrizeVault *PrizeVaultTransactorSession) Redeem0(_shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.Redeem0(&_PrizeVault.TransactOpts, _shares, _receiver, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrizeVault *PrizeVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrizeVault *PrizeVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrizeVault.Contract.RenounceOwnership(&_PrizeVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrizeVault *PrizeVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrizeVault.Contract.RenounceOwnership(&_PrizeVault.TransactOpts)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xcdfb5832.
//
// Solidity: function setClaimer(address _claimer) returns()
func (_PrizeVault *PrizeVaultTransactor) SetClaimer(opts *bind.TransactOpts, _claimer common.Address) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "setClaimer", _claimer)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xcdfb5832.
//
// Solidity: function setClaimer(address _claimer) returns()
func (_PrizeVault *PrizeVaultSession) SetClaimer(_claimer common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.SetClaimer(&_PrizeVault.TransactOpts, _claimer)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xcdfb5832.
//
// Solidity: function setClaimer(address _claimer) returns()
func (_PrizeVault *PrizeVaultTransactorSession) SetClaimer(_claimer common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.SetClaimer(&_PrizeVault.TransactOpts, _claimer)
}

// SetHooks is a paid mutator transaction binding the contract method 0xc78c72c1.
//
// Solidity: function setHooks((bool,bool,address) hooks) returns()
func (_PrizeVault *PrizeVaultTransactor) SetHooks(opts *bind.TransactOpts, hooks PrizeHooks) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "setHooks", hooks)
}

// SetHooks is a paid mutator transaction binding the contract method 0xc78c72c1.
//
// Solidity: function setHooks((bool,bool,address) hooks) returns()
func (_PrizeVault *PrizeVaultSession) SetHooks(hooks PrizeHooks) (*types.Transaction, error) {
	return _PrizeVault.Contract.SetHooks(&_PrizeVault.TransactOpts, hooks)
}

// SetHooks is a paid mutator transaction binding the contract method 0xc78c72c1.
//
// Solidity: function setHooks((bool,bool,address) hooks) returns()
func (_PrizeVault *PrizeVaultTransactorSession) SetHooks(hooks PrizeHooks) (*types.Transaction, error) {
	return _PrizeVault.Contract.SetHooks(&_PrizeVault.TransactOpts, hooks)
}

// SetLiquidationPair is a paid mutator transaction binding the contract method 0x25fa66e0.
//
// Solidity: function setLiquidationPair(address _liquidationPair) returns()
func (_PrizeVault *PrizeVaultTransactor) SetLiquidationPair(opts *bind.TransactOpts, _liquidationPair common.Address) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "setLiquidationPair", _liquidationPair)
}

// SetLiquidationPair is a paid mutator transaction binding the contract method 0x25fa66e0.
//
// Solidity: function setLiquidationPair(address _liquidationPair) returns()
func (_PrizeVault *PrizeVaultSession) SetLiquidationPair(_liquidationPair common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.SetLiquidationPair(&_PrizeVault.TransactOpts, _liquidationPair)
}

// SetLiquidationPair is a paid mutator transaction binding the contract method 0x25fa66e0.
//
// Solidity: function setLiquidationPair(address _liquidationPair) returns()
func (_PrizeVault *PrizeVaultTransactorSession) SetLiquidationPair(_liquidationPair common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.SetLiquidationPair(&_PrizeVault.TransactOpts, _liquidationPair)
}

// SetYieldFeePercentage is a paid mutator transaction binding the contract method 0xe4e243ac.
//
// Solidity: function setYieldFeePercentage(uint32 _yieldFeePercentage) returns()
func (_PrizeVault *PrizeVaultTransactor) SetYieldFeePercentage(opts *bind.TransactOpts, _yieldFeePercentage uint32) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "setYieldFeePercentage", _yieldFeePercentage)
}

// SetYieldFeePercentage is a paid mutator transaction binding the contract method 0xe4e243ac.
//
// Solidity: function setYieldFeePercentage(uint32 _yieldFeePercentage) returns()
func (_PrizeVault *PrizeVaultSession) SetYieldFeePercentage(_yieldFeePercentage uint32) (*types.Transaction, error) {
	return _PrizeVault.Contract.SetYieldFeePercentage(&_PrizeVault.TransactOpts, _yieldFeePercentage)
}

// SetYieldFeePercentage is a paid mutator transaction binding the contract method 0xe4e243ac.
//
// Solidity: function setYieldFeePercentage(uint32 _yieldFeePercentage) returns()
func (_PrizeVault *PrizeVaultTransactorSession) SetYieldFeePercentage(_yieldFeePercentage uint32) (*types.Transaction, error) {
	return _PrizeVault.Contract.SetYieldFeePercentage(&_PrizeVault.TransactOpts, _yieldFeePercentage)
}

// SetYieldFeeRecipient is a paid mutator transaction binding the contract method 0x63003b16.
//
// Solidity: function setYieldFeeRecipient(address _yieldFeeRecipient) returns()
func (_PrizeVault *PrizeVaultTransactor) SetYieldFeeRecipient(opts *bind.TransactOpts, _yieldFeeRecipient common.Address) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "setYieldFeeRecipient", _yieldFeeRecipient)
}

// SetYieldFeeRecipient is a paid mutator transaction binding the contract method 0x63003b16.
//
// Solidity: function setYieldFeeRecipient(address _yieldFeeRecipient) returns()
func (_PrizeVault *PrizeVaultSession) SetYieldFeeRecipient(_yieldFeeRecipient common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.SetYieldFeeRecipient(&_PrizeVault.TransactOpts, _yieldFeeRecipient)
}

// SetYieldFeeRecipient is a paid mutator transaction binding the contract method 0x63003b16.
//
// Solidity: function setYieldFeeRecipient(address _yieldFeeRecipient) returns()
func (_PrizeVault *PrizeVaultTransactorSession) SetYieldFeeRecipient(_yieldFeeRecipient common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.SetYieldFeeRecipient(&_PrizeVault.TransactOpts, _yieldFeeRecipient)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PrizeVault *PrizeVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PrizeVault *PrizeVaultSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.Transfer(&_PrizeVault.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PrizeVault *PrizeVaultTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.Transfer(&_PrizeVault.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PrizeVault *PrizeVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PrizeVault *PrizeVaultSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.TransferFrom(&_PrizeVault.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PrizeVault *PrizeVaultTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.TransferFrom(&_PrizeVault.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PrizeVault *PrizeVaultTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PrizeVault *PrizeVaultSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.TransferOwnership(&_PrizeVault.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PrizeVault *PrizeVaultTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.TransferOwnership(&_PrizeVault.TransactOpts, _newOwner)
}

// TransferTokensOut is a paid mutator transaction binding the contract method 0x7cc99d3f.
//
// Solidity: function transferTokensOut(address , address _receiver, address _tokenOut, uint256 _amountOut) returns(bytes)
func (_PrizeVault *PrizeVaultTransactor) TransferTokensOut(opts *bind.TransactOpts, arg0 common.Address, _receiver common.Address, _tokenOut common.Address, _amountOut *big.Int) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "transferTokensOut", arg0, _receiver, _tokenOut, _amountOut)
}

// TransferTokensOut is a paid mutator transaction binding the contract method 0x7cc99d3f.
//
// Solidity: function transferTokensOut(address , address _receiver, address _tokenOut, uint256 _amountOut) returns(bytes)
func (_PrizeVault *PrizeVaultSession) TransferTokensOut(arg0 common.Address, _receiver common.Address, _tokenOut common.Address, _amountOut *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.TransferTokensOut(&_PrizeVault.TransactOpts, arg0, _receiver, _tokenOut, _amountOut)
}

// TransferTokensOut is a paid mutator transaction binding the contract method 0x7cc99d3f.
//
// Solidity: function transferTokensOut(address , address _receiver, address _tokenOut, uint256 _amountOut) returns(bytes)
func (_PrizeVault *PrizeVaultTransactorSession) TransferTokensOut(arg0 common.Address, _receiver common.Address, _tokenOut common.Address, _amountOut *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.TransferTokensOut(&_PrizeVault.TransactOpts, arg0, _receiver, _tokenOut, _amountOut)
}

// VerifyTokensIn is a paid mutator transaction binding the contract method 0xc8576e61.
//
// Solidity: function verifyTokensIn(address _tokenIn, uint256 _amountIn, bytes ) returns()
func (_PrizeVault *PrizeVaultTransactor) VerifyTokensIn(opts *bind.TransactOpts, _tokenIn common.Address, _amountIn *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "verifyTokensIn", _tokenIn, _amountIn, arg2)
}

// VerifyTokensIn is a paid mutator transaction binding the contract method 0xc8576e61.
//
// Solidity: function verifyTokensIn(address _tokenIn, uint256 _amountIn, bytes ) returns()
func (_PrizeVault *PrizeVaultSession) VerifyTokensIn(_tokenIn common.Address, _amountIn *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _PrizeVault.Contract.VerifyTokensIn(&_PrizeVault.TransactOpts, _tokenIn, _amountIn, arg2)
}

// VerifyTokensIn is a paid mutator transaction binding the contract method 0xc8576e61.
//
// Solidity: function verifyTokensIn(address _tokenIn, uint256 _amountIn, bytes ) returns()
func (_PrizeVault *PrizeVaultTransactorSession) VerifyTokensIn(_tokenIn common.Address, _amountIn *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _PrizeVault.Contract.VerifyTokensIn(&_PrizeVault.TransactOpts, _tokenIn, _amountIn, arg2)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner, uint256 _maxShares) returns(uint256)
func (_PrizeVault *PrizeVaultTransactor) Withdraw(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address, _maxShares *big.Int) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "withdraw", _assets, _receiver, _owner, _maxShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner, uint256 _maxShares) returns(uint256)
func (_PrizeVault *PrizeVaultSession) Withdraw(_assets *big.Int, _receiver common.Address, _owner common.Address, _maxShares *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.Withdraw(&_PrizeVault.TransactOpts, _assets, _receiver, _owner, _maxShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner, uint256 _maxShares) returns(uint256)
func (_PrizeVault *PrizeVaultTransactorSession) Withdraw(_assets *big.Int, _receiver common.Address, _owner common.Address, _maxShares *big.Int) (*types.Transaction, error) {
	return _PrizeVault.Contract.Withdraw(&_PrizeVault.TransactOpts, _assets, _receiver, _owner, _maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_PrizeVault *PrizeVaultTransactor) Withdraw0(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _PrizeVault.contract.Transact(opts, "withdraw0", _assets, _receiver, _owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_PrizeVault *PrizeVaultSession) Withdraw0(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.Withdraw0(&_PrizeVault.TransactOpts, _assets, _receiver, _owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_PrizeVault *PrizeVaultTransactorSession) Withdraw0(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _PrizeVault.Contract.Withdraw0(&_PrizeVault.TransactOpts, _assets, _receiver, _owner)
}

// PrizeVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PrizeVault contract.
type PrizeVaultApprovalIterator struct {
	Event *PrizeVaultApproval // Event containing the contract specifics and raw log

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
func (it *PrizeVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultApproval)
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
		it.Event = new(PrizeVaultApproval)
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
func (it *PrizeVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultApproval represents a Approval event raised by the PrizeVault contract.
type PrizeVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PrizeVault *PrizeVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PrizeVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultApprovalIterator{contract: _PrizeVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PrizeVault *PrizeVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PrizeVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultApproval)
				if err := _PrizeVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PrizeVault *PrizeVaultFilterer) ParseApproval(log types.Log) (*PrizeVaultApproval, error) {
	event := new(PrizeVaultApproval)
	if err := _PrizeVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultClaimYieldFeeSharesIterator is returned from FilterClaimYieldFeeShares and is used to iterate over the raw logs and unpacked data for ClaimYieldFeeShares events raised by the PrizeVault contract.
type PrizeVaultClaimYieldFeeSharesIterator struct {
	Event *PrizeVaultClaimYieldFeeShares // Event containing the contract specifics and raw log

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
func (it *PrizeVaultClaimYieldFeeSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultClaimYieldFeeShares)
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
		it.Event = new(PrizeVaultClaimYieldFeeShares)
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
func (it *PrizeVaultClaimYieldFeeSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultClaimYieldFeeSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultClaimYieldFeeShares represents a ClaimYieldFeeShares event raised by the PrizeVault contract.
type PrizeVaultClaimYieldFeeShares struct {
	Recipient common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimYieldFeeShares is a free log retrieval operation binding the contract event 0x9b51aebdd88b3f321397bdf9848415ce15dedd1783c775b93b8036c6b0af0f56.
//
// Solidity: event ClaimYieldFeeShares(address indexed recipient, uint256 shares)
func (_PrizeVault *PrizeVaultFilterer) FilterClaimYieldFeeShares(opts *bind.FilterOpts, recipient []common.Address) (*PrizeVaultClaimYieldFeeSharesIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "ClaimYieldFeeShares", recipientRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultClaimYieldFeeSharesIterator{contract: _PrizeVault.contract, event: "ClaimYieldFeeShares", logs: logs, sub: sub}, nil
}

// WatchClaimYieldFeeShares is a free log subscription operation binding the contract event 0x9b51aebdd88b3f321397bdf9848415ce15dedd1783c775b93b8036c6b0af0f56.
//
// Solidity: event ClaimYieldFeeShares(address indexed recipient, uint256 shares)
func (_PrizeVault *PrizeVaultFilterer) WatchClaimYieldFeeShares(opts *bind.WatchOpts, sink chan<- *PrizeVaultClaimYieldFeeShares, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "ClaimYieldFeeShares", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultClaimYieldFeeShares)
				if err := _PrizeVault.contract.UnpackLog(event, "ClaimYieldFeeShares", log); err != nil {
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

// ParseClaimYieldFeeShares is a log parse operation binding the contract event 0x9b51aebdd88b3f321397bdf9848415ce15dedd1783c775b93b8036c6b0af0f56.
//
// Solidity: event ClaimYieldFeeShares(address indexed recipient, uint256 shares)
func (_PrizeVault *PrizeVaultFilterer) ParseClaimYieldFeeShares(log types.Log) (*PrizeVaultClaimYieldFeeShares, error) {
	event := new(PrizeVaultClaimYieldFeeShares)
	if err := _PrizeVault.contract.UnpackLog(event, "ClaimYieldFeeShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultClaimerSetIterator is returned from FilterClaimerSet and is used to iterate over the raw logs and unpacked data for ClaimerSet events raised by the PrizeVault contract.
type PrizeVaultClaimerSetIterator struct {
	Event *PrizeVaultClaimerSet // Event containing the contract specifics and raw log

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
func (it *PrizeVaultClaimerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultClaimerSet)
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
		it.Event = new(PrizeVaultClaimerSet)
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
func (it *PrizeVaultClaimerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultClaimerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultClaimerSet represents a ClaimerSet event raised by the PrizeVault contract.
type PrizeVaultClaimerSet struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimerSet is a free log retrieval operation binding the contract event 0xa04eca05f1c4f5674beaad80f345070124aa2192abced410b9b227c0c04e755a.
//
// Solidity: event ClaimerSet(address indexed claimer)
func (_PrizeVault *PrizeVaultFilterer) FilterClaimerSet(opts *bind.FilterOpts, claimer []common.Address) (*PrizeVaultClaimerSetIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "ClaimerSet", claimerRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultClaimerSetIterator{contract: _PrizeVault.contract, event: "ClaimerSet", logs: logs, sub: sub}, nil
}

// WatchClaimerSet is a free log subscription operation binding the contract event 0xa04eca05f1c4f5674beaad80f345070124aa2192abced410b9b227c0c04e755a.
//
// Solidity: event ClaimerSet(address indexed claimer)
func (_PrizeVault *PrizeVaultFilterer) WatchClaimerSet(opts *bind.WatchOpts, sink chan<- *PrizeVaultClaimerSet, claimer []common.Address) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "ClaimerSet", claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultClaimerSet)
				if err := _PrizeVault.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
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

// ParseClaimerSet is a log parse operation binding the contract event 0xa04eca05f1c4f5674beaad80f345070124aa2192abced410b9b227c0c04e755a.
//
// Solidity: event ClaimerSet(address indexed claimer)
func (_PrizeVault *PrizeVaultFilterer) ParseClaimerSet(log types.Log) (*PrizeVaultClaimerSet, error) {
	event := new(PrizeVaultClaimerSet)
	if err := _PrizeVault.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the PrizeVault contract.
type PrizeVaultDepositIterator struct {
	Event *PrizeVaultDeposit // Event containing the contract specifics and raw log

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
func (it *PrizeVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultDeposit)
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
		it.Event = new(PrizeVaultDeposit)
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
func (it *PrizeVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultDeposit represents a Deposit event raised by the PrizeVault contract.
type PrizeVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_PrizeVault *PrizeVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*PrizeVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultDepositIterator{contract: _PrizeVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_PrizeVault *PrizeVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PrizeVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultDeposit)
				if err := _PrizeVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_PrizeVault *PrizeVaultFilterer) ParseDeposit(log types.Log) (*PrizeVaultDeposit, error) {
	event := new(PrizeVaultDeposit)
	if err := _PrizeVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the PrizeVault contract.
type PrizeVaultEIP712DomainChangedIterator struct {
	Event *PrizeVaultEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *PrizeVaultEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultEIP712DomainChanged)
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
		it.Event = new(PrizeVaultEIP712DomainChanged)
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
func (it *PrizeVaultEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultEIP712DomainChanged represents a EIP712DomainChanged event raised by the PrizeVault contract.
type PrizeVaultEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_PrizeVault *PrizeVaultFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*PrizeVaultEIP712DomainChangedIterator, error) {

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &PrizeVaultEIP712DomainChangedIterator{contract: _PrizeVault.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_PrizeVault *PrizeVaultFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *PrizeVaultEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultEIP712DomainChanged)
				if err := _PrizeVault.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_PrizeVault *PrizeVaultFilterer) ParseEIP712DomainChanged(log types.Log) (*PrizeVaultEIP712DomainChanged, error) {
	event := new(PrizeVaultEIP712DomainChanged)
	if err := _PrizeVault.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultLiquidationPairSetIterator is returned from FilterLiquidationPairSet and is used to iterate over the raw logs and unpacked data for LiquidationPairSet events raised by the PrizeVault contract.
type PrizeVaultLiquidationPairSetIterator struct {
	Event *PrizeVaultLiquidationPairSet // Event containing the contract specifics and raw log

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
func (it *PrizeVaultLiquidationPairSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultLiquidationPairSet)
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
		it.Event = new(PrizeVaultLiquidationPairSet)
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
func (it *PrizeVaultLiquidationPairSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultLiquidationPairSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultLiquidationPairSet represents a LiquidationPairSet event raised by the PrizeVault contract.
type PrizeVaultLiquidationPairSet struct {
	TokenOut        common.Address
	LiquidationPair common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidationPairSet is a free log retrieval operation binding the contract event 0xb540015bf51edcbfd9e43df5316486615bf954f7b4c6cc0304eb3757f1601f95.
//
// Solidity: event LiquidationPairSet(address indexed tokenOut, address indexed liquidationPair)
func (_PrizeVault *PrizeVaultFilterer) FilterLiquidationPairSet(opts *bind.FilterOpts, tokenOut []common.Address, liquidationPair []common.Address) (*PrizeVaultLiquidationPairSetIterator, error) {

	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}
	var liquidationPairRule []interface{}
	for _, liquidationPairItem := range liquidationPair {
		liquidationPairRule = append(liquidationPairRule, liquidationPairItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "LiquidationPairSet", tokenOutRule, liquidationPairRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultLiquidationPairSetIterator{contract: _PrizeVault.contract, event: "LiquidationPairSet", logs: logs, sub: sub}, nil
}

// WatchLiquidationPairSet is a free log subscription operation binding the contract event 0xb540015bf51edcbfd9e43df5316486615bf954f7b4c6cc0304eb3757f1601f95.
//
// Solidity: event LiquidationPairSet(address indexed tokenOut, address indexed liquidationPair)
func (_PrizeVault *PrizeVaultFilterer) WatchLiquidationPairSet(opts *bind.WatchOpts, sink chan<- *PrizeVaultLiquidationPairSet, tokenOut []common.Address, liquidationPair []common.Address) (event.Subscription, error) {

	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}
	var liquidationPairRule []interface{}
	for _, liquidationPairItem := range liquidationPair {
		liquidationPairRule = append(liquidationPairRule, liquidationPairItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "LiquidationPairSet", tokenOutRule, liquidationPairRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultLiquidationPairSet)
				if err := _PrizeVault.contract.UnpackLog(event, "LiquidationPairSet", log); err != nil {
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

// ParseLiquidationPairSet is a log parse operation binding the contract event 0xb540015bf51edcbfd9e43df5316486615bf954f7b4c6cc0304eb3757f1601f95.
//
// Solidity: event LiquidationPairSet(address indexed tokenOut, address indexed liquidationPair)
func (_PrizeVault *PrizeVaultFilterer) ParseLiquidationPairSet(log types.Log) (*PrizeVaultLiquidationPairSet, error) {
	event := new(PrizeVaultLiquidationPairSet)
	if err := _PrizeVault.contract.UnpackLog(event, "LiquidationPairSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultOwnershipOfferedIterator is returned from FilterOwnershipOffered and is used to iterate over the raw logs and unpacked data for OwnershipOffered events raised by the PrizeVault contract.
type PrizeVaultOwnershipOfferedIterator struct {
	Event *PrizeVaultOwnershipOffered // Event containing the contract specifics and raw log

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
func (it *PrizeVaultOwnershipOfferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultOwnershipOffered)
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
		it.Event = new(PrizeVaultOwnershipOffered)
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
func (it *PrizeVaultOwnershipOfferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultOwnershipOfferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultOwnershipOffered represents a OwnershipOffered event raised by the PrizeVault contract.
type PrizeVaultOwnershipOffered struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipOffered is a free log retrieval operation binding the contract event 0x239a2ddded15777fa246aed5f7e1a9bc69a39d4eb4a397034d1d85766cca7d4c.
//
// Solidity: event OwnershipOffered(address indexed pendingOwner)
func (_PrizeVault *PrizeVaultFilterer) FilterOwnershipOffered(opts *bind.FilterOpts, pendingOwner []common.Address) (*PrizeVaultOwnershipOfferedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "OwnershipOffered", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultOwnershipOfferedIterator{contract: _PrizeVault.contract, event: "OwnershipOffered", logs: logs, sub: sub}, nil
}

// WatchOwnershipOffered is a free log subscription operation binding the contract event 0x239a2ddded15777fa246aed5f7e1a9bc69a39d4eb4a397034d1d85766cca7d4c.
//
// Solidity: event OwnershipOffered(address indexed pendingOwner)
func (_PrizeVault *PrizeVaultFilterer) WatchOwnershipOffered(opts *bind.WatchOpts, sink chan<- *PrizeVaultOwnershipOffered, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "OwnershipOffered", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultOwnershipOffered)
				if err := _PrizeVault.contract.UnpackLog(event, "OwnershipOffered", log); err != nil {
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

// ParseOwnershipOffered is a log parse operation binding the contract event 0x239a2ddded15777fa246aed5f7e1a9bc69a39d4eb4a397034d1d85766cca7d4c.
//
// Solidity: event OwnershipOffered(address indexed pendingOwner)
func (_PrizeVault *PrizeVaultFilterer) ParseOwnershipOffered(log types.Log) (*PrizeVaultOwnershipOffered, error) {
	event := new(PrizeVaultOwnershipOffered)
	if err := _PrizeVault.contract.UnpackLog(event, "OwnershipOffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PrizeVault contract.
type PrizeVaultOwnershipTransferredIterator struct {
	Event *PrizeVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PrizeVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultOwnershipTransferred)
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
		it.Event = new(PrizeVaultOwnershipTransferred)
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
func (it *PrizeVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultOwnershipTransferred represents a OwnershipTransferred event raised by the PrizeVault contract.
type PrizeVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrizeVault *PrizeVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PrizeVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultOwnershipTransferredIterator{contract: _PrizeVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrizeVault *PrizeVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PrizeVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultOwnershipTransferred)
				if err := _PrizeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrizeVault *PrizeVaultFilterer) ParseOwnershipTransferred(log types.Log) (*PrizeVaultOwnershipTransferred, error) {
	event := new(PrizeVaultOwnershipTransferred)
	if err := _PrizeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultSetHooksIterator is returned from FilterSetHooks and is used to iterate over the raw logs and unpacked data for SetHooks events raised by the PrizeVault contract.
type PrizeVaultSetHooksIterator struct {
	Event *PrizeVaultSetHooks // Event containing the contract specifics and raw log

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
func (it *PrizeVaultSetHooksIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultSetHooks)
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
		it.Event = new(PrizeVaultSetHooks)
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
func (it *PrizeVaultSetHooksIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultSetHooksIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultSetHooks represents a SetHooks event raised by the PrizeVault contract.
type PrizeVaultSetHooks struct {
	Account common.Address
	Hooks   PrizeHooks
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetHooks is a free log retrieval operation binding the contract event 0x5eaf51436f308830fc57b00bb6843675b383bce9b2c669e1b9ce60de8bbb9e8e.
//
// Solidity: event SetHooks(address indexed account, (bool,bool,address) hooks)
func (_PrizeVault *PrizeVaultFilterer) FilterSetHooks(opts *bind.FilterOpts, account []common.Address) (*PrizeVaultSetHooksIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "SetHooks", accountRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultSetHooksIterator{contract: _PrizeVault.contract, event: "SetHooks", logs: logs, sub: sub}, nil
}

// WatchSetHooks is a free log subscription operation binding the contract event 0x5eaf51436f308830fc57b00bb6843675b383bce9b2c669e1b9ce60de8bbb9e8e.
//
// Solidity: event SetHooks(address indexed account, (bool,bool,address) hooks)
func (_PrizeVault *PrizeVaultFilterer) WatchSetHooks(opts *bind.WatchOpts, sink chan<- *PrizeVaultSetHooks, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "SetHooks", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultSetHooks)
				if err := _PrizeVault.contract.UnpackLog(event, "SetHooks", log); err != nil {
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

// ParseSetHooks is a log parse operation binding the contract event 0x5eaf51436f308830fc57b00bb6843675b383bce9b2c669e1b9ce60de8bbb9e8e.
//
// Solidity: event SetHooks(address indexed account, (bool,bool,address) hooks)
func (_PrizeVault *PrizeVaultFilterer) ParseSetHooks(log types.Log) (*PrizeVaultSetHooks, error) {
	event := new(PrizeVaultSetHooks)
	if err := _PrizeVault.contract.UnpackLog(event, "SetHooks", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PrizeVault contract.
type PrizeVaultTransferIterator struct {
	Event *PrizeVaultTransfer // Event containing the contract specifics and raw log

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
func (it *PrizeVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultTransfer)
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
		it.Event = new(PrizeVaultTransfer)
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
func (it *PrizeVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultTransfer represents a Transfer event raised by the PrizeVault contract.
type PrizeVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PrizeVault *PrizeVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PrizeVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultTransferIterator{contract: _PrizeVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PrizeVault *PrizeVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PrizeVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultTransfer)
				if err := _PrizeVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PrizeVault *PrizeVaultFilterer) ParseTransfer(log types.Log) (*PrizeVaultTransfer, error) {
	event := new(PrizeVaultTransfer)
	if err := _PrizeVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultTransferYieldOutIterator is returned from FilterTransferYieldOut and is used to iterate over the raw logs and unpacked data for TransferYieldOut events raised by the PrizeVault contract.
type PrizeVaultTransferYieldOutIterator struct {
	Event *PrizeVaultTransferYieldOut // Event containing the contract specifics and raw log

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
func (it *PrizeVaultTransferYieldOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultTransferYieldOut)
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
		it.Event = new(PrizeVaultTransferYieldOut)
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
func (it *PrizeVaultTransferYieldOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultTransferYieldOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultTransferYieldOut represents a TransferYieldOut event raised by the PrizeVault contract.
type PrizeVaultTransferYieldOut struct {
	LiquidationPair common.Address
	TokenOut        common.Address
	Recipient       common.Address
	AmountOut       *big.Int
	YieldFee        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferYieldOut is a free log retrieval operation binding the contract event 0x64d8a6bd56c2c610cd5d2ed87ba4bde3659df2650c7ad8aabde397decf94e07d.
//
// Solidity: event TransferYieldOut(address indexed liquidationPair, address indexed tokenOut, address indexed recipient, uint256 amountOut, uint256 yieldFee)
func (_PrizeVault *PrizeVaultFilterer) FilterTransferYieldOut(opts *bind.FilterOpts, liquidationPair []common.Address, tokenOut []common.Address, recipient []common.Address) (*PrizeVaultTransferYieldOutIterator, error) {

	var liquidationPairRule []interface{}
	for _, liquidationPairItem := range liquidationPair {
		liquidationPairRule = append(liquidationPairRule, liquidationPairItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "TransferYieldOut", liquidationPairRule, tokenOutRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultTransferYieldOutIterator{contract: _PrizeVault.contract, event: "TransferYieldOut", logs: logs, sub: sub}, nil
}

// WatchTransferYieldOut is a free log subscription operation binding the contract event 0x64d8a6bd56c2c610cd5d2ed87ba4bde3659df2650c7ad8aabde397decf94e07d.
//
// Solidity: event TransferYieldOut(address indexed liquidationPair, address indexed tokenOut, address indexed recipient, uint256 amountOut, uint256 yieldFee)
func (_PrizeVault *PrizeVaultFilterer) WatchTransferYieldOut(opts *bind.WatchOpts, sink chan<- *PrizeVaultTransferYieldOut, liquidationPair []common.Address, tokenOut []common.Address, recipient []common.Address) (event.Subscription, error) {

	var liquidationPairRule []interface{}
	for _, liquidationPairItem := range liquidationPair {
		liquidationPairRule = append(liquidationPairRule, liquidationPairItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "TransferYieldOut", liquidationPairRule, tokenOutRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultTransferYieldOut)
				if err := _PrizeVault.contract.UnpackLog(event, "TransferYieldOut", log); err != nil {
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

// ParseTransferYieldOut is a log parse operation binding the contract event 0x64d8a6bd56c2c610cd5d2ed87ba4bde3659df2650c7ad8aabde397decf94e07d.
//
// Solidity: event TransferYieldOut(address indexed liquidationPair, address indexed tokenOut, address indexed recipient, uint256 amountOut, uint256 yieldFee)
func (_PrizeVault *PrizeVaultFilterer) ParseTransferYieldOut(log types.Log) (*PrizeVaultTransferYieldOut, error) {
	event := new(PrizeVaultTransferYieldOut)
	if err := _PrizeVault.contract.UnpackLog(event, "TransferYieldOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PrizeVault contract.
type PrizeVaultWithdrawIterator struct {
	Event *PrizeVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *PrizeVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultWithdraw)
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
		it.Event = new(PrizeVaultWithdraw)
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
func (it *PrizeVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultWithdraw represents a Withdraw event raised by the PrizeVault contract.
type PrizeVaultWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_PrizeVault *PrizeVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*PrizeVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultWithdrawIterator{contract: _PrizeVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_PrizeVault *PrizeVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PrizeVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultWithdraw)
				if err := _PrizeVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_PrizeVault *PrizeVaultFilterer) ParseWithdraw(log types.Log) (*PrizeVaultWithdraw, error) {
	event := new(PrizeVaultWithdraw)
	if err := _PrizeVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultYieldFeePercentageSetIterator is returned from FilterYieldFeePercentageSet and is used to iterate over the raw logs and unpacked data for YieldFeePercentageSet events raised by the PrizeVault contract.
type PrizeVaultYieldFeePercentageSetIterator struct {
	Event *PrizeVaultYieldFeePercentageSet // Event containing the contract specifics and raw log

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
func (it *PrizeVaultYieldFeePercentageSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultYieldFeePercentageSet)
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
		it.Event = new(PrizeVaultYieldFeePercentageSet)
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
func (it *PrizeVaultYieldFeePercentageSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultYieldFeePercentageSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultYieldFeePercentageSet represents a YieldFeePercentageSet event raised by the PrizeVault contract.
type PrizeVaultYieldFeePercentageSet struct {
	YieldFeePercentage *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterYieldFeePercentageSet is a free log retrieval operation binding the contract event 0xefb99600b8c2baadc10ee4fe77c625b379acb4c29fbb0b90a5f17652d7de0341.
//
// Solidity: event YieldFeePercentageSet(uint256 yieldFeePercentage)
func (_PrizeVault *PrizeVaultFilterer) FilterYieldFeePercentageSet(opts *bind.FilterOpts) (*PrizeVaultYieldFeePercentageSetIterator, error) {

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "YieldFeePercentageSet")
	if err != nil {
		return nil, err
	}
	return &PrizeVaultYieldFeePercentageSetIterator{contract: _PrizeVault.contract, event: "YieldFeePercentageSet", logs: logs, sub: sub}, nil
}

// WatchYieldFeePercentageSet is a free log subscription operation binding the contract event 0xefb99600b8c2baadc10ee4fe77c625b379acb4c29fbb0b90a5f17652d7de0341.
//
// Solidity: event YieldFeePercentageSet(uint256 yieldFeePercentage)
func (_PrizeVault *PrizeVaultFilterer) WatchYieldFeePercentageSet(opts *bind.WatchOpts, sink chan<- *PrizeVaultYieldFeePercentageSet) (event.Subscription, error) {

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "YieldFeePercentageSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultYieldFeePercentageSet)
				if err := _PrizeVault.contract.UnpackLog(event, "YieldFeePercentageSet", log); err != nil {
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

// ParseYieldFeePercentageSet is a log parse operation binding the contract event 0xefb99600b8c2baadc10ee4fe77c625b379acb4c29fbb0b90a5f17652d7de0341.
//
// Solidity: event YieldFeePercentageSet(uint256 yieldFeePercentage)
func (_PrizeVault *PrizeVaultFilterer) ParseYieldFeePercentageSet(log types.Log) (*PrizeVaultYieldFeePercentageSet, error) {
	event := new(PrizeVaultYieldFeePercentageSet)
	if err := _PrizeVault.contract.UnpackLog(event, "YieldFeePercentageSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeVaultYieldFeeRecipientSetIterator is returned from FilterYieldFeeRecipientSet and is used to iterate over the raw logs and unpacked data for YieldFeeRecipientSet events raised by the PrizeVault contract.
type PrizeVaultYieldFeeRecipientSetIterator struct {
	Event *PrizeVaultYieldFeeRecipientSet // Event containing the contract specifics and raw log

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
func (it *PrizeVaultYieldFeeRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeVaultYieldFeeRecipientSet)
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
		it.Event = new(PrizeVaultYieldFeeRecipientSet)
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
func (it *PrizeVaultYieldFeeRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeVaultYieldFeeRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeVaultYieldFeeRecipientSet represents a YieldFeeRecipientSet event raised by the PrizeVault contract.
type PrizeVaultYieldFeeRecipientSet struct {
	YieldFeeRecipient common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterYieldFeeRecipientSet is a free log retrieval operation binding the contract event 0x5d55a9f17ebec95846f4d400ad862a51f9564a7973f6622d3c99a34feb6a0aec.
//
// Solidity: event YieldFeeRecipientSet(address indexed yieldFeeRecipient)
func (_PrizeVault *PrizeVaultFilterer) FilterYieldFeeRecipientSet(opts *bind.FilterOpts, yieldFeeRecipient []common.Address) (*PrizeVaultYieldFeeRecipientSetIterator, error) {

	var yieldFeeRecipientRule []interface{}
	for _, yieldFeeRecipientItem := range yieldFeeRecipient {
		yieldFeeRecipientRule = append(yieldFeeRecipientRule, yieldFeeRecipientItem)
	}

	logs, sub, err := _PrizeVault.contract.FilterLogs(opts, "YieldFeeRecipientSet", yieldFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &PrizeVaultYieldFeeRecipientSetIterator{contract: _PrizeVault.contract, event: "YieldFeeRecipientSet", logs: logs, sub: sub}, nil
}

// WatchYieldFeeRecipientSet is a free log subscription operation binding the contract event 0x5d55a9f17ebec95846f4d400ad862a51f9564a7973f6622d3c99a34feb6a0aec.
//
// Solidity: event YieldFeeRecipientSet(address indexed yieldFeeRecipient)
func (_PrizeVault *PrizeVaultFilterer) WatchYieldFeeRecipientSet(opts *bind.WatchOpts, sink chan<- *PrizeVaultYieldFeeRecipientSet, yieldFeeRecipient []common.Address) (event.Subscription, error) {

	var yieldFeeRecipientRule []interface{}
	for _, yieldFeeRecipientItem := range yieldFeeRecipient {
		yieldFeeRecipientRule = append(yieldFeeRecipientRule, yieldFeeRecipientItem)
	}

	logs, sub, err := _PrizeVault.contract.WatchLogs(opts, "YieldFeeRecipientSet", yieldFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeVaultYieldFeeRecipientSet)
				if err := _PrizeVault.contract.UnpackLog(event, "YieldFeeRecipientSet", log); err != nil {
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

// ParseYieldFeeRecipientSet is a log parse operation binding the contract event 0x5d55a9f17ebec95846f4d400ad862a51f9564a7973f6622d3c99a34feb6a0aec.
//
// Solidity: event YieldFeeRecipientSet(address indexed yieldFeeRecipient)
func (_PrizeVault *PrizeVaultFilterer) ParseYieldFeeRecipientSet(log types.Log) (*PrizeVaultYieldFeeRecipientSet, error) {
	event := new(PrizeVaultYieldFeeRecipientSet)
	if err := _PrizeVault.contract.UnpackLog(event, "YieldFeeRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
