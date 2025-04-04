package commands

import (
	"errors"
	"fmt"

	"github.com/nudopnu/obsidian-cli/internal"
)

func ListFiles() error {
	data, err := internal.Call("/vault/")
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
