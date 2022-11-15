package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/getsentry/sentry-go"
	"github.com/yearn/ydaemon/common/logs"
)

// var chains = env.SUPPORTED_CHAIN_IDS
var chains = []uint64{1}

func waitGroupSummonDaemons(ctx context.Context, wg *sync.WaitGroup, chainID uint64) {
	span := sentry.StartSpan(ctx, "app.bootstrap.summon.daemon")
	span.SetTag("chainId", strconv.Itoa(int(chainID)))

	SummonDaemons(chainID)

	wg.Done()
	logs.Success(`Daemons for chainID ` + strconv.Itoa(int(chainID)) + ` summoned successfully!`)
	span.Finish()
}

func summonDaemonsForAllChains(ctx context.Context) {
	span := sentry.StartSpan(ctx, "app.bootstrap.summon.all")
	defer span.Finish()

	var wg sync.WaitGroup
	for _, chainID := range chains {
		wg.Add(1)
		go waitGroupSummonDaemons(span.Context(), &wg, chainID)
	}

	wg.Wait()
}

func waitGroupLoadDaemons(ctx context.Context, wg *sync.WaitGroup, chainID uint64) {
	span := sentry.StartSpan(ctx, "app.bootstrap.load_state.daemon")
	span.SetTag("chainId", strconv.Itoa(int(chainID)))
	defer span.Finish()

	LoadDaemons(chainID)

	wg.Done()
	logs.Success(`Store data loaded in yDaemon memory for chainID ` + strconv.Itoa(int(chainID)) + `!`)
}

func loadDaemonsForAllChains(ctx context.Context) {
	span := sentry.StartSpan(ctx, "app.bootstrap.load_state.all")
	defer span.Finish()

	var wg sync.WaitGroup
	for _, chainID := range chains {
		wg.Add(1)
		go waitGroupLoadDaemons(span.Context(), &wg, chainID)
	}
	wg.Wait()
}

func getSampleRate() float64 {
	SAMPLE_RATE, exists := os.LookupEnv("SENTRY_SAMPLE_RATE")
	if !exists {
		return 1.0
	}
	parsed, error := strconv.ParseFloat(SAMPLE_RATE, 64)
	if error == nil {
		return parsed
	} else {
		log.Fatalf("Sample rate is not a float32: %s", SAMPLE_RATE)
		return 0
	}
}

func setupSentry() {
	SENTRY_DSN, exists := os.LookupEnv("SENTRY_DSN")
	if exists {
		sampleRate := getSampleRate()
		logs.Info(`Sentry TracesSampleRate set to ` + strconv.FormatFloat(sampleRate, 'f', 2, 64))
		err := sentry.Init(sentry.ClientOptions{
			Dsn: SENTRY_DSN,
			// Set TracesSampleRate to 1.0 to capture 100%
			// of transactions for performance monitoring.
			// We recommend adjusting this value in production,
			TracesSampleRate: sampleRate,
			// As it's not uncommon to panic with a string, it's
			// recommended to use the AttachStacktrace option
			// during SDK initialization, which will try to
			// provide a useful stack trace for messages as well.
			AttachStacktrace: true,
		})
		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		} else {
			logs.Info("Sentry initialized")
		}
	} else {
		logs.Warning("SENTRY_DSN not set, Sentry not initialized")
	}
}

func main() {
	setupSentry()

	ctx := context.Background()
	span := sentry.StartSpan(ctx, "app.bootstrap",
		sentry.TransactionName("Start Daemons"))
	span.SetTag("subsystem", "main")

	logs.Info(`Loading store data to yDaemon memory...`)
	loadDaemonsForAllChains(span.Context())
	logs.Info(`Summoning yDaemon...`)
	summonDaemonsForAllChains(span.Context())

	span.Finish()
	logs.Success(`Server ready!`)
	NewRouter().Run()
}
