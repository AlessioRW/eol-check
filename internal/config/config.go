package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

func ParseConfig() (*EolConfig, error) {
	configFile, err := os.ReadFile("eol.yaml")
	if err != nil {
		slog.Error("error reading eol.yaml file", "error", err)
		return nil, err
	}

	config := &EolConfig{}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		slog.Error("error parsing config file", "error", err)
		return nil, err
	}

	return config, nil
}
