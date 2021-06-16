package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLavel string `loml:"log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLavel: "debug",
	}
}
