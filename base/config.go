package base

import (
	"github.com/ichaly/go-easy/base/logger"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Database struct {
	Type     string     `yaml:"type"`
	Url      string     `yaml:"url"`
	Host     string     `yaml:"host"`
	Port     int        `yaml:"port"`
	Name     string     `yaml:"name"`
	Username string     `yaml:"username"`
	Password string     `yaml:"password"`
	Sources  []Database `yaml:"sources"`
	Replicas []Database `yaml:"replicas"`
}

type Config struct {
	AutoMigrate bool     `yaml:"autoMigrate"`
	DataSource  Database `yaml:"dataSource"`
	CacheStore  Database `yaml:"cacheStore"`
}

func NewConfig() (cfg Config, err error) {
	file, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		logger.Panicf("Read config err #%v ", err)
		return
	}
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		logger.Panicf("Unmarshal: %v", err)
	}
	return
}
