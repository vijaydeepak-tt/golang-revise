package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputPath)

	if err != nil {
		return nil, errors.New("Failed to opening the file")
	}

	// go will close the file for ue on the actions are done
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("Reading the file content failed.")
	}
	// file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputPath)

	if err != nil {
		return errors.New("Failed to Create a file")
	}
	time.Sleep(3 * time.Second)

	// go will close the file for ue on the actions are done
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("Failed to convert data to JSON")
	}
	// file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}
