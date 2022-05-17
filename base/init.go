package base

import (
	"go.uber.org/fx"
)

var Initializer = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewCache),
	fx.Provide(NewConnect),
)
