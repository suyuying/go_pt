package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Website struct {
	TargetURL   string `yaml:"target_url"`
	Description string `yaml:"description"`
}
type Config struct {
	Websites []Website `yaml:"websites"`
}

func LoadConfig(route string, configPtr *Config) error {
	// Read the configuration file
	configData, err := os.ReadFile(route)
	if err != nil {
		return fmt.Errorf("failed to read configuration file: %w", err)
	}

	// Unmarshal the configuration data into the provided Config pointer
	err = yaml.Unmarshal(configData, configPtr)
	if err != nil {
		return fmt.Errorf("failed to unmarshal configuration data: %w", err)
	}

	return nil
}
