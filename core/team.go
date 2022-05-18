package core

import (
	"github.com/ichaly/go-easy/base"
	"github.com/ichaly/go-easy/base/logger"
)

func init() {
	logger.Debug("Team init ...")
	base.RegisterAutoMigrateModels(&Team{})
}

// TableName 自定义表名
func (Team) TableName() string {
	return "core_team"
}

type Team struct {
	base.Entity
}
