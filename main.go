package main

import (
	"fmt"
	"os"

	"github.com/nudopnu/obsidian-cli/internal/commands"
)

func parse_command(cmd string) {
	switch cmd {
	case "ls":
		commands.ListFiles()
		return
	case "exit":
		os.Exit(1)
	default:
		fmt.Printf("Unknown command '%s'\n", cmd)
	}
}

func main() {
	var cmd string
	for {
		fmt.Print("> ")
		fmt.Scan(&cmd)
		parse_command(cmd)
	}
}
