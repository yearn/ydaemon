package main

import "strings"

type TProcess string

const (
	ProcessServer          TProcess = "server"
	ProcessVaultMigrations TProcess = "vaultmigrations"
	ProcessInitDailyBlock  TProcess = "initdailyblock"
	ProcessAPY             TProcess = "apy"
)

func handleProcessInitialization(rawProcess *string) TProcess {
	switch strings.ToLower(*rawProcess) {
	case string(ProcessServer):
		process = ProcessServer
	case string(ProcessVaultMigrations):
		process = ProcessVaultMigrations
	case string(ProcessInitDailyBlock):
		process = ProcessInitDailyBlock
	case string(ProcessAPY):
		process = ProcessAPY
	default:
		process = ProcessServer
	}
	return process
}
