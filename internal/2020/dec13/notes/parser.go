package notes

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/floppyzedolfin/adventofcode/pkg/fileparser"
)

// ParseLine implements the LineParser interface
func (n *Notes) ParseLine(line string) error {
	var err error
	if !n.initialised {
		n.currentTime, err = strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("invalid line '%s', expected an number for the current time", line)
		}
		n.initialised = true
	} else {
		buses := strings.Split(line, ",")
		for i, b := range buses {
			if b == "X" {
				// this isn't a bus
				continue
			}
			busNumber, err := strconv.Atoi(b)
			if err != nil {
				return fmt.Errorf("invalid bus number '%s'", b)
			}
			n.buses = append(n.buses, bus{busNumber, busNumber - i})
		}
	}
	return nil
}

// ReadNotes reads the input file and returns a usable Notes object
func ReadNotes(inputPath string) (Notes, error) {
	var n Notes
	err := fileparser.ParseFile(inputPath, &n)
	if err != nil {
		return Notes{}, fmt.Errorf("unable to parser file '%s': %s", inputPath, err.Error())
	}
	return n, nil
}
