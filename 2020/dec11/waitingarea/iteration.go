package waitingarea

import (
	"fmt"
)

// run is the main loop, it will iterate as long as the waiting area's seating positions change
func (wa *WaitingArea) run() error {
	// this is a safety net. we don't want to iterate forever.
	const maxIterations = 1000000
	var i = 0
	for ; i < maxIterations; i++ {
		if !wa.iterate() {
			break
		}
	}
	if i == maxIterations {
		return fmt.Errorf("not enough iterations")
	}
	return nil
}

// iterate runs one round and returns whether the waiting area has evolved
func (wa *WaitingArea) iterate() bool {
	// this is performed in two steps:
	// first, checks which seats need be updated
	seatsToUpdate := wa.checkSeats()

	if len(seatsToUpdate) == 0 {
		// we're stalling
		return false
	}
	// second, update the seats accordingly
	wa.update(seatsToUpdate)

	return true
}

// checkStatus looks at the seats, and returns the seats to be updated
func (wa *WaitingArea) checkSeats() []seat {
	updatedSeats := make([]seat, 0)
	for y, row := range wa.seats {
		for x := range row {
			if s := wa.checkSeat(x, y); s != nil {
				updatedSeats = append(updatedSeats, *s)
			}
		}
	}
	return updatedSeats
}

// update sets a new status to each seat that needs correction
func (wa *WaitingArea) update(seats []seat) {
	for _, s := range seats {
		wa.seats[s.posY][s.posX].status = s.status
	}
}
