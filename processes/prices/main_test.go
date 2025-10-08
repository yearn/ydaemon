package prices

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
)

func TestApplyCandidatePricesInsertsWhenMissing(t *testing.T) {
	addr := common.HexToAddress("0x0000000000000000000000000000000000000011")
	candidate := models.TPrices{
		Address:        addr,
		Price:          bigNumber.NewInt(123),
		HumanizedPrice: bigNumber.NewFloat(0).SetFloat64(0.123),
		Source:         "lens",
	}

	target := map[common.Address]models.TPrices{}
	applyCandidatePrices(target, map[common.Address]models.TPrices{addr: candidate})

	stored, ok := target[addr]
	if !ok {
		t.Fatalf("expected candidate to be inserted")
	}
	if stored.Price.String() != candidate.Price.String() {
		t.Fatalf("expected stored price %s, got %s", candidate.Price.String(), stored.Price.String())
	}
}

func TestApplyCandidatePricesSkipsNonImproving(t *testing.T) {
	addr := common.HexToAddress("0x0000000000000000000000000000000000000022")
	existing := models.TPrices{
		Address:        addr,
		Price:          bigNumber.NewInt(456),
		HumanizedPrice: bigNumber.NewFloat(0).SetFloat64(0.456),
		Source:         "existing",
	}

	zeroCandidate := models.TPrices{
		Address: addr,
		Price:   bigNumber.NewInt(0),
		Source:  "lens",
	}

	target := map[common.Address]models.TPrices{addr: existing}
	applyCandidatePrices(target, map[common.Address]models.TPrices{addr: zeroCandidate})

	stored := target[addr]
	if stored.Source != existing.Source {
		t.Fatalf("expected existing price to remain, got source %s", stored.Source)
	}
}

func TestApplyCandidatePricesReplacesZero(t *testing.T) {
	addr := common.HexToAddress("0x0000000000000000000000000000000000000033")
	zeroPrice := models.TPrices{
		Address: addr,
		Price:   bigNumber.NewInt(0),
		Source:  "placeholder",
	}

	candidate := models.TPrices{
		Address:        addr,
		Price:          bigNumber.NewInt(789),
		HumanizedPrice: bigNumber.NewFloat(0).SetFloat64(0.789),
		Source:         "lens",
	}

	target := map[common.Address]models.TPrices{addr: zeroPrice}
	applyCandidatePrices(target, map[common.Address]models.TPrices{addr: candidate})

	stored := target[addr]
	if stored.Source != candidate.Source {
		t.Fatalf("expected candidate to replace zero price, got source %s", stored.Source)
	}
	if stored.Price.String() != candidate.Price.String() {
		t.Fatalf("expected stored price %s, got %s", candidate.Price.String(), stored.Price.String())
	}
}
