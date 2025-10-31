package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	Save() error
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("hello")

	fmt.Println("Good day User")
	title, content := getNotesData()
	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}
}

func printSomething(value interface{}) {
	// switch value.(type) {
	// case string:
	// 	fmt.Println("String:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// }

	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
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
