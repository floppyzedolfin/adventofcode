package notes

import (
	"math"
	"math/big"

	"github.com/floppyzedolfin/adventofcode/2020/dec13/bezout"
)

// FirstBus returns the number of the first bus multiplied by the wait before its departure
func FirstBus(n Notes) int64 {
	nextDeparture := math.MaxInt32
	nextBusNumber := math.MaxInt32
	// find the first bus to leave - or at least one of the firsts.
	for _, b := range n.buses {
		departureIn := b.number - n.currentTime%b.number
		if departureIn < nextDeparture {
			nextDeparture = departureIn
			nextBusNumber = b.number
		}
	}
	return int64(nextDeparture * nextBusNumber)
}

// ConsecutiveBuses returns the time of departure of the first bus to match all the conditions.
func ConsecutiveBuses(n Notes) int64 {
	// this is an implementation of the Chinese Remainder Theorem.
	// First, we need to pair the numbers with the remainders.

	// congruences lists all the equations the number X we look for verifies
	congruences := make([]bezout.BBPair, len(n.buses))
	// fill the equations - it is mandatory that the bus numbers are prime, otherwise we can't apply the theorem
	for i, b := range n.buses {
		congruences[i].Divisor = *big.NewInt(int64(b.number))
		congruences[i].Remainder = *big.NewInt(int64(b.position))
	}

	res := bezout.Solve(congruences)
	return res.Int64()
}
