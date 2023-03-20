package main

type TProcess string

const (
	ProcessServer      TProcess = "server"
	ProcessPartnerFees TProcess = "partnerFees"
)

func handleProcessInitialization(rawProcess *string) TProcess {
	switch *rawProcess {
	case string(ProcessServer):
		process = ProcessServer
	case string(ProcessPartnerFees):
		process = ProcessPartnerFees
	default:
		process = ProcessServer
	}
	return process
}
