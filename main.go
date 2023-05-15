package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "nosh",
		Usage: "nostr CLI tool",
		Commands: []*cli.Command{
			genKeyCmd,
			eventCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
