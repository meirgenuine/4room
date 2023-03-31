package config

import (
	"os"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type ServerConfig struct {
	Address string
}

var (
	DBConfig   = &DatabaseConfig{}
	ServerConf = &ServerConfig{}
)

func LoadConfig() {
	// Load configurations from environment variables or set defaults
	DBConfig.Driver = getEnv("DB_DRIVER", "sqlite3")
	DBConfig.Host = getEnv("DB_HOST", "")
	DBConfig.Port = getEnv("DB_PORT", "")
	DBConfig.User = getEnv("DB_USER", "")
	DBConfig.Password = getEnv("DB_PASSWORD", "")
	DBConfig.Database = getEnv("DB_DATABASE", "4room.db")

	ServerConf.Address = getEnv("SERVER_ADDRESS", ":8080")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
