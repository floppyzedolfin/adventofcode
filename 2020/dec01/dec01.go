package dec01

import (
	"fmt"
	"strings"

	"github.com/floppyzedolfin/adventofcode/common"
	"github.com/floppyzedolfin/adventofcode/door"
)

// New builds a solver that can solve the exercise of Dec 01.
func New(inputPath string) door.Solver {
	return dec01Solver{inputPath: inputPath}
}

// Implementation of the solver for dec01
type dec01Solver struct {
	inputPath string
}

// Solve implements the Solver interface
func (s dec01Solver) Solve(p door.Parts) (door.Result, error) {
	lines, err := readLines(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read lines: %s", err.Error())
	}

	var result dec01Result
	if p.Contains(door.Prima) {
		// Choose the implementation here
		matches, err := findMatch2(lines, door.Year)
		if err != nil {
			return nil, fmt.Errorf("unable to find match: %s", err.Error())
		}
		result.productPrima = common.IntPointer(product(matches))
	}
	if p.Contains(door.Secunda) {
		// Choose the implementation here
		matches, err := findMatch3(lines, door.Year)
		if err != nil {
			return nil, fmt.Errorf("unable to find match: %s", err.Error())
		}
		result.productSecunda = common.IntPointer(product(matches))
	}
	return result, nil
}

// Implementation of the result for dec01
type dec01Result struct {
	productPrima *int
	productSecunda *int
}

// String implements the Result interface
func (r dec01Result) String() string {
	if r.productPrima == nil && r.productSecunda == nil {
		return "No job done by the elves today."
	}
	output := strings.Builder{}
	if r.productPrima != nil {
		output.WriteString(fmt.Sprintf("The product for Part 1 is %d.\n", *r.productPrima))
	}
	if r.productSecunda != nil {
		output.WriteString(fmt.Sprintf("The product for Part 2 is %d.\n", *r.productSecunda))
	}
	return output.String()
}
