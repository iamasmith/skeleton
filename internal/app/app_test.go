package app

import (
	"testing"

	"github.com/iamasmith/skeleton/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	assert := assert.New(t)
	config.Config.ResetForTest()
	config.Config.ParseArgs([]string{})
	assert.NotNil(Setup())
}
