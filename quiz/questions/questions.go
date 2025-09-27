package questions

import (
	"encoding/json"
	"fmt"
	"os"
)

type Question struct {
	Country string `json:"country"`
	Capital string `json:"capital"`
}

func LoadQuestions() ([]Question, error) {
	jsonData, err := os.ReadFile("quiz.json")

	if err != nil {
		return nil, fmt.Errorf("error reading questions file: %w", err)
	}

	var question []Question

	err = json.Unmarshal(jsonData, &question)

	if err != nil {
		return nil, fmt.Errorf("error parsing JSON file: %w", err)
	}

	return question, nil
}
