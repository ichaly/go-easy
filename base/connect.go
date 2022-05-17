package base

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ichaly/go-easy/base/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"reflect"
	"strings"
	"time"
)

type Connect struct{}

type ctxTransactionKey struct{}

var (
	cfg       Config
	db        *gorm.DB
	injectors []func(db *gorm.DB)
)

func NewConnect(config Config, cache Cache) (conn Connect, err error) {
	cfg = config
	master := config.DataSource
	if db, err = gorm.Open(buildDialect(master), &gorm.Config{}); err != nil {
		logger.Panic(err)
		return
	}
	if err = db.Use(cache); err != nil {
		logger.Panic(err)
		return
	}
	if len(master.Sources) > 0 || len(master.Replicas) > 0 {
		sources := make([]gorm.Dialector, len(master.Sources))
		replicas := make([]gorm.Dialector, len(master.Replicas))
		for i, v := range master.Sources {
			sources[i] = buildDialect(v)
		}
		for i, v := range master.Replicas {
			replicas[i] = buildDialect(v)
		}
		if err = db.Use(dbresolver.Register(dbresolver.Config{
			Sources: sources, Replicas: replicas, Policy: dbresolver.RandomPolicy{},
		})); err != nil {
			logger.Panic(err)
			return
		}
	}
	var sqlDb *sql.DB
	if sqlDb, err = db.DB(); err != nil {
		logger.Panic(err)
		return
	}
	sqlDb.SetMaxIdleConns(10)                   //最大空闲连接数
	sqlDb.SetMaxOpenConns(30)                   //最大连接数
	sqlDb.SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时
	//执行回调
	for _, v := range injectors {
		v(db)
	}
	conn = Connect{}
	return
}

func buildDialect(ds Database) gorm.Dialector {
	dsn := strings.Trim(ds.Url, "")
	args := []interface{}{ds.Username, ds.Password, ds.Host, ds.Port, ds.Name}
	if ds.Type == "mysql" {
		dsn = map[bool]string{true: dsn, false: fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", args...,
		)}[len(dsn) > 0]
		return mysql.Open(dsn)
	} else if ds.Type == "pgsql" {
		dsn = map[bool]string{true: dsn, false: fmt.Sprintf(
			"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable TimeZone=Asia/Shanghai", args...,
		)}[len(dsn) > 0]
		return postgres.Open(dsn)
	}
	return nil
}

func (my *Connect) GetDB(ctx context.Context) *gorm.DB {
	val := ctx.Value(ctxTransactionKey{})
	if val != nil {
		tx, ok := val.(*gorm.DB)
		if !ok {
			logger.Panicf("unexpect context value type: %s", reflect.TypeOf(tx))
			return nil
		}
		return tx
	}
	return db.WithContext(ctx)
}

func WithTransaction(ctx context.Context, handle func(tx context.Context) error) error {
	return db.WithContext(ctx).Transaction(func(db *gorm.DB) error {
		return handle(context.WithValue(ctx, ctxTransactionKey{}, db))
	})
}

// RegisterInjector 注册回调
func RegisterInjector(f func(*gorm.DB)) {
	injectors = append(injectors, f)
}

// SetupTableModel 自动初始化表结构
func SetupTableModel(db *gorm.DB, models ...interface{}) {
	if !cfg.AutoMigrate {
		return
	}
	if err := db.AutoMigrate(models...); err != nil {
		logger.Fatal(err)
	}
}
