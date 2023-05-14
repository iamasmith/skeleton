// -build release
package main

import (
	"os"

	"github.com/iamasmith/skeleton/internal/app"
	"github.com/iamasmith/skeleton/internal/config"
)

func main() {
	config.Config.ParseArgs(os.Args[1:])
	app.Setup().Start()
}
