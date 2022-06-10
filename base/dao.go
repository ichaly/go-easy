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
	WithContext(ctx context.Context) *gorm.DB
	Save(ctx context.Context, t *T) (rows int64, err error)
	List(ctx context.Context, query Query) (Result[T], error)
	Delete(ctx context.Context, ids []uint64) (rows int64, err error)
}

func NewDao[T model](conn Connect) IDao[T] {
	return dao[T]{conn}
}

type dao[T model] struct {
	conn Connect
}

func (my dao[T]) WithContext(ctx context.Context) *gorm.DB {
	return my.conn.GetDB(ctx)
}

func (my dao[T]) Save(ctx context.Context, t *T) (rows int64, err error) {
	var r *gorm.DB
	if (*t).GetID() <= 0 {
		r = my.WithContext(ctx).Create(t)
	} else {
		r = my.WithContext(ctx).Model(t).Updates(t)
	}
	return r.RowsAffected, r.Error
}

func (my dao[T]) Delete(ctx context.Context, ids []uint64) (rows int64, err error) {
	model := new(T)
	r := my.WithContext(ctx).Delete(&model, ids)
	return r.RowsAffected, r.Error
}

func (my dao[T]) List(ctx context.Context, query Query) (Result[T], error) {
	var rows []T
	var count int64
	r := my.WithContext(ctx).Find(&rows).Count(&count)
	return Result[T]{
		Total: count,
		Page:  query.Page,
		Size:  query.Size,
		Data:  rows,
	}, r.Error
}
