package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatValueFromFile(fileName string) (float64, error) {
	valueData, error := os.ReadFile(fileName)
	if error != nil {
		return 0, errors.New("Failed to read the file")
	}
	valueString := string(valueData)
	value, error := strconv.ParseFloat(valueString, 64)
	if error != nil {
		return 0, errors.New("Failed to convert value from file to float value")
	}
	return value, nil
}

func WriteFloatValuetoFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
