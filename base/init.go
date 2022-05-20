package base

import (
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/fx"
)

var Initializer = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewCache),
	fx.Provide(NewConnect),
)
