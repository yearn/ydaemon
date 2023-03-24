package partnerFees

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	partnerTracker "github.com/yearn/ydaemon/internal/indexer.partnerTracker"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/vaults"
)

/**************************************************************************************************
** computeBlockData will compute the data for each partner at the given block based on the events
** we have collected.
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
** - partnerTierRatio: The share of each partner in the total TVL.
** - partnerDelegatedBalanceBreakdown: The amount each depositor delegated to each partner.
**************************************************************************************************/
func computeBlockData(
	chainID uint64,
	blockNumber uint64,
	partnersTrackerEvents map[common.Address]map[common.Address]map[common.Address][]partnerTracker.TEventReferredBalanceIncreased,
	allRefereesEvents map[common.Address]map[common.Address][]partnerTracker.TRefereeTransferEvent,
) (
	partnerTotalTVL map[common.Address]*bigNumber.Float,
	partnerTierRatio map[common.Address]*bigNumber.Float,
	vaultDelegatedAmountForPartner map[common.Address]map[common.Address]*bigNumber.Int,
	partnerDelegatedBalanceBreakdown map[common.Address]map[common.Address]map[common.Address]TBig,
) {
	// map -> [partnerAddress][totalAmountDelegated]
	partnerTotalTVL = make(map[common.Address]*bigNumber.Float)

	// map -> [partnerAddress][partnerTierRatio]
	partnerTierRatio = make(map[common.Address]*bigNumber.Float)

	// map -> [partnerAddress][vaultAddress][depositorAddress][TBig]
	partnerDelegatedBalanceBreakdown = make(map[common.Address]map[common.Address]map[common.Address]TBig)

	// map -> [vaultAddress][partnerAddress][amountDelegated]
	vaultDelegatedAmountForPartner = make(map[common.Address]map[common.Address]*bigNumber.Int)

	// map -> [vaultAddress][eventHash][events]
	partnerDelegateIncreaseEvent := make(map[common.Address]map[common.Hash]partnerTracker.TEventReferredBalanceIncreased)

	/**********************************************************************************************
	** First we sort all the events by block number in order to compute the TVL at each block.
	** Here, we got ALL the events from the start block to the end block, no matter the vault or
	** the partner.
	**********************************************************************************************/
	partnerTrackerEvents := partnerTracker.SortReferralBalanceIncreaseEvents(partnersTrackerEvents)

	/**********************************************************************************************
	** We want to know, for every harvest that happened impacting the partner:
	** - if a partner has a delegated balance in any vault
	** - the amount of delegated balance in each of theses vaults.
	**
	** We will then create a map that will contain each event mapped to it's hash. This will allow
	** us to quickly check if a transfer event is a deposit or not.
	** partnerDelegateIncreaseEvent -> [partnerAddress][vaultAddress][eventHash][event]
	**********************************************************************************************/
	for vaultAddress, eventsForThatVault := range partnerTrackerEvents {
		for partnerAddress, eventsForThatPartner := range eventsForThatVault {
			/**************************************************************************************
			** We need to init the vaultDelegatedAmountForPartner map to avoid nil assignment
			**************************************************************************************/
			if _, ok := vaultDelegatedAmountForPartner[vaultAddress]; !ok {
				vaultDelegatedAmountForPartner[vaultAddress] = make(map[common.Address]*bigNumber.Int)
			}
			if _, ok := vaultDelegatedAmountForPartner[vaultAddress][partnerAddress]; !ok {
				vaultDelegatedAmountForPartner[vaultAddress][partnerAddress] = bigNumber.NewInt(0)
			}

			/**************************************************************************************
			** For each of the events, we will assign the event in a map that will allow us to
			** quickly check if a transfer event is a deposit from the partner contract or not:
			** If we find a the hash of the transfer event we will be checking matches one hash
			** that we retrieved in the partnerDelegateIncreaseEvent map, then we know that this
			** transfer event is a deposit from the partner contract and we can delegate the
			** amount deposited to the corresponding partner.
			**************************************************************************************/
			for _, events := range eventsForThatPartner {
				/**********************************************************************************
				** We need to init the partnerDelegateIncreaseEvent map to avoid nil assignment.
				**********************************************************************************/
				if _, ok := partnerDelegateIncreaseEvent[vaultAddress]; !ok {
					partnerDelegateIncreaseEvent[vaultAddress] = make(map[common.Hash]partnerTracker.TEventReferredBalanceIncreased)
				}

				/**********************************************************************************
				** Process all the events for that partner/vault/depositor and assign them to the
				** corresponding maps.
				**********************************************************************************/
				for _, event := range events {
					partnerDelegateIncreaseEvent[vaultAddress][event.TxHash] = event
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
			/**************************************************************************************
			** We will proceed all the events for that depositor on that vault. We will need to
			** store two things across all the events:
			** - The actual balance of the depositor in the vault, aka how much he has
			** - The delegated balance of the depositor in the vault, aka how much he has delegated
			**   to each partner.
			** Theses two elements will allow us to know, at any point in time, the ratio of the
			** delegated balance to the actual balance of the depositor.
			**************************************************************************************/
			depositorDelegatedBalance := make(map[common.Address]*bigNumber.Int)
			depositorActualBalance := bigNumber.NewInt(0)

			/**************************************************************************************
			** For each transfer event, we check the step by step history to compute the delegated
			** balances and the actual balance.
			**************************************************************************************/
			for _, transfer := range eventForThisDepositor {
				if transfer.BlockNumber > blockNumber {
					continue
				}
				isReceiving := addresses.Equals(transfer.To, depositorAddress)
				isSending := addresses.Equals(transfer.From, depositorAddress)

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
					if _, ok := partnerDelegateIncreaseEvent[vaultAddress]; ok {
						if partnerTrackerEvent, ok := partnerDelegateIncreaseEvent[vaultAddress][transfer.TxHash]; ok {
							/**********************************************************************
							** We have a match, this is a deposit via the partner tracker.
							**********************************************************************/
							partnerAddress := partnerTrackerEvent.PartnerID
							if _, ok := depositorDelegatedBalance[partnerAddress]; !ok {
								depositorDelegatedBalance[partnerAddress] = bigNumber.NewInt(0)
							}
							depositorDelegatedBalance[partnerAddress] = bigNumber.NewInt(0).Add(depositorDelegatedBalance[partnerAddress], transfer.Value)
						}
					}

					/******************************************************************************
					** If we are here, it means that this is not a deposit via the partner tracker.
					** We need to update the actual balance of the user in the vault without
					** changing the delegated balance.
					******************************************************************************/
					depositorActualBalance = bigNumber.NewInt(0).Add(depositorActualBalance, transfer.Value)
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
					depositorActualBalance = bigNumber.NewInt(0).Sub(depositorActualBalance, transfer.Value)
					for partnerAddress, delegatedBalance := range depositorDelegatedBalance {
						delegatedRatio := bigNumber.NewFloat(0).Div(
							bigNumber.NewFloat().SetInt(delegatedBalance),
							bigNumber.NewFloat().SetInt(depositorActualBalanceBeforeEvent),
						)
						delegatedWithdraw := bigNumber.NewFloat(0).Mul(
							delegatedRatio,
							bigNumber.NewFloat(0).SetInt(transfer.Value),
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

			/**************************************************************************************
			** If the depositor actual balance is negative or zero, this means that this depositor
			** has withdrawn all his funds from the vault and thus we can ignore him and all the
			** partners he has delegated to.
			**************************************************************************************/
			if depositorActualBalance.Lte(bigNumber.NewInt(0)) {
				continue
			}

			/**************************************************************************************
			** For each partners with a delegated balance, we can add a new item in the mapping to
			** store that for this specific vault, this specific partner has a delegated balance
			** increased by X for this specific depositor.
			**************************************************************************************/
			for partnerAddress, delegatedBalance := range depositorDelegatedBalance {
				/**********************************************************************************
				** We need to init the delegatedVaultPartnerBalance maps to avoid nil assignment
				**********************************************************************************/
				if _, ok := delegatedVaultPartnerBalance[vaultAddress]; !ok {
					delegatedVaultPartnerBalance[vaultAddress] = make(map[common.Address]map[common.Address]*bigNumber.Int)
				}
				if _, ok := delegatedVaultPartnerBalance[vaultAddress][partnerAddress]; !ok {
					delegatedVaultPartnerBalance[vaultAddress][partnerAddress] = make(map[common.Address]*bigNumber.Int)
				}
				if _, ok := delegatedVaultPartnerBalance[vaultAddress][partnerAddress][depositorAddress]; !ok {
					delegatedVaultPartnerBalance[vaultAddress][partnerAddress][depositorAddress] = bigNumber.NewInt(0)
				}

				/**********************************************************************************
				** And we can now just store the delegated balance for this specific vault, partner
				** and depositor, adding the new balance to the previous one.
				**********************************************************************************/
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
	tvlForPartner := make(map[common.Address]*bigNumber.Float)
	vaultsWithTVL := []common.Address{}
	for vaultAddress, partnersForThatVault := range delegatedVaultPartnerBalance {
		for partner, depositorsForThatPartner := range partnersForThatVault {
			for _, delegatedBalance := range depositorsForThatPartner {
				if _, ok := tvlForPartner[partner]; !ok {
					tvlForPartner[partner] = bigNumber.NewFloat(0)
				}
				if delegatedBalance.Lte(bigNumber.NewInt(0)) {
					continue
				}
				thisVault, _ := vaults.GetVault(chainID, vaultAddress)
				vaultsWithTVL = append(vaultsWithTVL, thisVault.Token.Address)
			}
		}
	}
	prices.FetchPricesOnBlock(chainID, blockNumber, vaultsWithTVL)

	/**********************************************************************************************
	** Then, we can just add the TVL for each partner.
	**********************************************************************************************/
	for vaultAddress, partnersForThatVault := range delegatedVaultPartnerBalance {
		for partnerAddress, depositorsForThatPartner := range partnersForThatVault {
			for depositorAddress, delegatedBalance := range depositorsForThatPartner {
				if _, ok := tvlForPartner[partnerAddress]; !ok {
					tvlForPartner[partnerAddress] = bigNumber.NewFloat(0)
				}
				if delegatedBalance.Gt(bigNumber.NewInt(0)) {
					thisVault, _ := vaults.GetVault(chainID, vaultAddress)
					tvlForPartner[partnerAddress] = bigNumber.NewFloat(0).Add(
						tvlForPartner[partnerAddress],
						formatWithPriceOnBlock(chainID, blockNumber, thisVault.Token.Address, delegatedBalance, thisVault.Decimals),
					)
					vaultDelegatedAmountForPartner[vaultAddress][partnerAddress] = bigNumber.NewInt(0).Add(
						vaultDelegatedAmountForPartner[vaultAddress][partnerAddress],
						delegatedBalance,
					)

					if _, ok := partnerDelegatedBalanceBreakdown[partnerAddress]; !ok {
						partnerDelegatedBalanceBreakdown[partnerAddress] = make(map[common.Address]map[common.Address]TBig)
					}
					if _, ok := partnerDelegatedBalanceBreakdown[partnerAddress][vaultAddress]; !ok {
						partnerDelegatedBalanceBreakdown[partnerAddress][vaultAddress] = make(map[common.Address]TBig)
					}
					if _, ok := partnerDelegatedBalanceBreakdown[partnerAddress][vaultAddress][depositorAddress]; !ok {
						partnerDelegatedBalanceBreakdown[partnerAddress][vaultAddress][depositorAddress] = TBig{}
					}

					partnerDelegatedBalanceBreakdown[partnerAddress][vaultAddress][depositorAddress] = TBig{
						Raw:        delegatedBalance,
						Normalized: toNormalizedAmount(delegatedBalance, thisVault.Decimals),
						Value:      formatWithPriceOnBlock(chainID, blockNumber, thisVault.Token.Address, delegatedBalance, thisVault.Decimals),
					}
				}
			}
		}
	}

	for partner, tvl := range tvlForPartner {
		partnerTotalTVL[partner] = bigNumber.NewFloat(0).Clone(tvl)
		partnerTierRatio[partner] = bigNumber.NewFloat(0).Clone(partnerTVLTier(tvl))
	}
	return partnerTotalTVL, partnerTierRatio, vaultDelegatedAmountForPartner, partnerDelegatedBalanceBreakdown
}
