package commands

import (
	"fmt"
	"io"
	"log"
	"net/url"

	"github.com/nudopnu/obsidian-cli/internal"
)

func (state *State) Cat(path string) {
	if len(path) > 8 && path[:9] == "obsidian:" {
		location, err := url.ParseRequestURI(path)
		if err != nil {
			log.Fatal(err)
		}
		queries, err := url.ParseQuery(location.RawQuery)
		if err != nil {
			log.Fatal(err)
		}
		filepath := "/vault/" + queries["file"][0] + ".md"
		reader, err := state.Call(filepath)
		if err != nil {
			log.Fatal(err)
		}
		data, err := io.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		}
		content := string(data)
		if state.Plain {
			content = internal.Clean(content)
		}
		fmt.Println(content)
	} else if path == "" {
		reader, err := state.Call("/active/")
		if err != nil {
			log.Fatal(err)
		}
		data, err := io.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		}
		content := string(data)
		if state.Plain {
			content = internal.Clean(content)
		}
		fmt.Println(content)
	}
}
