package dec05

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
)

// New builds a solver that can solve the exercise of Dec 05.
func New(inputPath string) door.Solver {
	return dec05Solver{inputPath: inputPath}
}

// Implementation of the solver for dec05
type dec05Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func([]int) int{
		door.Prima:   findHighestPresent,
		door.Secunda: findLowestMissing,
	}
)

// Solve implements the Solver interface
func (s dec05Solver) Solve(p door.Parts) (door.Result, error) {
	seats, err := readSeats(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to parse input file for passports: %s", err.Error())
	}
	var result dec05Result
	result.availableSeats = make(map[door.Part]*int, len(p))
	for _, part := range p {
		solver, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
		result.availableSeats[part] = ptr.Int(solver(seats))
	}
	return result, nil
}

// Implementation of the result for dec05
type dec05Result struct {
	availableSeats map[door.Part]*int
}

// String implements the Result interface
func (r dec05Result) String() string {
	if len(r.availableSeats) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The highest available int for Part %d is %d.\n", k, *r.availableSeats[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec05Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.availableSeats))
	for k := range r.availableSeats {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
