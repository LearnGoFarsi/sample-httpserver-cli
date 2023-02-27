package config

import "github.com/caarlos0/env/v6"

type Server struct {
	Host string `env:"SERVERHOST"`
	Port int    `env:"SERVERPORT"`
}

type Config struct {
	Server
}

func LoadConfigOrPanic() Config {
	var config *Config = new(Config)
	if err := env.Parse(config); err != nil {
		panic(err)
	}

	return *config
}
