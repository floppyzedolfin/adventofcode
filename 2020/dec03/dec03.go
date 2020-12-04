package dec03

import (
	"fmt"
	"strings"

	"github.com/floppyzedolfin/adventofcode/common"
	"github.com/floppyzedolfin/adventofcode/door"
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
	slope, err := readLines(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to parse input file '%s': %s", s.inputPath, err.Error())
	}
	var result dec03Result
	if p.Contains(door.Part1) {
		result.treeCountPart1 = common.IntPointer(countTreesPart1(slope))
	}
	if p.Contains(door.Part2) {
		result.treeCountPart2 = common.IntPointer(countTreesPart2(slope))
	}
	return result, nil
}

// Implementation of the result for dec03
type dec03Result struct {
	treeCountPart1 *int
	treeCountPart2 *int
}

// String implements the Result interface
func (r dec03Result) String() string {
	if r.treeCountPart1 == nil && r.treeCountPart2 == nil {
		return "No job done by the elves today."
	}
	output := strings.Builder{}
	if r.treeCountPart1 != nil {
		output.WriteString(fmt.Sprintf("The number of trees hit for Part 1 is %d.\n", *r.treeCountPart1))
	}
	if r.treeCountPart2 != nil {
		output.WriteString(fmt.Sprintf("The number of trees hit for Part 2 is %d.\n", *r.treeCountPart2))
	}
	return output.String()
}
