package apr

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// APY trace logging helper (opt-in via env vars):
//
//   - Enable tracing:
//     APY_TRACE=1
//
//   - Write to a dedicated file (optional):
//     APY_TRACE_PATH=/tmp/apy_trace.log
//
//   - If this is a relative path, it is relative to the process working directory.
//
//   - If not set, output goes to stderr.
//
//   - Filter by chain, vault, and/or strategy (comma-separated):
//     APY_TRACE_CHAIN=1,10
//     APY_TRACE_VAULT=0xabc...,0xdef...
//     APY_TRACE_STRATEGY=0x123...,0x456...
//
// Notes:
// - Filters are case-insensitive and address strings should be hex with 0x prefix.
// - If a filter is set, only matching entries are emitted.
type apyTraceConfig struct {
	enabled    bool
	path       string
	chains     map[string]struct{}
	vaults     map[string]struct{}
	strategies map[string]struct{}
}

var (
	apyTraceConfigOnce sync.Once
	apyTraceCfg        apyTraceConfig

	apyTraceFileOnce sync.Once
	apyTraceFile     *os.File
	apyTraceFileErr  error
	apyTraceFileWarn sync.Once

	apyTraceWriteMu sync.Mutex
)

func apyTrace(scope string, chainID uint64, vaultAddr common.Address, strategyAddr common.Address, step string, value interface{}) {
	cfg := apyTraceLoadConfig()
	if !apyTraceMatches(cfg, chainID, vaultAddr, strategyAddr) {
		return
	}

	line := fmt.Sprintf(
		"%s APY_TRACE scope=%s chain=%d vault=%s strategy=%s step=%s value=%v",
		time.Now().UTC().Format(time.RFC3339Nano),
		scope,
		chainID,
		vaultAddr.Hex(),
		strategyAddr.Hex(),
		step,
		value,
	)

	if cfg.path != "" {
		if file, err := apyTraceOpenFile(cfg.path); err == nil && file != nil {
			apyTraceWriteMu.Lock()
			fmt.Fprintln(file, line)
			apyTraceWriteMu.Unlock()
			return
		}
	}

	fmt.Fprintln(os.Stderr, line)
}

func apyTraceLoadConfig() apyTraceConfig {
	apyTraceConfigOnce.Do(func() {
		enabled := false
		if value, ok := os.LookupEnv("APY_TRACE"); ok {
			enabled = apyTraceParseBool(value)
		} else if _, ok := os.LookupEnv("APY_TRACE_PATH"); ok {
			enabled = true
		}

		apyTraceCfg = apyTraceConfig{
			enabled:    enabled,
			path:       strings.TrimSpace(os.Getenv("APY_TRACE_PATH")),
			chains:     apyTraceParseList("APY_TRACE_CHAIN"),
			vaults:     apyTraceParseList("APY_TRACE_VAULT"),
			strategies: apyTraceParseList("APY_TRACE_STRATEGY"),
		}
	})

	return apyTraceCfg
}

func apyTraceParseBool(value string) bool {
	if strings.TrimSpace(value) == "" {
		return true
	}
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "true", "t", "yes", "y", "on":
		return true
	default:
		return false
	}
}

func apyTraceParseList(envKey string) map[string]struct{} {
	raw := strings.TrimSpace(os.Getenv(envKey))
	if raw == "" {
		return nil
	}

	items := strings.Split(raw, ",")
	values := make(map[string]struct{}, len(items))
	for _, item := range items {
		value := strings.ToLower(strings.TrimSpace(item))
		if value == "" {
			continue
		}
		values[value] = struct{}{}
	}

	return values
}

func apyTraceMatches(cfg apyTraceConfig, chainID uint64, vaultAddr common.Address, strategyAddr common.Address) bool {
	if !cfg.enabled {
		return false
	}

	if len(cfg.chains) > 0 {
		chainKey := strings.ToLower(strconv.FormatUint(chainID, 10))
		if _, ok := cfg.chains[chainKey]; !ok {
			return false
		}
	}

	if len(cfg.vaults) > 0 {
		vaultKey := strings.ToLower(vaultAddr.Hex())
		if _, ok := cfg.vaults[vaultKey]; !ok {
			return false
		}
	}

	if len(cfg.strategies) > 0 {
		strategyKey := strings.ToLower(strategyAddr.Hex())
		if _, ok := cfg.strategies[strategyKey]; !ok {
			return false
		}
	}

	return true
}

func apyTraceOpenFile(path string) (*os.File, error) {
	apyTraceFileOnce.Do(func() {
		if path == "" {
			return
		}
		apyTraceFile, apyTraceFileErr = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if apyTraceFileErr != nil {
			apyTraceFileWarn.Do(func() {
				fmt.Fprintf(os.Stderr, "%s APY_TRACE error opening %s: %v\n", time.Now().UTC().Format(time.RFC3339Nano), path, apyTraceFileErr)
			})
		}
	})

	return apyTraceFile, apyTraceFileErr
}
