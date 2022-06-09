package base

import (
	"context"
)

type IService[T model] interface {
	Save(ctx context.Context, t *T) (rows int64, err error)
	List(ctx context.Context, query Query) (Result[T], error)
	Delete(ctx context.Context, ids []uint64) (rows int64, err error)
}

func NewService[T model](dao IDao[T]) IService[T] {
	return service[T]{dao}
}

type service[T model] struct {
	dao IDao[T]
}

func (my service[T]) Save(ctx context.Context, t *T) (rows int64, err error) {
	return my.dao.Save(ctx, t)
}

func (my service[T]) Delete(ctx context.Context, ids []uint64) (rows int64, err error) {
	return my.dao.Delete(ctx, ids)
}

func (my service[T]) List(ctx context.Context, query Query) (res Result[T], err error) {
	return my.dao.List(ctx, query)
}
