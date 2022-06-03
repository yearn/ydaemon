package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/majorfi/ydaemon/internal/controllers"
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/utils"
	"github.com/urfave/cli"
)

var commands = []cli.Command{
	startAPI(),
}

func startAPI() cli.Command {
	var env string

	return cli.Command{
		Name: "start-api",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env",
				Value:       ".env",
				Usage:       "which env file do you want to use",
				Destination: &env,
			},
		},
		Action: func(c *cli.Context) error {
			_ = godotenv.Load(env)
			utils.InitEnvironment()
			// utils.InitDB()

			go utils.RefreshPrices()
			go ethereum.InitDaemon(true)
			return controllers.NewRouter().Run()
		},
	}
}

func main() {
	api := cli.NewApp()
	api.Commands = commands
	if err := api.Run(os.Args); err != nil {
		log.Fatalf("failed to run the command: %v\n", err)
	}
}
