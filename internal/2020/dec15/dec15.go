package dec15

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/internal/2020/dec15/game"
	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
)

// New builds a solver that can solve the exercise of Dec 15.
func New(inputPath string) door.Solver {
	return dec15Solver{inputPath: inputPath}
}

// dec15Solver implements the solver for dec15
type dec15Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func(string) (int, error){
		door.Prima:   game.Solve2020,
		door.Secunda: game.Solve30000000,
	}
)

// Solve implements the Solver interface
func (s dec15Solver) Solve(p door.Parts) (door.Result, error) {
	result := dec15Result{map[door.Part]*int{}}
	for _, part := range p {
		solver, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
		res, err := solver(s.inputPath)
		if err != nil {
			return nil, fmt.Errorf("eror while solving problem: %s", err.Error())
		}
		result.data[part] = ptr.Int(res)
	}
	return result, nil
}

// dec15Result implements the results for dec15
type dec15Result struct {
	data map[door.Part]*int
}

// String implements the Result interface
func (r dec15Result) String() string {
	if len(r.data) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The answer for Part %d is %d.\n", k, *r.data[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec15Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.data))
	for k := range r.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
