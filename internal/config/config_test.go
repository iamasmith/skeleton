package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestLogLevelFlag(t *testing.T) {
	assert := assert.New(t)
	ResetForTest()
	ParseArgs([]string{})
	assert.Equal(zapcore.InfoLevel, Config.LogLevel)
	ResetForTest()
	ParseArgs([]string{"-level=debug"})
	assert.Equal(zapcore.DebugLevel, Config.LogLevel)
	ResetForTest()
	ParseArgs([]string{"-level=unknown"})
	assert.Equal(zapcore.DebugLevel, Config.LogLevel)
}
