package main

import (
	"github.com/majorfi/ydaemon/internal/controllers"
	"github.com/majorfi/ydaemon/internal/daemons"
)

func main() {
	go daemons.SummonDaemons(1, 0)
	go daemons.SummonDaemons(250, 0)
	go daemons.SummonDaemons(42161, 0)

	controllers.NewRouter().Run()
}
