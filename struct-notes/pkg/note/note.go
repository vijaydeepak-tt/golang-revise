package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// use the upper case even to available in the 'Marshal' method
// use "`" to define the metadata for a property
// there is no property key is fixed, we can give ny names
// json: is the property key in the metadata the json package looks for
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// No needed the pointe as we are not manipulating
func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonData, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

// No needed the pointe as we are not manipulating
func (n Note) Display() {
	fmt.Println("Note Title", n.Title)
	fmt.Println("Note Content", n.Content)
	fmt.Println("Note Created At", n.CreatedAt)
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input.")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
