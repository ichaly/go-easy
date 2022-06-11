package core

import (
	"context"
	"github.com/ichaly/go-easy/base"
)

type IUserDao interface {
	base.IDao[User]
	ListAll(ctx context.Context) ([]User, error)
}

func NewUserDao(conn base.Connect) IUserDao {
	return userDao{IDao: base.NewDao[User](conn)}
}

type userDao struct {
	base.IDao[User]
}

func (my userDao) ListAll(ctx context.Context) ([]User, error) {
	var users []User
	r := my.WithContext(ctx).Find(&users)
	return users, r.Error
}
