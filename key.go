package main

import (
	"fmt"

	"github.com/nbd-wtf/go-nostr"
	"github.com/urfave/cli/v2"
)

var genKeyCmd = &cli.Command{
	Name:   "genkey",
	Usage:  "Generate private key",
	Action: generatePrivateKey,
}

func generatePrivateKey(*cli.Context) error {
	pk := nostr.GeneratePrivateKey()
	fmt.Printf("private key: %v\n", pk)
	return nil
}
