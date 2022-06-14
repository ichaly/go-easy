package base

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Direction string

const (
	ASC  Direction = "ASC"
	DESC Direction = "DESC"
)

func (my Direction) String() (val string, err error) {
	switch my {
	case ASC:
		val = "ASC"
		return
	case DESC:
		val = "DESC"
		return
	default:
		err = errors.New("unsupported direction")
		return
	}
}

type Predicate string

const (
	EQ       = Predicate("=")
	NE       = Predicate("<>")
	GT       = Predicate(">")
	GE       = Predicate(">=")
	LT       = Predicate("<")
	LE       = Predicate("<=")
	IN       = Predicate("IN")
	LIKE     = Predicate("LIKE")
	NOT_IN   = Predicate("NOT IN")
	NOT_LIKE = Predicate("NOT LIKE")
)

type Connector string

const (
	AND Connector = "AND"
	OR  Connector = "OR"
)

func (my Connector) String() (val string, err error) {
	switch my {
	case AND:
		val = "AND"
		return
	case OR:
		val = "OR"
		return
	default:
		err = errors.New("unsupported connector")
		return
	}
}

type Where struct {
	Column    string
	Operator  Predicate
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
	Size      int
	Page      int
	Offset    interface{}
}

type QueryOption func(*Query)

func WithSize(size int) QueryOption {
	return func(q *Query) {
		q.Size = size
	}
}

func WithPage(page int) QueryOption {
	return func(q *Query) {
		q.Page = page
	}
}

func WithOffset(offset int) QueryOption {
	return func(q *Query) {
		q.Offset = offset
	}
}

func BuildQuery(db *gorm.DB, options ...QueryOption) *gorm.DB {
	query := &Query{Size: 10}
	for _, o := range options {
		o(query)
	}
	for _, c := range query.Condition {
		db.Where(fmt.Sprintf("%s %s ?", c.Column, c.Operator), c.Value)
	}
	// 处理分页
	if query.Size >= 1000 {
		query.Size = 1000
	}
	db.Limit(query.Size)
	if query.Page > 0 {
		offset := (query.Page - 1) * query.Size
		db.Offset(offset).Limit(query.Size)
	}
	// 处理排序
	for _, v := range query.Sort {
		if d, e := v.Direction.String(); e == nil {
			db.Order(fmt.Sprintf("%s %s", v.Column, d))
		}
	}
	return db
}
