package adapter

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/floppyzedolfin/adventofcode/pkg/fileparser"
)

// ParseBag reads the contents of a cabin baggage.
func ParseBag(inputFile string) (Adapters, error) {
	as := Adapters{list: make([]adapter, 0)}
	err := fileparser.ParseFile(inputFile, &as)
	if err != nil {
		return Adapters{}, fmt.Errorf("unable to parse file '%s': %s", inputFile, err.Error())
	}
	// for the sake of it, we might as well sort the list right now:
	sort.Ints(as.list)
	return as, nil
}


// ParseLine implements the LineParser interface
func (as *Adapters) ParseLine(line string) error {
	a, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("unable to parse line '%s': %s", line, err.Error())
	}

	as.list = append(as.list, a)
	return nil
}
