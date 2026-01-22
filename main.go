package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/roidaradal/fn/io"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/pson/internal/pson"
)

const currentVersion string = "v0.1.3"
const usage string = "Usage: pson <align|compress|indent|validate|update> <file.json> (output=path) (overwrite) (indent=2) (flatlist)"

var optionalFile = map[string]bool{
	"version": true,
}

func main() {
	var err error
	command, inputPath, outputPath := getArgs()
	switch command {
	case "compress":
		err = pson.CompressJSON(inputPath, outputPath)
	case "indent":
		err = pson.IndentJSON(inputPath, outputPath)
	case "align":
		err = pson.AlignJSON(inputPath, outputPath)
	case "validate":
		err = pson.ValidateJSON(inputPath)
	case "version":
		fmt.Println("pson", currentVersion)
	case "update":
		path := "github.com/roidaradal/pson@latest"
		err = io.RunGoInstall(path)
	default:
		fmt.Println("Unknown command: ", command)
		fmt.Println(usage)
	}
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Get command and JSON file path from command-line args
func getArgs() (command, inputPath, outputPath string) {
	args := os.Args[1:]
	numArgs := len(args)
	if numArgs < 1 {
		fmt.Println(usage)
		os.Exit(1)
	}
	command, inputPath = args[0], ""
	if numArgs >= 2 {
		inputPath = args[1]
	}

	if optionalFile[command] {
		// End early if files are optional
		return command, "", ""
	}

	if !strings.HasSuffix(inputPath, ".json") {
		fmt.Println("File path needs to be a .json file")
		os.Exit(1)
	}

	// Default output path
	filename, _ := strings.CutSuffix(inputPath, ".json")
	outputPath = fmt.Sprintf("%s.%s.json", filename, command)

	for _, arg := range args[2:] {
		if arg == "overwrite" {
			outputPath = inputPath
		} else if arg == "flatlist" {
			pson.FLAT_LIST = true
		} else if strings.HasPrefix(arg, "indent=") {
			customIndent, ok := str.TryGetPart(arg, "=", 1)
			if !ok {
				continue
			}
			pson.INDENT_SPACE = max(pson.INDENT_SPACE, number.ParseInt(customIndent))
			str.SetJSONIndentLength(pson.INDENT_SPACE)
		} else if strings.HasPrefix(arg, "output=") {
			path, ok := str.TryGetPart(arg, "=", 1)
			if !ok {
				continue
			}
			outputPath = path
		}
	}
	return command, inputPath, outputPath
}
