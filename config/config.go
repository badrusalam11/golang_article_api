package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// DatabaseConfig holds the database configuration values
type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
}

// Config holds the entire configuration structure
type Config struct {
	Database DatabaseConfig `json:"database"`
}

// AppConfig is a global variable to store the loaded configuration
var AppConfig Config

// LoadConfig loads configuration from the config.json file
func LoadConfig() error {
	// Open the config.json file
	file, err := os.Open("config/config.json")
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	// Read and parse the JSON file
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	// Unmarshal JSON into AppConfig
	err = json.Unmarshal(bytes, &AppConfig)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config data: %v", err)
	}

	return nil
}
