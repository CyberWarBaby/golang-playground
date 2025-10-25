package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
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
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getNotesData() (string, string) {
	title := getUserInput("Title of the note:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
