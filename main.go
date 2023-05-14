package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var key string

func main() {
	app := &cli.App{
		Name:  "nosh",
		Usage: "nostr CLI helper",
		Action: func(*cli.Context) error {
			fmt.Print("hello\n")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
