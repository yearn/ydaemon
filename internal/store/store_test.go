package store

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/stretchr/testify/assert"
)

// Not much to test here, but we are doing them for the coverage
func TestLogs(t *testing.T) {
	assert.IsType(t, TokenList, map[uint64][]common.Address{})
	assert.IsType(t, TokenPrices, map[uint64]map[common.Address]uint64{})
	assert.IsType(t, VaultsFromMeta, map[uint64]map[string]models.TVaultFromMeta{})
	assert.IsType(t, TokensFromMeta, map[uint64]map[string]models.TTokenFromMeta{})
	assert.IsType(t, StrategiesFromMeta, map[uint64]map[string]models.TStrategyFromMeta{})
	assert.IsType(t, VaultsFromAPIV1, map[uint64]map[string]models.TAPIV1Vault{})
}
