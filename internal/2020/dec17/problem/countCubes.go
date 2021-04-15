package problem

import (
	"fmt"

	"github.com/floppyzedolfin/adventofcode/internal/2020/dec17/pocketdimension"
	"github.com/floppyzedolfin/adventofcode/pkg/fileparser"
)

// CountCubes3D returns the number of cubes in a 3D environment
func CountCubes3D(inputPath string, iterations int) (int, error) {
	p := pocketdimension.New(pocketdimension.Neighbours3D)
	err := fileparser.ParseFile(inputPath, p)
	if err != nil {
		return 0, fmt.Errorf("unable to parse file: %w", err)
	}

	return p.Iterate(iterations).CountCubes(), nil
}

// CountCubes4D returns the number of cubes in a 4D environment
func CountCubes4D(inputPath string, iterations int) (int, error) {
	p := pocketdimension.New(pocketdimension.Neighbours4D)
	err := fileparser.ParseFile(inputPath, p)
	if err != nil {
		return 0, fmt.Errorf("unable to parse file: %w", err)
	}

	return p.Iterate(iterations).CountCubes(), nil
}
