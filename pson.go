package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/dyn"
	"github.com/roidaradal/fn/io"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
)

var (
	errInvalidJSON = errors.New("invalid JSON file")
)

// Read JSON from inputPath, and saves the compressed JSON in outputPath
func compressJSON(inputPath, outputPath string) error {
	return transferJSON(inputPath, outputPath, false)
}

// Read JSON from inputPath, and saves indented JSON in outputPath
func indentJSON(inputPath, outputPath string) error {
	return transferJSON(inputPath, outputPath, true)
}

// Read JSON from inputPath, and saves aligned JSON in outputPath
func alignJSON(inputPath, outputPath string) error {
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

// Align map data
func alignMap(data map[string]any, level int) string {
	out := make([]string, 0)
	indent := strings.Repeat("  ", level)
	keys := dict.Keys(data)
	slices.Sort(keys)
	maxLength := slices.Max(list.Map(keys, str.Length)) + 2
	template := fmt.Sprintf("%%-%dq : %%s", maxLength)
	for _, key := range keys {
		value := data[key]
		if isList(value) {
			dataList, ok := value.([]any)
			if ok {
				body := alignList(dataList, level+1)
				valueString := fmt.Sprintf("[\n%s%s]", body, indent)
				line := indent + fmt.Sprintf(template, key, valueString)
				out = append(out, line)
				continue
			}
		} else if isMap(value) {
			dataMap, ok := value.(map[string]any)
			if ok {
				body := alignMap(dataMap, level+1)
				valueString := fmt.Sprintf("{\n%s%s}", body, indent)
				line := indent + fmt.Sprintf(template, key, valueString)
				out = append(out, line)
				continue
			}
		}

		line := indent + fmt.Sprintf(template, key, toString(data[key]))
		out = append(out, line)
	}
	return strings.Join(out, ",\n") + "\n"
}

// Align list data
func alignList(data []any, level int) string {
	out := make([]string, 0)
	indent := strings.Repeat("  ", level)
	for _, item := range data {
		if isList(item) {
			dataList, ok := item.([]any)
			if ok {
				body := alignList(dataList, level+1)
				line := indent + "[\n" + body + indent + "]"
				out = append(out, line)
				continue
			}
		} else if isMap(item) {
			dataMap, ok := item.(map[string]any)
			if ok {
				body := alignMap(dataMap, level+1)
				line := indent + "{\n" + body + indent + "}"
				out = append(out, line)
				continue
			}
		}

		line := indent + toString(item)
		out = append(out, line)
	}
	return strings.Join(out, ",\n") + "\n"
}

// Convert item of any type to string
func toString(item any) string {
	switch dyn.TypeOf(item) {
	case "string":
		return fmt.Sprintf("%q", item)
	default:
		return fmt.Sprintf("%v", item)
	}
}

// Check if list
func isList(item any) bool {
	return reflect.TypeOf(item).Kind() == reflect.Slice
}

// Check if map
func isMap(item any) bool {
	return reflect.TypeOf(item).Kind() == reflect.Map
}
