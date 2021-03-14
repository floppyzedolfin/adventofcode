package dec16

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/internal/2020/dec16/traintickets"
	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
)

// New builds a solver that can solve the exercise of Dec 16.
func New(inputPath string) door.Solver {
	return dec16Solver{inputPath: inputPath}
}

// dec16Solver implements the solver for dec16
type dec16Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func(string) (uint64, error){
		door.Prima: traintickets.SumInvalidValues,
		door.Secunda: traintickets.MultiplyDepartures,
	}
)

// Solve implements the Solver interface
func (s dec16Solver) Solve(p door.Parts) (door.Result, error) {
	result := dec16Result{map[door.Part]*uint64{}}
	for _, part := range p {
		solver, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
		res, err := solver(s.inputPath)
		if err != nil {
			return nil, fmt.Errorf("error while running part %d: %s", part, err.Error())
		}
		result.data[part] = ptr.Uint64(res)
	}
	return result, nil
}

// dec16Result implements the results for dec16
type dec16Result struct {
	data map[door.Part]*uint64
}

// String implements the Result interface
func (r dec16Result) String() string {
	if len(r.data) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The answser for Part %d is %d .\n", k, *r.data[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec16Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.data))
	for k := range r.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
