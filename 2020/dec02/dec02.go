package dec02

import (
	"fmt"
	"strings"

	"github.com/floppyzedolfin/adventofcode/common"
	"github.com/floppyzedolfin/adventofcode/door"
)

// New builds a solver that can solve the exercise of Dec 02.
func New(inputPath string) door.Solver {
	return dec02Solver{inputPath: inputPath}
}

// Implementation of the solver for dec02
type dec02Solver struct {
	inputPath string
}

// Solve implements the Solver interface
func (s dec02Solver) Solve(p door.Parts) (door.Result, error) {
	passwords, err := readLines(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to parse input file '%s': %s", s.inputPath, err.Error())
	}
	var result dec02Result
	if p.Contains(door.Part1) {
		result.validPasswordsPart1 = common.IntPointer(countValidPasswordsPart1(passwords))
	}
	if p.Contains(door.Part2) {
		result.validPasswordsPart2 = common.IntPointer(countValidPasswordsPart2(passwords))
	}
	return result, nil
}

// Implementation of the result for dec02
type dec02Result struct {
	validPasswordsPart1 *int
	validPasswordsPart2 *int
}

// String implements the Result interface
func (r dec02Result) String() string {
	if r.validPasswordsPart1 == nil && r.validPasswordsPart2 == nil {
		return "No job done by the elves today."
	}
	output := strings.Builder{}
	if r.validPasswordsPart1 != nil {
		output.WriteString(fmt.Sprintf("The number of valid passwords for Part 1 is %d.\n", *r.validPasswordsPart1))
	}
	if r.validPasswordsPart2 != nil {
		output.WriteString(fmt.Sprintf("The number of valid passwords for Part 2 is %d.\n", *r.validPasswordsPart2))
	}
	return output.String()
}
