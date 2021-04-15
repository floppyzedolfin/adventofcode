package dec10

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/internal/2020/dec10/adapter"
	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
)

// New builds a solver that can solve the exercise of Dec 10.
func New(inputPath string) door.Solver {
	return dec10Solver{inputPath: inputPath}
}

// Implementation of the solver for dec10
type dec10Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func(adapter.Adapters) int64{
		door.Prima:   adapter.Jolts1By3,
		door.Secunda: adapter.DifferentCombinations,
	}
)

// Solve implements the Solver interface
func (s dec10Solver) Solve(p door.Parts) (door.Result, error) {
	adapters, err := adapter.ParseBag(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read adapters from file '%s': %s", s.inputPath, err.Error())
	}
	result := dec10Result{data: map[door.Part]*int64{}}
	for _, part := range p {
		s, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
		result.data[part] = ptr.Int64(s(adapters))
	}
	return result, nil
}

// Implementation of the result for dec10
type dec10Result struct {
	data map[door.Part]*int64
}

// String implements the Result interface
func (r dec10Result) String() string {
	if len(r.data) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The joltage computation for Part %d is %d.\n", k, *r.data[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec10Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.data))
	for k := range r.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
