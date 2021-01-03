package dec11

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/2020/dec11/waitingarea"
	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
)

// New builds a solver that can solve the exercise of Dec 11.
func New(inputPath string) door.Solver {
	return dec11Solver{inputPath: inputPath}
}

// dec11Solver implements the solver for dec11
type dec11Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func(*waitingarea.WaitingArea) int{
		door.Prima: waitingarea.OccupiedSeatsCloseRange,
		door.Secunda: waitingarea.OccupiedSeatsLongRange,
	}
)

// Solve implements the Solver interface
func (s dec11Solver) Solve(p door.Parts) (door.Result, error) {
	wa, err := waitingarea.ReadSeats(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read seats from file '%s': %s", s.inputPath, err.Error())
	}
	result := dec11Result{data: map[door.Part]*int{}}
	for _, part := range p {
		s, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
		result.data[part] = ptr.Int(s(wa))
	}
	return result, nil
}

// Implementation of the result for dec11
type dec11Result struct {
	data map[door.Part]*int
}

// String implements the Result interface
func (r dec11Result) String() string {
	if len(r.data) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The number of occupied seats for Part %d is %d.\n", k, *r.data[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec11Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.data))
	for k := range r.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
