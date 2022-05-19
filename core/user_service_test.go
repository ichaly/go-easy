package core

import (
	"context"
	"github.com/ichaly/go-easy/base"
	"go.uber.org/fx"
	"testing"
)

func TestListAll(t *testing.T) {
	app := fx.New(
		//禁用fx 默认logger
		fx.NopLogger,
		base.Initializer,
		Initializer,
		fx.Invoke(func(s IUserService) {
			s.ListAll(context.Background())
		}),
	)
	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}
	if err := app.Stop(context.Background()); err != nil {
		panic(err)
	}
}
