package config

import "fmt"

type Config struct {
	Port string
	Env string
	Db *DBConfig
}

type DBConfig struct {
	Dsn string
}

func NewConfig(port , env string , DB *DBConfig) *Config {
	addr := fmt.Sprintf(":%s", port)
	return &Config{
		Port: addr,
		Env: "development",
		Db: DB,
	}
}

func NewDBConfig(dsn string) *DBConfig {
	return &DBConfig{
		Dsn: dsn,
	}
}

