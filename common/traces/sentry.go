package traces

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/yearn/ydaemon/common/logs"
)

var IsEnabled = false

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
	if !IsEnabled {
		return 0
	}
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
	if !IsEnabled {
		return
	}
	SENTRY_DSN, exists := os.LookupEnv("SENTRY_DSN")
	if exists {
		SERVER_NAME, exists := os.LookupEnv("SERVER_NAME")
		if !exists {
			SERVER_NAME = `unknown`
		}

		sampleRate := getSampleRate()
		logs.Info(`Sentry TracesSampleRate set to ` + strconv.FormatFloat(sampleRate, 'f', 2, 64))
		if err := sentry.Init(sentry.ClientOptions{
			ServerName:       SERVER_NAME,
			Dsn:              SENTRY_DSN,
			TracesSampleRate: sampleRate,
			AttachStacktrace: true,
		}); err != nil {
			logs.Error(err)
		} else {
			logs.Success("Sentry initialized")
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
	if !IsEnabled {
		return nil
	}
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
func (s *TTrace) SetTag(key string, value string) *TTrace {
	if !IsEnabled {
		return &TTrace{}
	}
	s.span.SetTag(key, value)
	return s
}

func (s *TTrace) Child(key string, tags ...TTags) *TTrace {
	if !IsEnabled {
		return &TTrace{}
	}
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

func (s TTrace) Finish() {
	if !IsEnabled {
		return
	}
	startedAt := s.span.StartTime
	endedAt := time.Now()
	duration := endedAt.Sub(startedAt)

	logs.Trace(s.key, 0, duration.String())
	s.span.Finish()
}

var captureLevel = map[string]sentry.Level{
	`debug`: sentry.LevelDebug,
	`info`:  sentry.LevelInfo,
	`warn`:  sentry.LevelWarning,
	`error`: sentry.LevelError,
}

type TCapturedEvent struct {
	Message string
	Level   sentry.Level
	Tags    map[string]string
	Event   *sentry.Event
}

func Capture(level string, msg string, tags ...TTags) *TCapturedEvent {
	if !IsEnabled {
		return &TCapturedEvent{}
	}
	event := sentry.NewEvent()
	event.Level = captureLevel[level]
	event.Message = msg
	event.Fingerprint = []string{msg}
	if len(tags) > 0 {
		for _, tag := range tags {
			event.Tags[tag.Name] = tag.Value
		}
	}

	return &TCapturedEvent{
		Message: msg,
		Level:   event.Level,
		Tags:    event.Tags,
		Event:   event,
	}
}

func (c *TCapturedEvent) SetEntity(value string) *TCapturedEvent {
	if !IsEnabled {
		return &TCapturedEvent{}
	}
	c.Event.Type = value
	return c
}
func (c *TCapturedEvent) SetTag(key string, value string) *TCapturedEvent {
	if !IsEnabled {
		return &TCapturedEvent{}
	}
	c.Event.Tags[key] = value
	return c
}
func (c *TCapturedEvent) SetTags(tags ...TTags) *TCapturedEvent {
	if !IsEnabled {
		return &TCapturedEvent{}
	}
	if len(tags) > 0 {
		for _, tag := range tags {
			c.Event.Tags[tag.Name] = tag.Value
		}
	}
	return c
}
func (c *TCapturedEvent) SetExtra(key string, value interface{}) *TCapturedEvent {
	if !IsEnabled {
		return &TCapturedEvent{}
	}
	c.Event.Extra[key] = value
	return c
}

func (c *TCapturedEvent) Send() *TCapturedEvent {
	if !IsEnabled {
		return &TCapturedEvent{}
	}
	sentry.CurrentHub().Clone().CaptureEvent(c.Event)
	switch c.Level {
	case sentry.LevelError:
		logs.Error(c.Message)
	case sentry.LevelInfo:
		logs.Info(c.Message)
	case sentry.LevelDebug:
		logs.Debug(c.Message)
	case sentry.LevelWarning:
		logs.Warning(c.Message)
	default:
		logs.Info(c.Message)
	}
	return c
}
