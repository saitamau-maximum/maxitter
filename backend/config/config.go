package config

import (
	"os"
)

type (
	Config struct {
		MYSQL
	}

	MYSQL struct {
		MYSQL_USER,
		MYSQL_PASSWORD,
		MYSQL_HOST,
		MYSQL_PORT,
		MYSQL_DATABASE string
	}
)

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		value = fallback
	}
	return value
}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.MYSQL.MYSQL_USER = getEnv("MYSQL_USER", "user")
	cfg.MYSQL.MYSQL_PASSWORD = getEnv("MYSQL_PASSWORD", "password")
	cfg.MYSQL.MYSQL_HOST = getEnv("MYSQL_HOST", "database")
	cfg.MYSQL.MYSQL_PORT = getEnv("MYSQL_PORT", "3306")
	cfg.MYSQL.MYSQL_DATABASE = getEnv("MYSQL_DATABASE", "db")

	return cfg
}
