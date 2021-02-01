package decXX

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/door"
)

// New builds a solver that can solve the exercise of Dec XX.
func New(inputPath string) door.Solver {
	return decXXSolver{inputPath: inputPath}
}

// decXXSolver implements the solver for decXX
type decXXSolver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func([]int) int{
	}
)

// Solve implements the Solver interface
func (s decXXSolver) Solve(p door.Parts) (door.Result, error) {
	var result decXXResult
	for _, part := range p {
		_, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
	}
	return result, nil
}

// decXXResult implements the results for decXX
type decXXResult struct {
	data map[door.Part]*int
}

// String implements the Result interface
func (r decXXResult) String() string {
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
func (r decXXResult) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.data))
	for k := range r.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
