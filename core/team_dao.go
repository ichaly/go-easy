package core

import (
	"context"
	"github.com/ichaly/go-easy/base"
)

type ITeamDao interface {
	base.IDao[Team]
	ListAll(ctx context.Context) ([]Team, error)
}

func NewTeamDao(conn base.Connect) ITeamDao {
	return teamDao{IDao: base.NewDao[Team](conn)}
}

type teamDao struct {
	base.IDao[Team]
}

func (my teamDao) ListAll(ctx context.Context) ([]Team, error) {
	var teams []Team
	r := my.WithContext(ctx).Find(&teams)
	return teams, r.Error
}
