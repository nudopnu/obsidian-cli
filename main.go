package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nudopnu/obsidian-cli/internal"
	"github.com/nudopnu/obsidian-cli/internal/commands"
)

func parse_command(cmd string) {
	parts := strings.Split(cmd, " ")
	arg := ""
	if len(parts) > 1 {
		arg = parts[1]
	}
	switch parts[0] {
	case "ls":
		commands.ListFiles(arg)
		return
	case "cat":
		commands.Cat(arg)
		return
	case "get":
		internal.Call(arg)
	case "exit":
		os.Exit(1)
	default:
		fmt.Printf("Unknown command '%s'\n", cmd)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fmt.Println(input)
		parse_command(input)
	}
}
