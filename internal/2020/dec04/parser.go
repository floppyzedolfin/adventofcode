package dec04

import (
	"fmt"
	"strings"

	"github.com/floppyzedolfin/adventofcode/pkg/fileparser"
)

// passports is a local struct we use to parse a file.
type passports struct {
	list []passport
}

// ParseLine implements the LineParser interface
// Rather than adding passports for each line, this func will mainly update the last passport, and create one on empty lines
func (ps *passports) ParseLine(line string) error {
	// if this line is blank, create a new passport to add to the list of passports
	if len(strings.TrimSpace(line)) == 0 {
		ps.list = append(ps.list, passport{})
		// nothing else to do on empty lines
		return nil
	}

	passportCount := len(ps.list)
	// if this is the first line, populate the slice with a passport
	if passportCount == 0 {
		// add an empty passport
		ps.list = []passport{{}}
		passportCount = 1
	}

	const fieldSeparator = " "
	// at this point, there is always at least 1 passport in the list of passports, and this line we're reading contains data
	fields := strings.Split(line, fieldSeparator)
	for _, field := range fields {
		// update the last item
		err := ps.list[passportCount-1].setField(field)
		if err != nil {
			return fmt.Errorf("error while parsing field '%s'", field)
		}
	}
	return nil
}

// readPassports parses a file and returns a list of passports. Or an error, should that ever happen.
func readPassports(inputPath string) ([]passport, error) {
	var ps passports
	err := fileparser.ParseFile(inputPath, &ps)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file '%s': %s", inputPath, err.Error())
	}
	return ps.list, nil
}
