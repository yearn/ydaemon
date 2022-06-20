package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Not much to test here, but we are doing them for the coverage
func TestLogs(t *testing.T) {
	APIV1Fees := TAPIV1Fees{}
	APIV1APY := TAPIV1APY{}
	APIV1Vault := TAPIV1Vault{}
	assert.IsType(t, APIV1Fees, TAPIV1Fees{})
	assert.IsType(t, APIV1APY, TAPIV1APY{})
	assert.IsType(t, APIV1Vault, TAPIV1Vault{})

	VaultFromGraphToken := TVaultFromGraphToken{}
	VaultFromGraphStrategy := TVaultFromGraphStrategy{}
	VaultFromGraphVaultDayData := TVaultFromGraphVaultDayData{}
	VaultFromGraph := TVaultFromGraph{}
	GraphQueryResponseForVaults := TGraphQueryResponseForVaults{}
	assert.IsType(t, VaultFromGraphToken, TVaultFromGraphToken{})
	assert.IsType(t, VaultFromGraphStrategy, TVaultFromGraphStrategy{})
	assert.IsType(t, VaultFromGraphVaultDayData, TVaultFromGraphVaultDayData{})
	assert.IsType(t, VaultFromGraph, TVaultFromGraph{})
	assert.IsType(t, GraphQueryResponseForVaults, TGraphQueryResponseForVaults{})

	VaultFromMeta := TVaultFromMeta{}
	TokenFromMeta := TTokenFromMeta{}
	StrategyFromMeta := TStrategyFromMeta{}
	assert.IsType(t, VaultFromMeta, TVaultFromMeta{})
	assert.IsType(t, TokenFromMeta, TTokenFromMeta{})
	assert.IsType(t, StrategyFromMeta, TStrategyFromMeta{})

	Token := TToken{}
	TVL := TTVL{}
	APYFees := TAPYFees{}
	APYPoints := TAPYPoints{}
	APYComposite := TAPYComposite{}
	APY := TAPY{}
	Strategy := TStrategy{}
	Migration := TMigration{}
	Vault := TVault{}
	assert.IsType(t, Token, TToken{})
	assert.IsType(t, TVL, TTVL{})
	assert.IsType(t, APYFees, TAPYFees{})
	assert.IsType(t, APYPoints, TAPYPoints{})
	assert.IsType(t, APYComposite, TAPYComposite{})
	assert.IsType(t, APY, TAPY{})
	assert.IsType(t, Strategy, TStrategy{})
	assert.IsType(t, Migration, TMigration{})
	assert.IsType(t, Vault, TVault{})
}
