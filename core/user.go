package core

import (
	"github.com/ichaly/go-easy/base"
	"github.com/ichaly/go-easy/base/logger"
	"time"
)

func init() {
	logger.Debug("User init ...")
	base.RegisterMigrateModels(&User{})
}

// TableName 自定义表名
func (User) TableName() string {
	return "core_user"
}

type User struct {
	Username      string     `gorm:"size:50;unique_index;comment:用户名;"`
	Mobile        string     `gorm:"size:50;unique_index;comment:手机号;"`
	Email         string     `gorm:"size:50;unique_index;comment:邮箱;"`
	Avatar        string     `gorm:"size:500;comment:头像;"`
	Nickname      string     `gorm:"size:100;comment:昵称;"`
	Password      string     `gorm:"size:100;comment:密码;"`
	Gender        Gender     `gorm:"size:50;default:SECRET;comment:性别;"`
	Province      string     `gorm:"size:100;comment:省份;"`
	City          string     `gorm:"size:100;comment:城市;"`
	District      string     `gorm:"size:100;comment:区县;"`
	Idcard        string     `gorm:"size:50;comment:身份证号;"`
	Name          string     `gorm:"size:50;comment:真实姓名;"`
	Enabled       bool       `gorm:"comment:是否可用;"`
	Locked        bool       `gorm:"comment:是否锁定;"`
	Birthday      *time.Time `gorm:"comment:生日;"`
	ExpireTime    *time.Time `gorm:"comment:过期时间;"`
	LastLoginTime *time.Time `gorm:"comment:最后登录;"`
	base.Entity
}

type Gender string

const (
	SECRET Gender = "SECRET" //保密
	MALE   Gender = "MALE"   //男性
	FEMALE Gender = "FEMALE" //女性
)
