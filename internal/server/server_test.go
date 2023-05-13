package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartStop(t *testing.T) {
	server := New(":8000")
	go server.Start()
	server.Stop()
}

func TestStartFail(t *testing.T) {
	assert := assert.New(t)
	const binding = ":8000"
	a := New(binding)
	go a.Start()
	defer a.Stop()
	time.Sleep(time.Duration(1) * time.Second)
	b := New(binding)
	assert.Error(b.Start(), "should error")
	b.Stop()
}
