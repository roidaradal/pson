package pson

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/dyn"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/str"
)

// Align map data
func alignMap(data map[string]any, level int) string {
	out := make([]string, 0)
	indent := strings.Repeat(" ", INDENT_SPACE*level)
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
				body = bodyList(body, indent)
				valueString := fmt.Sprintf("[%s]", body)
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
	indent := strings.Repeat(" ", INDENT_SPACE*level)
	for _, item := range data {
		if isList(item) {
			dataList, ok := item.([]any)
			if ok {
				body := alignList(dataList, level+1)
				body = bodyList(body, indent)
				line := indent + "[" + body + "]"
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

// Body list
func bodyList(body, indent string) string {
	if FLAT_LIST {
		body = strings.ReplaceAll(body, "\n", " ")
		return strings.Join(strings.Fields(body), " ")
	} else {
		return "\n" + body + indent
	}
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
