package commands

import (
	"fmt"
	"io"
	"log"
)

func (state *State) LocateCurrentFile() {
	reader, err := state.CallObsidian("/active/", nil)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
