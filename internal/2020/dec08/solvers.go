package dec08

import (
	"github.com/floppyzedolfin/adventofcode/internal/2020/dec08/program"
	"github.com/floppyzedolfin/adventofcode/pkg/door"
)

var (
	solvers = map[door.Part]func(program.Program) int{
		door.Prima:   program.RunTillLoop,
		door.Secunda: program.FixOnce,
	}
)
