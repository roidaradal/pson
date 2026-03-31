# pson 
A tool for indenting, aligning, and compressing JSON files, written in Go

## Installation 
`go install github.com/zeroibot/pson@latest` 

OR 

Download `pson.exe` from the [releases](https://github.com/zeroibot/pson/releases) page. Add the folder where you saved `pson.exe` to your system PATH.

## Usage 
`pson <align|compress|indent|validate|rawdiff|version|update> <file.json> (output=path) (with=path) (overwrite) (indent=2) (flatlist)`

### Commands
* `align` - Saves the JSON with object keys aligned
* `compress` - Saves compressed JSON (no whitespace)
* `indent` - Saves indented JSON 
* `validate` - Validate JSON
* `rawdiff` - Compares two JSON files per line
* `version` - Output current version
* `update` - Update pson to the latest version

### Options 
* `output=path` - Set custom output file path 
* `with=path` - Set file path for `diff` and `rawdiff`
* `overwrite` - Makes the output file path same as input file (overwrite)
* `indent=X` - Sets the number of spaces as indent level for the `indent` command
* `flatlist` - Enables lists (except main list) to be displayed inline, instead of indented

Note: If both `output=path` and `overwrite` are used, the last option will be followed.

## TODO
* Rawdiff - compare line per line
    - Advanced: find the best alignment of lines (detect insertions / deletions)
* Diff with another JSON (analyze both contents to compare)
* Validation: duplicate keys, inconsistent key types, etc
* Schema validation: given an object schema, check if JSON follows it 
* Generate Go struct / TypeScript object from JSON 
* Convert to another format (CSV, TOML, ENV, etc)