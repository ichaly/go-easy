package base

import (
	"context"
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

func NewDao[T model](conn Connect) IDao[T] {
	return dao[T]{
		db: func(ctx context.Context) *gorm.DB {
			return conn.GetDB(ctx)
		},
	}
}

type dao[T model] struct {
	db func(ctx context.Context) *gorm.DB
}

func (my dao[T]) Save(ctx context.Context, t *T) (rows int64, err error) {
	var r *gorm.DB
	if (*t).GetID() <= 0 {
		r = my.db(ctx).Create(t)
	} else {
		r = my.db(ctx).Model(t).Updates(t)
	}
	return r.RowsAffected, r.Error
}

func (my dao[T]) Delete(ctx context.Context, ids []uint64) (rows int64, err error) {
	model := new(T)
	r := my.db(ctx).Delete(&model, ids)
	return r.RowsAffected, r.Error
}

func (my dao[T]) List(ctx context.Context, query Query) (Result[T], error) {
	var rows []T
	var count int64
	r := my.db(ctx).Find(&rows).Count(&count)
	return Result[T]{
		Total: count,
		Page:  query.Page,
		Size:  query.Size,
		Data:  rows,
	}, r.Error
}
