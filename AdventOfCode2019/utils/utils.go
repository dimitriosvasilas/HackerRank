package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// ReadFileStruct is a struct used of reading input from a file.
type ReadFileStruct struct {
	file    *os.File
	scanner *bufio.Scanner
}

// ReadFile opens a file in the given path, and returns a ReadFileStruct.
func ReadFile(filePath string) (*ReadFileStruct, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return &ReadFileStruct{}, err
	}
	scanner := bufio.NewScanner(file)
	readFileStruct := &ReadFileStruct{file: file, scanner: scanner}
	return readFileStruct, nil
}

// Close closes a previously opened file.
func (f *ReadFileStruct) Close() {
	f.file.Close()
}

// ReadLines reads a given number of lines from an open file.
func (f *ReadFileStruct) ReadLines(lineCount int) ([]string, error) {
	lines := make([]string, lineCount)
	for i := 0; i < lineCount; i++ {
		if f.scanner.Scan() {
			lines[i] = f.scanner.Text()
		} else {
			err := f.scanner.Err()
			if err != nil {
				return nil, err
			} else {
				return lines, io.EOF
			}
		}
	}
	return lines, nil
}

// ParseInputLine parses a given string based on a template of datatypes
// and performs the necessary casts from strings to the given types.
// sep is a separator used to split the given string to multiple segments
// repeat indicated wether the given template should be repeated for the length of the splitted input
// template describes the target datatypes
func ParseInputLine(inputString string, sep string, repeat bool, template ...interface{}) ([]interface{}, error) {
	input := strings.Split(inputString, sep)
	parsedInput := make([]interface{}, len(input))
	if repeat {
		switch template[0].(type) {
		case int:
			for i, elem := range input {
				res, err := strconv.ParseInt(elem, 10, 64)
				if err != nil {
					return nil, errors.New(fmt.Sprint("error converting ", input[i], " to int"))
				}
				parsedInput[i] = int(res)
			}
		case string:
			for i, elem := range input {
				parsedInput[i] = elem
			}
		default:
			return nil, errors.New("unknown data type")
		}
		return parsedInput, nil
	}
	if len(input) != len(template) {
		return nil, errors.New("input string and template lengths do not match")
	}
	for i, arg := range template {
		switch arg.(type) {
		case int:
			res, err := strconv.ParseInt(input[i], 10, 64)
			if err != nil {
				return nil, errors.New(fmt.Sprint("error converting ", input[i], " to int"))
			}
			parsedInput[i] = int(res)
		default:
			return nil, errors.New("unknown data type")
		}
	}
	return parsedInput, nil
}

// ReadAndParseLines combines ReadLines and ParseInputLine by reading and parsing a given number of lines.
func (f *ReadFileStruct) ReadAndParseLines(lineCount int, sep string, repeat bool, template ...interface{}) ([][]interface{}, error) {
	lines, err := f.ReadLines(lineCount)
	if err != nil {
		return nil, err
	}
	parseLines := make([][]interface{}, len(lines))
	for i, line := range lines {
		parseLines[i], err = ParseInputLine(line, sep, repeat, template...)
		if err != nil {
			return nil, err
		}
	}
	return parseLines, nil
}
