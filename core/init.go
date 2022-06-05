package core

import (
	"github.com/ichaly/go-easy/base"
	"go.uber.org/fx"
)

var Initializer = fx.Options(
	base.Initializer,
	fx.Provide(NewUserDao),
	fx.Provide(NewUserService),
	fx.Provide(NewSchema),
)
