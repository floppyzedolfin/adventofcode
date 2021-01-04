package instructions

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/floppyzedolfin/adventofcode/fileparser"
)

// ReadNavigationInstructions returns the instructions read from the source
func ReadNavigationInstructions(inputFile string) (Instructions, error) {
	i := Instructions{Commands: make([]Instruction, 0)}

	err := fileparser.ParseFile(inputFile, &i)
	if err != nil {
		return Instructions{}, fmt.Errorf("unable to parse file '%s': %s", inputFile, err.Error())
	}

	return i, nil
}

// ParseLine implements the LineParser interface
func (is *Instructions) ParseLine(line string) error {
	const instructionRE = `([NSFWLER])([0-9]+)`
	re := regexp.MustCompile(instructionRE)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 3 {
		return fmt.Errorf("invalid line '%s', doesn't support the syntax", line)
	}
	v, err := strconv.Atoi(matches[2])
	if err != nil {
		return fmt.Errorf("unable to extract the value from the line '%s', found '%s'", line, matches[2])
	}
	i := Instruction{
		Action: Action(matches[1]),
		Value:  v,
	}
	is.Commands = append(is.Commands, i)
	return nil
}
