package commands

import (
	"errors"
	"io"
	"log"
	"net/url"
	"strings"

	"github.com/nudopnu/obsidian-cli/internal"
)

func (state *State) Curl(path string) (title, content string, err error) {
	if len(path) > 8 && path[:9] == "obsidian:" {
		location, err := url.ParseRequestURI(path)
		if err != nil {
			log.Fatal(err)
		}
		queries, err := url.ParseQuery(location.RawQuery)
		if err != nil {
			log.Fatal(err)
		}
		title := queries["file"][0]
		filepath := "/vault/" + title + ".md"
		reader, err := state.CallObsidian(filepath, nil)
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
		return title, content, nil
	} else if path == "" {
		reader, err := state.CallObsidian("/active/", map[string]string{
			"accept": "application/vnd.olrapi.note+json",
		})
		if err != nil {
			log.Fatal(err)
		}
		dict, err := internal.ToDict(reader)
		if err != nil {
			log.Fatal(err)
		}
		path := dict["path"].(string)
		content := dict["content"].(string)
		parts := strings.Split(path, "/")
		title := parts[len(parts)-1]
		title = title[:len(title)-3]
		if state.Plain {
			content = internal.Clean(content)
		}
		return title, content, nil
	}
	return "", "", errors.New("invalid path")
}
