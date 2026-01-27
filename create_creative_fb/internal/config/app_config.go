package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Server   ServerConfig
	Facebook FacebookConfig
}

type ServerConfig struct {
	Port int
	Host string
}

type FacebookConfig struct {
	Accounts map[string]AccountConfig `mapstructure:"accounts"`
}

type AccountConfig struct {
	Account           string `mapstructure:"account"`
	FacebookAccountID string `mapstructure:"facebook_account_id"`
	AccessToken       string `mapstructure:"access_token"`
}

func LoadConfig() *AppConfig {
	cfg := &AppConfig{
		Server: ServerConfig{
			Port: 8080,
			Host: "localhost",
		},
	}

	// Load from environment variables first
	if port := os.Getenv("PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			cfg.Server.Port = p
		}
	}

	if host := os.Getenv("HOST"); host != "" {
		cfg.Server.Host = host
	}

	// Load from config file
	viper := viper.New()
	viper.SetConfigFile("facebook_config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Could not read config file: %v", err)
	}

	var fbCfg FacebookConfig
	if err := viper.Unmarshal(&fbCfg); err != nil {
		log.Fatal(err)
	}
	cfg.Facebook = fbCfg

	return cfg
}
