package main

import (
	"fmt"

	"example.com/notes/note"
)

func main() {
	fmt.Println("Good day User")
	title, content := getNotesData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.DisplayNotes()
}

func getNotesData() (string, string) {
	title := getUserInput("Title of the note:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)

	return value
}
