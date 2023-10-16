package main

import "github.com/yearn/ydaemon/common/logs"

func handleSentryInitialization() {
	if *shouldEnableSentry {
		logs.SetupSentry()
		logs.Init(`app.bootstrap`).SetTag(`subsystem`, `main`)
	} else {
		logs.IsEnabled = false
	}
}
