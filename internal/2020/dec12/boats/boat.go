package boats

import (
	"github.com/floppyzedolfin/adventofcode/internal/2020/dec12/instructions"
	"github.com/floppyzedolfin/adventofcode/internal/2020/dec12/position"
)

// boat holds the current state of the boat that navigates with 4 directions only
type boat struct {
	coordinates position.Position
	orientation int // angle of 0, by default, is facing East.
}

// BoatL1 operates a boat till the end of the instructions list, and returns its Manhattan distance to origin
func BoatL1(is instructions.Instructions) int {
	b := boat{}
	for _, i := range is.Commands {
		b.operate(i)
	}
	return b.coordinates.L1()
}

// operate performs an instruction
func (b *boat) operate(inst instructions.Instruction) {
	switch inst.Action {
	case instructions.MoveNorth:
		b.coordinates.Move(0, inst.Value)
	case instructions.MoveEast:
		b.coordinates.Move(inst.Value, 0)
	case instructions.MoveSouth:
		b.coordinates.Move(0, -inst.Value)
	case instructions.MoveWest:
		b.coordinates.Move(-inst.Value, 0)
	case instructions.TurnPort:
		b.orientation += inst.Value
		b.orientation = cleanOrientation(b.orientation)
	case instructions.TurnStarboard:
		b.orientation -= inst.Value
		b.orientation = cleanOrientation(b.orientation)
	case instructions.Forward:
		b.advance(inst.Value)
	}
}

// advance moves the ship in the direction of its orientation
func (b *boat) advance(v int) {
	switch b.orientation {
	case 0:
		// eastwards
		b.coordinates.Move(v, 0)
	case 90:
		// northwards
		b.coordinates.Move(0, v)
	case 180:
		// westwards
		b.coordinates.Move(-v, 0)
	case 270:
		// southwards
		b.coordinates.Move(0, -v)
	default:
		// stall. unknown direction.
	}
}
