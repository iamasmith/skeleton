package app

import (
	"testing"

	"github.com/iamasmith/skeleton/internal/server"
)

func TestSetup(t *testing.T) {
	s := server.New(":8000")
	Setup(s)
}
