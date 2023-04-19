package main

import "strings"

type TProcess string

const (
	ProcessServer      TProcess = "server"
	ProcessPartnerFees TProcess = "partnerfees"
	ProcessTokenList   TProcess = "tokenlist"
)

func handleProcessInitialization(rawProcess *string) TProcess {
	switch strings.ToLower(*rawProcess) {
	case string(ProcessServer):
		process = ProcessServer
	case string(ProcessPartnerFees):
		process = ProcessPartnerFees
	case string(ProcessTokenList):
		process = ProcessTokenList
	default:
		process = ProcessServer
	}
	return process
}
