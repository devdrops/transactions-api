package app

import (
	"os"
)

type Config struct {
	DbUser string
	DbPass string
	DbName string
	DbHost string
}

func (c *Config) NewConfig() *Config {
	return &Config{
		DbUser: os.Getenv("POSTGRES_USER"),
		DbPass: os.Getenv("POSTGRES_PASSWORD"),
		DbName: os.Getenv("DATABASE_NAME"),
		DbHost: os.Getenv("DATABASE_HOST"),
	}
}