package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/roidaradal/fn/io"
)

// Read JSON from inputPath, and saves the compressed JSON in outputPath
func compressJSON(inputPath, outputPath string) error {
	if !io.PathExists(inputPath) {
		return fmt.Errorf("path %q does not exist", inputPath)
	}

	jsonString, err := io.ReadFile(inputPath)
	if err != nil {
		return err
	}
	jsonString = strings.TrimSpace(jsonString)

	if strings.HasPrefix(jsonString, "{") {
		var data map[string]any
		err = json.Unmarshal([]byte(jsonString), &data)
		if err != nil {
			return err
		}
		err = io.SaveJSON(data, outputPath)
	} else if strings.HasPrefix(jsonString, "[") {
		var data []any
		err = json.Unmarshal([]byte(jsonString), &data)
		if err != nil {
			return err
		}
		err = io.SaveJSON(data, outputPath)
	} else {
		err = fmt.Errorf("invalid JSON file")
	}
	if err != nil {
		return err
	}

	fmt.Println("Saved:", outputPath)
	return nil
}
