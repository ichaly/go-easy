package base

import (
	"context"
	"github.com/ichaly/go-easy/base/logger"
	"testing"
)

func TestNewConnect(t *testing.T) {
	cfg, err := NewConfig()
	if err != nil {
		logger.Panic(err)
		return
	}
	cache := NewCache(cfg)
	conn, err := NewConnect(cfg, cache)
	var rows []interface{}
	conn.GetDB(context.Background()).Raw("SELECT 1").Scan(&rows)
	logger.Debugf(">>>>>>>>>>>>>>>>>>>>%v", rows)
}
