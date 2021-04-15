package dec13

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/internal/2020/dec13/notes"
	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
)

// New builds a solver that can solve the exercise of Dec 13.
func New(inputPath string) door.Solver {
	return dec13Solver{inputPath: inputPath}
}

// dec13Solver implements the solver for dec13
type dec13Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func(notes.Notes) int64{
		door.Prima: notes.FirstBus,
		door.Secunda: notes.ConsecutiveBuses,
	}
)

// Solve implements the Solver interface
func (s dec13Solver) Solve(p door.Parts) (door.Result, error) {
	n, err := notes.ReadNotes(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read notes: %s", err.Error())
	}
	result := dec13Result{map[door.Part]*int64{}}
	for _, part := range p {
		s, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
		result.data[part] = ptr.Int64(s(n))
	}
	return result, nil
}

// dec13Result implements the results for dec13
type dec13Result struct {
	data map[door.Part]*int64
}

// String implements the Result interface
func (r dec13Result) String() string {
	if len(r.data) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The highest available int for Part %d is %d.\n", k, *r.data[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec13Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.data))
	for k := range r.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
