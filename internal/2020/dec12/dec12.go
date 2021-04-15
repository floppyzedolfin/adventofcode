package dec12

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/internal/2020/dec12/boats"
	"github.com/floppyzedolfin/adventofcode/internal/2020/dec12/instructions"
	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
)

// New builds a solver that can solve the exercise of Dec 12.
func New(inputPath string) door.Solver {
	return dec12Solver{inputPath: inputPath}
}

// dec12Solver implements the solver for dec12
type dec12Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func(instructions.Instructions) int{
		door.Prima: boats.BoatL1,
		door.Secunda: boats.WaypointBoatL1,
	}
)

// Solve implements the Solver interface
func (s dec12Solver) Solve(p door.Parts) (door.Result, error) {
	is, err := instructions.ReadNavigationInstructions(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to parse input file '%s': %s", s.inputPath, err.Error())
	}
	result := dec12Result{data: map[door.Part]*int{}}
	for _, part := range p {
		s, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
		result.data[part] = ptr.Int(s(is))
	}
	return result, nil
}

// Implementation of the result for dec12
type dec12Result struct {
	data map[door.Part]*int
}

// String implements the Result interface
func (r dec12Result) String() string {
	if len(r.data) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The distance for Part %d is %d.\n", k, *r.data[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec12Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.data))
	for k := range r.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
