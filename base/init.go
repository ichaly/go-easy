package base

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var Initializer = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewCache),
	fx.Provide(NewConnect),
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
