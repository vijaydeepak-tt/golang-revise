package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/struct-notes/pkg/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving to file is Failed!")
		return
	}

	fmt.Println("Saving to file is Success!")
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
