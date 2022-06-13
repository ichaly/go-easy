package base

import (
	"errors"
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

func (my Operator) String() (val string, err error) {
	switch my {
	case EQ:
		val = "="
		return
	case NE:
		val = "<>"
		return
	case GT:
		val = ">"
		return
	case GE:
		val = ">="
		return
	case LT:
		val = "<"
		return
	case LE:
		val = "<="
		return
	case IN:
		val = "IN"
		return
	case LIKE:
		val = "LIKE"
		return
	case NOT_IN:
		val = "NOT IN"
		return
	case NOT_LIKE:
		val = "NOT LIKE"
		return
	default:
		err = errors.New("unsupported operator")
		return
	}
}

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

	if query.Size >= 1000 {
		query.Size = 1000
	}
	db.Limit(query.Size)
	if query.Page > 0 {
		offset := (query.Page - 1) * query.Size
		db.Offset(offset).Limit(query.Size)
	}
	return db
}
