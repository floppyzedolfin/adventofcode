package dec02

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/floppyzedolfin/adventofcode/fileparser"
)

// readLines reads a whole file into memory
// and returns a slice of passwords - its lines.
func readLines(path string) ([]password, error) {
	var ps passwords
	err := fileparser.ParseFile(path, &ps)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file '%s': %s", path, err.Error())
	}
	return ps.p, nil
}

// passwords is a utility we use to parse a file
type passwords struct {
	p []password
}

// ParseLine implements the LineParser interface
func (ps *passwords) ParseLine(line string) error {
	pw, err := buildPassword(line)
	if err != nil {
		return fmt.Errorf("error while building password from '%s': %s", line, err.Error())
	}
	ps.p = append(ps.p, pw)
	return nil
}

// buildPassword parses a line and returns a password
func buildPassword(line string) (password, error) {
	const validLineRegExp = `(\d+)-(\d+) ([[:alpha:]]): ([[:alnum:]]+)$`

	// parse the line to extract the information we need
	r := regexp.MustCompile(validLineRegExp)
	res := r.FindStringSubmatch(line)
	if len(res) != 5 {
		return password{}, fmt.Errorf("unable to extract password from line %s", line)
	}

	// convert the values to integers for later use
	min, err := strconv.Atoi(res[1])
	if err != nil {
		return password{}, fmt.Errorf("unable to build password, min %s is not an integer", res[1])
	}
	max, err := strconv.Atoi(res[2])
	if err != nil {
		return password{}, fmt.Errorf("unable to build password, max %s is not an integer", res[2])
	}

	// finally build the password
	return password{
		value: res[4],
		policy: policy{
			min:    min,
			max:    max,
			letter: res[3],
		},
	}, nil
}
