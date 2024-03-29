package dec01

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec01Result_String(t *testing.T) {
	tt := map[string]struct {
		result dec01Result
		output string
	}{
		"part 1": {
			result: dec01Result{productPrima: ptr.Int(4)},
			output: "The product for Part 1 is 4.\n",
		},
		"part 2": {
			result: dec01Result{productSecunda: ptr.Int(6)},
			output: "The product for Part 2 is 6.\n",
		},
		"parts 1 & 2": {
			result: dec01Result{productPrima: ptr.Int(4), productSecunda: ptr.Int(6)},
			output: "The product for Part 1 is 4.\nThe product for Part 2 is 6.\n",
		},
		"empty result": {
			result: dec01Result{},
			output: "No job done by the elves today.",
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
		parts     door.Parts
		output    dec01Result
		errMsg    string
	}{
		"nominal Prima": {
			inputPath: "./input",
			parts:     door.Parts{door.Prima},
			output:    dec01Result{productPrima: ptr.Int(381699)},
		},
		"nominal Seconda": {
			inputPath: "./input",
			parts:     door.Parts{door.Secunda},
			output:    dec01Result{productSecunda: ptr.Int(111605670)},
		},
		"nominal Prima and Seconda": {
			inputPath: "./input",
			parts:     door.Parts{door.Prima, door.Secunda},
			output:    dec01Result{productPrima: ptr.Int(381699), productSecunda: ptr.Int(111605670)},
		},
		"no parts - no job for the elves": {
			inputPath: "./input",
			output:    dec01Result{},
		},
		"no input file": {
			inputPath: "./testdata/file_doesnt_exist",
			parts:     door.Parts{door.Prima},
			errMsg:    "unable to read lines",
		},
		"aberrant input file": {
			inputPath: "./testdata/aberrant_data",
			parts:     door.Parts{door.Prima},
			errMsg:    "unable to read lines",
		},
		"no match found": {
			inputPath: "./testdata/no_match",
			parts:     door.Parts{door.Prima},
			errMsg:    "unable to find match",
		},
		"1010 once": {
			inputPath: "./testdata/1010",
			parts:     door.Parts{door.Prima},
			errMsg:    "unable to find match",
		},
		"one liner": {
			inputPath: "./testdata/oneliner",
			parts:     door.Parts{door.Prima},
			errMsg:    "unable to find match",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := New(tc.inputPath)
			res, err := s.Solve(tc.parts)
			if tc.errMsg != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.output, res)
			}
		})
	}
}
