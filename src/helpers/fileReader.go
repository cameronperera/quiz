package helpers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func Open(csvFilename string) (*os.File, error) {
	return os.Open(csvFilename)
}

func Read(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func ReadFile(csvFilename string) ([][]string, error) {
	file, err := Open(csvFilename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to open the CSV file: %s\n", csvFilename))
	}

	lines, err := Read(file)
	if err != nil {
		return nil, errors.New("Failed to parse the provided CSV file.")
	}
	return lines, nil
}
