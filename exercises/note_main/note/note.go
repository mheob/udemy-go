package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

const noteStorePath = "exercises/note_main/data/"

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Note) Display() {
	fmt.Printf("Your note titled \"%v\" has the following content:\n\n%v\n\n", n.Title, n.Content)
}

func (n Note) Save() error {
	fileName := strings.ToLower(n.Title)
	fileName = regexp.MustCompile(`[\s\.]+`).ReplaceAllString(fileName, "_")
	fileName = regexp.MustCompile(`[^a-zA-Z0-9_]+`).ReplaceAllString(fileName, "") + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	storingPath := noteStorePath + fileName

	return os.WriteFile(storingPath, json, 0644)
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		fmt.Print(title, content)
		return &Note{}, errors.New("Input cannot be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
