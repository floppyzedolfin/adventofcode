package pocketdimension

import (
	"github.com/floppyzedolfin/adventofcode/pkg/coords"
)

// PocketDimension represents the space at a given time
type PocketDimension struct {
	cubes      map[coords.Point4DI]int
	neighbours []coords.Point4DI

	currentLine int // required for parsing only
}

// New returns a ready-to-use PocketDimension
func New(neighbours []coords.Point4DI) *PocketDimension {
	points := make(map[coords.Point4DI]int, 0)
	return &PocketDimension{cubes: points, neighbours: neighbours, currentLine: 0}
}

// ParseLine implements the LineParser interface
func (p *PocketDimension) ParseLine(line string) error {
	const cubeChar = '#'
	for i, c := range line {
		if c == cubeChar {
			// we've seen a cube; its X position is the position on the line
			// its Y position is the line number, and its Z and W positions are 0 (all input are flat)
			cube := coords.Point4DI{X: i, Y: p.currentLine, Z: 0, W: 0}
			p.cubes[cube] = 0 // this value doesn't really matter.
		}
	}
	p.currentLine++
	return nil
}

// CountCubes returns the number of cubes
func (p PocketDimension) CountCubes() int {
	return len(p.cubes)
}

// Iterate returns the state of the space after N iterations
func (p PocketDimension) Iterate(cycles int) PocketDimension {
	n := p
	for i := 0; i < cycles; i++ {
		n = *n.iterate()
	}
	return n
}

// iterate returns the state of the space after one iteration
func (p PocketDimension) iterate() *PocketDimension {
	// prepare the future PocketDimension
	next := New(p.neighbours)

	// propagate cubes from current to next version
	for cube := range p.cubes {
		next.visitNeighbours(cube)
	}

	// cleanup, using the parent to know which cubes were active
	next.sieve(p)
	return next
}

// sieve removes cubes with not enough / too many neighbours from the pocket dimension
func (p *PocketDimension) sieve(parent PocketDimension) {
	for cube, neighboursCount := range p.cubes {
		switch neighboursCount {
		case 3:
			// do nothing, we are happy. This cube is active.
		case 2:
			// check whether the parent had this cube activated
			if _, ok := parent.cubes[cube]; !ok {
				// 2 cubes is not enough to activate an inactive cube
				delete(p.cubes, cube)
			}
		default:
			// deactivate the cube.
			delete(p.cubes, cube)
		}
	}
}

// visitNeighbours populates all neighbours of a cube
func (p *PocketDimension) visitNeighbours(cube coords.Point4DI) {
	for _, n := range p.neighbours {
		neighbour := coords.Point4DI{
			X: cube.X + n.X,
			Y: cube.Y + n.Y,
			Z: cube.Z + n.Z,
			W: cube.W + n.W,
		}
		p.cubes[neighbour]++
	}
}
