package core

import (
	"context"
	"github.com/ichaly/go-easy/base"
)

type ITeamDao interface {
	base.IDao[Team]
	ListAll(ctx context.Context) ([]Team, error)
}

func NewTeamDao() ITeamDao {
	return teamDao{base.NewDao[Team]()}
}

type teamDao struct {
	*base.Dao[Team]
}

func (my teamDao) ListAll(ctx context.Context) ([]Team, error) {
	var teams []Team
	r := my.DB(ctx).Find(&teams)
	return teams, r.Error
}
