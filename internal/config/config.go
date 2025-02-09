package config

import (
	"fmt"
	"os"
)

func GetPort() string {
	if v := os.Getenv("PORT"); v != "" {
		return ":" + v
	}
	return ":8080"
}

func GetEnviroment() string {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		return "dev"
	}
	return env
}

type postgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func GetPostgresConfig() postgresConfig {
	sslMode := os.Getenv("POSTGRES_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	return postgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"), // TODO
		Port:     os.Getenv("POSTGRES_PORT"), // TODO
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  sslMode,
	}
}

func (v postgresConfig) DataSourceName() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", v.User, v.Password, v.DBName, v.SSLMode)
}
