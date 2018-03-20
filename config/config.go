package config

import "github.com/BurntSushi/toml"

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
	Encoding string `toml:"encoding"`
}
