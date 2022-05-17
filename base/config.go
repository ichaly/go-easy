package base

import (
	"encoding/json"
	"github.com/ichaly/go-easy/base/logger"
	"os"
)

type Database struct {
	Type     string     `json:"type"`
	Url      string     `json:"url"`
	Host     string     `json:"host"`
	Port     int        `json:"port"`
	Name     string     `json:"name"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Sources  []Database `json:"sources"`
	Replicas []Database `json:"replicas"`
}

type Config struct {
	AutoMigrate bool     `json:"autoMigrate"`
	DataSource  Database `json:"dataSource"`
	CacheStore  Database `json:"cacheStore"`
}

func NewConfig() (cfg Config, err error) {
	var f *os.File
	if f, err = os.Open("./cfg/config.json"); err != nil {
		logger.Panic(err)
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	decoder := json.NewDecoder(f)
	if err = decoder.Decode(&cfg); err != nil {
		logger.Panic(err)
		return
	}
	return
}
