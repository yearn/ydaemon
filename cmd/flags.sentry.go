package main

import (
	"github.com/yearn/ydaemon/common/traces"
)

func handleSentryInitialization() {
	if *shouldEnableSentry {
		traces.SetupSentry()
		traces.Init(`app.bootstrap`).SetTag(`subsystem`, `main`)
	} else {
		traces.IsEnabled = false
	}
}
