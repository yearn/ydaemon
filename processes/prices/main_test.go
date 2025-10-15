package prices

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
)

func captureOutput(t *testing.T, fn func()) string {
	t.Helper()
	oldStdout := os.Stdout
	oldStderr := os.Stderr

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	os.Stdout = w
	os.Stderr = w

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("failed to close writer: %v", err)
	}

	os.Stdout = oldStdout
	os.Stderr = oldStderr

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed to read pipe: %v", err)
	}
	if err := r.Close(); err != nil {
		t.Fatalf("failed to close reader: %v", err)
	}

	return buf.String()
}

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

func TestLogZeroPricesWarnsForNilOrZero(t *testing.T) {
	zeroAddr := common.HexToAddress("0x00000000000000000000000000000000000000AA")
	nilAddr := common.HexToAddress("0x00000000000000000000000000000000000000BB")

	zeroPrice := models.TPrices{
		Address: zeroAddr,
		Price:   bigNumber.NewInt(0),
		Source:  "test-zero",
	}
	nilPrice := models.TPrices{
		Address: nilAddr,
		Source:  "test-nil",
	}

	output := captureOutput(t, func() {
		logZeroPrices(1, map[common.Address]models.TPrices{
			zeroAddr: zeroPrice,
			nilAddr:  nilPrice,
		})
	})

	if count := strings.Count(output, "[PRICE ZERO]"); count != 2 {
		t.Fatalf("expected 2 zero-price warnings, got %d output %q", count, output)
	}
	if !strings.Contains(output, zeroAddr.Hex()) || !strings.Contains(output, nilAddr.Hex()) {
		t.Fatalf("expected output to include both token addresses, got %q", output)
	}
}

func TestLogZeroPricesSkipsNonZero(t *testing.T) {
	addr := common.HexToAddress("0x00000000000000000000000000000000000000CC")
	positive := models.TPrices{
		Address: addr,
		Price:   bigNumber.NewInt(123),
		Source:  "test-positive",
	}

	output := captureOutput(t, func() {
		logZeroPrices(1, map[common.Address]models.TPrices{addr: positive})
	})

	if strings.Contains(output, "[PRICE ZERO]") {
		t.Fatalf("expected no zero-price warnings, got %q", output)
	}
}
