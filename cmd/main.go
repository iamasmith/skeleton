// -build release
package main

import (
	"github.com/iamasmith/skeleton/internal/app"
	"github.com/iamasmith/skeleton/internal/server"
)

func main() {
	s := server.New(":8000")
	app.Setup(s)
	s.Start()
}
