package dec07

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/common"
	"github.com/floppyzedolfin/adventofcode/door"
)

// New builds a solver that can solve the exercise of Dec 07.
func New(inputPath string) door.Solver {
	return dec07Solver{inputPath: inputPath}
}

// Implementation of the solver for dec07
type dec07Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func(bagDictionary, string) int{
		door.Prima: countUniqueAncestorColours,
		door.Secunda: countContainedBags,
	}
)

// Solve implements the Solver interface
func (s dec07Solver) Solve(p door.Parts) (door.Result, error) {
	dic, err := readBags(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to parse input file '%s': %s", s.inputPath, err.Error())
	}
	result := dec07Result{data: make(map[door.Part]*int, len(p))}
	for _, part := range p {
		solver, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d'", part)
		}
		// call the solver for that part and store a pointer to the result
		const shinyGoldBags = "shiny gold"
		result.data[part] = common.IntPointer(solver(dic, shinyGoldBags))
	}
	return result, nil
}

// Implementation of the result for dec07
type dec07Result struct {
	data map[door.Part]*int
}

// String implements the Result interface
func (r dec07Result) String() string {
	if len(r.data) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The bag count for Part %d is %d.\n", k, *r.data[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec07Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.data))
	for k := range r.data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
