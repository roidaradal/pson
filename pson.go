package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/roidaradal/fn/io"
)

var (
	errInvalidJSON = errors.New("invalid JSON file")
)

// Read JSON from inputPath, and saves the compressed JSON in outputPath
func compressJSON(inputPath, outputPath string) error {
	jsonString, err := readJSON(inputPath)
	if err != nil {
		return err
	}

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
		err = errInvalidJSON
	}
	if err != nil {
		return err
	}

	fmt.Println("Saved:", outputPath)
	return nil
}

// Read JSON from inputPath, and saves indented JSON in outputPath
func indentJSON(inputPath, outputPath string) error {
	jsonString, err := readJSON(inputPath)
	if err != nil {
		return err
	}

	if strings.HasPrefix(jsonString, "{") {
		var data map[string]any
		err = json.Unmarshal([]byte(jsonString), &data)
		if err != nil {
			return err
		}
		err = io.SaveIndentedJSON(data, outputPath)
	} else if strings.HasPrefix(jsonString, "[") {
		var data []any
		err = json.Unmarshal([]byte(jsonString), &data)
		if err != nil {
			return err
		}
		err = io.SaveIndentedJSON(data, outputPath)
	} else {
		err = errInvalidJSON
	}
	if err != nil {
		return err
	}

	fmt.Println("Saved:", outputPath)
	return nil
}

// Read JSON from inputPath
func readJSON(inputPath string) (string, error) {
	var jsonString string
	if !io.PathExists(inputPath) {
		return "", fmt.Errorf("path %q does not exist", inputPath)
	}

	jsonString, err := io.ReadFile(inputPath)
	if err != nil {
		return "", err
	}
	jsonString = strings.TrimSpace(jsonString)

	return jsonString, nil
}
