package apiserver

import "github.com/SidorkinAlex/go-rest-api/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLavel string `loml:"log_level"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLavel: "debug",
		Store:    store.NewConfig(),
	}
}
