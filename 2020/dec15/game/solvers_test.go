package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	tt := map[string]struct {
		inputFile string
		duration  int
		result    int
		errMsg    string
	}{
		"0-3-6, 4": {
			inputFile: "../testdata/0-3-6",
			duration:  4,
			result:    0,
		},
		"0-3-6, 5": {
			inputFile: "../testdata/0-3-6",
			duration:  5,
			result:    3,
		},
		"0-3-6, 6": {
			inputFile: "../testdata/0-3-6",
			duration:  6,
			result:    3,
		},
		"0-3-6, 7": {
			inputFile: "../testdata/0-3-6",
			duration:  7,
			result:    1,
		},
		"0-3-6, 8": {
			inputFile: "../testdata/0-3-6",
			duration:  8,
			result:    0,
		},
		"0-3-6, 9": {
			inputFile: "../testdata/0-3-6",
			duration:  9,
			result:    4,
		},
		"0-3-6, 10": {
			inputFile: "../testdata/0-3-6",
			duration:  10,
			result:    0,
		},
		"file doesn't exist": {
			inputFile: "file_doesnt_exist",
			errMsg:    "unable to parse file",
		},
		"file not eligible": {
			inputFile: "game.go",
			errMsg:    "unable to parse file",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			res, err := solve(tc.inputFile, tc.duration)
			if tc.errMsg != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				assert.Equal(t, tc.result, res)
			}
		})
	}
}
