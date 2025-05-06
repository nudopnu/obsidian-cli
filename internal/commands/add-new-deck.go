package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (*State) AddNewDeck(deckName string) {
	// Prepare the AnkiConnect payload
	payload := map[string]interface{}{
		"action":  "createDeck",
		"version": 6,
		"params": map[string]interface{}{
			"deck": deckName,
		},
	}

	// Convert the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		os.Exit(1)
	}

	// Send the request to AnkiConnect
	resp, err := http.Post("http://localhost:8765", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Decode the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		os.Exit(1)
	}
}
