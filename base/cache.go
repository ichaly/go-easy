package base

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	"github.com/ichaly/go-easy/base/logger"
	"github.com/ichaly/go-easy/base/utils"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"reflect"
	"time"
)

var (
	requestGroup singleflight.Group
)

type keyCacheContext struct{}

type Cache struct {
	Cache       *cache.Cache
	exp         time.Duration
	keyGenerate func(*gorm.DB) string
}

// Name `gorm.Plugin` implements.
func (my Cache) Name() string { return "gorm-cache" }

func NewCache(c Config) Cache {
	ds := c.CacheStore
	redisStore := store.NewRedis(redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", ds.Host, ds.Port),
		Password: ds.Password,
		DB:       0,
	}), nil)
	return Cache{
		cache.New(redisStore),
		30 * time.Minute,
		func(db *gorm.DB) string {
			return fmt.Sprintf(
				"%s:%s",
				db.Statement.Table,
				utils.MD5(db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)),
			)
		},
	}
}

// Initialize `gorm.Plugin` implements.
func (my Cache) Initialize(db *gorm.DB) error {
	if err := db.Callback().Query().Replace("gorm:query", my.query); err != nil {
		return err
	}
	if err := db.Callback().Query().After("gorm:query").Register(my.Name()+":after_query", my.afterQuery); err != nil {
		return err
	}
	if err := db.Callback().Create().After("gorm:create").Register(my.Name()+":after_create", my.afterUpdate); err != nil {
		return err
	}
	if err := db.Callback().Update().After("gorm:update").Register(my.Name()+":after_update", my.afterUpdate); err != nil {
		return err
	}
	if err := db.Callback().Delete().After("gorm:delete").Register(my.Name()+":after_delete", my.afterUpdate); err != nil {
		return err
	}
	return nil
}

// query replace gorm:query
func (my Cache) query(db *gorm.DB) {
	if db.DryRun || db.Error != nil {
		return
	}
	callbacks.BuildQuerySQL(db)
	cacheKey := my.keyGenerate(db)
	val, err := my.Cache.Get(db.Statement.Context, cacheKey)
	if err != nil {
		my.queryFromDB(db, cacheKey)
		return
	}
	_ = my.unmarshalToDB(val.(string), db)
}

func (my Cache) afterQuery(db *gorm.DB) {
	if v := db.Statement.Context.Value(keyCacheContext{}); v == nil {
		return
	}
	if db.Error != nil || db.Statement.Schema == nil {
		return
	}
	ok, value, err := my.queryResult(db)
	if !ok {
		return
	}
	cacheKey := my.keyGenerate(db)
	if err != nil {
		_ = db.AddError(err)
		return
	}
	if err := my.Cache.Set(db.Statement.Context, cacheKey, value, &store.Options{Expiration: my.exp, Tags: []string{db.Statement.Table}}); err != nil {
		_ = db.AddError(err)
		return
	}
}

func (my Cache) afterUpdate(db *gorm.DB) {
	total := db.Statement.RowsAffected
	if total <= 0 {
		return
	}
	if err := my.Cache.Invalidate(db.Statement.Context, store.InvalidateOptions{Tags: []string{db.Statement.Table}}); err != nil {
		_ = db.AddError(err)
	}
}

func (my Cache) unmarshalToDB(value string, db *gorm.DB) (err error) {
	reflectValue := db.Statement.ReflectValue
	if err = json.Unmarshal([]byte(value), db.Statement.Dest); err != nil {
		return err
	}
	elem := reflect.ValueOf(db.Statement.Dest)
	if !elem.CanSet() {
		elem = elem.Elem()
	}
	switch reflectValue.Kind() {
	case reflect.Struct:
		db.RowsAffected = 1
	case reflect.Slice, reflect.Array:
		db.RowsAffected = int64(elem.Len())
	default:
		logger.Infof("%s: %s, type not support", my.Name(), reflectValue.Kind().String())
	}
	reflectValue.Set(elem)
	return
}

func (my Cache) queryFromDB(db *gorm.DB, cacheKey string) {
	var (
		rows *sql.Rows
		err  error
	)
	var any interface{}
	any, err, _ = requestGroup.Do(cacheKey, func() (interface{}, error) {
		db.Statement.Context = context.WithValue(db.Statement.Context, keyCacheContext{}, 1)
		return db.Statement.ConnPool.QueryContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
	})
	rows = any.(*sql.Rows)
	if err != nil {
		_ = db.AddError(err)
		return
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	gorm.Scan(rows, db, 0)
}

func (my Cache) queryResult(db *gorm.DB) (bool, string, error) {
	var (
		fields       = db.Statement.Schema.Fields
		reflectValue = db.Statement.ReflectValue
	)
	switch reflectValue.Kind() {
	case reflect.Struct:
		valueData := make(map[string]interface{})
		for _, field := range fields {
			if fieldValue, isZero := field.ValueOf(db.Statement.Context, reflectValue); !isZero {
				valueData[field.Name] = fieldValue
			}
		}
		val, err := json.Marshal(valueData)
		return true, string(val), err
	case reflect.Slice, reflect.Array:
		lens := reflectValue.Len()
		valueArrData := make([]map[string]interface{}, lens)
		for _, field := range fields {
			for i := 0; i < lens; i++ {
				if fieldValue, isZero := field.ValueOf(db.Statement.Context, reflectValue.Index(i)); !isZero {
					if valueArrData[i] == nil {
						valueArrData[i] = make(map[string]interface{})
					}
					valueArrData[i][field.Name] = fieldValue
				}
			}
		}
		val, err := json.Marshal(valueArrData)
		return true, string(val), err
	default:
		logger.Infof("gorm-cache: %s, type not support", reflectValue.Kind().String())
	}
	return false, "", nil
}
