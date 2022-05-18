package core

import (
	"context"
	"github.com/ichaly/go-easy/base"
)

type IUserService interface {
	ListAll(ctx context.Context) []User
	SignIn(ctx context.Context, username string, password string) User
	Update(ctx context.Context, user User) (rows int64, err error)
	Delete(ctx context.Context, ids []uint64) (rows int64, err error)
}

func NewUserService(dao IUserDao) IUserService {
	return userService{dao}
}

type userService struct {
	dao IUserDao
}

func (my userService) SignIn(ctx context.Context, username string, password string) User {
	var user User
	_ = base.Transaction(ctx, func(tx context.Context) error {
		user = User{Username: username, Password: password}
		_, err := my.dao.Save(tx, &user)
		if err != nil {
			return err
		}
		return nil
	})
	return user
}

func (my userService) ListAll(ctx context.Context) []User {
	return my.dao.List(ctx)
}

func (my userService) Update(ctx context.Context, user User) (rows int64, err error) {
	return my.dao.Update(ctx, user)
}

func (my userService) Delete(ctx context.Context, ids []uint64) (rows int64, err error) {
	return my.dao.Delete(ctx, ids)
}
