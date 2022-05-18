package core

import (
	"go.uber.org/fx"
)

var Initializer = fx.Options(
	fx.Provide(NewUserDao),
	fx.Provide(NewUserService),
)
