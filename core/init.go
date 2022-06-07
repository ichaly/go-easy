package core

import (
	"github.com/ichaly/go-easy/base"
	"go.uber.org/fx"
)

var Initializer = fx.Options(
	base.Initializer,
	fx.Provide(NewSchema),
	fx.Provide(NewUserDao),
	fx.Provide(NewUserService),
	fx.Provide(NewTeamDao),
	fx.Provide(NewTeamService),
)
