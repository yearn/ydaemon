package main

import "strings"

type TProcess string

const (
	ProcessServer          TProcess = "server"
	ProcessTokenList       TProcess = "tokenlist"
	ProcessVaultMigrations TProcess = "vaultmigrations"
	ProcessInitDailyBlock  TProcess = "initdailyblock"
	ProcessAPY             TProcess = "apy"
	ProcessPrice           TProcess = "price"
)

func handleProcessInitialization(rawProcess *string) TProcess {
	switch strings.ToLower(*rawProcess) {
	case string(ProcessServer):
		process = ProcessServer
	case string(ProcessTokenList):
		process = ProcessTokenList
	case string(ProcessVaultMigrations):
		process = ProcessVaultMigrations
	case string(ProcessInitDailyBlock):
		process = ProcessInitDailyBlock
	case string(ProcessAPY):
		process = ProcessAPY
	case string(ProcessPrice):
		process = ProcessPrice
	default:
		process = ProcessServer
	}
	return process
}
