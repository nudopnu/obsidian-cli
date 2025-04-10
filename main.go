package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nudopnu/obsidian-cli/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	_ = godotenv.Load(".env")
	state := commands.State{}
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "cat",
				Aliases: []string{"c"},
				Usage:   "Print out the content of an obsidian md file by providing its URL. If no URL is provided, the current opened file is taken.",
				Action: func(ctx *cli.Context) error {
					title, content, err := state.Curl(ctx.Args().First())
					if err != nil {
						return err
					}
					fmt.Printf("%s\n%s\n", title, content)
					return nil
				},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "plain",
						Usage:       "Remove links and metadata from output",
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
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Adds the current opened file in obsidian to anki",
				Action: func(ctx *cli.Context) error {
					state.Plain = true
					title, content, err := state.Curl("")
					if err != nil {
						return err
					}
					deckname := ctx.Args().First()
					if deckname == "" {
						fmt.Println("the deck name must not be empty")
						os.Exit(1)
					}
					state.AddNote(deckname, title, content)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "api_key",
						Usage:       "Provide your API key via flag or environment variable `API_KEY`. You can obtain it in the plugin settings in obsidian.",
						EnvVars:     []string{"API_KEY"},
						Destination: &state.ApiKey,
					},
				},
			},
			{
				Name:    "deck",
				Aliases: []string{"d"},
				Action: func(ctx *cli.Context) error {
					name := ctx.Args().First()
					if name == "" {
						fmt.Println("the deck name must not be empty")
						os.Exit(1)
					}
					state.AddDeck(name)
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
