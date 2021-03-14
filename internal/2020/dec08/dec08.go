package dec08

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
)

// New builds a solver that can solve the exercise of Dec 08.
func New(inputPath string) door.Solver {
	return dec08Solver{inputPath: inputPath}
}

// Implementation of the solver for dec08
type dec08Solver struct {
	inputPath string
}

// Solve implements the Solver interface
func (s dec08Solver) Solve(p door.Parts) (door.Result, error) {
	prog, err := readInstructions(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read instructions from file '%s': %s", s.inputPath, err.Error())
	}
	result := dec08Result{acc: make(map[door.Part]*int, len(p))}
	for _, part := range p {
		solver, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
		result.acc[part] = ptr.Int(solver(prog))
	}
	return result, nil
}

// Implementation of the result for dec08
type dec08Result struct {
	acc map[door.Part]*int
}

// String implements the Result interface
func (r dec08Result) String() string {
	if len(r.acc) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The accumulated value for Part %d is %d.\n", k, *r.acc[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec08Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.acc))
	for k := range r.acc {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
