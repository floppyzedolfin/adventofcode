package traintickets

import (
	"fmt"

	"github.com/floppyzedolfin/adventofcode/fileparser"
)

// SumInvalidValues returns the sum of values that can't be mapped to a field
func SumInvalidValues(inputFile string) (uint64, error) {
	p := newProblem()
	err := fileparser.ParseFile(inputFile, p)
	if err != nil {
		return 0, fmt.Errorf("unable to parse file '%s': %s", inputFile, err.Error())
	}

	return p.sumInvalidValues(), nil
}

// MultiplyDepartures returns the product of values of the "*departures*" fields
func MultiplyDepartures(inputFile string) (uint64, error) {
	p := newProblem()
	err := fileparser.ParseFile(inputFile, p)
	if err != nil {
		return 0, fmt.Errorf("unable to parse file '%s': %s", inputFile, err.Error())
	}

	return p.multiplyDepartures(), nil
}
