package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFileName string
}

func New(inputFilePath, outputFileName string) *FileManager {
	return &FileManager{
		InputFilePath:  inputFilePath,
		OutputFileName: outputFileName,
	}
}

func (fileManager FileManager) WriteJSON(data any) error {
	file, err := os.Create(fileManager.OutputFileName)
	if err != nil {
		return errors.New("error creating file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		//file.Close()
		return errors.New("Failed to convert data to json")
	}
	//file.Close()

	return nil
}

func (fileManager FileManager) Readlines() ([]string, error) {
	file, err := os.Open(fileManager.InputFilePath)
	if err != nil {
		fmt.Println("Could not open file")
		return nil, err
	}

	// this is like finally clause in java and javascript
	// this will not be called right away but go will call it
	// at the end no matter what is outcome of Readlines error or success
	// so below in this function we commented file.Close()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file failed")
		fmt.Println(err)
		//file.Close()
		return nil, err
	}
	//file.Close()
	return lines, nil
}
