package main

import (
	"log"
	"os"

	"github.com/nudopnu/obsidian-cli/internal/commands"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func main() {
	plain := false
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "curl",
				Aliases: []string{"c"},
				Usage:   "Print out the content of an obsidian md file by providing its url",
				Action: func(cCtx *cli.Context) error {
					commands.Cat(cCtx.Args().First(), plain)
					return nil
				},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "plain",
						Usage:       "remove links from output",
						Aliases:     []string{"p"},
						Destination: &plain,
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
