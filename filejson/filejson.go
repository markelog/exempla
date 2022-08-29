package filejson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/** Example model struct to hold JSON data. You won't use this in your actual program. */
type Entry struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func (e *Entry) JsonString() (string, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ReadFromJsonContent(bytes []byte) ([]Entry, error) {
	var result []Entry
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, fmt.Errorf("failed to deserialize with error %v", err)
	}
	return result, nil
}

func ReadFromJsonFile(name string) ([]Entry, error) {
	jsonFile, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s with error: %v", name, err)
	}
	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read file at %s with error %v", name, err)
	}
	return ReadFromJsonContent(bytes)
}