package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
)

const (
	ConfigFileDir  = ".lci"
	ConfigFileName = "config.json"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
	filename string         `json:"-"`
}

type DatabaseConfig struct {
	DBName   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func Load(path string) (*Config, error) {
	if path == "" {
		dir, err := defaultConfigDir()
		if err != nil {
			return nil, err
		}

		return load(dir)
	}

	return load(path)
}

func (c *Config) Save() error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(c.filename, data, 0600)
}

func defaultConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ConfigFileDir)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return configDir, nil
}

func load(configDir string) (*Config, error) {
	filePath := filepath.Join(configDir, ConfigFileName)
	var config = new(Config)

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}

		return nil, err
	}

	defer file.Close()

	if err := json.NewDecoder(file).Decode(config); err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	return config, nil
}
