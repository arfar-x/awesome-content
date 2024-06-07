package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config *AppConfig

type AppConfig struct {
	Name string
	Mode string
	DB   DBConfig
	Http HttpConfig
}

type HttpConfig struct {
	Host string
	Port int
}

// DBConfig Database configuration.
type DBConfig struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
	Charset  string
}

// InitConfig Initialize application configurations by environment variables.
func InitConfig() *AppConfig {
	if err := godotenv.Load(".env.app"); err != nil {
		panic("Could not load environment variables.")
	}

	if Config == nil {
		Config = &AppConfig{
			Name: os.Getenv("APP_NAME"),
			Mode: os.Getenv("APP_MODE"),
			DB: DBConfig{
				Name:     os.Getenv("DB_NAME"),
				Host:     os.Getenv("DB_HOST"),
				Port:     getIntEnv("DB_PORT"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
				Charset:  os.Getenv("DB_CHARSET"),
			},
			Http: HttpConfig{
				Host: os.Getenv("APP_HOST"),
				Port: getIntEnv("APP_PORT"),
			},
		}
	}

	return Config
}

func getIntEnv(key string, defaultValue ...int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}
