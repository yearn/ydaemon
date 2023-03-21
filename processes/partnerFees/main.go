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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/harvests"
	"github.com/yearn/ydaemon/internal/indexer"
	partnerTracker "github.com/yearn/ydaemon/internal/indexer.partnerTracker"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/registries"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
)

var AllHarvests = make(map[ethcommon.Address][]vaults.THarvest)
var STRATLIST = []strategies.TStrategy{}
var LEDGER = common.HexToAddress(`0x558247e365be655f9144e1a0140D793984372Ef3`).ToAddress()

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
	allHarvests map[ethcommon.Address]map[ethcommon.Address][]harvests.THarvest,
) map[ethcommon.Address]map[uint64]*big.Int {
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

	vaultsSupplyAtBlock := make(map[ethcommon.Address]map[uint64]*big.Int)
	syncMap.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddress := ethcommon.HexToAddress(eventKey[0])
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
	harvest harvests.THarvest,
	currentVault ethcommon.Address,
	vaultDecimals uint64,
	partnersTrackerEvents map[ethcommon.Address]map[ethcommon.Address][]partnerTracker.TEventReferredBalanceIncreased,
	allRefereesEvents map[ethcommon.Address][]partnerTracker.TRefereeTransferEvent,
) (map[ethcommon.Address]*bigNumber.Int, map[ethcommon.Address]map[ethcommon.Address]*bigNumber.Int) {
	delegatedTVL := make(map[ethcommon.Address]map[ethcommon.Address]*bigNumber.Int)
	partnerTVL := make(map[ethcommon.Address]*bigNumber.Int)

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
		harvest.BlockNumber,
	)

	/**********************************************************************************************
	** We have the list of deposits that happened before the harvest. Now, we want to check the
	** balance of the depositor in the vault before the harvest to see if he had other
	** deposits that were not done by the partner, or if he had done some withdrawals.
	**********************************************************************************************/
	for partner, eventsDepositorLevel := range relevantPartnerTrackerEvents {
		if _, ok := delegatedTVL[partner]; !ok {
			delegatedTVL[partner] = make(map[ethcommon.Address]*bigNumber.Int)
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
			userDepositViaPartnerMap := map[ethcommon.Hash]partnerTracker.TEventReferredBalanceIncreased{}
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
				if transfer.BlockNumber > harvest.BlockNumber {
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
**************************************************************************************************/
func computePartnerTotalTVL(
	chainID uint64,
	blockNumber uint64,
	partnersTrackerEvents map[ethcommon.Address]map[ethcommon.Address]map[ethcommon.Address][]partnerTracker.TEventReferredBalanceIncreased,
	allRefereesEvents map[ethcommon.Address]map[ethcommon.Address][]partnerTracker.TRefereeTransferEvent,
) map[ethcommon.Address]map[ethcommon.Address]map[uint64]*bigNumber.Int {
	// cumulatedPartnerTVL -> map[vaultAddress][partnerAddress][blockNumber][TVLAtThatBlock]
	cumulatedPartnerTVL := make(map[ethcommon.Address]map[ethcommon.Address]map[uint64]*bigNumber.Int)
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
	partnerVaults := make(map[ethcommon.Address]map[ethcommon.Address][]partnerTracker.TEventReferredBalanceIncreased)
	eventsHash := make(map[ethcommon.Address]map[ethcommon.Hash]partnerTracker.TEventReferredBalanceIncreased)
	for vaultAddress, eventsForThatVault := range relevantPartnerTrackerEvents {
		for partnerAddress, eventsForThanPartner := range eventsForThatVault {
			for _, events := range eventsForThanPartner {
				if _, ok := partnerVaults[partnerAddress]; !ok {
					partnerVaults[partnerAddress] = make(map[ethcommon.Address][]partnerTracker.TEventReferredBalanceIncreased)
					eventsHash[partnerAddress] = make(map[ethcommon.Hash]partnerTracker.TEventReferredBalanceIncreased)
				}
				partnerVaults[partnerAddress][vaultAddress] = append(partnerVaults[partnerAddress][vaultAddress], events...)
				for _, event := range events {
					if _, ok := eventsHash[vaultAddress]; !ok {
						eventsHash[vaultAddress] = make(map[ethcommon.Hash]partnerTracker.TEventReferredBalanceIncreased)
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
	** delegatedVaultPartnerBalance -> map[vaultAddress][partnerAddress][delegatedBalance]
	**********************************************************************************************/
	delegatedVaultPartnerBalance := make(map[ethcommon.Address]map[ethcommon.Address]*bigNumber.Int)
	for vaultAddress, eventsForThatVault := range allRefereesEvents {
		vault, _ := vaults.GetVault(chainID, common.FromAddress(vaultAddress))

		for depositorAddress, eventForThisDepositor := range eventsForThatVault {
			depositorDelegatedBalance := make(map[ethcommon.Address]*bigNumber.Int)
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
							partnerAddress := partnerTrackerEvent.PartnerID.ToAddress()
							if _, ok := depositorDelegatedBalance[partnerAddress]; !ok {
								depositorDelegatedBalance[partnerAddress] = bigNumber.NewInt(0)
							}
							depositorDelegatedBalance[partnerAddress] = bigNumber.NewInt(0).Add(depositorDelegatedBalance[partnerAddress], event.Value)
							depositorActualBalance = bigNumber.NewInt(0).Add(depositorActualBalance, event.Value)
							if partnerAddress.Hex() == LEDGER.Hex() {
								logs.Success("Ledger: +", formatWithPriceOnBlock(chainID, event.BlockNumber, common.FromAddress(vault.Address), event.Value, vault.Decimals))
							}

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
						if partnerAddress.Hex() == LEDGER.Hex() {
							logs.Success("Ledger: -", formatWithPriceOnBlock(chainID, event.BlockNumber, common.FromAddress(vault.Address), delegatedWithdraw.Int(), vault.Decimals))
						}

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
					delegatedVaultPartnerBalance[vaultAddress] = make(map[ethcommon.Address]*bigNumber.Int)
				}
				if _, ok := delegatedVaultPartnerBalance[vaultAddress][partnerAddress]; !ok {
					delegatedVaultPartnerBalance[vaultAddress][partnerAddress] = bigNumber.NewInt(0)
				}
				delegatedVaultPartnerBalance[vaultAddress][partnerAddress] = bigNumber.NewInt(0).Add(
					delegatedVaultPartnerBalance[vaultAddress][partnerAddress],
					delegatedBalance,
				)
			}
		}
	}

	/**********************************************************************************************
	** Now that we got that, we want to know, for a given partner, how much TVL he has in total.
	** As we have the delegated balance for each vault, we just need to grab this balance, the
	** vault token price and sum it up.
	**********************************************************************************************/
	partnersTVL := make(map[ethcommon.Address]*bigNumber.Float)
	for vaultAddress, partnersForThatVault := range delegatedVaultPartnerBalance {
		for partner, delegatedBalance := range partnersForThatVault {
			if _, ok := partnersTVL[partner]; !ok {
				partnersTVL[partner] = bigNumber.NewFloat(0)
			}
			thisVault, _ := vaults.GetVault(chainID, common.FromAddress(vaultAddress))

			partnersTVL[partner] = bigNumber.NewFloat(0).Add(
				partnersTVL[partner],
				formatWithPrice(
					chainID,
					common.FromAddress(vaultAddress),
					delegatedBalance,
					thisVault.Decimals,
				),
			)
		}
	}

	for partner, tvl := range partnersTVL {
		logs.Success(`Partner: ` + partner.Hex() + ` has a TVL of $` + tvl.Text('f', 4) + ` and has a PartnerTier of ` + partnerTVLTier(tvl).Text('f', 2) + `%`)
	}

	return cumulatedPartnerTVL
}

type TBalanceSheet struct {
	VaultName      string
	BlockNumber    uint64
	VaultDecimals  uint64
	VaultAddress   ethcommon.Address
	PartnerAddress ethcommon.Address
	VaultTVL       *bigNumber.Int
	VaultFees      *bigNumber.Int
	PartnerTVL     *bigNumber.Int
	DelegatedTVL   map[ethcommon.Address]*bigNumber.Int
}

func getPartnerFees(chainID uint64) {
	fromBlock := uint64(0)
	toBlock := uint64(16391125)

	_ = toBlock

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
	// ledger := common.HexToAddress(`0x558247e365be655f9144e1a0140D793984372Ef3`).ToAddress() //Ledger
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
	// 	totalValue = bigNumber.NewFloat(0).Add(totalValue, formatWithPrice(chainID, common.FromAddress(vault.Address), tvl, vault.Decimals))
	// }
	// logs.Success(`Total value for Ledger: $` + totalValue.Text('f', 4))
	/**************************************************************************************************
	** END_OF_BLOCK
	**************************************************************************************************/

	// partnerToTest := common.HexToAddress(`0xFEB4acf3df3cDEA7399794D0869ef76A6EfAff52`).ToAddress() //Yearn

	balanceSheets := make(map[ethcommon.Address][]TBalanceSheet)
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
					harvest,
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

	computePartnerTotalTVL(
		chainID,
		toBlock,
		allPartnerTrackerEvents,
		allRefereesEvents,
	)

	os.Exit(0)

	for _, sheets := range balanceSheets {
		for _, sheet := range sheets {
			if sheet.PartnerTVL.Cmp(big.NewInt(0)) == 0 {
				continue //Skip, no balance in the vault
			}
			if sheet.PartnerAddress.Hex() != LEDGER.Hex() {
				continue //Skip, not ledger
			}
			logs.Info(`------------------------------------------------------------------------------------------------`)
			logs.Info(`Vault `+sheet.VaultAddress.Hex()+` - `+sheet.VaultName+` at block`, sheet.BlockNumber)
			logs.Info(`Partner ` + sheet.PartnerAddress.Hex())
			logs.Info(`TVL ` + toNormalizedAmount(sheet.VaultTVL, sheet.VaultDecimals).Text('f', -1) + ` ($` + formatWithPrice(chainID, common.FromAddress(sheet.VaultAddress), sheet.VaultTVL, sheet.VaultDecimals).Text('f', 4) + `)`)
			logs.Info(`Fees ` + toNormalizedAmount(sheet.VaultFees, sheet.VaultDecimals).Text('f', -1) + ` ($` + formatWithPrice(chainID, common.FromAddress(sheet.VaultAddress), sheet.VaultFees, sheet.VaultDecimals).Text('f', 4) + `)`)
			logs.Info(`Partner TVL ` + toNormalizedAmount(sheet.PartnerTVL, sheet.VaultDecimals).Text('f', -1) + ` ($` + formatWithPrice(chainID, common.FromAddress(sheet.VaultAddress), sheet.PartnerTVL, sheet.VaultDecimals).Text('f', 4) + `) | (` + toNormalizedRatio(sheet.PartnerTVL, sheet.VaultTVL, sheet.VaultDecimals).Text('f', 2) + `%)`)
			logs.Info(`Delegated TVL`)
			for depositor, delegatedBalance := range sheet.DelegatedTVL {
				logs.Info("- " + depositor.Hex() + `: ` + toNormalizedAmount(delegatedBalance, sheet.VaultDecimals).Text('f', -1) + ` ($` + formatWithPrice(chainID, common.FromAddress(sheet.VaultAddress), delegatedBalance, sheet.VaultDecimals).Text('f', 4) + `) | (` + toNormalizedRatio(delegatedBalance, sheet.VaultTVL, sheet.VaultDecimals).Text('f', 2) + `%)`)

				// allEventsBeforeThisBlock := partnerTracker.FilterRefereeTransferEvent(allRefereesEvents, sheet.BlockNumber)
				allEventsForThisDepositor := partnerTracker.GroupRefereeTransferEventPerVault(allRefereesEvents[sheet.VaultAddress][depositor], sheet.BlockNumber)
				for vaultAddress, events := range allEventsForThisDepositor {
					thisVault, _ := vaults.GetVault(chainID, common.FromAddress(vaultAddress))
					for _, event := range events {
						if event.Referee == event.From {
							logs.Success(`-- ` + thisVault.DisplaySymbol + `: -` + toNormalizedAmount(event.Value, sheet.VaultDecimals).Text('f', -1) + ` ($` + formatWithPrice(chainID, common.FromAddress(sheet.VaultAddress), event.Value, sheet.VaultDecimals).Text('f', 4) + `)`)
						} else {
							logs.Success(`-- ` + thisVault.DisplaySymbol + `: ` + toNormalizedAmount(event.Value, sheet.VaultDecimals).Text('f', -1) + ` ($` + formatWithPrice(chainID, common.FromAddress(sheet.VaultAddress), event.Value, sheet.VaultDecimals).Text('f', 4) + `)`)
						}
					}
				}
			}
			// logs.Success(`Total delegated to ` + sheet.PartnerAddress.Hex() + `: $` + totalDelegatedToPartnerAtBlock[sheet.PartnerAddress].Text('f', 4))
			// for partnerAddress, delegatedBalance := range totalDelegatedToPartnerAtBlock {
			// 	logs.Info(`Total delegated to ` + partnerAddress.Hex() + `: $` + delegatedBalance.Text('f', 4))
			// }
			logs.Info(`------------------------------------------------------------------------------------------------`)
			fmt.Println()
			// os.Exit(0)
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

// UTILS FUNCTION
func formatWithPrice(chainID uint64, tokenAddress common.Address, amount *bigNumber.Int, decimals uint64) *bigNumber.Float {
	price, ok := prices.FindPrice(chainID, tokenAddress)
	if ok {
		return toNormalizedValue(amount, price, decimals)
	}
	return toNormalizedValue(amount, bigNumber.NewInt(0), 1)
}
func formatWithPriceOnBlock(chainID uint64, blockNumber uint64, tokenAddress common.Address, amount *bigNumber.Int, decimals uint64) *bigNumber.Float {
	allPrices := prices.FetchPricesOnBlock(chainID, blockNumber, []common.Address{tokenAddress})
	price := allPrices[tokenAddress]
	if price == nil || price.IsZero() {
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
