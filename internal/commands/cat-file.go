package commands

import (
	"fmt"
	"log"
	"net/url"

	"github.com/nudopnu/obsidian-cli/internal"
)

func Cat(file string) {
	if len(file) > 8 && file[:9] == "obsidian:" {
		location, err := url.ParseRequestURI(file)
		queries, err := url.ParseQuery(location.RawQuery)
		filepath := "/vault/" + queries["file"][0] + ".md"
		data, err := internal.Call(filepath)
		fmt.Println(data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
