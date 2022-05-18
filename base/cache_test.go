package base

import (
	"github.com/ichaly/go-easy/base/logger"
	"testing"
)

func TestNewCache(t *testing.T) {
	var cfg, err = NewConfig()
	if err != nil {
		logger.Panic(err)
		return
	}
	cache := NewCache(cfg)
	logger.Debug(cache)
}
