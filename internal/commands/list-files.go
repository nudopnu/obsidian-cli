package commands

import (
	"errors"
	"fmt"

	"github.com/nudopnu/obsidian-cli/internal"
)

func ListFiles(path string) error {
	data, err := internal.Call("/vault/" + path)
	if err != nil {
		return err
	}
	files, ok := data["files"].([]interface{})
	if !ok {
		return errors.New("unexpected type")
	}
	for _, item := range files {
		fmt.Println(item)
	}
	return nil
}
