package dec02

import (
	"fmt"
	"strings"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
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
	if p.Contains(door.Prima) {
		result.validPasswordsPrima = ptr.Int(countValidPasswordsPrima(passwords))
	}
	if p.Contains(door.Secunda) {
		result.validPasswordsSecunda = ptr.Int(countValidPasswordsSecunda(passwords))
	}
	return result, nil
}

// Implementation of the result for dec02
type dec02Result struct {
	validPasswordsPrima *int
	validPasswordsSecunda *int
}

// String implements the Result interface
func (r dec02Result) String() string {
	if r.validPasswordsPrima == nil && r.validPasswordsSecunda == nil {
		return "No job done by the elves today."
	}
	output := strings.Builder{}
	if r.validPasswordsPrima != nil {
		output.WriteString(fmt.Sprintf("The number of valid passwords for Part 1 is %d.\n", *r.validPasswordsPrima))
	}
	if r.validPasswordsSecunda != nil {
		output.WriteString(fmt.Sprintf("The number of valid passwords for Part 2 is %d.\n", *r.validPasswordsSecunda))
	}
	return output.String()
}
