package partnerFees

import (
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
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
				totalSupply, err := vaultContract.TotalSupply(&bind.CallOpts{BlockNumber: big.NewInt(int64(blockNumber - 1))})
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

func getPartnerFees(chainID uint64) {
	fromBlock := 16650000
	logs.Info(`Retrieving all harvests for chain: ` + strconv.Itoa(int(chainID)))
	allHarvestPerStrategyPerVault := harvests.RetrieveAllHarvestPerStrategyPerVault(chainID)

	logs.Info(`Retrieving all partner tracker events for chain: ` + strconv.Itoa(int(chainID)))
	allPartnerTrackerEvents, allRefereesEvents := partnerTracker.RetrieveAllPartnerTrackerEvents(chainID)

	logs.Info(`Retrieving all vaults for chain: ` + strconv.Itoa(int(chainID)))
	allVaults := vaults.ListVaults(chainID)

	logs.Info(`Retrieving totalSupply for vaults for each harvest for chain: ` + strconv.Itoa(int(chainID)))
	vaultsSupplies := getVaultsSupplyAtBlock(allVaults, allHarvestPerStrategyPerVault)
	// partnerTrackerTree[event.Vault.ToAddress()][event.PartnerID.ToAddress()][event.Depositer.ToAddress()] = bigNumber.NewInt(0)

	// partnerToTest := common.HexToAddress(`0x558247e365be655f9144e1a0140D793984372Ef3`).ToAddress() //Ledger
	partnerToTest := common.HexToAddress(`0xFEB4acf3df3cDEA7399794D0869ef76A6EfAff52`).ToAddress() //Yearn

	logs.Info(`Processing partner fees for chain: ` + strconv.Itoa(int(chainID)))
	totalDelegatedValue := float64(0)
	totalFeesCollected := float64(0)
	for _, vault := range allVaults {
		vaultToTest := vault.Address
		allHarvestForVault := allHarvestPerStrategyPerVault[vaultToTest]
		vaultToTestName := vault.DisplaySymbol
		vaultDecimals := vault.Decimals

		for _, harvests := range allHarvestForVault {
			vaultDelegatedFees := float64(0)
			vaultDelegatedValue := float64(0)
			for _, harvest := range harvests {
				// if harvest.BlockNumber != 16695075 {
				// 	continue
				// }
				totalFees, _ := helpers.FormatAmount(harvest.Fees.TreasuryCollectedFee.String(), int(vaultDecimals))
				if totalFees == 0 { //Skip, no gain or loss, no fees
					continue
				}

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
					allPartnerTrackerEvents,
					vaultToTest,
					harvest.BlockNumber,
				)

				/**********************************************************************************************
				** We need to connect to the vault contract to be able to check the balance of the depositor
				** in the vault before the harvest, as well as the total supply of the vault tokens when the
				** harvest happened.
				**********************************************************************************************/
				totalSupply := vaultsSupplies[vaultToTest][harvest.BlockNumber]
				humanizedTotalSupply, bigHumanizedTotalSupply := helpers.FormatAmount(totalSupply.String(), int(vaultDecimals))
				_ = bigHumanizedTotalSupply

				/**********************************************************************************************
				** We have the list of deposits that happened before the harvest. Now, we want to check the
				** balance of the depositor in the vault before the harvest to see if he had other
				** deposits that were not done by the partner, or if he had done some withdrawals.
				**********************************************************************************************/
				// logs.Info(`For the harvest on block`, harvest.BlockNumber, `we have`, len(relevantPartnerTrackerEvents), `partners that should receive fees on a`, totalFees, `total fees collected`)
				for partner, eventsDepositorLevel := range relevantPartnerTrackerEvents {
					if partner.Hex() != partnerToTest.Hex() {
						continue
					}
					totalAllocatedForThisPartner := bigNumber.NewInt(0)
					// fmt.Println()
					for depositor, events := range eventsDepositorLevel {
						// logs.Info(`Working on partner`, partner.Hex(), `depositor`, depositor.Hex(), `with`, len(events), `events`)
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
						allTransfersForThisReferee, ok := allRefereesEvents[vaultToTest][depositor]
						if !ok {
							// logs.Info(`User`, depositor.Hex(), `has no transfers`)
							continue
						}

						delegatedBalance := bigNumber.NewInt(0)
						actualBalance := bigNumber.NewInt(0)
						// logs.Info(`User has a balance of 0`)
						for _, transfer := range allTransfersForThisReferee {
							// We need to filter the transfers that happened before the harvest.
							if transfer.BlockNumber > harvest.BlockNumber {
								continue
							}
							// Transfers that happened before the earliest deposit via partner are not relevant.
							if transfer.BlockNumber < earliestDepositBlock {
								continue
							}
							isReceiving := transfer.To.Hex() == depositor.Hex()
							isSending := transfer.From.Hex() == depositor.Hex()

							// transferValue := bigNumber.NewInt(0).Clone(transfer.Value)
							// humanizedTransferValue, _ := helpers.FormatAmount(transferValue.String(), int(vaultDecimals))
							/**************************************************************************
							** For each transfer IN, we check if the transfer hash matches one of the
							** deposit hash that we have computed above. If it does, we add the
							** amount to the accounting balance.
							** Otherwise, we just skip that deposit as it is not relevant for our
							** current partner accountability.
							**************************************************************************/
							if isReceiving {
								actualBalance = bigNumber.NewInt(0).Add(actualBalance, bigNumber.NewInt(0).Clone(transfer.Value))
								// humanizedActualBalance, _ := helpers.FormatAmount(actualBalance.String(), int(vaultDecimals))

								// It's not a deposit via this partner, we skip it.
								if _, ok := userDepositViaPartnerMap[transfer.TxHash]; !ok {
									// humanizedDelegatedBalance, _ := helpers.FormatAmount(delegatedBalance.String(), int(vaultDecimals))
									// delegatedRatioPercent := strconv.FormatFloat(humanizedDelegatedBalance/humanizedActualBalance*100, 'f', 2, 64)
									// logs.Warning(`User received`, humanizedTransferValue, `and now has a balance of`, humanizedActualBalance, `and a delegated balance of`, humanizedDelegatedBalance, `(`+delegatedRatioPercent+`%)`)
									continue
								}

								delegatedBalance = bigNumber.NewInt(0).Add(delegatedBalance, bigNumber.NewInt(0).Clone(transfer.Value))
								// humanizedDelegatedBalance, _ := helpers.FormatAmount(delegatedBalance.String(), int(vaultDecimals))
								// delegatedRatioPercent := strconv.FormatFloat(humanizedDelegatedBalance/humanizedActualBalance*100, 'f', 2, 64)
								// logs.Info(`User received`, humanizedTransferValue, `and now has a balance of`, humanizedActualBalance, `and a delegated balance of`, humanizedDelegatedBalance, `(`+delegatedRatioPercent+`%)`)
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
								if actualBalance.IsZero() {
									delegatedBalance = bigNumber.NewInt(0)
									// logs.Warning(`User sent`, humanizedTransferValue, `and now has a balance of 0 and a delegated balance of 0 (0%)`)
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

								// humanizedActualBalance, _ := helpers.FormatAmount(actualBalance.String(), int(vaultDecimals))
								// humanizedDelegatedBalance, _ := helpers.FormatAmount(delegatedBalance.String(), int(vaultDecimals))
								// delegatedRatioPercent := strconv.FormatFloat(humanizedDelegatedBalance/humanizedActualBalance*100, 'f', 2, 64)

								// logs.Warning(`User sent`, humanizedTransferValue, `and now has a balance of`, humanizedActualBalance, `and a delegated balance of`, humanizedDelegatedBalance, `(`+delegatedRatioPercent+`%)`)
							}
						}
						// humanizedDelegatedBalance, _ := helpers.FormatAmount(delegatedBalance.String(), int(vaultDecimals))

						/**********************************************************************************************
						** We can now check the actual balance of the depositor in the vault before the harvest.
						** If the balance is 0, we just skip that depositor for that partner for that harvest.
						** If the balance is greater than 0, that's not a problem, we can use the expected balance
						** If the balance is lower than the expected balance, that's a problem and we should scale
						** the expected balance down.
						**********************************************************************************************/
						if delegatedBalance.Cmp(big.NewInt(0)) == 0 {
							//Skip, no balance in the vault
							// logs.Success(`User `+depositor.Hex()+` has`, humanizedDelegatedBalance, vaultToTestName, `delegated to `+partner.Hex())
							// fmt.Println()
							continue
						}
						totalAllocatedForThisPartner = bigNumber.NewInt(0).Add(totalAllocatedForThisPartner, delegatedBalance)
						// logs.Success(`User `+depositor.Hex()+` has`, humanizedDelegatedBalance, vaultToTestName, `delegated to `+partner.Hex())

						// if balance.Cmp(&userBalanceViaPartner.Int) >= 0 {
						// 	//No problem, the balance is greater than the expected balance
						// 	logs.Info(`Partner `+partner.Hex()+` has`, humanizedUserBalanceViaPartner, vaultToTestName+` for user `+depositor.Hex())
						// 	totalAllocatedForThisPartner = bigNumber.NewInt(0).Add(totalAllocatedForThisPartner, userBalanceViaPartner)
						// }
						// if balance.Cmp(&userBalanceViaPartner.Int) < 0 {
						// 	//Get the reduced expected balance, aka half of what is left
						// 	halfUserBalanceViaPartner := bigNumber.NewInt(0).Div(balance, bigNumber.NewInt(2))
						// 	humanizedHalfUserBalanceViaPartner, _ := helpers.FormatAmount(halfUserBalanceViaPartner.String(), int(vaultDecimals))

						// 	logs.Info(`Partner `+partner.Hex()+` is supposed to have`, humanizedUserBalanceViaPartner, vaultToTestName+` for user `+depositor.Hex(), `, but only`, humanizedDelegatedBalance, vaultToTestName, `are in the vault. Using half, aka`, humanizedHalfUserBalanceViaPartner, vaultToTestName)
						// 	totalAllocatedForThisPartner = bigNumber.NewInt(0).Add(totalAllocatedForThisPartner, bigNumber.NewInt(0).Clone(halfUserBalanceViaPartner))
						// }
						// fmt.Println()
					}
					/**********************************************************************************************
					** Once we got that, we know how much tokens the partner has under his "balance sheet" in the
					** vault. We can now compute the percentage of the total supply that this partner has.
					**********************************************************************************************/
					humanizedTotalAllocatedForThisPartner, _ := helpers.FormatAmount(totalAllocatedForThisPartner.String(), int(vaultDecimals))
					// logs.Info(`Vault has a total supply of `+bigHumanizedTotalSupply.Text('f', -1), vaultToTestName)
					// logs.Info(`Partner`, partner.Hex(), `has`, humanizedTotalAllocatedForThisPartner, vaultToTestName, ` in the vault`)
					// logs.Info(`Partner`, partner.Hex(), `has`, humanizedTotalAllocatedForThisPartner/humanizedTotalSupply*100, `% of the total supply of`, vaultToTestName)
					// logs.Info(`--------------`)
					// logs.Info(`Total fees collected are`, totalFees, vaultToTestName)
					bigFloat64 := bigNumber.NewFloat(0).SetFloat64((totalFees * humanizedTotalAllocatedForThisPartner / humanizedTotalSupply) * 0.40)
					// logs.Info(`Fees for this partner are`, bigFloat64.Text('f', -1), vaultToTestName, `over`, totalFees, vaultToTestName, `collected`)
					asFloat64, _ := bigFloat64.Float64()
					vaultDelegatedFees += asFloat64
					vaultDelegatedValue += humanizedTotalAllocatedForThisPartner
				}
			}

			_, price := buildTokenPrice(chainID, common.FromAddress(vault.Address))
			if vaultDelegatedFees > 0 {
				asBigFloatStr := bigNumber.NewFloat(vaultDelegatedFees*price).Text('f', -1)
				logs.Info(vaultToTestName, `collected:`, vaultDelegatedFees, `($`, asBigFloatStr, `)`)
				totalFeesCollected += vaultDelegatedFees
			}
			if vaultDelegatedValue > 0 {
				asBigFloatStr := bigNumber.NewFloat(vaultDelegatedValue*price).Text('f', -1)
				logs.Info(vaultToTestName, `delegated value:`, vaultDelegatedValue, `($`, asBigFloatStr, `)`)
				totalDelegatedValue += vaultDelegatedValue
			}
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

// DEBUG
func buildTokenPrice(chainID uint64, tokenAddress common.Address) (*bigNumber.Float, float64) {
	fPrice := new(bigNumber.Float)
	price, ok := prices.FindPrice(chainID, tokenAddress)
	if ok {
		fPrice.SetInt(price)
		humanizedPrice := new(bigNumber.Float).Quo(fPrice, bigNumber.NewFloat(math.Pow10(int(6))))
		fHumanizedPrice, _ := humanizedPrice.Float64()
		return humanizedPrice, fHumanizedPrice
	}
	return bigNumber.NewFloat(), 0.0
}
