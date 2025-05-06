package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (*State) CallAnki(payload map[string]interface{}) (io.ReadCloser, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling json: %w", err)
	}

	// Send the request to AnkiConnect
	resp, err := http.Post("http://localhost:8765", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	return resp.Body, nil
}

func (*State) AddNote(deckName, front, back string) {
	// Define the JSON payload for AnkiConnect
	payload := map[string]interface{}{
		"action":  "addNote",
		"version": 6,
		"params": map[string]interface{}{
			"note": map[string]interface{}{
				"deckName":  deckName,
				"modelName": "Basic",
				"fields": map[string]string{
					"Front": front,
					"Back":  back,
				},
				"options": map[string]bool{
					"allowDuplicate": false,
				},
			},
		},
	}

	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
	}

	// Send the request to AnkiConnect
	resp, err := http.Post("http://localhost:8765", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Decode response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		os.Exit(1)
	}

	fmt.Println("Response from AnkiConnect:", result)
}
