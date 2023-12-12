package note_main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mheob/udemy-go/exercises/note_main/note"
)

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

func Run() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Note saved successfully")
}
