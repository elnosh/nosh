package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nbd-wtf/go-nostr"
	"github.com/urfave/cli/v2"
)

var eventCmd = &cli.Command{
	Name:  "event",
	Usage: "generate nostr event",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "content",
			Aliases: []string{"c"},
			Usage:   "set content for event",
		},
		&cli.IntFlag{
			Name:    "kind",
			Aliases: []string{"k"},
			Usage:   "set event kind",
			Value:   1,
		},
		&cli.StringFlag{
			Name:  "pk",
			Usage: "private key (does not have to be specified if key was set with 'key' command)",
		},
		&cli.StringSliceFlag{
			Name:    "tags",
			Aliases: []string{"t"},
			Usage:   "set tags for event: -t e=<eventId>,p=<pubkey>",
		},
		&cli.BoolFlag{
			Name:               "wrap",
			Aliases:            []string{"w"},
			Usage:              `wrap event for sending to relay ["EVENT", <json event>]`,
			DisableDefaultText: true,
		},
	},
	Action: generateEvent,
}

func generateEvent(ctx *cli.Context) error {
	privKey := ""

	if ctx.IsSet("pk") {
		privKey = ctx.String("pk")
	} else {
		config := getConfig()
		if config == nil || config.PrivateKey == "" {
			printErr("private key not set")
		}

		privKey = config.PrivateKey
	}

	pubkey, err := nostr.GetPublicKey(privKey)
	if err != nil {
		return err
	}

	tags := []nostr.Tag{}

	tagsFlag := ctx.StringSlice("tags")
	for _, t := range tagsFlag {
		if len(t) < 2 {
			printErr("specify tags in format: -t e=<eventId>,p=<pubkey>")
			return nil
		}
		prefix := t[:2]

		if prefix == "e=" || prefix == "p=" {
			tagStr, _ := strings.CutPrefix(t, prefix)

			if prefix == "p=" {
				if !nostr.IsValidPublicKeyHex(tagStr) {
					printErr("invalid pubkey")
				}
			}

			tag := nostr.Tag{string(prefix[0]), tagStr}
			tags = append(tags, tag)
		} else {
			printErr("specify tags in format: -t e=<eventId>,p=<pubkey>")
		}
	}

	evt := &nostr.Event{
		PubKey:    pubkey,
		CreatedAt: nostr.Now(),
		Kind:      ctx.Int("kind"),
		Tags:      tags,
		Content:   ctx.String("content"),
	}

	err = evt.Sign(privKey)
	if err != nil {
		printErr("error signing event")
	}

	eventStr := evt.String()
	// wrap event in ["EVENT", <json event>] if flag is present
	if ctx.Bool("wrap") {
		fmt.Printf("[\"EVENT\", %v]\n", eventStr)
	} else {
		fmt.Println(eventStr)
	}

	return nil
}

func printErr(err string) {
	fmt.Println(err)
	os.Exit(0)
}
