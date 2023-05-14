package app

import (
	"testing"

	"github.com/iamasmith/skeleton/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	assert := assert.New(t)
	config.ResetForTest()
	config.ParseArgs([]string{})
	assert.NotNil(Setup())
}
