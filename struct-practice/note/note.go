package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note with title ---> %s has content of --->%s created at --->%s \n",
		note.Title, note.Content, note.CreatedAt.Local().Format("2006-01-02 15:04:05"))
}

func (note Note) Save() error {
	fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_")) + ".json"
	jsonData, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("notetitle or notecontent should not be empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
