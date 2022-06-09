package base

import (
	"context"
)

type IService[T model] interface {
	Save(ctx context.Context, t *T) (rows int64, err error)
	Delete(ctx context.Context, ids []uint64) (rows int64, err error)
	List(ctx context.Context, query Query) (Result[T], error)
}

type Service[T model] struct {
	dao IDao[T]
}

func NewService[T model](dao IDao[T]) *Service[T] {
	return &Service[T]{dao}
}

func (my Service[T]) Save(ctx context.Context, t *T) (rows int64, err error) {
	return my.dao.Save(ctx, t)
}

func (my Service[T]) Delete(ctx context.Context, ids []uint64) (rows int64, err error) {
	return my.dao.Delete(ctx, ids)
}

func (my Service[T]) List(ctx context.Context, query Query) (res Result[T], err error) {
	return my.dao.List(ctx, query)
}
