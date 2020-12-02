package dec01

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec01Result_String(t *testing.T) {
	tt := map[string]struct {
		result dec01Result
		output string
	}{
		"part 1": {
			result: dec01Result{productPart1: 4},
			output: "The product for Part 1 is 4.\n",
		},
		"part 2": {
			result: dec01Result{productPart2: 6},
			output: "The product for Part 2 is 6.\n",
		},
		"parts 1 & 2": {
			result: dec01Result{productPart1: 4, productPart2: 6},
			output: "The product for Part 1 is 4.\nThe product for Part 2 is 6.\n",
		},
		"empty result": {
			result: dec01Result{},
			output: "No match found.",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.output, tc.result.String())
		})
	}
}

func TestDec01Solver_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		output    string
		errMsg    string
	}{
		"nominal": {
			inputPath: "./input",
			output:    "The product for Part 2 is 111605670.\n",
		},
		"no input file": {
			inputPath: "./test_data/file_doesnt_exist",
			errMsg:    "unable to read lines",
		},
		"aberrant input file": {
			inputPath: "./test_data/aberrant_data",
			errMsg:    "unable to read lines",
		},
		"no match found": {
			inputPath: "./test_data/no_match",
			errMsg:    "unable to find match",
		},
		"1010 once": {
			inputPath: "./test_data/1010",
			errMsg:    "unable to find match",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := New(tc.inputPath)
			res, err := s.Solve()
			if tc.errMsg != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.output, res.String())
			}
		})
	}
}
