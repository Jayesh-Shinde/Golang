package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.org/note/note"
	"example.org/note/todo"
)

type saver interface {
	Save() error
}

type outputabble interface {
	saver
	Display()
}

func outPutData(o outputabble) error {
	o.Display()
	return o.Save()
}
func saveData(s saver) error {
	err := s.Save()
	if err != nil {
		fmt.Println("Error while saving the file")
		return err
	}
	fmt.Println("Success saving the file")
	return nil
}

// printAnything prints "any" type of data passed to it.
func printAnything(data interface{}) {

	intVal, ok := data.(int)
	if ok {
		fmt.Println("Integer value:", intVal)
		return
	}

	floatValue, ok := data.(float64)
	if ok {
		fmt.Println("Float64 value:", floatValue)
		return
	}

	switch data.(type) {
	case int:
		fmt.Println("Integer value:", data)
	case float64:
		fmt.Println("Float64 value:", data)
	case string:
		fmt.Println("String value:", data)
	default:
		fmt.Println("Can not handle this type of data")
	}
	fmt.Println(data)
}

func add(a, b interface{}) interface{} {
	aInt, aOk := a.(int)
	bInt, bOk := b.(int)

	if aOk && bOk {
		return aInt + bInt
	}
	return nil
}

func addGeneric[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {

	result := add(1, 2)

	fmt.Println("Result:", result)

	result1 := addGeneric(1, 2)
	fmt.Println("Result1:", result1)

	printAnything(1)
	printAnything(1.5)
	printAnything("Tokyo")

	noteTitle, noteContent := getNoteData()

	todoText := getUserInput("Enter todo text:")

	todoObj, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = outPutData(todoObj)
	if err != nil {
		return
	}

	noteObj, err := note.New(noteTitle, noteContent)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = outPutData(noteObj)
}

func getNoteData() (string, string) {
	noteTitle := getUserInput("Please input note title:")
	notecontent := getUserInput("Please input note content:")
	return noteTitle, notecontent
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	input = strings.TrimSpace(strings.TrimSuffix(strings.TrimSuffix(input, "\n"), "\r"))
	return input
}
