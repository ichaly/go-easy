package core

import (
	"github.com/ichaly/go-easy/base"
	"github.com/ichaly/go-easy/base/logger"
)

func init() {
	logger.Debug("Role init ...")
	base.RegisterMigrateModels(&Role{})
}

// TableName 自定义表名
func (Role) TableName() string {
	return "core_role"
}

type Role struct {
	base.Entity
}
