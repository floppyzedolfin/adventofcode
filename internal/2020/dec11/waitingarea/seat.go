package waitingarea

type seatStatus = rune

const (
	seatEmptyChar    = 76 // 'L'
	seatOccupiedChar = 35 // '#'
	floorChar        = 46 // '.'
)

type seat struct {
	status seatStatus
	posX   int
	posY   int
}

// checkSeat returns the new seat if the seat at pos x,y needs to be updated
func (wa *WaitingArea) checkSeat(x, y int) *seat {
	switch wa.seats[y][x].status {
	case seatEmptyChar:
		// check whether someone can sit here
		if wa.inRange(x, y) == 0 {
			wa.occupied++
			return &seat{seatOccupiedChar, x, y}
		}
	case seatOccupiedChar:
		// are there too many passengers around
		if wa.inRange(x, y) >= wa.maxNeighbours {
			wa.occupied--
			return &seat{seatEmptyChar, x, y}
		}
	}
	// the floor is lava, don't sit there
	return nil
}

// seatStatusAt returns the status of a seat at pos x,y. The outer world is made of floor.
func (wa *WaitingArea) seatStatusAt(x, y int) seatStatus {
	if !wa.withinBounds(x, y) {
		// by convention, ground exists around somewhere.
		return floorChar
	}
	return wa.seats[y][x].status
}
