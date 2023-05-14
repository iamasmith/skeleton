// -build release
package main

import (
	"github.com/iamasmith/skeleton/internal/app"
)

func main() {
	app.Setup().Start()
}
