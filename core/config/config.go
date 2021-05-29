package config

import (
	"errors"
	"log"
	"os"
	"strconv"
)

type Config struct{}

var config = Config{}

func GetConfig() *Config {
	return &config
}

func (config *Config) GetDSN() string {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal(errors.New("you have to set DSN environment variable to be able to store all settings"))
	}

	return dsn
}

func (config *Config) GetHttpPort() int {
	stringPort := os.Getenv("PORT")
	if stringPort == "" {
		return 8080
	}

	port, err := strconv.Atoi(stringPort)
	if err != nil {
		log.Fatal(errors.New("environment variable PORT is not a number"))
	}

	return port
}
