package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	assert := assert.New(t)
	logger := Setup()
	assert.NotNil(logger, "must create")
}
