package dec09

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/pkg/door"
)

// New builds a solver that can solve the exercise of Dec 09.
func New(inputPath string) door.Solver {
	return dec09Solver{inputPath: inputPath}
}

// Implementation of the solver for dec09
type dec09Solver struct {
	inputPath string
}

// this var can be overridden by tests
var preambleLength = 25

// Solve implements the Solver interface
func (s dec09Solver) Solve(p door.Parts) (door.Result, error) {
	result := dec09Result{invalidNumber: make(map[door.Part]*int, len(p))}
	for _, part := range p {
		solver, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		var err error
		// call the solver for that part and store a pointer to the result
		result.invalidNumber[part], err = solver(s.inputPath, preambleLength)
		if err != nil {
			return nil, fmt.Errorf("error while solving part %v: %s", part, err.Error())
		}
	}
	return result, nil
}

// Implementation of the result for dec09
type dec09Result struct {
	invalidNumber map[door.Part]*int
}

// String implements the Result interface
func (r dec09Result) String() string {
	if len(r.invalidNumber) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		if r.invalidNumber[k] == nil {
			output.WriteString(fmt.Sprintf("No invalid number for Part %d.\n", k))
		} else {
			output.WriteString(fmt.Sprintf("The first invalid number for Part %d is %d.\n", k, *r.invalidNumber[k]))
		}
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec09Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.invalidNumber))
	for k := range r.invalidNumber {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
