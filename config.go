package main

import (
	"encoding/json"
	"io"
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

	var err error
	// os.UserConfigDir() returns $HOME/Library/Application Support for darwin
	// set config path manually to $HOME/.config
	if runtime.GOOS == "darwin" {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		configPath = filepath.Join(home, ".config")
	} else {
		configPath, err = os.UserConfigDir()
		if err != nil {
			log.Fatal(err)
		}
	}

	configPath = filepath.Join(configPath, "nosh")
	return configPath
}

func getConfig() *Config {
	var config *Config

	configPath := getConfigPath()
	configFilePath := filepath.Join(configPath, "config.json")

	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	defer configFile.Close()

	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	if err != nil && err != io.EOF {
		log.Fatalf("error reading config: %v", err)
	}

	return config
}
