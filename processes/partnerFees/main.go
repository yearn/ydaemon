package partnerFees

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gocarina/gocsv"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/harvests"
	"github.com/yearn/ydaemon/internal/indexer"
	partnerTracker "github.com/yearn/ydaemon/internal/indexer.partnerTracker"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/registries"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
)

var AllHarvests = make(map[common.Address][]vaults.THarvest)
var STRATLIST = []strategies.TStrategy{}
var LEDGER = common.HexToAddress(`0x558247e365be655f9144e1a0140D793984372Ef3`)

type TMainCSVExport struct {
	Block               uint64 `csv:"block"`
	Timestamp           uint64 `csv:"timestamp"`
	Wrapper             string `csv:"partner"`
	Vault               string `csv:"vault"`
	Tier                string `csv:"tier"`
	VaultPrice          string `csv:"vault_price"`
	Balance             string `csv:"delegated_balance"`
	BalanceUSD          string `csv:"delegated_balance_usd"`
	TotalSupply         string `csv:"total_supply"`
	TotalSupplyUSD      string `csv:"total_supply_usd"`
	TotalDelegatedValue string `csv:"total_delegated_value"`
	CollectedFee        string `csv:"collected_fee"`
	CollectedFeeUSD     string `csv:"collected_fee_usd"`
	PartnerFeeShare     string `csv:"partner_fee"`
	PartnerFeeShareUSD  string `csv:"partner_fee_usd"`
}

type TPartnerBreakdownCSVExport struct {
	Block      uint64 `csv:"block"`
	Wrapper    string `csv:"wrapper"`
	Vault      string `csv:"vault"`
	Depositor  string `csv:"depositor"`
	Balance    string `csv:"balance"`
	BalanceUSD string `csv:"balance_usd"`
	Share      string `csv:"share"`
}

type TWithNormalized struct {
	Raw        *bigNumber.Int
	Normalized *bigNumber.Float
	Value      *bigNumber.Float
}

func loadVaultAndStrategies(chainID uint64) {
	vaultsMap := registries.RetrieveAllVaults(chainID, 0)
	tokens.RetrieveAllTokens(chainID, vaultsMap)
	prices.RetrieveAllPrices(chainID)
	vaults.RetrieveAllVaults(chainID, vaultsMap)
	strategiesAddedList := strategies.RetrieveAllStrategiesAdded(chainID, vaultsMap)
	strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
	indexer.PostProcessStrategies(chainID)
}

func getVaultsSupplyAtBlock(
	allVaults []*vaults.TVault,
	allHarvests map[common.Address]map[common.Address][]harvests.THarvest,
) map[common.Address]map[uint64]*big.Int {
	syncMap := sync.Map{}
	for _, vault := range allVaults {
		vaultContract, _ := contracts.NewERC20(vault.Address, ethereum.GetRPC(vault.ChainID))

		blockNumbers := []uint64{}
		for _, harvests := range allHarvests[vault.Address] {
			for _, harvest := range harvests {
				blockNumbers = append(blockNumbers, harvest.BlockNumber)
			}
		}

		innerWg := sync.WaitGroup{}
		for _, blockNumber := range blockNumbers {
			innerWg.Add(1)
			go func(blockNumber uint64) {
				defer innerWg.Done()
				totalSupply, err := vaultContract.TotalSupply(&bind.CallOpts{BlockNumber: big.NewInt(int64(blockNumber))})
				if err != nil {
					logs.Error(err)
					return
				}
				eventKey := vault.Address.Hex() + `-` + strconv.FormatUint(blockNumber, 10)
				syncMap.Store(eventKey, totalSupply)
			}(blockNumber)
		}
		innerWg.Wait()
	}

	vaultsSupplyAtBlock := make(map[common.Address]map[uint64]*big.Int)
	syncMap.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddress := common.HexToAddress(eventKey[0])
		blockNumber, _ := strconv.ParseUint(eventKey[1], 10, 64)

		if vaultsSupplyAtBlock[vaultAddress] == nil {
			vaultsSupplyAtBlock[vaultAddress] = make(map[uint64]*big.Int)
		}
		vaultsSupplyAtBlock[vaultAddress][blockNumber] = value.(*big.Int)
		return true
	})

	return vaultsSupplyAtBlock
}

