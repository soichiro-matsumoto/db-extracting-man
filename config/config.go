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
	Default Default `toml:"default"`
	Databases []Database `toml:"database"`
}

func (this *Config) GetDatabase(name string) (*Database, error){

	where := linq.From(this.Databases).Where(func(c interface{}) bool{
		return c.(Database).Name == name
	})

	if !where.Any() {
		return nil, errors.New(fmt.Sprintf("config name not defined : %s", name))
	}

	db := where.First().(Database)
	return &db, nil
}

type Default struct {
	Db     string `toml:"db"`
	Input  string `toml:"input"`
	Output string `toml:"output"`
}

type Database struct {
	Name     string `toml:"name"`
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Pass     string `toml:"pass"`
	System   string `toml:"system"`
	Port     int    `toml:"port"`
	Schema   string `toml:"schema"`
	Charset string `toml:"charset"`
	Timeout  int    `toml:"timeout"`
}

func (this *Database) ToString() string {
	return fmt.Sprintf(`
	Name    : %s
	Host    : %s
	System  : %s
	Port    : %d
	Schema  : %s
	Charset : %s
	Timeout : %d
	`, this.Name, this.Host, this.System, this.Port, this.Schema, this.Charset, this.Timeout);
}
