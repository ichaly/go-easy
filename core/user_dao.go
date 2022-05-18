package core

import (
	"context"
	"errors"
	"github.com/ichaly/go-easy/base"
	"gorm.io/gorm"
)

type IUserDao interface {
	List(ctx context.Context) []User
	Save(ctx context.Context, user *User) (rows int64, err error)
	Update(ctx context.Context, user User) (rows int64, err error)
	Delete(ctx context.Context, ids []uint64) (rows int64, err error)
}

func NewUserDao(conn base.Connect) IUserDao {
	return userDao{func(ctx context.Context) *gorm.DB {
		return conn.GetDB(ctx)
	}}
}

type userDao struct {
	db func(ctx context.Context) *gorm.DB
}

func (my userDao) Delete(ctx context.Context, ids []uint64) (rows int64, err error) {
	r := my.db(ctx).Delete(&User{}, ids)
	if err = r.Error; err != nil {
		return
	}
	rows = r.RowsAffected
	return
}

func (my userDao) Update(ctx context.Context, user User) (rows int64, err error) {
	if user.ID <= 0 {
		err = errors.New("ID can't be Zero")
		return
	}
	r := my.db(ctx).Model(&user).Updates(user)
	if err = r.Error; err != nil {
		return
	}
	rows = r.RowsAffected
	return
}

func (my userDao) List(ctx context.Context) []User {
	var users []User
	my.db(ctx).Find(&users)
	return users
}

func (my userDao) Save(ctx context.Context, user *User) (rows int64, err error) {
	r := my.db(ctx).Save(user)
	if err = r.Error; err != nil {
		return
	}
	rows = r.RowsAffected
	return
}
