package config

import (
	"os"
)

type DBConfig struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

func NewDBConfig() *DBConfig {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	return &DBConfig{
		dbName,
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
	}
}
