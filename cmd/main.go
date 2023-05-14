// -build release
package main

import (
	"flag"
	"os"

	"github.com/iamasmith/skeleton/internal/app"
	"github.com/iamasmith/skeleton/internal/config"
)

func main() {
	config.Config.FlagSet = flag.NewFlagSet("standard", flag.ExitOnError)
	config.ParseArgs(os.Args[1:])
	app.Setup().Start()
}
