package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

// Config ...
type Config struct {
	ConnectionString string
	Port             string
}

// NewConfig ...
func NewConfig(filenames ...string) (config *Config, err error) {
	err = godotenv.Load(filenames...)
	if err != nil {
		return nil, err
	}

	connectionString := getVarOrDefault("CONNECTION_STRING", "")
	if connectionString == "" {
		return nil, errors.New("unable to connect to database: ConnectionString is empty")
	}

	return &Config{
		ConnectionString: connectionString,
		Port:             ":" + getVarOrDefault("PORT", "5000"),
	}, nil
}

// getVarOrDefault ...
func getVarOrDefault(name string, d string) string {
	value, exists := os.LookupEnv(name)

	if exists == false {
		return d
	}

	return value
}
