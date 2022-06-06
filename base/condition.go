package base

//https://github.com/micrease/gorme/blob/master/gorme.go
//https://studygolang.com/articles/26966
//https://juejin.cn/post/7078279187471679518
type Condition struct {
	Field     string
	Operator  string
	Connector string
	Value     interface{}
}

type Pagination struct {
	page  uint64
	size  uint64
	total uint64
}

type Order struct {
	Field     string
	Direction Direction
}

type Sort struct {
	Orders []Order
}

type Direction uint32

const (
	ASC Direction = 1 << (32 - 1 - iota)
	DESC
)
