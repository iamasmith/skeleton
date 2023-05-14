// -build release
package main

import (
	"os"
	"os/signal"

	"github.com/iamasmith/skeleton/internal/app"
	"github.com/iamasmith/skeleton/internal/config"
)

func main() {
	config.Config.ParseArgs(os.Args[1:])
	server, app := app.Setup()
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
	)
	go func() {
		<-c
		app.Stop()
		server.Stop()
	}()
	server.Start()
}
