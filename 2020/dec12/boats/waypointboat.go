package boats

import (
	"github.com/floppyzedolfin/adventofcode/2020/dec12/instructions"
	"github.com/floppyzedolfin/adventofcode/2020/dec12/position"
)

// waypointBoat is a boat that navigates using a waypoint
type waypointBoat struct {
	waypoint    position.Position
	coordinates position.Position
	orientation int
}

// WaypointBoatL1 operates a waypoint boat till the end of the instructions list, and returns its Manhattan distance to origin
func WaypointBoatL1(is instructions.Instructions) int {
	w := waypointBoat{
		waypoint: position.Position{X: 10, Y: 1},
	}
	for _, i := range is.Commands {
		w.operate(i)
	}
	return w.coordinates.L1()
}

// operate performs an instruction on the ship
func (w *waypointBoat) operate(inst instructions.Instruction) {
	switch inst.Action {
	case instructions.MoveNorth:
		w.waypoint.Move(0, inst.Value)
	case instructions.MoveEast:
		w.waypoint.Move(inst.Value, 0)
	case instructions.MoveSouth:
		w.waypoint.Move(0, -inst.Value)
	case instructions.MoveWest:
		w.waypoint.Move(-inst.Value, 0)
	case instructions.TurnPort:
		// turn in the trigonometric fashion
		w.waypoint.Rotate(cleanOrientation(inst.Value))
	case instructions.TurnStarboard:
		// turn in the weird clockwise direction
		w.waypoint.Rotate(cleanOrientation(-inst.Value))
	case instructions.Forward:
		w.coordinates.Move(w.waypoint.X*inst.Value, w.waypoint.Y*inst.Value)
	}
}
