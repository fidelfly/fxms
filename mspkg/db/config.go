package db

import "fmt"

type Config struct {
	Host     string
	Port     int64
	Schema   string
	User     string
	Password string
}

func (db Config) getUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", db.User, db.Password, db.Host, db.Port, db.Schema)
}

func (db Config) getTarget() string {
	return fmt.Sprintf("%s:%d", db.Host, db.Port)
}

type Configurable interface {
	GetDbConfig() *Config
}
