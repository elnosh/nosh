package main

import (
	"fmt"

	"github.com/nbd-wtf/go-nostr"
	"github.com/urfave/cli/v2"
)

var eventCmd *cli.Command = &cli.Command{
	Name:  "event",
	Usage: "Generate nostr events",
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
			Usage: "private key",
		},
	},
	Action: generateEvent,
}

func generateEvent(ctx *cli.Context) error {
	privKey := ctx.String("pk")
	if privKey == "" {
		return fmt.Errorf("private key not specified")
	}
	pubkey, err := nostr.GetPublicKey(privKey)
	if err != nil {
		return err
	}

	evt := &nostr.Event{
		PubKey:    pubkey,
		CreatedAt: nostr.Now(),
		Kind:      ctx.Int("kind"),
		Tags:      []nostr.Tag{},
		Content:   ctx.String("content"),
	}

	err = evt.Sign(privKey)
	if err != nil {
		return err
	}
	valid, err := evt.CheckSignature()
	if !valid || err != nil {
		return fmt.Errorf("error with event signature")
	}

	fmt.Println(evt.String())
	return nil
}
