package dec09

import (
	"github.com/floppyzedolfin/adventofcode/pkg/door"
)

var (
	solvers = map[door.Part]func(string, int) (*int, error){
		door.Prima: xmas1Solver,
		door.Secunda: xmas2Solver,
	}
)
