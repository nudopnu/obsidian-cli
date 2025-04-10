package commands

import (
	"errors"
	"fmt"

	"github.com/nudopnu/obsidian-cli/internal"
)

func (state *State) ListFiles(path string) error {
	reader, err := state.CallObsidian("/vault/"+path, nil)
	if err != nil {
		return err
	}
	data, err := internal.ToDict(reader)
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
