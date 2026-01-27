package config

import (
	"os"
)

type AppConfig struct {
	FacebookAccessToken string
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		FacebookAccessToken: os.Getenv("FACEBOOK_ACCESS_TOKEN"),
	}
}
