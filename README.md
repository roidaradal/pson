# pson 
A tool for indenting, aligning, and compressing JSON files, written in Go

## Installation 
`go install github.com/roidaradal/pson@latest` 

OR 

Download `pson.exe` from the [releases](https://github.com/roidaradal/pson/releases) page. Add the folder where you saved `pson.exe` to your system PATH.

## Usage 
`pson <align|compress|indent|validate|version> <file.json> (output=path) (overwrite) (indent=2) (flatlist)`

### Commands
* `align` - Saves the JSON with object keys aligned
* `compress` - Saves compressed JSON (no whitespace)
* `indent` - Saves indented JSON 
* `validate` - Validate JSON
* `version` - Output current version

### Options 
* `output=path` - Set custom output file path 
* `overwrite` - Makes the output file path same as input file (overwrite)
* `indent=X` - Sets the number of spaces as indent level for the `indent` command
* `flatlist` - Enables lists (except main list) to be displayed inline, instead of indented

Note: If both `output=path` and `overwrite` are used, the last option will be followed.

## TODO
* Use io.GetCommandOptions
* Validation: check for syntax errors (missing commas, dangling commas), duplicate keys, inconsistent key types, etc
* Schema validation: given an object schema, check if JSON follows it 
* Generate Go struct / TypeScript object from JSON 
* Diff with another JSON (analyze both contents to compare)
* Convert to another format (CSV, TOML, ENV, etc)