package config

import (
	"github.com/joho/godotenv"
	"os"
)

// Config ...
type Config struct {
	Port string
}

// NewConfig ...
func NewConfig(filenames ...string) (config *Config, err error) {
	err = godotenv.Load(filenames...)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port: ":" + getVarOrDefault("PORT", "5000"),
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