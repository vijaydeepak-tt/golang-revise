package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// use the upper case even to available in the 'Marshal' method
// use "`" to define the metadata for a property
// there is no property key is fixed, we can give ny names
// json: is the property key in the metadata the json package looks for
type Todo struct {
	Text string `json:"text"`
}

// No needed the pointe as we are not manipulating
func (todo Todo) Save() error {
	fileName := "todo.json"

	jsonData, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

// No needed the pointe as we are not manipulating
func (todo Todo) Display() {
	fmt.Println("Todo: ", todo.Text)
}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("Invalid Input.")
	}

	return Todo{
		Text: content,
	}, nil
}
