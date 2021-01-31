package game

import (
	"fmt"

	"github.com/floppyzedolfin/adventofcode/fileparser"
)

// Solve2020 returns the 2020th number spoken
func Solve2020(inputFile string) (int, error) {
	return solve(inputFile, 2020)
}

// Solve30000000 returns the 3e7th number spoken
func Solve30000000(inputFile string) (int, error) {
	return solve(inputFile, 30000000)
}

// solve a problem given a game duration
func solve(inputFile string, gameDuration int) (int, error) {
	g := newGame()
	err := fileparser.ParseFile(inputFile, g)
	if err != nil {
		return 0, fmt.Errorf("unable to parse file '%s': %s", inputFile, err.Error())
	}

	for g.clock < gameDuration {
		// feed the last result to g
		g.turn(g.lastResult)
	}

	return g.lastResult, nil
}
