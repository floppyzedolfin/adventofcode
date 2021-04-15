package waitingarea

import (
	"fmt"
	"strings"
)

// WaitingArea holds the state of the waiting area. It also holds some trickeries regarding the way it will be updated.
type WaitingArea struct {
	seats         [][]seat
	occupied      int
	closeRange    bool
	maxNeighbours int
	inRange       func(int, int) int
}

// OccupiedSeatsCloseRange counts the number of occupied seats when the evolution stalls, using the close range radius
func OccupiedSeatsCloseRange(wa *WaitingArea) int {
	// set close range parameters
	wa.closeRange = true
	wa.maxNeighbours = 4
	wa.inRange = wa.neighbours

	if err := wa.run(); err != nil {
		panic(err.Error())
	}
	return wa.occupied
}

// OccupiedSeatsLongRange counts the number of occupied seats when the evolution stalls, using the long range radius
func OccupiedSeatsLongRange(wa *WaitingArea) int {
	// set long range parameters
	wa.closeRange = false
	wa.maxNeighbours = 5
	wa.inRange = wa.inSight

	if err := wa.run(); err != nil {
		panic(err.Error())
	}
	return wa.occupied
}

// withinBounds lets us check whether a point is still inside the waiting area
func (wa *WaitingArea) withinBounds(x, y int) bool {
	// we'll assume all lines have the same width as the first one.
	return x >= 0 && y >= 0 && x < len(wa.seats[0]) && y < len(wa.seats)
}

// print is a debugging tool
func (wa *WaitingArea) print() {
	var sb strings.Builder
	for _, row := range wa.seats {
		for _, s := range row {
			sb.WriteString(string(s.status))
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}