func computePartnerVaultTVL(
	blockNumber uint64,
	currentVault common.Address,
	vaultDecimals uint64,
	partnersTrackerEvents map[common.Address]map[common.Address][]partnerTracker.TEventReferredBalanceIncreased,
	allRefereesEvents map[common.Address][]partnerTracker.TRefereeTransferEvent,
) (map[common.Address]*bigNumber.Int, map[common.Address]map[common.Address]*bigNumber.Int) {
	delegatedTVL := make(map[common.Address]map[common.Address]*bigNumber.Int)
	partnerTVL := make(map[common.Address]*bigNumber.Int)

	/**********************************************************************************************
	** We have one harvest. Now, we want to filter the deposits events that happened before the
	** harvest from the partner tracker in order to be able to compute, for that harvest, the
	** amount of fees that each partner should receive.
	** We will create a map of partner -> depositor -> listOfDeposits.
	** This list of deposit will not be the only thing used to compute the fees as we may check
	** the balance of the depositor in the vault before the harvest to see if he had other
	** deposits that were not done by the partner, or if he had done some withdrawals.
	**********************************************************************************************/
	relevantPartnerTrackerEvents := partnerTracker.FilterReferralBalanceIncreaseEventsForVault(
		partnersTrackerEvents,
		blockNumber,
	)

	/**********************************************************************************************
	** We have the list of deposits that happened before the harvest. Now, we want to check the
	** balance of the depositor in the vault before the harvest to see if he had other
	** deposits that were not done by the partner, or if he had done some withdrawals.
	**********************************************************************************************/
	for partner, eventsDepositorLevel := range relevantPartnerTrackerEvents {
		if _, ok := delegatedTVL[partner]; !ok {
			delegatedTVL[partner] = make(map[common.Address]*bigNumber.Int)
			partnerTVL[partner] = bigNumber.NewInt(0)
		}
		totalAllocatedForThisPartner := bigNumber.NewInt(0)
		for depositor, events := range eventsDepositorLevel {
			/******************************************************************************
			** We can compute all the deposits for the depositor that happened before the
			** harvest for that specific partner in order to know how much fees that
			** partner should receive.
			******************************************************************************/
			earliestDepositBlock := uint64(math.MaxUint64)
			userDepositViaPartnerMap := map[common.Hash]partnerTracker.TEventReferredBalanceIncreased{}
			for _, event := range events {
				userDepositViaPartnerMap[event.TxHash] = event
				if event.BlockNumber < earliestDepositBlock {
					earliestDepositBlock = event.BlockNumber
				}
			}

			/******************************************************************************
			** However, we also want to check the balance of the depositor in the vault
			** before the harvest to ensure that:
			** - The depositor did not have other deposits not done by the partner
			** - The depositor still have some tokens in the vault.
			** If no token are left on the vault, we just skip that depositor for that
			** partner for that harvest.
			******************************************************************************/
			allTransfersForThisReferee, ok := allRefereesEvents[depositor]
			if !ok {
				continue //no transfers for this referee
			}

			delegatedBalance := bigNumber.NewInt(0)
			actualBalance := bigNumber.NewInt(0)
			for _, transfer := range allTransfersForThisReferee {
				if transfer.BlockNumber > blockNumber {
					continue //transfer happened after the harvest
				}
				if transfer.BlockNumber < earliestDepositBlock {
					continue //transfer happened before the earliest deposit via partner
				}
				isReceiving := transfer.To.Hex() == depositor.Hex()
				isSending := transfer.From.Hex() == depositor.Hex()

				/**************************************************************************
				** For each transfer IN, we check if the transfer hash matches one of the
				** deposit hash that we have computed above. If it does, we add the
				** amount to the accounting balance.
				** Otherwise, we just skip that deposit as it is not relevant for our
				** current partner accountability.
				**************************************************************************/
				if isReceiving {
					actualBalance = bigNumber.NewInt(0).Add(actualBalance, bigNumber.NewInt(0).Clone(transfer.Value))
					if _, ok := userDepositViaPartnerMap[transfer.TxHash]; !ok {
						continue // It's not a deposit via this partner, we skip it.
					}
					delegatedBalance = bigNumber.NewInt(0).Add(delegatedBalance, bigNumber.NewInt(0).Clone(transfer.Value))
				}

				/**************************************************************************
				** For each transfer OUT, we need to remove the amount from the balance in
				** proportion to the amount that the partner deposited. If the accounting
				** balance is lower than the amount that the user is trying to withdraw,
				** we just set the balance to 0.
				**************************************************************************/
				if isSending {
					actualBalanceBefore := bigNumber.NewInt(0).Clone(actualBalance)
					actualBalance = bigNumber.NewInt(0).Sub(actualBalance, bigNumber.NewInt(0).Clone(transfer.Value))
					/**************************************************************************
					** If the user is withdrawing all the tokens, we can just skip all the
					** subsequent calculations and set his delegated balance to 0.
					**************************************************************************/
					if actualBalance.Lte(bigNumber.NewInt(0)) {
						delegatedBalance = bigNumber.NewInt(0)
						continue
					}

					/**************************************************************************
					** The user is withdrawing X. Only Y% of X is delegated to the partner. We
					** first need to calculate Y% based on the delegated balance and the actual
					** balance.
					** Then we need to calculate the amount that the user is withdrawing from
					** the delegated balance.
					**************************************************************************/
					delegatedRatio := bigNumber.NewFloat(0).Div(bigNumber.NewFloat().SetInt(delegatedBalance), bigNumber.NewFloat().SetInt(actualBalanceBefore))
					delegatedWithdraw := bigNumber.NewFloat(0).Mul(delegatedRatio, bigNumber.NewFloat(0).SetInt(transfer.Value))
					delegatedBalance = bigNumber.NewInt(0).Sub(delegatedBalance, delegatedWithdraw.Int())
				}
			}

			/**********************************************************************************************
			** We can now check the actual balance of the depositor in the vault before the harvest.
			** If the balance is 0, we just skip that depositor for that partner for that harvest.
			** If the balance is greater than 0, that's not a problem, we can use the expected balance
			** If the balance is lower than the expected balance, that's a problem and we should scale
			** the expected balance down.
			**********************************************************************************************/
			if delegatedBalance.Cmp(big.NewInt(0)) == 0 {
				continue //Skip, no balance in the vault
			}
			delegatedTVL[partner][depositor] = delegatedBalance
			totalAllocatedForThisPartner = bigNumber.NewInt(0).Add(totalAllocatedForThisPartner, delegatedBalance)
		}
		partnerTVL[partner] = bigNumber.NewInt(0).Add(partnerTVL[partner], totalAllocatedForThisPartner)
	}
	return partnerTVL, delegatedTVL
}

