package apr

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
)

func ComputeForwardStrategyAPR(strategy models.TStrategy) (*bigNumber.Float, error) {
	oracleAPR := bigNumber.NewFloat(0)
	chain, ok := env.GetChain(strategy.ChainID)
	if !ok {
		return nil, errors.New(`chain not found`)
	}
	oracleContract := chain.APROracleContract.Address
	if oracleContract == common.HexToAddress(``) {
		return nil, errors.New(`oracle not found`)
	}
	oracle, err := contracts.NewYVaultsV3APROracleCaller(oracleContract, ethereum.GetRPC(strategy.ChainID))
	if err != nil {
		return nil, err
	}

	/**********************************************************************************************
	** If the vault is a single strategy vault, we can use the oracle directly to get the APR of
	** the vault as expected APR
	**********************************************************************************************/
	var hasError error
	expected, err := oracle.GetStrategyApr(nil, strategy.Address, big.NewInt(0))
	if err == nil {
		oracleAPR = helpers.ToNormalizedAmount(bigNumber.SetInt(expected), 18)
	} else {
		hasError = err
	}

	if hasError != nil || oracleAPR.IsZero() {
		expected, newErr := oracle.GetCurrentApr(nil, strategy.VaultAddress)
		err = newErr
		if newErr == nil {
			oracleAPR = helpers.ToNormalizedAmount(bigNumber.SetInt(expected), 18)
		} else {
			return nil, err
		}
	}

	/**********************************************************************************************
	** Define which APR we want to use as "Net APR".
	**********************************************************************************************/
	primaryAPR := oracleAPR
	primaryAPRFloat64, _ := primaryAPR.Float64()
	primaryAPY := bigNumber.NewFloat(0).SetFloat64(convertFloatAPRToAPY(primaryAPRFloat64, 52))

	return primaryAPY, nil
}
