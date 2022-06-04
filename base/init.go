package base

import (
	_ "github.com/ichaly/go-env/auto"
	"go.uber.org/fx"
)

var Initializer = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewCache),
	fx.Provide(NewConnect),
	fx.Provide(NewEngine),
)
