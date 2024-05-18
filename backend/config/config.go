package config

import "fmt"

type Config struct {
	Port string
	Env string
	Db *DBConfig
}

func NewConfig(port , env , DB *DBConfig) *Config {
	addr := fmt.Sprintf(":%s", port)
	return &Config{
		Port: addr,
		Env: "development",
		Db: DB,
	}
}

type DBConfig struct {
	Dsn string
}

func NewDBConfig(dsn string) *DBConfig {
	return &DBConfig{
		Dsn: dsn,
	}
}

