package waitingarea

// vector represents... a vector.
type vector struct {
	dx int
	dy int
}

// inSight counts the number of occupied seats one can see in all 8 directions
func (wa *WaitingArea) inSight(x, y int) int {
	peopleSeen := 0
	for _, d := range directions {
		peopleSeen += wa.countInDirection(x, y, d)
	}
	return peopleSeen
}

// countInDirection returns the number of people we see when looking in a direction, stopping at either a seat or a wall
func (wa *WaitingArea) countInDirection(x, y int, d vector) int {
	ix, iy := x+d.dx, y+d.dy
	// loop as long as we remain within the area
	for wa.withinBounds(ix, iy) {
		switch wa.seatStatusAt(ix, iy) {
		case seatOccupiedChar:
			// there's someone there.
			return 1
		case seatEmptyChar:
			// we see an empty seat, that's enough for our own Lebensraum
			return 0
		}
		// keep looking over floor tiles
		ix += d.dx
		iy += d.dy
	}
	// we've reached the edge
	return 0
}
