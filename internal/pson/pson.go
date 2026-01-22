package pson

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/roidaradal/fn/io"
	"github.com/roidaradal/nlp"
)

var (
	errInvalidJSON = errors.New("invalid JSON file")
)

var (
	INDENT_SPACE int  = 2
	FLAT_LIST    bool = false
)

// Read JSON from inputPath, and saves the compressed JSON in outputPath
func CompressJSON(inputPath, outputPath string) error {
	return transferJSON(inputPath, outputPath, false)
}

// Read JSON from inputPath, and saves indented JSON in outputPath
func IndentJSON(inputPath, outputPath string) error {
	return transferJSON(inputPath, outputPath, true)
}

// Read JSON from inputPath, and saves aligned JSON in outputPath
func AlignJSON(inputPath, outputPath string) error {
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
		body := alignMap(data, 1)
		body = fmt.Sprintf("{\n%s}", body)
		err = io.SaveString(body, outputPath)
	} else if strings.HasPrefix(jsonString, "[") {
		var data []any
		err = json.Unmarshal([]byte(jsonString), &data)
		if err != nil {
			return err
		}
		body := alignList(data, 1)
		body = fmt.Sprintf("[\n%s]", body)
		err = io.SaveString(body, outputPath)
	} else {
		err = errInvalidJSON
	}
	if err != nil {
		return err
	}

	fmt.Println("Saved:", outputPath)
	return nil
}

// Validate JSON from inputPath
func ValidateJSON(inputPath string) error {
	parser, err := nlp.NewJSONParser()
	if err != nil {
		return err
	}
	lines, err := nlp.ReadLineBytes(inputPath)
	if err != nil {
		return err
	}
	// for _, line := range lines[:10] {
	// 	fmt.Println(line, string(line))
	// }
	err = parser.Parse(lines, nil)
	if err == nil {
		fmt.Println("Valid JSON")
	}
	return err
}

// Common: Read JSON from inputPath
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

// Common: Transfer JSON from inputPath to outputPath
func transferJSON(inputPath, outputPath string, indent bool) error {
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
		if indent {
			err = io.SaveIndentedJSON(data, outputPath)
		} else {
			err = io.SaveJSON(data, outputPath)
		}
	} else if strings.HasPrefix(jsonString, "[") {
		var data []any
		err = json.Unmarshal([]byte(jsonString), &data)
		if err != nil {
			return err
		}
		if indent {
			err = io.SaveIndentedJSON(data, outputPath)
		} else {
			err = io.SaveJSON(data, outputPath)
		}
	} else {
		err = errInvalidJSON
	}
	if err != nil {
		return err
	}

	fmt.Println("Saved:", outputPath)
	return nil
}
