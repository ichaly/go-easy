package core

import (
	"context"
	"github.com/ichaly/go-easy/base"
)

type ITeamService interface {
	base.IService[Team]
	ListAll(ctx context.Context) ([]Team, error)
}

func NewTeamService(dao ITeamDao) ITeamService {
	return teamService{dao: dao}
}

type teamService struct {
	dao ITeamDao
	base.Service[Team]
}

func (my teamService) ListAll(ctx context.Context) ([]Team, error) {
	return my.dao.ListAll(ctx)
}
