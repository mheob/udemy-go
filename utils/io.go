package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	const DEFAULT_VALUE = 0

	data, err := os.ReadFile(fileName)
	if err != nil {
		return DEFAULT_VALUE, errors.New("Failed to find file.")
	}

	valueText := string(data)
	valueAmount, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return DEFAULT_VALUE, errors.New("Failed to parse stored value.")
	}

	return valueAmount, nil
}

func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprintf("%.2f", value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
