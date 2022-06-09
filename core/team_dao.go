package core

import (
	"context"
	"github.com/ichaly/go-easy/base"
	"gorm.io/gorm"
)

type ITeamDao interface {
	base.IDao[Team]
	ListAll(ctx context.Context) ([]Team, error)
}

func NewTeamDao(conn base.Connect) ITeamDao {
	return teamDao{
		IDao: base.NewDao[Team](conn),
		db: func(ctx context.Context) *gorm.DB {
			return conn.GetDB(ctx)
		},
	}
}

type teamDao struct {
	base.IDao[Team]
	db func(ctx context.Context) *gorm.DB
}

func (my teamDao) ListAll(ctx context.Context) ([]Team, error) {
	var teams []Team
	r := my.db(ctx).Find(&teams)
	return teams, r.Error
}
