package main

import (
	"log"

	"github.com/pelletier/go-toml"
)

type ServerCfg struct {
	IP           string `toml:"ip"`
	Port         string `toml:"port"`
	InProduction bool   `toml:"in_production"`
}

type Config struct {
	Server ServerCfg `toml:"server"`
}

func getConfig(file string) Config {
	var cfg Config

	data, err := toml.LoadFile(file)
	if err != nil {
		log.Panic()
	}

	err = data.Unmarshal(&cfg)
	if err != nil {
		log.Panic()
	}

	return cfg
}