/**************************************************************************************************
** computePartnerTotalTVL will compute the total TVL for each partner at the given block based on
** the events that we have collected.
**
** Arguments:
** - chainID: The chain ID of the chain that we are computing the TVL for.
** - blockNumber: The block number at which we want to compute the TVL.
** - partnersTrackerEvents: The events that we have collected from the partners tracker.
**   It's a map of: map[vault][partner][depositor][events[]]
** - allRefereesEvents: The events that we have collected from the referees tracker.
**   It's a map of: map[vault][depositor][events[]]
**
** Returns:
** - totalTVL: The total TVL for each partner at the given block.
** - tierShare: The share of each partner in the total TVL.
** - delegateValue: The amount each depositor delegated to each partner.
**************************************************************************************************/
func computePartnerTotalTVL(
	chainID uint64,
	blockNumber uint64,
	partnersTrackerEvents map[common.Address]map[common.Address]map[common.Address][]partnerTracker.TEventReferredBalanceIncreased,
	allRefereesEvents map[common.Address]map[common.Address][]partnerTracker.TRefereeTransferEvent,
) (
	totalTVL map[common.Address]*bigNumber.Float,
	tierShare map[common.Address]*bigNumber.Float,
	delegateValue map[common.Address]map[common.Address]map[common.Address]TWithNormalized,
) {
	totalTVL = make(map[common.Address]*bigNumber.Float)
	tierShare = make(map[common.Address]*bigNumber.Float)
	delegateValue = make(map[common.Address]map[common.Address]map[common.Address]TWithNormalized)

	/**********************************************************************************************
	** First we sort all the events by block number in order to compute the TVL at each block
	**********************************************************************************************/
	relevantPartnerTrackerEvents := partnerTracker.SortReferralBalanceIncreaseEvents(partnersTrackerEvents)

	/**********************************************************************************************
	** We want to know, for every harvest that happened impacting the partner:
	** - if a partner has a delegated balance in any vault
	** - the amount of delegated balance in each of theses vaults.
	**
	** For this we first need to list the vaults in which the partner has a delegated balance.
	** We will create a partnerVaults map that will look like this:
	** partnerVaults -> map[partnerAddress][vaultAddress][events[]]
	**
	** We will also create a map that will contain each event mapped to it's hash. This will allow
	** us to quickly check if a transfer event is a deposit or not.
	** eventsHash -> [vaultAddress][eventHash][event]
	**********************************************************************************************/
	partnerVaults := make(map[common.Address]map[common.Address][]partnerTracker.TEventReferredBalanceIncreased)
	eventsHash := make(map[common.Address]map[common.Hash]partnerTracker.TEventReferredBalanceIncreased)
	for vaultAddress, eventsForThatVault := range relevantPartnerTrackerEvents {
		for partnerAddress, eventsForThanPartner := range eventsForThatVault {
			for _, events := range eventsForThanPartner {
				if _, ok := partnerVaults[partnerAddress]; !ok {
					partnerVaults[partnerAddress] = make(map[common.Address][]partnerTracker.TEventReferredBalanceIncreased)
					eventsHash[partnerAddress] = make(map[common.Hash]partnerTracker.TEventReferredBalanceIncreased)
				}
				partnerVaults[partnerAddress][vaultAddress] = append(partnerVaults[partnerAddress][vaultAddress], events...)
				for _, event := range events {
					if _, ok := eventsHash[vaultAddress]; !ok {
						eventsHash[vaultAddress] = make(map[common.Hash]partnerTracker.TEventReferredBalanceIncreased)
					}
					eventsHash[vaultAddress][event.TxHash] = event
				}
			}
		}
	}

	/**********************************************************************************************
	** Once we got that, we got all the INCREASE events for each partner in each vault, aka when
	** a user delegates to the partner. As the DECREASE events are not emitted, we will need to
	** filter the TRANSFER events from the delegator to somewhere else in order to know the, after
	** each transfer, the amount of delegated balance that the partner has in the vault.
	** If the user receive new funds in a vault, it will
	** - be to the same delegated partner, in which case we should have a matching INCREASE hash
	** - be to a different/none delegated partner, in which case the delegated balance should not
	**   change.
	** If the user withdraws funds from a vault, we will need to decrease the delegated balance by
	** the ratio of the amount withdrawn to the total balance of the user in the vault vs the
	** delegated balance.
	**
	** We will create a map that will contain the delegated balance for each partner in each vault
	** delegatedVaultPartnerBalance -> map[vaultAddr][partnerAddr][depositorAddr][delegatedBalance]
	**********************************************************************************************/
	delegatedVaultPartnerBalance := make(map[common.Address]map[common.Address]map[common.Address]*bigNumber.Int)
	for vaultAddress, eventsForThatVault := range allRefereesEvents {
		for depositorAddress, eventForThisDepositor := range eventsForThatVault {
			depositorDelegatedBalance := make(map[common.Address]*bigNumber.Int)
			depositorActualBalance := bigNumber.NewInt(0)

			for _, event := range eventForThisDepositor {
				if event.BlockNumber > blockNumber {
					continue
				}
				isReceiving := event.To.Hex() == depositorAddress.Hex()
				isSending := event.From.Hex() == depositorAddress.Hex()

				/**********************************************************************************
				** If the user is receiving fund, this means that he has either:
				** - deposited via partner tracker, in which case we should have a matching hash
				** - deposited via other way, in which case we should not have a matching hash
				** - received a transfer from another user, so no matching hash either
				**********************************************************************************/
				if isReceiving {
					/******************************************************************************
					** First we need to check if this specific deposit event is linked to a partner
					** deposit. All deposits via the partner tracker will have a hash that matches.
					** All other transfers will be ignored.
					******************************************************************************/
					if _, ok := eventsHash[vaultAddress]; ok {
						if partnerTrackerEvent, ok := eventsHash[vaultAddress][event.TxHash]; ok {
							/**********************************************************************
							** We have a match, this is a deposit via the partner tracker.
							**********************************************************************/
							partnerAddress := partnerTrackerEvent.PartnerID
							if _, ok := depositorDelegatedBalance[partnerAddress]; !ok {
								depositorDelegatedBalance[partnerAddress] = bigNumber.NewInt(0)
							}
							depositorDelegatedBalance[partnerAddress] = bigNumber.NewInt(0).Add(depositorDelegatedBalance[partnerAddress], event.Value)
							depositorActualBalance = bigNumber.NewInt(0).Add(depositorActualBalance, event.Value)
							continue
						}
					}

					/******************************************************************************
					** If we are here, it means that this is not a deposit via the partner tracker.
					** We need to update the actual balance of the user in the vault without
					** changing the delegated balance.
					******************************************************************************/
					depositorActualBalance = bigNumber.NewInt(0).Add(depositorActualBalance, event.Value)
					continue
				}

				/**********************************************************************************
				** If the user is sending fund, we need to adjust the delegated balance by the
				** ratio of the amount withdrawn to the total balance of the user in the vault vs
				** the delegated balance.
				** Aka, if the user has 1000 DAI in the vault, and 500 DAI delegated to partner A,
				** and he withdraws 100 DAI, we need to decrease the delegated balance of partner A
				** by 50 DAI.
				** As we are not working with a specific partner, we need to do this for all
				** partners.
				**********************************************************************************/
				if isSending {
					depositorActualBalanceBeforeEvent := bigNumber.NewInt(0).Clone(depositorActualBalance)
					depositorActualBalance = bigNumber.NewInt(0).Sub(depositorActualBalance, event.Value)
					for partnerAddress, delegatedBalance := range depositorDelegatedBalance {
						delegatedRatio := bigNumber.NewFloat(0).Div(
							bigNumber.NewFloat().SetInt(delegatedBalance),
							bigNumber.NewFloat().SetInt(depositorActualBalanceBeforeEvent),
						)
						delegatedWithdraw := bigNumber.NewFloat(0).Mul(
							delegatedRatio,
							bigNumber.NewFloat(0).SetInt(event.Value),
						)
						expectedNewDelegatedBalance := bigNumber.NewInt(0).Sub(delegatedBalance, delegatedWithdraw.Int())
						if expectedNewDelegatedBalance.Lte(bigNumber.NewInt(0)) {
							depositorDelegatedBalance[partnerAddress] = bigNumber.NewInt(0)
						} else {
							depositorDelegatedBalance[partnerAddress] = bigNumber.NewInt(0).Sub(delegatedBalance, delegatedWithdraw.Int())
						}
					}
				}
			}

			if depositorActualBalance.Lte(bigNumber.NewInt(0)) {
				continue
			}
			for partnerAddress, delegatedBalance := range depositorDelegatedBalance {
				if _, ok := delegatedVaultPartnerBalance[vaultAddress]; !ok {
					delegatedVaultPartnerBalance[vaultAddress] = make(map[common.Address]map[common.Address]*bigNumber.Int)
				}
				if _, ok := delegatedVaultPartnerBalance[vaultAddress][partnerAddress]; !ok {
					delegatedVaultPartnerBalance[vaultAddress][partnerAddress] = make(map[common.Address]*bigNumber.Int)
				}
				if _, ok := delegatedVaultPartnerBalance[vaultAddress][partnerAddress][depositorAddress]; !ok {
					delegatedVaultPartnerBalance[vaultAddress][partnerAddress][depositorAddress] = bigNumber.NewInt(0)
				}

				delegatedVaultPartnerBalance[vaultAddress][partnerAddress][depositorAddress] = bigNumber.NewInt(0).Add(
					delegatedVaultPartnerBalance[vaultAddress][partnerAddress][depositorAddress],
					delegatedBalance,
				)
			}
		}
	}

	/**********************************************************************************************
	** Now that we got that, we want to know, for a given partner, how much TVL he has in total.
	** As we have the delegated balance for each vault, we just need to grab this balance, the
	** vault token price and sum it up.
	** This first step if to fetch the vault token prices. We will batch the request for the block
	** in order to make it faster.
	**********************************************************************************************/
	partnersTVL := make(map[common.Address]*bigNumber.Float)
	vaultsWithTVL := []common.Address{}
	//First grab the missing prices
	for vaultAddress, partnersForThatVault := range delegatedVaultPartnerBalance {
		for partner, depositorsForThatPartner := range partnersForThatVault {
			for _, delegatedBalance := range depositorsForThatPartner {
				if _, ok := partnersTVL[partner]; !ok {
					partnersTVL[partner] = bigNumber.NewFloat(0)
				}
				if delegatedBalance.Lte(bigNumber.NewInt(0)) {
					continue
				}
				vaultsWithTVL = append(vaultsWithTVL, vaultAddress)
			}
		}
	}
	prices.FetchPricesOnBlock(chainID, blockNumber, vaultsWithTVL)

	/**********************************************************************************************
	** Then, we can just add the TVL for each partner.
	**********************************************************************************************/
	for vaultAddress, partnersForThatVault := range delegatedVaultPartnerBalance {
		for partner, depositorsForThatPartner := range partnersForThatVault {
			for depositorAddress, delegatedBalance := range depositorsForThatPartner {
				if _, ok := partnersTVL[partner]; !ok {
					partnersTVL[partner] = bigNumber.NewFloat(0)
				}
				if delegatedBalance.Gt(bigNumber.NewInt(0)) {
					thisVault, _ := vaults.GetVault(chainID, vaultAddress)
					if blockNumber == 16680455 && addresses.Equals(partner, LEDGER) {
						value := bigNumber.NewFloat(0).Add(
							partnersTVL[partner],
							formatWithPriceOnBlock(chainID, blockNumber, vaultAddress, delegatedBalance, thisVault.Decimals),
						)
						logs.Success(`Adding`, formatWithPriceOnBlock(chainID, blockNumber, vaultAddress, delegatedBalance, thisVault.Decimals).Text('f', 4), `for vault`, vaultAddress.Hex(), `for a total of`, value.Text('f', 4), `for partner`, partner.Hex())
					}
					partnersTVL[partner] = bigNumber.NewFloat(0).Add(
						partnersTVL[partner],
						formatWithPriceOnBlock(chainID, blockNumber, vaultAddress, delegatedBalance, thisVault.Decimals),
					)
					if _, ok := delegateValue[partner]; !ok {
						delegateValue[partner] = make(map[common.Address]map[common.Address]TWithNormalized)
					}
					if _, ok := delegateValue[partner][vaultAddress]; !ok {
						delegateValue[partner][vaultAddress] = make(map[common.Address]TWithNormalized)
					}
					if _, ok := delegateValue[partner][vaultAddress][depositorAddress]; !ok {
						delegateValue[partner][vaultAddress][depositorAddress] = TWithNormalized{}
					}

					delegateValue[partner][vaultAddress][depositorAddress] = TWithNormalized{
						Raw:        delegatedBalance,
						Normalized: toNormalizedAmount(delegatedBalance, thisVault.Decimals),
						Value:      formatWithPriceOnBlock(chainID, blockNumber, vaultAddress, delegatedBalance, thisVault.Decimals),
					}
				}
			}
		}
	}

	for partner, tvl := range partnersTVL {
		totalTVL[partner] = bigNumber.NewFloat(0).Clone(tvl)
		tierShare[partner] = bigNumber.NewFloat(0).Clone(partnerTVLTier(tvl))
	}
	return totalTVL, tierShare, delegateValue
}

