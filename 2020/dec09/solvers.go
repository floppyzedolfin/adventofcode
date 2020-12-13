package dec09

import (
	"github.com/floppyzedolfin/adventofcode/door"
)

var (
	solvers = map[door.Part]func(string, int) (*int, error){
		door.Prima: xmas1Solver,
		door.Secunda: xmas2Solver,
	}
)
