package game

import (
	"fmt"
	"strconv"
	"strings"
)

// game holds the state of the game.
type game struct {
	// previouslySeen holds the position we've last seen each int
	previouslySeen map[int]int

	// clock is the time-tracker
	clock int

	// lastResult holds the last returned result. It's used as a work-around to solve "dealing with input" vs "playing the game"
	lastResult int
}

// newGame returns a properly initialised game
func newGame() *game {
	// time starts at 1, I guess it's some 0 AD joke the elves have
	return &game{clock: 1, previouslySeen: make(map[int]int)}
}

// turn plays a turn - we receive a number and return one depending on that number
func (g *game) turn(entry int) {
	// increase clock
	g.clock++

	// check if we've already seen this one
	if g.previouslySeen[entry] == 0 {
		// it's new - set its last-time seen to now
		g.previouslySeen[entry] = g.clock
	}

	// compute the number of turns apart since last seen
	g.lastResult = g.clock - g.previouslySeen[entry]

	// override the latest time we've seen this value
	g.previouslySeen[entry] = g.clock
}

// ParseLine implements the LineParser interface
func (g *game) ParseLine(line string) error {
	const separator = ","
	ints := strings.Split(line, separator)
	for _, i := range ints {
		entry, err := strconv.Atoi(i)
		if err != nil {
			return fmt.Errorf("unable to parse int '%s': %s", i, err.Error())
		}
		// we don't care about the last result here, as we're only reading input values
		g.turn(entry)
	}
	return nil
}
