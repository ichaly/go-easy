package base

import (
	"github.com/ichaly/go-easy/base/logger"
	"testing"
)

func TestNewConfig(t *testing.T) {
	if cfg, err := NewConfig(); err != nil {
		logger.Panic(err)
	} else {
		logger.Debug(cfg)
	}
}
