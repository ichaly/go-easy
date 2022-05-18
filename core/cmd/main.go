package main

import (
	"context"
	"github.com/ichaly/go-easy/base"
	"github.com/ichaly/go-easy/core"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		//禁用fx 默认logger
		fx.NopLogger,
		base.Initializer,
		core.Initializer,
	)
	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}
	if err := app.Stop(context.Background()); err != nil {
		panic(err)
	}
}
