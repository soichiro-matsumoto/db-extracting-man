package config

import (
	"strconv"

	"github.com/BurntSushi/toml"
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

type Database struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Pass     string `toml:"pass"`
	System   string `toml:"system"`
	Port     int    `toml:"port"`
	Schema   string `toml:"schema"`
	Encoding string `toml:"encoding"`
}

func (this *Database) ToString() string {
	return "	Host	: " + this.Host + "\n" +
		"	System	: " + this.System + "\n" +
		"	Port	: " + strconv.Itoa(this.Port) + "\n" +
		"	Schema	: " + this.Schema + "\n" +
		"	Encoding: " + this.Encoding
}
