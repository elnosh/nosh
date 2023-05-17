package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	PrivateKey string `json:"privateKey"`
}

func InitConfig() {
	configPath := getConfigPath()

	err := os.MkdirAll(configPath, 0700)
	if err != nil {
		log.Fatalf("error creating config: %v", err)
	}
}

func getConfigPath() string {
	var configPath string

	// os.UserConfigDir() returns $HOME/Library/Application Support for darwin
	// set config path manually to $HOME/.config
	if runtime.GOOS == "darwin" {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		configPath = filepath.Join(home, ".config")
	} else {
		configPath, _ = os.UserConfigDir()
	}

	configPath = filepath.Join(configPath, "nosh")
	return configPath
}
