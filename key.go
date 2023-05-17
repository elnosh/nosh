package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/nbd-wtf/go-nostr"
	"github.com/urfave/cli/v2"
)

var genKeyCmd = &cli.Command{
	Name:   "genkey",
	Usage:  "generate private key",
	Action: generatePrivateKey,
}

func generatePrivateKey(*cli.Context) error {
	pk := nostr.GeneratePrivateKey()
	fmt.Printf("private key: %v\n", pk)
	return nil
}

func savePrivateKey(key string) {
	configPath := getConfigPath()
	configFilePath := filepath.Join(configPath, "config.json")

	configFile, err := os.OpenFile(configFilePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("error opening '%v' file: %v", configFilePath, err)
	}

	config := Config{PrivateKey: key}

	jsonBytes, err := json.MarshalIndent(config, "", "   ")
	if err != nil {
		log.Fatalf("error saving key: %v\n", err)
	}
	_, err = configFile.Write(jsonBytes)
	if err != nil {
		log.Fatalf("error saving key: %v\n", err)
	}
}
