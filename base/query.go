package base

//https://github.com/micrease/gorme/blob/master/gorme.go
//https://studygolang.com/articles/26966
//https://juejin.cn/post/7078279187471679518
type Direction string

const (
	ASC  Direction = "ASC"
	DESC Direction = "DESC"
)

type Operator string

const (
	EQ       Operator = "EQ"
	NE       Operator = "NE"
	GT       Operator = "GT"
	GE       Operator = "GE"
	LT       Operator = "LT"
	LE       Operator = "LE"
	IN       Operator = "IN"
	LIKE     Operator = "LIKE"
	NOT_IN   Operator = "NOT_IN"
	NOT_LIKE Operator = "NOT_LIKE"
)

type Connector string

const (
	AND Connector = "AND"
	OR  Connector = "OR"
)

type Where struct {
	Column    string
	Operator  Operator
	Connector Connector
	Value     interface{}
}

type Order struct {
	Column    string
	Direction Direction
}

type Query struct {
	Condition []Where
	Sort      []Order
	Size      int64
	Page      int64
	Offset    uint64
}
