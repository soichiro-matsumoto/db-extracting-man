package config

import (
	"github.com/BurntSushi/toml"
	"github.com/ahmetb/go-linq"
	"fmt"
	"errors"
)

func GetConfig() *Config {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}
	return &config
}

type Config struct {
	Databases []Database `toml:"database"`
}

func (this *Config) GetDatabase(key string) (*Database, error){

	where := linq.From(this.Databases).Where(func(c interface{}) bool{
		return c.(Database).Key == key
	})

	if !where.Any() {
		return nil, errors.New(fmt.Sprintf("config key not defined : %s", key))
	}

	db := where.First().(Database)
	return &db, nil
}

type Database struct {
	Key      string `toml:"key"`
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Pass     string `toml:"pass"`
	System   string `toml:"system"`
	Port     int    `toml:"port"`
	Schema   string `toml:"schema"`
	Encoding string `toml:"encoding"`
}

func (this *Database) ToString() string {
	return fmt.Sprintf(`
	Key     : %s
	Host    : %s
	System  : %s
	Port    : %d
	Schema  : %s
	Encoding: %s
	`, this.Key, this.Host, this.System, this.Port, this.Schema, this.Encoding);
}
