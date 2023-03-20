package main

import (
	"github.com/yearn/ydaemon/common/traces"
)

var trace *traces.TTrace

func handleSentryInitialization() {
	if *shouldEnableSentry {
		traces.SetupSentry()
		trace = traces.Init(`app.bootstrap`).SetTag(`subsystem`, `main`)
	} else {
		traces.IsEnabled = false
	}
}
