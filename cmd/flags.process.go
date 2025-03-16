package main

import "strings"

type TProcess string

const (
	ProcessServer          TProcess = "server"
	ProcessVaultMigrations TProcess = "vaultmigrations"
	ProcessAPY             TProcess = "apy"
)

func handleProcessInitialization(rawProcess *string) TProcess {
	switch strings.ToLower(*rawProcess) {
	case string(ProcessServer):
		process = ProcessServer
	case string(ProcessVaultMigrations):
		process = ProcessVaultMigrations
	case string(ProcessAPY):
		process = ProcessAPY
	default:
		process = ProcessServer
	}
	return process
}
