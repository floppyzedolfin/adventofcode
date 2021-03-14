package dec06

import (
	"fmt"
	"strings"

	"github.com/floppyzedolfin/adventofcode/pkg/fileparser"
)

// groupList is a local struct we use to parse a file.
type groupList struct {
	list []group
}

// ParseLine implements the LineParser interface
// Rather than adding answers for each line, this func will mainly update the last group, and create one on empty lines
// Notice that the response is always nil, but we must comply with the interface.
func (gsa *groupList) ParseLine(line string) error {
	// if this line is blank, create a new group of answers
	if len(strings.TrimSpace(line)) == 0 {

		gsa.list = append(gsa.list, group{})
		// nothing else to do on empty lines
		return nil
	}

	groupCount := len(gsa.list)
	// if this is the first line, populate the slice with a group
	if groupCount == 0 {
		// add a person with no response so far.
		gsa.list = []group{{people: make([]person, 0)}}
		groupCount = 1
	}

	// each line represents a person
	p := person{response: make(answers)}
	for _, answer := range []byte(line) {
		// update this person's answers
		p.response[answer] = struct{}{}
	}
	// add the person to the list
	gsa.list[groupCount-1].people = append(gsa.list[groupCount-1].people, p)
	return nil
}

// readGroups parses a file and returns a list of groups. Or an error, should that ever happen.
func readGroups(inputPath string) ([]group, error) {
	var gsa groupList
	err := fileparser.ParseFile(inputPath, &gsa)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file '%s': %s", inputPath, err.Error())
	}
	return gsa.list, nil
}
