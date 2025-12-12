package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func newConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, errors.New("error loading .env file")
	}
	key := os.Getenv("KEY")
	return &Config{
		Key: key,
	}, nil
}

func Init() (*Config, error) {
	config, err := newConfig()
	if err != nil {

		return nil, err
	}
	return config, nil
}
