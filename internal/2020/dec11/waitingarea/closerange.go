package waitingarea

// neighbours scans the area around pos (x,y) and returns the number of neighbours
func (wa *WaitingArea) neighbours(x, y int) int {
	neighbours := 0
	for _, d := range directions {
		if wa.seatStatusAt(x+d.dx, y+d.dy) == seatOccupiedChar {
			neighbours++
		}
	}
	return neighbours
}
