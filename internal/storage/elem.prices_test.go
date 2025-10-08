package storage

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/internal/models"
)

func captureOutput(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	_ = w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	_ = r.Close()
	return buf.String()
}

func resetZeroPriceLogState() {
	zeroPriceLogMu.Lock()
	zeroPriceLogCache = make(map[uint64]map[common.Address]time.Time)
	zeroPriceLogMu.Unlock()
}

func storeZeroPrice(chainID uint64, addr common.Address) {
	StorePrice(chainID, models.TPrices{
		Address:        addr,
		Price:          bigNumber.NewInt(0),
		HumanizedPrice: bigNumber.NewFloat(0),
		Source:         "test",
	})
}

func TestGetPriceLogsZeroPriceOnce(t *testing.T) {
	resetZeroPriceLogState()
	originalTTL := zeroPriceLogTTL
	zeroPriceLogTTL = time.Minute
	defer func() { zeroPriceLogTTL = originalTTL }()

	chainID := uint64(1)
	addr := common.HexToAddress("0x0000000000000000000000000000000000000011")
	storeZeroPrice(chainID, addr)

	first := captureOutput(func() {
		if _, ok := GetPrice(chainID, addr); !ok {
			t.Fatalf("expected price to be found")
		}
	})
	if !strings.Contains(first, "[PRICE ZERO]") {
		t.Fatalf("expected zero-price warning, got %q", first)
	}

	second := captureOutput(func() {
		GetPrice(chainID, addr)
	})
	if strings.Contains(second, "[PRICE ZERO]") {
		t.Fatalf("expected warning to be deduplicated, got %q", second)
	}
}

func TestGetPriceZeroLogRespectTTL(t *testing.T) {
	resetZeroPriceLogState()
	originalTTL := zeroPriceLogTTL
	zeroPriceLogTTL = 5 * time.Millisecond
	defer func() { zeroPriceLogTTL = originalTTL }()

	chainID := uint64(1)
	addr := common.HexToAddress("0x0000000000000000000000000000000000000012")
	storeZeroPrice(chainID, addr)

	_ = captureOutput(func() {
		GetPrice(chainID, addr)
	})

	time.Sleep(10 * time.Millisecond)

	relog := captureOutput(func() {
		GetPrice(chainID, addr)
	})
	if !strings.Contains(relog, "[PRICE ZERO]") {
		t.Fatalf("expected warning after TTL expiration, got %q", relog)
	}
}

func TestListPricesLogsZeroPrice(t *testing.T) {
	resetZeroPriceLogState()
	originalTTL := zeroPriceLogTTL
	zeroPriceLogTTL = time.Minute
	defer func() { zeroPriceLogTTL = originalTTL }()

	chainID := uint64(1)
	addr := common.HexToAddress("0x0000000000000000000000000000000000000013")
	storeZeroPrice(chainID, addr)

	output := captureOutput(func() {
		ListPrices(chainID)
	})
	if !strings.Contains(output, "[PRICE ZERO]") {
		t.Fatalf("expected zero-price warning during ListPrices, got %q", output)
	}
}
