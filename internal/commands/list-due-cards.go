package commands

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type FindCardsResult struct {
	Result []int64     `json:"result"`
	Error  interface{} `json:"error"`
}

type CardsInfoResult struct {
	Result []CardInfo  `json:"result"`
	Error  interface{} `json:"error"`
}

type CardInfo struct {
	Answer     string `json:"answer"`
	Question   string `json:"question"`
	DeckName   string `json:"deckName"`
	ModelName  string `json:"modelName"`
	FieldOrder int64  `json:"fieldOrder"`
	Fields     Fields `json:"fields"`
	CSS        string `json:"css"`
	CardID     int64  `json:"cardId"`
	Interval   int64  `json:"interval"`
	Note       int64  `json:"note"`
	Ord        int64  `json:"ord"`
	Type       int64  `json:"type"`
	Queue      int64  `json:"queue"`
	Due        int64  `json:"due"`
	Reps       int64  `json:"reps"`
	Lapses     int64  `json:"lapses"`
	Left       int64  `json:"left"`
	Mod        *int64 `json:"mod,omitempty"`
}

type Fields struct {
	Front Back `json:"Front"`
	Back  Back `json:"Back"`
}

type Back struct {
	Value string `json:"value"`
	Order int64  `json:"order"`
}

func (state *State) ListDueCards(deckName string, numCards int) {
	// Get cards of current deck
	payload := map[string]interface{}{
		"action":  "findCards",
		"version": 6,
		"params": map[string]interface{}{
			"query": "deck:current",
		},
	}
	resp, err := state.CallAnki(payload)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var result FindCardsResult
	if err := json.NewDecoder(resp).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		os.Exit(1)
	}

	// Get cards info
	payload = map[string]interface{}{
		"action":  "cardsInfo",
		"version": 6,
		"params": map[string]interface{}{
			"cards": result.Result,
		},
	}
	resp, err = state.CallAnki(payload)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var result2 CardsInfoResult
	if err = json.NewDecoder(resp).Decode(&result2); err != nil {
		fmt.Println("Error decoding respnse: ", err)
		os.Exit(1)
	}

	// sort by descending Due Date
	slices.SortFunc(result2.Result, func(a, b CardInfo) int {
		return cmp.Compare(b.Due, a.Due)
	})

	// list as csv
	for _, item := range result2.Result[:numCards] {
		front := item.Fields.Front.Value
		back := item.Fields.Back.Value
		fmt.Printf("%s;%s\n", front, back)
	}
}
