package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const todoStorePath = "exercises/note_main/data/"

type Todo struct {
	Text string `json:"text"`
}

func (todo *Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo *Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	storingPath := todoStorePath + fileName

	return os.WriteFile(storingPath, json, 0644)
}

func New(text string) (*Todo, error) {
	if text == "" {
		return &Todo{}, errors.New("Input cannot be empty")
	}

	return &Todo{Text: text}, nil
}
