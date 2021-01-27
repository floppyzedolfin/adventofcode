package dec14

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/door"
)

// New builds a solver that can solve the exercise of Dec 14.
func New(inputPath string) door.Solver {
	return dec14Solver{inputPath: inputPath}
}

// dec14Solver implements the solver for dec14
type dec14Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func(inputFile string) (*uint64, error){
		door.Prima:   sumValuesV1,
		door.Secunda: sumValuesV2,
	}
)

// Solve implements the Solver interface
func (s dec14Solver) Solve(p door.Parts) (door.Result, error) {
	result := dec14Result{map[door.Part]*uint64{}}
	var err error
	for _, part := range p {
		solver, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
		result.data[part], err = solver(s.inputPath)
		if err != nil {
			return nil, fmt.Errorf("error while solving part %d: %s", part, err.Error())
		}
	}
	return result, nil
}

// dec14Result implements the results for dec14
type dec14Result struct {
	data map[door.Part]*uint64
}

// String implements the Result interface
func (r dec14Result) String() string {
	if len(r.data) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The result for Part %d is %d.\n", k, *r.data[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec14Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.data))
	for k := range r.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
