package base

import "context"

type IService[T Entity] interface {
	Save(ctx context.Context, t *T) (rows int64, err error)
	Delete(ctx context.Context, ids []uint64) (rows int64, err error)
	Update(ctx context.Context, t T) (rows int64, err error)
	List(ctx context.Context, query Query) (Result[T], error)
}

type Service[T Entity] struct {
	dao IDao[T]
}

func (my Service[T]) Save(ctx context.Context, t *T) (rows int64, err error) {
	return my.dao.Save(ctx, t)
}

func (my Service[T]) Delete(ctx context.Context, ids []uint64) (rows int64, err error) {
	return my.dao.Delete(ctx, ids)
}

func (my Service[T]) Update(ctx context.Context, t T) (rows int64, err error) {
	return my.dao.Update(ctx, t)
}

func (my Service[T]) List(ctx context.Context, query Query) (res Result[T], err error) {
	return my.dao.List(ctx, query)
}
