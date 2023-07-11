package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	Server  ServerConfig
	MongoDB MongoDBConfig
	Logger  Logger
}

// ServerConfig is a struct that contains the configuration of the server
type ServerConfig struct {
	ServiceName string
	Port        string
}

// Logger is a struct that contains the configuration of the logger
type Logger struct {
	Environment string
	Level       logrus.Level
	Namespace   string
}

// MongoDBConfig is a struct that contains the configuration of the MongoDB
type MongoDBConfig struct {
	URI      string
	User     string
	Password string
}

var ErrInvalidConfig = errors.New("invalid config")

func LoadConfig() (config *Config, err error) {
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	config = &Config{
		Server: ServerConfig{
			ServiceName: getEnvString("SERVICE_NAME"),
			Port:        getEnvString("PORT"),
		},
		MongoDB: MongoDBConfig{
			URI:      getEnvString("URI"),
			User:     getEnvString("USER"),
			Password: getEnvString("PASSWORD"),
		},
		Logger: Logger{
			Environment: getEnvString("LOGGER_ENVIRONMENT"),
			Level:       logrus.Level(getEnvInt("LOGGER_LEVEL")),
			Namespace:   getEnvString("LOGGER_NAMESPACE"),
		},
	}
	return config, nil
}

func getEnvString(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(ErrInvalidConfig)
	}
	return value
}

func getEnvBool(key string, defaultValue ...bool) bool {
	valueStr := getEnvString(key)
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(ErrInvalidConfig)
	}
	return value
}

func getEnvInt(key string, defaultValue ...int) int {
	valueStr := getEnvString(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(ErrInvalidConfig)
	}
	return value
}
