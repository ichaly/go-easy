package core

import (
	"context"
	"github.com/ichaly/go-easy/base"
)

type IUserService interface {
	base.IService[User]
	ListAll(ctx context.Context) ([]User, error)
	SignIn(ctx context.Context, username string, password string) (user *User, err error)
}

func NewUserService(dao IUserDao) IUserService {
	return userService{dao, base.NewService[User](dao)}
}

type userService struct {
	dao IUserDao
	base.IService[User]
}

func (my userService) ListAll(ctx context.Context) ([]User, error) {
	return my.dao.ListAll(ctx)
}

func (my userService) SignIn(ctx context.Context, username string, password string) (user *User, err error) {
	err = base.Transaction(ctx, func(tx context.Context) error {
		user = &User{Username: username, Password: password}
		_, err = my.dao.Save(tx, user)
		return err
	})
	return
}
