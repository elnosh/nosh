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

var keyCmd = &cli.Command{
	Name:   "key",
	Usage:  "generate, set private key",
	Action: keyAction,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:               "gen",
			Usage:              "generate private key",
			DisableDefaultText: true,
		},
		&cli.StringFlag{
			Name:  "set",
			Usage: "set private key",
		},
		&cli.BoolFlag{
			Name:               "get",
			Usage:              "get current public key set",
			DisableDefaultText: true,
		},
	},
}

func keyAction(ctx *cli.Context) error {
	if ctx.NumFlags() < 1 {
		fmt.Printf("specify flag for 'key' command\n\n")
		cli.ShowSubcommandHelp(ctx)
		return nil
	}

	if ctx.IsSet("gen") {
		generatePrivateKey()
	} else if ctx.IsSet("set") {
		savePrivateKey(ctx.String("set"))
	} else if ctx.IsSet("get") {
		getPubKey()
	}

	return nil
}

func generatePrivateKey() {
	pk := nostr.GeneratePrivateKey()
	fmt.Printf("private key: %v\n", pk)
}

func savePrivateKey(key string) {
	configPath := getConfigPath()
	configFilePath := filepath.Join(configPath, "config.json")

	configFile, err := os.OpenFile(configFilePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("error opening '%v' file: %v", configFilePath, err)
	}
	defer configFile.Close()

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

func getPubKey() {
	config := getConfig()
	if config == nil || config.PrivateKey == "" {
		printErr("key not set")
	}

	pubkey, err := nostr.GetPublicKey(config.PrivateKey)
	if err != nil {
		printErr("error getting public key")
	}
	fmt.Println(pubkey)
}
