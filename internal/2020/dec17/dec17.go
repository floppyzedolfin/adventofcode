package dec17

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/internal/2020/dec17/problem"
	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
)

// New builds a solver that can solve the exercise of Dec 17.
func New(inputPath string) door.Solver {
	return dec17Solver{inputPath: inputPath}
}

// dec17Solver implements the solver for dec17
type dec17Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func(inputPath string, iterations int) (int, error){
		door.Prima:   problem.CountCubes3D,
		door.Secunda: problem.CountCubes4D,
	}
)

// Solve implements the Solver interface
func (s dec17Solver) Solve(p door.Parts) (door.Result, error) {
	const iterationCount = 6

	result := dec17Result{data: make(map[door.Part]*int, len(p))}
	for _, part := range p {
		solver, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
		res, err := solver(s.inputPath, iterationCount)
		if err != nil {
			return nil, fmt.Errorf("unable to solve part %d: %w", p, err)
		}
		result.data[part] = ptr.Int(res)
	}
	return result, nil
}

// dec17Result implements the results for dec17
type dec17Result struct {
	data map[door.Part]*int
}

// String implements the Result interface
func (r dec17Result) String() string {
	if len(r.data) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The answer for Part %d is %d .\n", k, *r.data[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec17Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.data))
	for k := range r.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
