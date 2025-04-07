package main

import (
	"log"
	"os"

	"github.com/nudopnu/obsidian-cli/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	state := commands.State{}
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "curl",
				Aliases: []string{"c"},
				Usage:   "Print out the content of an obsidian md file by providing its url",
				Action: func(cCtx *cli.Context) error {
					state.Cat(cCtx.Args().First())
					return nil
				},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "plain",
						Usage:       "remove links from output",
						Aliases:     []string{"p"},
						Destination: &state.Plain,
					},
					&cli.StringFlag{
						Name:        "api_key",
						Usage:       "Provide your API key via flag or environment variable `API_KEY`. You can obtain it in the plugin settings in obsidian.",
						EnvVars:     []string{"API_KEY"},
						Destination: &state.ApiKey,
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
