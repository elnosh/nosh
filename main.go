package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "nosh",
		Usage: "nostr CLI helper",
		Commands: []*cli.Command{
			genKeyCmd,
			eventCmd,
			// {
			// 	Name:   "genkey",
			// 	Usage:  "Generate private key",
			// 	Action: generatePrivateKey,
			// },
			// {
			// 	Name:  "event",
			// 	Usage: "Generate nostr events",
			// 	Flags: []cli.Flag{
			// 		&cli.StringFlag{
			// 			Name:    "content",
			// 			Aliases: []string{"c"},
			// 			Usage:   "set content for event",
			// 		},
			// 		&cli.IntFlag{
			// 			Name:    "kind",
			// 			Aliases: []string{"k"},
			// 			Usage:   "set event kind",
			// 			Value:   1,
			// 		},
			// 		&cli.StringFlag{
			// 			Name:  "pk",
			// 			Usage: "private key",
			// 		},
			// 	},
			// 	Action: generateEvent,
			// },
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
