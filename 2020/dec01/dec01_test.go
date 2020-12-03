package dec01

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/common"
	"github.com/floppyzedolfin/adventofcode/door"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec01Result_String(t *testing.T) {
	tt := map[string]struct {
		result dec01Result
		output string
	}{
		"part 1": {
			result: dec01Result{productPart1: common.IntPointer(4)},
			output: "The product for Part 1 is 4.\n",
		},
		"part 2": {
			result: dec01Result{productPart2: common.IntPointer(6)},
			output: "The product for Part 2 is 6.\n",
		},
		"parts 1 & 2": {
			result: dec01Result{productPart1: common.IntPointer(4), productPart2: common.IntPointer(6)},
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
		"nominal Part1": {
			inputPath: "./input",
			parts:     door.Parts{door.Part1},
			output:    dec01Result{productPart1: common.IntPointer(381699)},
		},
		"nominal Part2": {
			inputPath: "./input",
			parts:     door.Parts{door.Part2},
			output:    dec01Result{productPart2: common.IntPointer(111605670)},
		},
		"nominal Part1 and Part2": {
			inputPath: "./input",
			parts:     door.Parts{door.Part1, door.Part2},
			output:    dec01Result{productPart1: common.IntPointer(381699), productPart2: common.IntPointer(111605670)},
		},
		"no parts - no job for the elves": {
			inputPath: "./input",
			output:    dec01Result{},
		},
		"no input file": {
			inputPath: "./test_data/file_doesnt_exist",
			parts:     door.Parts{door.Part1},
			errMsg:    "unable to read lines",
		},
		"aberrant input file": {
			inputPath: "./test_data/aberrant_data",
			parts:     door.Parts{door.Part1},
			errMsg:    "unable to read lines",
		},
		"no match found": {
			inputPath: "./test_data/no_match",
			parts:     door.Parts{door.Part1},
			errMsg:    "unable to find match",
		},
		"1010 once": {
			inputPath: "./test_data/1010",
			parts:     door.Parts{door.Part1},
			errMsg:    "unable to find match",
		},
		"one liner": {
			inputPath: "./test_data/oneliner",
			parts:     door.Parts{door.Part1},
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
