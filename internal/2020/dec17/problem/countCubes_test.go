package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCountCubes3D(t *testing.T) {
	tt := map[string]struct {
		inputFile string
		cycles    int
		cubes     int
	}{
		"0 cycle": {
			inputFile: "../testdata/example1",
			cycles:    0,
			cubes:     5,
		},
		"1 cycle": {
			inputFile: "../testdata/example1",
			cycles:    1,
			cubes:     11,
		},
		"2 cycles": {
			inputFile: "../testdata/example1",
			cycles:    2,
			cubes:     21,
		},
		"6 cycles": {
			inputFile: "../testdata/example1",
			cycles:    6,
			cubes:     112,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			cubeCount, err := CountCubes3D(tc.inputFile, tc.cycles)
			require.NoError(t, err)
			assert.Equal(t, tc.cubes, cubeCount)
		})
	}
}

func TestCountCubes4D(t *testing.T) {
	tt := map[string]struct {
		inputFile string
		cycles    int
		cubes     int
	}{
		"0 cycle": {
			inputFile: "../testdata/example1",
			cycles:    0,
			cubes:     5,
		},
		"1 cycle": {
			inputFile: "../testdata/example1",
			cycles:    1,
			cubes:     29,
		},
		"6 cycle": {
			inputFile: "../testdata/example1",
			cycles:    6,
			cubes:     848,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			cubeCount, err := CountCubes4D(tc.inputFile, tc.cycles)
			require.NoError(t, err)
			assert.Equal(t, tc.cubes, cubeCount)
		})
	}
}
