package dec03

import (
	"fmt"
	"strings"

	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
)

// New builds a solver that can solve the exercise of Dec 03.
func New(inputPath string) door.Solver {
	return dec03Solver{inputPath: inputPath}
}

// Implementation of the solver for dec03
type dec03Solver struct {
	inputPath string
}

// Solve implements the Solver interface
func (s dec03Solver) Solve(p door.Parts) (door.Result, error) {
	f, err := readForest(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to parse input file '%s': %s", s.inputPath, err.Error())
	}
	var result dec03Result
	if p.Contains(door.Prima) {
		result.treeCountPrima = ptr.Int(f.countTreesPrima())
	}
	if p.Contains(door.Secunda) {
		result.treeCountSecunda = ptr.Int(f.countTreesSecunda())
	}
	return result, nil
}

// Implementation of the result for dec03
type dec03Result struct {
	treeCountPrima *int
	treeCountSecunda *int
}

// String implements the Result interface
func (r dec03Result) String() string {
	if r.treeCountPrima == nil && r.treeCountSecunda == nil {
		return "No job done by the elves today."
	}
	output := strings.Builder{}
	if r.treeCountPrima != nil {
		output.WriteString(fmt.Sprintf("The number of trees hit for Part 1 is %d.\n", *r.treeCountPrima))
	}
	if r.treeCountSecunda != nil {
		output.WriteString(fmt.Sprintf("The number of trees hit for Part 2 is %d.\n", *r.treeCountSecunda))
	}
	return output.String()
}
