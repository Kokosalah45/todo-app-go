package config

import "fmt"

type Config struct {
	Port string
	Env string

}

func NewConfig(port , env string ) *Config {
	addr := fmt.Sprintf(":%s", port)
	return &Config{
		Port: addr,
		Env: env,
	}
}