type TBalanceSheet struct {
	VaultName      string
	BlockNumber    uint64
	VaultDecimals  uint64
	VaultAddress   common.Address
	PartnerAddress common.Address
	VaultTVL       *bigNumber.Int
	VaultFees      *bigNumber.Int
	PartnerTVL     *bigNumber.Int
	DelegatedTVL   map[common.Address]*bigNumber.Int
}

func getPartnerFees(chainID uint64) {
	fromBlock := uint64(16036123)
	toBlock := uint64(16794005)

	logs.Info(`Retrieving all harvests for chain: ` + strconv.Itoa(int(chainID)))
	allHarvestPerStrategyPerVault := harvests.RetrieveAllHarvestPerStrategyPerVault(chainID, fromBlock, nil)
	logs.Info(`Retrieving all partner tracker events for chain: ` + strconv.Itoa(int(chainID)))
	allPartnerTrackerEvents, allRefereesEvents := partnerTracker.RetrieveAllPartnerTrackerEvents(chainID)
	logs.Info(`Retrieving all vaults for chain: ` + strconv.Itoa(int(chainID)))
	allVaults := vaults.ListVaults(chainID)
	logs.Info(`Retrieving totalSupply for vaults for each harvest for chain: ` + strconv.Itoa(int(chainID)))
	vaultsSupplies := getVaultsSupplyAtBlock(allVaults, allHarvestPerStrategyPerVault)

	/**************************************************************************************************
	** This block of code is used to get the TVL for a partner right now
	**************************************************************************************************/
	// ledger := common.HexToAddress(`0x558247e365be655f9144e1a0140D793984372Ef3`) //Ledger
	// allEvents := partnerTracker.ListReferralBalanceIncrease(chainID)
	// vaultTVL := make(map[common.Address]*bigNumber.Int)
	// for _, event := range allEvents {
	// 	if event.PartnerID.Hex() == ledger.Hex() {
	// 		if _, ok := vaultTVL[event.Vault]; !ok {
	// 			vaultTVL[event.Vault] = bigNumber.NewInt(0)
	// 		}
	// 		vaultTVL[event.Vault] = bigNumber.NewInt(0).Add(vaultTVL[event.Vault], event.Amount)
	// 	}
	// }

	// totalValue := bigNumber.NewFloat(0)
	// for vault, tvl := range vaultTVL {
	// 	vault, _ := vaults.GetVault(chainID, vault)
	// 	totalValue = bigNumber.NewFloat(0).Add(totalValue, formatWithPrice(chainID, vault.Address, tvl, vault.Decimals))
	// }
	// logs.Success(`Total value for Ledger: $` + totalValue.Text('f', 4))
	/**************************************************************************************************
	** END_OF_BLOCK
	**************************************************************************************************/

	// partnerToTest := common.HexToAddress(`0xFEB4acf3df3cDEA7399794D0869ef76A6EfAff52`) //Yearn

	balanceSheets := make(map[common.Address][]TBalanceSheet)
	totalDelegatedValue := float64(0)
	totalFeesCollected := float64(0)
	for _, vault := range allVaults {
		vaultToTest := vault.Address
		allHarvestForVault := allHarvestPerStrategyPerVault[vaultToTest]
		for _, harvests := range allHarvestForVault {
			for _, harvest := range harvests {
				if (harvest.Fees.TreasuryCollectedFee).IsZero() { //Skip, no gain or loss, no fees
					continue
				}
				partnerTVL, delegatedTVL := computePartnerVaultTVL(
					harvest.BlockNumber,
					vault.Address,
					vault.Decimals,
					allPartnerTrackerEvents[vault.Address],
					allRefereesEvents[vault.Address],
				)
				for partnerAddress, partner := range partnerTVL {
					if _, ok := balanceSheets[partnerAddress]; !ok {
						balanceSheets[partnerAddress] = []TBalanceSheet{}
					}
					newItem := TBalanceSheet{
						PartnerAddress: partnerAddress,
						VaultAddress:   vault.Address,
						VaultTVL:       bigNumber.NewInt(0).Set(vaultsSupplies[vault.Address][harvest.BlockNumber]),
						VaultFees:      bigNumber.NewInt(0).Clone(harvest.Fees.TreasuryCollectedFee),
						PartnerTVL:     partner,
						VaultName:      vault.DisplaySymbol,
						BlockNumber:    harvest.BlockNumber,
						VaultDecimals:  vault.Decimals,
						DelegatedTVL:   delegatedTVL[partnerAddress],
					}
					balanceSheets[partnerAddress] = append(balanceSheets[partnerAddress], newItem)
				}
			}
		}
	}

	CSVExport := []TMainCSVExport{}
	CSVExportBreakdown := []TPartnerBreakdownCSVExport{}
	for _, sheets := range balanceSheets {
		for _, sheet := range sheets {
			if sheet.BlockNumber < fromBlock || sheet.BlockNumber > toBlock {
				continue //Skip, not in the range
			}
			if sheet.PartnerTVL.Cmp(big.NewInt(0)) == 0 {
				continue //Skip, no balance in the vault
			}
			if sheet.PartnerAddress.Hex() != LEDGER.Hex() {
				continue //Skip, not ledger
			}

			partnersTvl, tierShare, delegateValue := computePartnerTotalTVL(
				chainID,
				sheet.BlockNumber,
				allPartnerTrackerEvents,
				allRefereesEvents,
			)

			tier := bigNumber.NewFloat(0)
			partnerTVL := bigNumber.NewFloat(0)
			if _, ok := tierShare[sheet.PartnerAddress]; ok {
				tier = tierShare[sheet.PartnerAddress]
				partnerTVL = partnersTvl[sheet.PartnerAddress]
			}

			prices := prices.FetchPricesOnBlock(chainID, sheet.BlockNumber, []common.Address{sheet.VaultAddress})
			price := prices[sheet.VaultAddress]
			PartnerFeeShareBase := bigNumber.NewFloat(0).Mul(
				toNormalizedAmount(sheet.VaultFees, sheet.VaultDecimals),
				bigNumber.NewFloat(0).Div(
					toNormalizedRatio(sheet.PartnerTVL, sheet.VaultTVL, sheet.VaultDecimals),
					bigNumber.NewFloat(100),
				),
			)
			PartnerFeeShareTier := bigNumber.NewFloat(0).Mul(
				bigNumber.NewFloat(0).Mul(tier, PartnerFeeShareBase),
				bigNumber.NewFloat(0.55), //OPEX
			)
			CSVExport = append(CSVExport, TMainCSVExport{
				Block:      sheet.BlockNumber,
				Timestamp:  0,
				Wrapper:    sheet.PartnerAddress.Hex(),                 //Partner address
				Vault:      sheet.VaultAddress.Hex(),                   //Vault address
				Tier:       tier.Text('f', -1),                         //Tier share
				VaultPrice: toNormalizedAmount(price, 6).Text('f', -1), //Vault price

				//
				Balance:    toNormalizedAmount(sheet.PartnerTVL, sheet.VaultDecimals).Text('f', -1),
				BalanceUSD: formatWithPriceOnBlock(chainID, sheet.BlockNumber, sheet.VaultAddress, sheet.PartnerTVL, sheet.VaultDecimals).Text('f', -1),
				//
				TotalSupply:    toNormalizedAmount(sheet.VaultTVL, sheet.VaultDecimals).Text('f', -1),
				TotalSupplyUSD: formatWithPriceOnBlock(chainID, sheet.BlockNumber, sheet.VaultAddress, sheet.VaultTVL, sheet.VaultDecimals).Text('f', -1),
				//
				TotalDelegatedValue: partnerTVL.Text('f', -1),
				//
				// Share:               toNormalizedRatio(sheet.PartnerTVL, sheet.VaultTVL, sheet.VaultDecimals).Text('f', -1),                                      //Share of the partner in the vault
				// PartnerFeeShareBase: PartnerFeeShareBase.Text('f', -1),                                                                                           //PartnerFeeShareBase = fees * share
				CollectedFee:       toNormalizedAmount(sheet.VaultFees, sheet.VaultDecimals).Text('f', -1), //Fees collected by Yearn
				CollectedFeeUSD:    formatWithPriceOnBlock(chainID, sheet.BlockNumber, sheet.VaultAddress, sheet.VaultFees, sheet.VaultDecimals).Text('f', -1),
				PartnerFeeShare:    PartnerFeeShareTier.Text('f', -1), //PartnerFeeShare = fees * share * tier
				PartnerFeeShareUSD: bigNumber.NewFloat(0).Mul(PartnerFeeShareTier, toNormalizedAmount(price, 6)).Text('f', -1),
			})

			logs.Info(`------------------------------------------------------------------------------------------------`)
			logs.Info(`Vault `+sheet.VaultAddress.Hex()+` - `+sheet.VaultName+` at block`, sheet.BlockNumber)
			logs.Info(`Partner ` + sheet.PartnerAddress.Hex())

			normalizedTVL := toNormalizedAmount(sheet.VaultTVL, sheet.VaultDecimals)
			normalizedTVLValue := formatWithPriceOnBlock(chainID, sheet.BlockNumber, sheet.VaultAddress, sheet.VaultTVL, sheet.VaultDecimals)
			logs.Info(`TVL ` + normalizedTVL.Text('f', -1) + ` ($` + normalizedTVLValue.Text('f', 4) + `)`)

			normalizedFees := toNormalizedAmount(sheet.VaultFees, sheet.VaultDecimals)
			normalizedFeesValue := formatWithPriceOnBlock(chainID, sheet.BlockNumber, sheet.VaultAddress, sheet.VaultFees, sheet.VaultDecimals)
			logs.Info(`Fees ` + normalizedFees.Text('f', -1) + ` ($` + normalizedFeesValue.Text('f', 4) + `)`)

			normalizedPartnerTVL := toNormalizedAmount(sheet.PartnerTVL, sheet.VaultDecimals)
			normalizedPartnerTVLValue := formatWithPriceOnBlock(chainID, sheet.BlockNumber, sheet.VaultAddress, sheet.PartnerTVL, sheet.VaultDecimals)
			normalizedPartnerRatio := toNormalizedRatio(sheet.PartnerTVL, sheet.VaultTVL, sheet.VaultDecimals)
			logs.Info(`Partner TVL ` + normalizedPartnerTVL.Text('f', -1) + ` ($` + normalizedPartnerTVLValue.Text('f', 4) + `) | (` + normalizedPartnerRatio.Text('f', 2) + `%)`)

			normalizedPartnerFees := bigNumber.NewFloat(0).Mul(toNormalizedAmount(sheet.VaultFees, sheet.VaultDecimals), tier)
			normalizedPartnerFeesValue := bigNumber.NewFloat(0).Mul(formatWithPriceOnBlock(chainID, sheet.BlockNumber, sheet.VaultAddress, sheet.VaultFees, sheet.VaultDecimals), tier)
			logs.Info(`Partner Fees ` + normalizedPartnerFees.Text('f', -1) + ` ($` + normalizedPartnerFeesValue.Text('f', 4) + `)`)

			logs.Info(`Delegated TVL`)
			for vault, depositorsForThatVault := range delegateValue[sheet.PartnerAddress] {
				for depositor, delegatedBalance := range depositorsForThatVault {
					logs.Info("- " + depositor.Hex() + `: ` + delegatedBalance.Normalized.Text('f', -1) + ` ($` + delegatedBalance.Value.Text('f', 4) + `) | (` + toNormalizedRatio(delegatedBalance.Raw, sheet.VaultTVL, sheet.VaultDecimals).Text('f', 2) + `%)`)

					CSVExportBreakdown = append(CSVExportBreakdown, TPartnerBreakdownCSVExport{
						Block:      sheet.BlockNumber,
						Wrapper:    sheet.PartnerAddress.Hex(),
						Vault:      vault.Hex(),
						Depositor:  depositor.Hex(),
						Balance:    delegatedBalance.Normalized.Text('f', -1),
						BalanceUSD: delegatedBalance.Value.Text('f', -1),
						Share:      toNormalizedRatio(delegatedBalance.Raw, sheet.PartnerTVL, sheet.VaultDecimals).Text('f', -1),
					})
				}
			}
			logs.Info(`------------------------------------------------------------------------------------------------`)
			fmt.Println()
		}

		//save in a csv file
		outputFile, err := os.OpenFile("output.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()
		if err := gocsv.MarshalFile(&CSVExport, outputFile); err != nil {
			panic(err)
		}

		outputBreakdownFile, err := os.OpenFile("outputBreakdown.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer outputBreakdownFile.Close()
		if err := gocsv.MarshalFile(&CSVExportBreakdown, outputBreakdownFile); err != nil {
			panic(err)
		}

	}

	totalFeesCollectedStr := bigNumber.NewFloat(totalFeesCollected).Text('f', -1)
	totalDelegatedValueStr := bigNumber.NewFloat(totalDelegatedValue).Text('f', -1)
	logs.Success(`Total fees collected are`, totalFeesCollectedStr, `over`, totalDelegatedValueStr, `delegated`)
}

func Run(chainID uint64) {
	loadVaultAndStrategies(chainID)
	getPartnerFees(chainID)
}

func partnerTVLTier(value *bigNumber.Float) *bigNumber.Float {
	if value.Lt(bigNumber.NewFloat(1_000_000)) {
		return bigNumber.NewFloat(0)
	}
	if value.Lt(bigNumber.NewFloat(5_000_000)) {
		return bigNumber.NewFloat(0.1)
	}
	if value.Lt(bigNumber.NewFloat(10_000_000)) {
		return bigNumber.NewFloat(0.15)
	}
	if value.Lt(bigNumber.NewFloat(50_000_000)) {
		return bigNumber.NewFloat(0.2)
	}
	if value.Lt(bigNumber.NewFloat(100_000_000)) {
		return bigNumber.NewFloat(0.25)
	}
	if value.Lt(bigNumber.NewFloat(200_000_000)) {
		return bigNumber.NewFloat(0.3)
	}
	if value.Lt(bigNumber.NewFloat(400_000_000)) {
		return bigNumber.NewFloat(0.35)
	}
	if value.Lt(bigNumber.NewFloat(700_000_000)) {
		return bigNumber.NewFloat(0.4)
	}
	if value.Lt(bigNumber.NewFloat(1_000_000_000)) {
		return bigNumber.NewFloat(0.45)
	}
	return bigNumber.NewFloat(0.5)

}

func formatWithPriceOnBlock(chainID uint64, blockNumber uint64, tokenAddress common.Address, amount *bigNumber.Int, decimals uint64) *bigNumber.Float {
	price, ok := prices.FindPriceOnBlock(chainID, blockNumber, tokenAddress)
	if !ok {
		logs.Warning(`No price found for`, tokenAddress.Hex(), `on block`, blockNumber)
	}
	return toNormalizedValue(amount, price, decimals)
}
func toNormalizedAmount(amount *bigNumber.Int, decimals uint64) *bigNumber.Float {
	return bigNumber.NewFloat(0).Quo(
		bigNumber.NewFloat(0).SetInt(amount),
		bigNumber.NewFloat(0).SetInt(
			bigNumber.NewInt(0).Exp(bigNumber.NewInt(10), bigNumber.NewInt(int64(decimals)), nil),
		),
	)
}
func toNormalizedRatio(amountA *bigNumber.Int, amountB *bigNumber.Int, decimals uint64) *bigNumber.Float {
	normalizedA := toNormalizedAmount(amountA, decimals)
	normalizedB := toNormalizedAmount(amountB, decimals)
	normalizedRatio := bigNumber.NewFloat(0).Quo(normalizedA, normalizedB)
	percent := bigNumber.NewFloat(0).Mul(normalizedRatio, bigNumber.NewFloat(100))
	return percent
}
func toNormalizedValue(amount *bigNumber.Int, price *bigNumber.Int, decimals uint64) *bigNumber.Float {
	normalizedAmount := toNormalizedAmount(amount, decimals)
	normalizedPrice := toNormalizedAmount(price, 6)
	normalizedValue := bigNumber.NewFloat(0).Mul(normalizedAmount, normalizedPrice)
	return normalizedValue
}
