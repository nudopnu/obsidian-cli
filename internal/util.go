package internal

import (
	"encoding/json"
	"io"
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
