package base

import (
	"context"
	"errors"
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
	Delete(ctx context.Context, ids []uint64) (rows int64, err error)
	Update(ctx context.Context, t T) (rows int64, err error)
	List(ctx context.Context, query Query) (Result[T], error)
}

type Dao[T model] struct {
}

func (my Dao[T]) DB(ctx context.Context) *gorm.DB {
	var conn Connect
	fx.Populate(conn)
	return conn.GetDB(ctx)
}

func (my Dao[T]) Save(ctx context.Context, t *T) (rows int64, err error) {
	r := my.DB(ctx).Save(t)
	return r.RowsAffected, r.Error
}

func (my Dao[T]) Delete(ctx context.Context, ids []uint64) (rows int64, err error) {
	model := new(T)
	r := my.DB(ctx).Delete(&model, ids)
	return r.RowsAffected, r.Error
}

func (my Dao[T]) Update(ctx context.Context, t T) (rows int64, err error) {
	if t.GetID() <= 0 {
		err = errors.New("ID can't be Zero")
		return
	}
	r := my.DB(ctx).Model(&t).Updates(t)
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
