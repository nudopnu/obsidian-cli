package commands

import (
	"fmt"
	"io"
	"log"
)

func (state *State) LocateCurrentFile() {
	reader, err := state.Call("/active/")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
