package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println("text:", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("notetitle or notecontent should not be empty")
	}
	return Todo{
		Text: content,
	}, nil
}
