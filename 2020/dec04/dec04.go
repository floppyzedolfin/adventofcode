package dec04

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/common"
	"github.com/floppyzedolfin/adventofcode/door"
)

// New builds a solver that can solve the exercise of Dec 04.
func New(inputPath string) door.Solver {
	return dec04Solver{inputPath: inputPath}
}

// Implementation of the solver for dec04
type dec04Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func([]passport) int{
		door.Prima:   validatePassports,
		door.Secunda: validatePassportsThoroughly,
	}
)

// Solve implements the Solver interface
func (s dec04Solver) Solve(p door.Parts) (door.Result, error) {
	ps, err := readPassports(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to parse input file for passports: %s", err.Error())
	}
	var result dec04Result
	result.validPassports = make(map[door.Part]*int, len(p))
	for _, part := range p {
		// call the solver for that part and store a pointer to the result
		result.validPassports[part] = common.IntPointer(solvers[part](ps))
	}
	return result, nil
}

// Implementation of the result for dec04
type dec04Result struct {
	validPassports map[door.Part]*int
}

// String implements the Result interface
func (r dec04Result) String() string {
	if len(r.validPassports) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The number of valid passports for Part %d is %d.\n", k, *r.validPassports[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec04Result)sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.validPassports))
	for k := range r.validPassports {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
