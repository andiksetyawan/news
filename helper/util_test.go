package helper

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnvDefault(t *testing.T) {
	os.Setenv("TEST_ENV", "")
	assert.Equal(t, GetEnv("TEST_ENV", "DEFAULT"), "DEFAULT")
}

func TestGetEnvFound(t *testing.T) {
	os.Setenv("TEST_ENV", "VALUE_ENV_FROM_OS")
	assert.Equal(t, GetEnv("TEST_ENV", "DEFAULT"), "VALUE_ENV_FROM_OS")
}
