package main

type TProcess string

const (
	ProcessServer TProcess = "server"
)

func handleProcessInitialization(rawProcess *string) TProcess {
	return ProcessServer
}
