package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func (*State) AddDeck(deckName string) {
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
				"tags": []string{"geography", "example"},
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
