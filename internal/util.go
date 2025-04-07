package internal

import (
	"encoding/json"
	"io"
	"regexp"
	"strings"
)

func ToDict(reader io.ReadCloser) (map[string]interface{}, error) {
	var result map[string]interface{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Clean(text string) string {
	lines := strings.Split(text, "\n")
	var resultlines []string
	isMetadata := false
	var re1 = regexp.MustCompile(`\[\[([^|]+)\]\]`)
	var re2 = regexp.MustCompile(`\[\[[^\]]*?\|(.*?)\]\]`)
	if lines[0][:3] == "---" {
		isMetadata = true
	}
	for idx, line := range lines {
		if isMetadata {
			if idx > 0 && line[:3] == "---" {
				isMetadata = false
			}
			continue
		}
		line = re1.ReplaceAllString(line, `$1`)
		line = re2.ReplaceAllString(line, `$1`)
		resultlines = append(resultlines, line)
	}
	return strings.Join(resultlines, "\n")
}
