package pocketdimension

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/pkg/coords"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tt := map[string]struct {
		neighbours     []coords.Point4DI
		neighbourCount int
	}{
		"3D points": {
			neighbours:     Neighbours3D,
			neighbourCount: 26,
		},
		"4D points": {
			neighbours:     Neighbours4D,
			neighbourCount: 80,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			p := New(tc.neighbours)
			assert.NotNil(t, p.cubes)
			assert.Equal(t, tc.neighbourCount, len(p.neighbours))
		})
	}
}

func TestPocketDimension_CountCubes(t *testing.T) {
	tt := map[string]struct {
		p     *PocketDimension
		count int
	}{
		"one dimension": {
			// .#.
			// ..#
			// ###
			p: &PocketDimension{
				cubes: map[coords.Point4DI]int{
					{1, 0, 0, 0}: 0,
					{2, 1, 0, 0}: 0,
					{0, 2, 0, 0}: 0,
					{1, 2, 0, 0}: 0,
					{2, 2, 0, 0}: 0,
				},
			},
			count: 5,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			c := tc.p.CountCubes()
			assert.Equal(t, tc.count, c)
		})
	}
}

func TestPocketDimension_ParseLine(t *testing.T) {
	tt := map[string]struct {
		p        *PocketDimension
		line     string
		cubes    map[coords.Point4DI]int
		nextLine int
	}{
		"one point": {
			p: &PocketDimension{
				cubes:       map[coords.Point4DI]int{{3, 0, 0, 0}: 0},
				currentLine: 1,
			},
			line:     "..#..",
			cubes:    map[coords.Point4DI]int{{3, 0, 0, 0}: 0, {2, 1, 0, 0}: 0},
			nextLine: 2,
		},
		"several points": {
			p: &PocketDimension{
				cubes:       map[coords.Point4DI]int{{3, 0, 0, 0}: 0},
				currentLine: 4,
			},
			line:     "#..##",
			cubes:    map[coords.Point4DI]int{{3, 0, 0, 0}: 0, {0, 4, 0, 0}: 0, {3, 4, 0, 0}: 0, {4, 4, 0, 0}: 0},
			nextLine: 5,
		},
		"no point": {
			p: &PocketDimension{
				cubes:       map[coords.Point4DI]int{{3, 0, 0, 0}: 0},
				currentLine: 3,
			},
			line:     ".......",
			cubes:    map[coords.Point4DI]int{{3, 0, 0, 0}: 0},
			nextLine: 4,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.p.ParseLine(tc.line)
			require.NoError(t, err)
			assert.Equal(t, tc.cubes, tc.p.cubes)
			assert.Equal(t, tc.nextLine, tc.p.currentLine)
		})
	}
}

func TestPocketDimension_Iterate(t *testing.T) {
	tt := map[string]struct {
		before PocketDimension
		count  int
		after  PocketDimension
	}{
		"3D, t=0 to t=1": {
			before: PocketDimension{
				cubes: map[coords.Point4DI]int{
					{1, 0, 0, 0}: 0,
					{2, 1, 0, 0}: 0,
					{0, 2, 0, 0}: 0,
					{1, 2, 0, 0}: 0,
					{2, 2, 0, 0}: 0,
				},
				neighbours: Neighbours3D,
			},
			count: 1,
			after: PocketDimension{
				cubes: map[coords.Point4DI]int{
					{0, 1, -1, 0}: 3,
					{2, 2, -1, 0}: 3,
					{1, 3, -1, 0}: 3,
					{0, 1, 0, 0}:  3,
					{2, 1, 0, 0}:  3,
					{1, 2, 0, 0}:  3,
					{2, 2, 0, 0}:  2,
					{1, 3, 0, 0}:  3,
					{0, 1, 1, 0}:  3,
					{2, 2, 1, 0}:  3,
					{1, 3, 1, 0}:  3,
				},
				neighbours: Neighbours3D,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			res := tc.before.Iterate(tc.count)
			assert.Equal(t, tc.after, res)
		})
	}
}
