package commands

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"regexp"
	"strings"
)

func clean(text string) string {
	lines := strings.Split(text, "\n")
	var resultlines []string
	isMetadata := false
	var re1 = regexp.MustCompile(`\[\[([^|]+)\]\]`)
	var re2 = regexp.MustCompile(`\[\[[^\]]*?\|(.*?)\]\]`)
	if lines[0][:3] == "---" {
		isMetadata = true
	}
	for idx, line := range lines {
		if isMetadata && idx > 0 && line[:3] != "---" {
			continue
		}
		isMetadata = false
		line = re1.ReplaceAllString(line, `$1`)
		line = re2.ReplaceAllString(line, `$1`)
		resultlines = append(resultlines, line)
	}
	return strings.Join(resultlines, "\n")
}

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
			content = clean(content)
		}
		fmt.Println(content)
	}
}
