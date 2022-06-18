package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"log-agent/internal/util"
)

func TestLoad(t *testing.T) {
	config, err := Load(util.TestDataPath("config.ini"))
	assert.NoError(t, err)
	assert.Equal(t, "日志收集处理", config.AppName)
	assert.True(t, len(config.Handler) > 0)
	laravel, ok := config.Handler["log.laravel"]
	assert.True(t, ok)
	assert.Equal(t, "/var/log/laravel.log", laravel.Path)
}
