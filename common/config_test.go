package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	defer func() {
		if rec := recover(); rec != nil {
			t.Errorf("Loading config failed: %v\n", rec)
		}
	}()

	config := LoadConfig("../.env")

	assert.IsType(t, Config{}, config)
}

func TestLoadConfigException(t *testing.T) {
	defer func() {
		if rec := recover(); rec != nil {
			assert.True(t, true, "Load config throw an expected panic")
		}
	}()

	LoadConfig(".env")
	assert.True(t, false, "Load config must throw panic if .env file not found")
}

func TestGetConfig(t *testing.T) {
	config := GetConfig()
	appName := config.AppName
	assert.IsType(t, Config{}, config)

	config.AppName = "undefined"
	config = GetConfig()
	assert.Equal(t, appName, config.AppName)
}
