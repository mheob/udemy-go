package note_main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mheob/udemy-go/exercises/note_main/note"
	"github.com/mheob/udemy-go/exercises/note_main/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getTodoData() string {
	return getUserInput("Todo text:")
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}

	fmt.Println("Data saved successfully")
	return nil
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func Run() {
	title, content := getNoteData()
	todoText := getTodoData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
}
