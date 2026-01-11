package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
)

const usage string = "Usage: pson <indent|align|compress> <file.json> (--overwrite) (--indent=2) (--flatlist)"

var (
	indentSpace int  = 2
	flatList    bool = false
)

func main() {
	var err error
	command, inputPath, outputPath := getArgs()
	switch command {
	case "compress":
		err = compressJSON(inputPath, outputPath)
	case "indent":
		err = indentJSON(inputPath, outputPath)
	case "align":
		err = alignJSON(inputPath, outputPath)
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
	if len(args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	command, inputPath = args[0], args[1]
	if !strings.HasSuffix(inputPath, ".json") {
		fmt.Println("File path needs to be a .json file")
		os.Exit(1)
	}

	// Default output path
	filename, _ := strings.CutSuffix(inputPath, ".json")
	outputPath = fmt.Sprintf("%s.%s.json", filename, command)

	for _, arg := range args[2:] {
		if arg == "--overwrite" {
			outputPath = inputPath
		} else if arg == "--flatlist" {
			flatList = true
		} else if strings.HasPrefix(arg, "--indent=") {
			parts := strings.Split(arg, "=")
			if len(parts) == 2 {
				customIndent := number.ParseInt(parts[1])
				indentSpace = max(indentSpace, customIndent)
				str.SetJSONIndentLength(indentSpace)
			}
		}
	}
	return command, inputPath, outputPath
}
