package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config holds the configuration values
type Config struct {
	ServerRunning bool `json:"serverRunning"`
	PortNumber    int  `json:"portNumber"`
}

const Configfile = "config.json"

// LoadConfig reads the configuration from a JSON file
func LoadConfig(filePath string) *Config {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil
	}

	return &config
}

// SaveConfig writes the configuration to a JSON file
func SaveConfig(config *Config, filePath string) {
	byteValue, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return
	}

	ioutil.WriteFile(filePath, byteValue, 0644)
}
