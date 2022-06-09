package base

import (
	"context"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type model interface {
	GetID() uint64
}

type Result[T model] struct {
	Data  []T
	Page  int64
	Size  int64
	Total int64
}

type IDao[T model] interface {
	Save(ctx context.Context, t *T) (rows int64, err error)
	List(ctx context.Context, query Query) (Result[T], error)
	Delete(ctx context.Context, ids []uint64) (rows int64, err error)
}

type Dao[T model] struct {
	DB func(ctx context.Context) *gorm.DB
}

func NewDao[T model]() *Dao[T] {
	var conn Connect
	fx.Populate(conn)
	return &Dao[T]{
		DB: func(ctx context.Context) *gorm.DB {
			return conn.GetDB(ctx)
		},
	}
}

func (my Dao[T]) Save(ctx context.Context, t *T) (rows int64, err error) {
	var r *gorm.DB
	if (*t).GetID() <= 0 {
		r = my.DB(ctx).Create(t)
	} else {
		r = my.DB(ctx).Model(t).Updates(t)
	}
	return r.RowsAffected, r.Error
}

func (my Dao[T]) Delete(ctx context.Context, ids []uint64) (rows int64, err error) {
	model := new(T)
	r := my.DB(ctx).Delete(&model, ids)
	return r.RowsAffected, r.Error
}

func (my Dao[T]) List(ctx context.Context, query Query) (Result[T], error) {
	var rows []T
	var count int64
	r := my.DB(ctx).Find(&rows).Count(&count)
	return Result[T]{
		Total: count,
		Page:  query.Page,
		Size:  query.Size,
		Data:  rows,
	}, r.Error
}
