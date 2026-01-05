package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const usage string = "Usage: pson <pretty|align|compress> <file.json> (--overwrite)"

func main() {
	var err error
	command, inputPath, outputPath := getArgs()
	switch command {
	case "compress":
		err = compressJSON(inputPath, outputPath)
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
	if slices.Contains(args, "--overwrite") {
		outputPath = inputPath
	} else {
		filename, _ := strings.CutSuffix(inputPath, ".json")
		outputPath = fmt.Sprintf("%s.%s.json", filename, command)
	}
	return command, inputPath, outputPath
}
