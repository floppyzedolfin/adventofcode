package dec01

import (
	"fmt"
	"strconv"

	"github.com/floppyzedolfin/adventofcode/pkg/fileparser"
)

// readLines reads a whole file into memory
// and returns a slice of ints - its lines.
func readLines(path string) ([]int, error) {
	var i integers
	err := fileparser.ParseFile(path, &i)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file %s: %s", path, err.Error())
	}
	return i.ints, nil
}

// integers is a local utility
type integers struct {
	ints []int
}

// ParseLines implements the LineParser interface
func (i *integers) ParseLine(line string) error {
	value, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("unable to parse line, expected an int, read '%s'", line)
	}
	// store value
	i.ints = append(i.ints, value)
	return nil
}
