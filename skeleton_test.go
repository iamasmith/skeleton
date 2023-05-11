package skeleton

import (
	"testing"
	"time"
)

func TestStartStop(t *testing.T) {
	server := New(":8000")
	go server.Start()
	server.Stop()
}

func TestStartFail(t *testing.T) {
	const binding = ":8000"
	a := New(binding)
	go a.Start()
	defer a.Stop()
	time.Sleep(time.Duration(1) * time.Second)
	b := New(binding)
	if err := b.Start(); err == nil {
		t.Error("Expected startup error for 2nd server on same binding")
	}
	b.Stop()
}
