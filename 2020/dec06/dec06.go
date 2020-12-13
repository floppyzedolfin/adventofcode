package dec06

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
)

// New builds a solver that can solve the exercise of Dec 06.
func New(inputPath string) door.Solver {
	return dec06Solver{inputPath: inputPath}
}

// Implementation of the solver for dec06
type dec06Solver struct {
	inputPath string
}

var (
	solvers = map[door.Part]func([]group) int{
		door.Prima: countPart1,
		door.Secunda: countPart2,
	}
)

// Solve implements the Solver interface
func (s dec06Solver) Solve(p door.Parts) (door.Result, error) {
	groups, err := readGroups(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the input file '%s': %s", s.inputPath, err.Error())
	}
	var result dec06Result
	result.countSum = make(map[door.Part]*int, len(p))
	for _, part := range p {
		s, ok := solvers[part]
		if !ok {
			return nil, fmt.Errorf("invalid part '%d', no solver found", part)
		}
		// call the solver for that part and store a pointer to the result
		result.countSum[part] = ptr.Int(s(groups))
	}
	return result, nil
}

// Implementation of the result for dec06
type dec06Result struct {
	countSum map[door.Part]*int
}

// String implements the Result interface
func (r dec06Result) String() string {
	if len(r.countSum) == 0 {
		return fmt.Sprint("No job done by the elves today.")
	}
	output := strings.Builder{}
	for _, k := range r.sortParts() {
		output.WriteString(fmt.Sprintf("The sum of answer group count for Part %d is %d.\n", k, *r.countSum[k]))
	}
	return output.String()
}

// sortParts helps making things a bit more deterministic.
func (r dec06Result) sortParts() []door.Part {
	keys := make([]door.Part, 0, len(r.countSum))
	for k := range r.countSum {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
