package dec05

import (
	"sort"
)

const (
	// logPlaneCapacity describes the number of pieces of information we need to locate a seat
	// we've got 7 rows and 3 columns, for a total of 10 binary decisions
	logPlaneCapacity = 10
)

func findHighestPresent(occupiedSeats []int) int {
	sort.Ints(occupiedSeats)
	return occupiedSeats[len(occupiedSeats)-1]
}

func findLowestMissing(occupiedSeats []int) int {
	freeSeats := availableSeats(occupiedSeats)
	// the first empty one is mine !
	if freeSeats[0] == 0 {
		// we need to have one person on each side
		return freeSeats[1]
	} else {
		return freeSeats[0]
	}
}

// this func is overly complex. 4 loops is way too many. It's still O(n) though.
func availableSeats(occupiedSeats []int) []int {
	// let's build a complete aeroplane -- it's not THAT difficult, as long as you don't expect it to fly.
	aeroplane := map[int]struct{}{}
	// add seats to our aeroplane
	planeCapacity :=  1 << logPlaneCapacity
	for i := 0; i < planeCapacity; i++ {
		aeroplane[i] = struct{}{}
	}

	minOccupied := occupiedSeats[0]
	maxOccupied := minOccupied

	// We now have an aeroplane full of empty seats. Let's discard used seats
	for _, v := range occupiedSeats {
		if v < minOccupied {
			minOccupied = v
		}
		if v > maxOccupied {
			maxOccupied = v
		}
		delete(aeroplane, v)
	}

	// however, some seats are missing on this plane, fore and aft.
	for i := 0; i < minOccupied; i++ {
		delete(aeroplane, i)
	}
	for i := maxOccupied + 1; i < planeCapacity; i++ {
		delete(aeroplane, i)
	}

	// retrieve unused seats
	unusedSeats := make([]int, 0, len(aeroplane))
	for k := range aeroplane {
		unusedSeats = append(unusedSeats, k)
	}
	sort.Ints(unusedSeats)

	return unusedSeats
}
