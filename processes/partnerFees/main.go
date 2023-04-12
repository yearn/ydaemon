package partnerFees

import (
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/harvests"
	"github.com/yearn/ydaemon/internal/indexer"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/registries"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
)

func Run(chainID uint64, fromBlock uint64, toBlock *uint64) {
	/**********************************************************************************************
	** Depending on the chainID, the partner contract was deployed at different block numbers. Any
	** fromBlock before that block number is a non-sense, so we need to adjust the fromBlock to
	** the block number of the deployment of the partner contract.
	**********************************************************************************************/
	chainStartBlock := getChainStartBlock(chainID)
	if fromBlock < chainStartBlock {
		fromBlock = chainStartBlock
	}
	previousTime := time.Now()

	/**********************************************************************************************
	** The first steps are to init out environment. We need to fetch all our vaults, strategies,
	** tokens, prices, etc. This will allow us to have an exhaustive list of all the data we need
	** to process.
	**********************************************************************************************/
	initYearnEcosystem(chainID)
	allVaults := vaults.ListVaults(chainID)

	/**********************************************************************************************
	** Once this env is initialized, we will need some specific data for the partner fees. We first
	** need to fetch all the harvests as we will compute the fees per harvest.
	**********************************************************************************************/
	allHarvests := harvests.RetrieveAllHarvests(chainID, fromBlock, toBlock)

	/**********************************************************************************************
	** Then we need to fetch all the partnerTrackerEvents and all the refereesEvents. The event we
	** are interested in for the partnerTracker contract is the ReferredBalanceIncreased which
	** indicates when a user deposited funds in a vault with a referral ID.
	** The refereesEvents correspond to all the transfers for the tokens the referees deposited in
	** a vault via the partner contract:
	** a user deposit DAI via the partner contract, the partner, we will then fetch all the DAI
	** transfers in or out of the vault from this user to get, at any block, it's balance in the
	** vault and the delegated amount he has.
	**********************************************************************************************/
	allPartnerTrackerEvents, allRefereesEvents := retrieveAllPartnerTrackerEvents(chainID, fromBlock, toBlock)

	/**********************************************************************************************
	** We finally need to fetch the supply of all the vaults at all the blocks we are interested
	** in, as we will use this to calculate the ratio of delegatedAmount / totalSupply for each
	** partner and determine the fee tier.
	**********************************************************************************************/
	vaultsSupplies := getVaultsSupplyAtBlock(allVaults, allHarvests)

	CSV_EXPORT_SYNCMAP := sync.Map{}
	CSV_BREAKDOWN_EXPORT_SYNCMAP := sync.Map{}
	// CSV_EXPORT := []TMainCSVExport{}
	// CSV_BREAKDOWN_EXPORT := []TPartnerBreakdownCSVExport{}
	payoutPerPartner := map[common.Address]map[common.Address]TBig{}

	/**********************************************************************************************
	** Below is the main loop, running in parallel for each harvest to compute all the relevant
	** elements.
	** In tldr; for each harvest we need to know:
	** - the partners that have delegated tokens in the vault
	** - the amount of tokens they have delegated
	** - the total TVL of all the partners across all the vaults
	** - the ratio of the TVL of each partner to the total TVL
	** - the fee tier of each partner
	** - the share of the fees that each partner is entitled to
	**********************************************************************************************/
	wg := sync.WaitGroup{}
	for _, vault := range allVaults {
		allHarvestForVault := allHarvests[vault.Address]
		for _, harvestList := range allHarvestForVault {
			wg.Add(len(harvestList))
			for _, harvest := range harvestList {
				go func(harvest models.THarvest, vaultAddress common.Address, vaultTokenAddress common.Address, vaultDecimals uint64) {
					defer wg.Done()
					if (harvest.BlockNumber < fromBlock) || (toBlock != nil && harvest.BlockNumber > *toBlock) {
						return // Skip, not in the range
					}
					if (harvest.Fees.TreasuryCollectedFee).IsZero() { //Skip, no gain or loss, no fees
						return
					}

					totalTVLForPartner, partnerTierRatio, vaultDelegatedAmountForPartner, partnerDelegationBreakdown := computeBlockData(
						chainID,
						harvest.BlockNumber,
						allPartnerTrackerEvents,
						allRefereesEvents,
					)
					for partnerAddress, partnerDelegatedAmount := range vaultDelegatedAmountForPartner[vaultAddress] {
						if partnerDelegatedAmount.IsZero() {
							continue // Skip, no TVL for this partner
						}

						/******************************************************************************
						** If we are here, this means that this partner has some delegated tokens in
						** that vault. We now need to know the value of those tokens and determine
						** the share of the fees that the partner is entitled to.
						******************************************************************************/
						partnerTier := bigNumber.NewFloat(0)
						if _, ok := partnerTierRatio[partnerAddress]; ok {
							partnerTier = partnerTierRatio[partnerAddress]
						}
						partnerTVL := bigNumber.NewFloat(0)
						if _, ok := partnerTierRatio[partnerAddress]; ok {
							partnerTVL = totalTVLForPartner[partnerAddress]
						}
						totalSupplyForVault := bigNumber.NewInt(0).Set(vaultsSupplies[vaultAddress][harvest.BlockNumber])
						vaultFees := bigNumber.NewInt(0).Clone(harvest.Fees.TreasuryCollectedFee)

						price, _ := prices.GetPricesOnBlock(chainID, harvest.BlockNumber, vaultTokenAddress)
						partnerFeeShareBase := bigNumber.NewFloat(0).Mul(
							toNormalizedAmount(vaultFees, vaultDecimals),
							bigNumber.NewFloat(0).Div(
								toNormalizedRatio(partnerDelegatedAmount, totalSupplyForVault, vaultDecimals),
								bigNumber.NewFloat(100),
							),
						)
						partnerFeeShareTier := bigNumber.NewFloat(0).Mul(
							bigNumber.NewFloat(0).Mul(partnerTier, partnerFeeShareBase),
							bigNumber.NewFloat(0.55), //OPEX
						)

						/******************************************************************************
						** Add this item to the CSV export
						******************************************************************************/
						if _, ok := payoutPerPartner[partnerAddress]; !ok {
							payoutPerPartner[partnerAddress] = map[common.Address]TBig{}
						}
						if _, ok := payoutPerPartner[partnerAddress][vaultAddress]; !ok {
							payoutPerPartner[partnerAddress][vaultAddress] = TBig{
								Raw:        bigNumber.NewInt(0),
								Normalized: bigNumber.NewFloat(0),
								Value:      bigNumber.NewFloat(0),
							}
						}
						payoutPerPartner[partnerAddress][vaultAddress] = TBig{
							Raw:        bigNumber.NewInt(0).Add(payoutPerPartner[partnerAddress][vaultAddress].Raw, partnerFeeShareTier.Int()),
							Normalized: bigNumber.NewFloat(0).Add(payoutPerPartner[partnerAddress][vaultAddress].Normalized, partnerFeeShareTier),
							Value: bigNumber.NewFloat(0).Add(
								payoutPerPartner[partnerAddress][vaultAddress].Value,
								bigNumber.NewFloat(0).Mul(partnerFeeShareTier, toNormalizedAmount(price, 6)),
							),
						}

						/******************************************************************************
						** Add this item to the CSV export
						******************************************************************************/
						eventKey := strconv.FormatUint(harvest.BlockNumber, 10) + "_" + partnerAddress.Hex() + "_" + vaultAddress.Hex() + "_" + partnerTVL.String() + "_" + partnerDelegatedAmount.String()
						CSV_EXPORT_SYNCMAP.Store(eventKey, TMainCSVExport{
							Block:               harvest.BlockNumber,
							Partner:             partnerAddress.Hex(),                       //Partner address
							Vault:               vaultAddress.Hex(),                         //Vault address
							Tier:                partnerTier.Text('f', -1),                  //Tier share
							VaultPrice:          toNormalizedAmount(price, 6).Text('f', -1), //Vault price
							Balance:             toNormalizedAmount(partnerDelegatedAmount, vaultDecimals).Text('f', -1),
							BalanceUSD:          formatWithPriceOnBlock(chainID, harvest.BlockNumber, vaultTokenAddress, partnerDelegatedAmount, vaultDecimals).Text('f', -1),
							TotalSupply:         toNormalizedAmount(totalSupplyForVault, vaultDecimals).Text('f', -1),
							TotalSupplyUSD:      formatWithPriceOnBlock(chainID, harvest.BlockNumber, vaultTokenAddress, totalSupplyForVault, vaultDecimals).Text('f', -1),
							TotalDelegatedValue: partnerTVL.Text('f', -1),
							CollectedFee:        toNormalizedAmount(vaultFees, vaultDecimals).Text('f', -1),
							CollectedFeeUSD:     formatWithPriceOnBlock(chainID, harvest.BlockNumber, vaultTokenAddress, vaultFees, vaultDecimals).Text('f', -1),
							PartnerFeeShare:     partnerFeeShareTier.Text('f', -1),
							PartnerFeeShareUSD:  bigNumber.NewFloat(0).Mul(partnerFeeShareTier, toNormalizedAmount(price, 6)).Text('f', -1),
						})
					}

					for partnerAddress, delegatedForParner := range partnerDelegationBreakdown {
						for vaultAddress, delegatedForVault := range delegatedForParner {
							for depositorAddress, delegatedBalance := range delegatedForVault {
								if delegatedBalance.Raw.IsZero() {
									continue // Skip, no TVL for this partner
								}
								eventKey := strconv.FormatUint(harvest.BlockNumber, 10) + "_" + partnerAddress.Hex() + "_" + vaultAddress.Hex() + "_" + depositorAddress.Hex() + "_" + delegatedBalance.Raw.String()
								CSV_BREAKDOWN_EXPORT_SYNCMAP.Store(eventKey, TPartnerBreakdownCSVExport{
									Block:      harvest.BlockNumber,
									Partner:    partnerAddress.Hex(),
									Vault:      vaultAddress.Hex(),
									Depositor:  depositorAddress.Hex(),
									Balance:    delegatedBalance.Normalized.Text('f', -1),
									BalanceUSD: delegatedBalance.Value.Text('f', -1),
								})
							}
						}
					}
				}(harvest, vault.Address, vault.Token.Address, vault.Decimals)
			}
		}
	}
	wg.Wait()

	/**********************************************************************************************
	** By default, our map are not sorted by blocknumber as we ran the script in parallel. To ease
	** things, we will sort them. and export them
	**********************************************************************************************/
	CSV_EXPORT := []TMainCSVExport{}
	CSV_BREAKDOWN_EXPORT := []TPartnerBreakdownCSVExport{}
	CSV_EXPORT_SYNCMAP.Range(func(key, value interface{}) bool {
		CSV_EXPORT = append(CSV_EXPORT, value.(TMainCSVExport))
		return true
	})
	CSV_BREAKDOWN_EXPORT_SYNCMAP.Range(func(key, value interface{}) bool {
		CSV_BREAKDOWN_EXPORT = append(CSV_BREAKDOWN_EXPORT, value.(TPartnerBreakdownCSVExport))
		return true
	})

	sort.Slice(CSV_EXPORT, func(i, j int) bool {
		return CSV_EXPORT[i].Block < CSV_EXPORT[j].Block
	})
	sort.Slice(CSV_BREAKDOWN_EXPORT, func(i, j int) bool {
		return CSV_BREAKDOWN_EXPORT[i].Block < CSV_BREAKDOWN_EXPORT[j].Block
	})

	//loop over CSV_BREAKDOWN_EXPORT and check if we got some duplicates
	for i, item := range CSV_BREAKDOWN_EXPORT {
		if i == 0 {
			continue
		}
		if item.Partner == CSV_BREAKDOWN_EXPORT[i-1].Partner &&
			item.Vault == CSV_BREAKDOWN_EXPORT[i-1].Vault &&
			item.Depositor == CSV_BREAKDOWN_EXPORT[i-1].Depositor &&
			item.Balance == CSV_BREAKDOWN_EXPORT[i-1].Balance &&
			item.BalanceUSD == CSV_BREAKDOWN_EXPORT[i-1].BalanceUSD &&
			item.Block == CSV_BREAKDOWN_EXPORT[i-1].Block {
			logs.Error(`Duplicate found in CSV_BREAKDOWN_EXPORT`, item)
		}
	}
	logs.Success(`NO DUPLICATES FOUND IN CSV_BREAKDOWN_EXPORT`)

	exportCSV(`output.csv`, CSV_EXPORT)
	exportCSV(`outputBreakdown.csv`, CSV_BREAKDOWN_EXPORT)

	/**********************************************************************************************
	** Finally, just because we are nice, we will print some logs to the console, aka for each
	** partner, how much they should get paid both in USD and in the vault token
	**********************************************************************************************/
	logs.Success(`Here are the expected payments for each partner`)
	totalForPartnersUSD := make(map[common.Address]*bigNumber.Float)
	for partnerAddress, vaultsForThatPartner := range payoutPerPartner {
		for vaultAddress, payout := range vaultsForThatPartner {
			if _, ok := totalForPartnersUSD[partnerAddress]; !ok {
				totalForPartnersUSD[partnerAddress] = bigNumber.NewFloat(0)
			}
			vault, _ := vaults.GetVault(chainID, vaultAddress)
			if payout.Normalized.Gt(bigNumber.NewFloat(0)) && payout.Value.Gt(bigNumber.NewFloat(0)) {
				logs.Info(`Partner`, partnerAddress.Hex(), ` for vault `, vaultAddress.Hex(), `:`, payout.Normalized.Text('f', -1), vault.Symbol, ` ($`+payout.Value.Text('f', -1)+`)`)
			}
			totalForPartnersUSD[partnerAddress] = bigNumber.NewFloat(0).Add(totalForPartnersUSD[partnerAddress], payout.Value)
		}
	}

	logs.Success(`----------------------------------------------------------------------------------`)
	for partnerAddress, total := range totalForPartnersUSD {
		if total.Gt(bigNumber.NewFloat(0)) {
			logs.Success(`Partner`, partnerAddress.Hex(), `: $`, total.Text('f', -1))
		}
	}
	logs.Success(`----------------------------------------------------------------------------------`)
	logs.Success(`Total time`, time.Since(previousTime))
}

/**************************************************************************************************
** initYearnEcosystem is used to initialize the yearn ecosystem data, aka fetching all the vaults,
** strategies, tokens, prices, etc.
** Based on that, we have everything ready to compute the fees for each partner.
**************************************************************************************************/
func initYearnEcosystem(chainID uint64) {
	vaultsMap := registries.RegisterAllVaults(chainID, 0, nil)
	tokens.RetrieveAllTokens(chainID, vaultsMap)
	prices.RetrieveAllPrices(chainID)
	vaults.RetrieveAllVaults(chainID, vaultsMap)
	strategiesAddedList := events.HandleStrategyAdded(chainID, vaultsMap, 0, nil)
	strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
	indexer.PostProcessStrategies(chainID)
}
