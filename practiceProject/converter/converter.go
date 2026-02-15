package converter

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floatValues := make([]float64, len(strings))
	for stringIndex, stringValue := range strings {
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			return nil, errors.New("error while converting float to string")
		}
		floatValues[stringIndex] = floatValue
	}
	return floatValues, nil
}
