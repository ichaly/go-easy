package base

import (
	"context"
	"errors"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Result[T Entity] struct {
	Data  []T
	Page  int64
	Size  int64
	Total int64
}

type IDao[T Entity] interface {
	Save(ctx context.Context, t *T) (rows int64, err error)
	Delete(ctx context.Context, ids []uint64) (rows int64, err error)
	Update(ctx context.Context, t T) (rows int64, err error)
	List(ctx context.Context, query Query) (Result[T], error)
}

type Dao[T Entity] struct {
}

func (my Dao[T]) db(ctx context.Context) *gorm.DB {
	var conn Connect
	fx.Populate(conn)
	return conn.GetDB(ctx)
}

func (my Dao[T]) Save(ctx context.Context, t *T) (rows int64, err error) {
	r := my.db(ctx).Save(t)
	return r.RowsAffected, r.Error
}

func (my Dao[T]) Delete(ctx context.Context, ids []uint64) (rows int64, err error) {
	model := new(T)
	r := my.db(ctx).Delete(&model, ids)
	return r.RowsAffected, r.Error
}

func (my Dao[T]) Update(ctx context.Context, t T) (rows int64, err error) {
	if t.ID <= 0 {
		err = errors.New("ID can't be Zero")
		return
	}
	r := my.db(ctx).Model(&t).Updates(t)
	return r.RowsAffected, r.Error
}

func (my Dao[T]) List(ctx context.Context, query Query) (Result[T], error) {
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
