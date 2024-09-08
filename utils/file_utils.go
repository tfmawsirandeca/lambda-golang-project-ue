package utils

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"go-lambda-sagemaker/types"
)

// ReadInputFile reads the JSON file located in the data folder and unmarshals it into a slice of InputItem.
func ReadInputFile() ([]types.InputItem, error) {
	// Get the absolute path to the data folder
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("failed to get executable path: %v", err)
		return nil, err
	}

	// Determine the directory of the executable and construct the path to the JSON file
	execDir := filepath.Dir(execPath)
	dataFilePath := filepath.Join(execDir, "data", "input.json")

	// Read the JSON file
	data, err := os.ReadFile(dataFilePath)
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
		return nil, err
	}

	// Unmarshal the JSON data into a slice of InputItem
	var items []types.InputItem
	if err := json.Unmarshal(data, &items); err != nil {
		log.Fatalf("failed to unmarshal input data: %v", err)
		return nil, err
	}

	return items, nil
}
