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

const currentVersion string = "v0.1.4"
const usage string = "Usage: pson <align|compress|indent|validate|update> <file.json> (output=path) (overwrite) (indent=2) (flatlist)"

var optionalFile = map[string]bool{
	"version": true,
	"update":  true,
}

func main() {
	var err error
	command, path1, path2 := getArgs()
	switch command {
	case "compress":
		err = pson.CompressJSON(path1, path2)
	case "indent":
		err = pson.IndentJSON(path1, path2)
	case "align":
		err = pson.AlignJSON(path1, path2)
	case "validate":
		err = pson.ValidateJSON(path1)
	case "rawdiff":
		err = pson.RawDiffJSON(path1, path2)
	case "version":
		fmt.Println(str.Green(fmt.Sprintf("pson %s", currentVersion)))
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
func getArgs() (command, path1, path2 string) {
	args := os.Args[1:]
	numArgs := len(args)
	if numArgs < 1 {
		fmt.Println(usage)
		os.Exit(1)
	}
	command, path1 = args[0], ""
	if numArgs >= 2 {
		path1 = args[1]
	}

	if optionalFile[command] {
		// End early if files are optional
		return command, "", ""
	}

	if !strings.HasSuffix(path1, ".json") {
		fmt.Println("File path needs to be a .json file")
		os.Exit(1)
	}

	// Default output path
	filename, _ := strings.CutSuffix(path1, ".json")
	path2 = fmt.Sprintf("%s.%s.json", filename, command)

	for _, arg := range args[2:] {
		if arg == "overwrite" {
			path2 = path1
		} else if arg == "flatlist" {
			pson.FLAT_LIST = true
		} else if strings.HasPrefix(arg, "indent=") {
			customIndent, ok := str.TryGetPart(arg, "=", 1)
			if !ok || customIndent == "" {
				continue
			}
			pson.INDENT_SPACE = max(pson.INDENT_SPACE, number.ParseInt(customIndent))
			str.SetJSONIndentLength(pson.INDENT_SPACE)
		} else if strings.HasPrefix(arg, "output=") || strings.HasPrefix(arg, "with=") {
			path, ok := str.TryGetPart(arg, "=", 1)
			if !ok || path == "" {
				continue
			}
			path2 = path
		}
	}
	return command, path1, path2
}
