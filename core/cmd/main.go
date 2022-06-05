package main

import (
	"github.com/ichaly/go-easy/core"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		//禁用fx 默认logger
		fx.NopLogger,
		core.Initializer,
	).Run()
}
