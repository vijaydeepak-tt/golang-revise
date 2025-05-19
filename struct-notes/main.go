package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/struct-notes/pkg/note"
	"example.com/struct-notes/pkg/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// Interface embeddings
type outputtable interface {
	saver
	Display()
	// displayer
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	title, content := getNoteData()
	todoText := getUserData("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(todo)

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)

}

func outputData(data outputtable) {
	data.Display()

	err := saveData(data)
	if err != nil {
		return
	}
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving Note to file is Failed!")
		return err
	}

	fmt.Println("Saving Note to file is Success!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserData("Note Title: ")
	content := getUserData("Note Content: ")

	return title, content
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	// using bufio package for long test's
	// and os.Stdin for reading text from cmd line
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n') // using single quote for mentioning a single character

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// we can also use 'any' instead of 'interface{}'
func printSomething(value interface{}) {

	intVal, ok := value.(int) // return ok as true if value is int

	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64) // return ok as true if value is float64

	if ok {
		fmt.Println("Float64: ", floatVal)
		return
	}

	stringVal, ok := value.(string) // return ok as true if value is string

	if ok {
		fmt.Println("String: ", stringVal)
		return
	}

	// // This is one way of using the type of the value
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float64: ", value)
	// case string:
	// 	fmt.Println("String: ", value)
	// }
}
