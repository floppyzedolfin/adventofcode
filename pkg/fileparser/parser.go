package fileparser

import (
	"bufio"
	"fmt"
	"os"
)

// LineParser defines the interface we need when parsing lines
type LineParser interface {
	ParseLine(string) error
}

// ParseFile reads a file line by line, and applies something on each line. Check tests for examples
func ParseFile(inputPath string, lineParser LineParser) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("unable to read file '%s': %s", inputPath, err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		err := lineParser.ParseLine(line)
		if err != nil {
			return fmt.Errorf("an error occurred on line %s: %s", line, err.Error())
		}
	}
	return scanner.Err()
}
