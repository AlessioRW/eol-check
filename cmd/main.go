package main

import (
	"eol-checker/internal/config"
	"eol-checker/internal/handler"
	"log/slog"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		slog.Error("not enough arguments passed, required CONFIG_PATH")
		os.Exit(1)
	}
	configPath := os.Args[1] // path to config file

	config, err := config.ParseConfig(configPath)
	if err != nil {
		os.Exit(1)
	}

	err = handler.Run(config)
	if err != nil {
		os.Exit(1)
	}
}
