package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	InitConfig()
	app := &cli.App{
		Name:  "nosh",
		Usage: "nostr CLI tool for generating events",
		Commands: []*cli.Command{
			keyCmd,
			eventCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
