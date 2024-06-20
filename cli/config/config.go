package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	DBName   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func LoadConfig(path string) *Config {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var config *Config

	if err := json.Unmarshal(dat, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

func SaveConfig(path string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
