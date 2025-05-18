package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatValue(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("Failed to find the file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0, errors.New("Failed to parse the value from file.")
	}

	return value, nil
}

func SaveFloatValue(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
