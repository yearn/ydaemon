package traces

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/yearn/ydaemon/common/logs"
)

var SPANS = map[string]string{
	//app.bootstrap
	`app.bootstrap`:                   `Bootstrap`,
	`app.bootstrap.summon.daemon`:     `Summon Daemons`,
	`app.bootstrap.summon.all`:        `Summon All`,
	`app.bootstrap.load_state.daemon`: `Load Daemons`,
	`app.bootstrap.load_state.all`:    `Load All`,

	//app.indexer
	`app.indexer.bribes.reward_added`:          `Fetch RewardAdded Events`,
	`app.indexer.registry.new_vaults_events`:   `Fetch Vaults From Registry`,
	`app.indexer.vaults.multicall_data`:        `Fetch Vaults Multicall Data`,
	`app.indexer.vaults.harvest_events`:        `Fetch Vaults Harvets Events`,
	`app.indexer.vaults.activation_events`:     `Fetch Vaults Activation Events`,
	`app.indexer.tokens.multicall_data`:        `Fetch Tokens Multicall Data`,
	`app.indexer.prices`:                       `Fetch Prices`,
	`app.indexer.strategies.activation_events`: `Fetch Strategies Activation Events`,
	`app.indexer.strategies.multicall_data`:    `Fetch Strategies Multicall Data`,
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

func SetupSentry() {
	SENTRY_DSN, exists := os.LookupEnv("SENTRY_DSN")
	if exists {
		SERVER_NAME, exists := os.LookupEnv("SERVER_NAME")
		if !exists {
			SERVER_NAME = `unknown`
		}

		sampleRate := getSampleRate()
		logs.Info(`Sentry TracesSampleRate set to ` + strconv.FormatFloat(sampleRate, 'f', 2, 64))
		err := sentry.Init(sentry.ClientOptions{
			ServerName:       SERVER_NAME,
			Dsn:              SENTRY_DSN,
			TracesSampleRate: sampleRate,
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

type TTags struct {
	Name  string
	Value string
}

type TTrace struct {
	key  string
	span *sentry.Span
}

func Init(key string, tags ...TTags) *TTrace {
	logs.Trace(key, 1, ``)
	span := sentry.StartSpan(context.Background(), key, sentry.TransactionName(SPANS[key]))
	if len(tags) > 0 {
		for _, tag := range tags {
			span.SetTag(tag.Name, tag.Value)
		}
	}

	return &TTrace{
		key:  key,
		span: span,
	}
}

func (s TTrace) Child(key string, tags ...TTags) *TTrace {
	logs.Trace(key, 1, ``)
	span := s.span.StartChild(key, sentry.TransactionName(SPANS[key]))
	if len(tags) > 0 {
		for _, tag := range tags {
			span.SetTag(tag.Name, tag.Value)
		}
	}

	return &TTrace{
		key:  key,
		span: span,
	}
}

func (s TTrace) Capture(msg string) {
	sentry.CurrentHub().CaptureException(errors.New(msg))
}

func (s TTrace) Finish() {
	startedAt := s.span.StartTime
	endedAt := time.Now()
	duration := endedAt.Sub(startedAt)

	logs.Trace(s.key, 0, duration.String())
	s.span.Finish()
}
