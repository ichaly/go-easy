package base

type Result[T any] struct {
	Data  []T
	Page  int64
	Size  int64
	Total int64
}
