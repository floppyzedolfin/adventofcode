package dec02

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/common"
	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestDec02Result_String(t *testing.T) {
	tt := map[string]struct {
		result dec02Result
		output string
	}{
		"part 1": {
			result: dec02Result{validPasswordsPart1: common.IntPointer(4)},
			output: "The number of valid passwords for Part 1 is 4.\n",
		},
		"part 2": {
			result: dec02Result{validPasswordsPart2: common.IntPointer(6)},
			output: "The number of valid passwords for Part 2 is 6.\n",
		},
		"parts 1 & 2": {
			result: dec02Result{validPasswordsPart1: common.IntPointer(4), validPasswordsPart2: common.IntPointer(6)},
			output: "The number of valid passwords for Part 1 is 4.\nThe number of valid passwords for Part 2 is 6.\n",
		},
		"empty result": {
			result: dec02Result{},
			output: "No job done by the elves today.",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.output, tc.result.String())
		})
	}
}

func TestDec02Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts door.Parts
		output dec02Result
		errMsg string
	} {
		"nominal Part1": {
			inputPath: "./input",
			parts: door.Parts{door.Part1},
			output: dec02Result{validPasswordsPart1: common.IntPointer(556)},
		},
		"nominal Part2": {
			inputPath: "./input",
			parts: door.Parts{door.Part2},
			output: dec02Result{validPasswordsPart2: common.IntPointer(605)},
		},
		"nominal Part1 & Part2": {
			inputPath: "./input",
			parts: door.Parts{door.Part1, door.Part2},
			output: dec02Result{validPasswordsPart1: common.IntPointer(556), validPasswordsPart2: common.IntPointer(605)},
		},
		"no parts": {
			inputPath: "./input",
			parts: door.Parts{},
			output: dec02Result{},
		},
		"invalid contents": {
			inputPath: "./test_data/invalid_contents",
			parts: door.Parts{door.Part1},
			errMsg: "unable to parse input file './test_data/invalid_contents'",
		},
		"dangerous range": {
			inputPath: "./test_data/dangerous_range",
			parts: door.Parts{door.Part1, door.Part2},
			output: dec02Result{validPasswordsPart1: common.IntPointer(0), validPasswordsPart2: common.IntPointer(0)},
		},
		"invalid range": {
			inputPath: "./test_data/invalid_range",
			parts: door.Parts{door.Part1, door.Part2},
			errMsg: "unable to parse input file './test_data/invalid_range'",
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
