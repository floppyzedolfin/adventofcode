package dec05

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec05Result_String(t *testing.T) {
	tt := map[string]struct {
		result dec05Result
		output string
	}{
		"part 1": {
			result: dec05Result{availableSeats: map[door.Part]*int{door.Prima: ptr.Int(4)}},
			output: "The highest available int for Part 1 is 4.\n",
		},
		"part 2": {
			result: dec05Result{availableSeats: map[door.Part]*int{door.Secunda: ptr.Int(6)}},
			output: "The highest available int for Part 2 is 6.\n",
		},
		"parts 1 & 2": {
			result: dec05Result{availableSeats: map[door.Part]*int{door.Prima: ptr.Int(4), door.Secunda: ptr.Int(6)}},
			output: "The highest available int for Part 1 is 4.\nThe highest available int for Part 2 is 6.\n",
		},
		"empty result": {
			result: dec05Result{},
			output: "No job done by the elves today.",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.output, tc.result.String())
		})
	}
}

func TestDec05Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts     door.Parts
		output    dec05Result
		errMsg    string
	}{
		"nominal Parts 1&2": {
			inputPath: "input",
			parts:     door.Parts{door.Secunda, door.Prima},
			output:    dec05Result{map[door.Part]*int{door.Prima: ptr.Int(828), door.Secunda: ptr.Int(565)}},
		},
		"example 1": {
			inputPath: "./testdata/example1",
			parts:     door.Parts{door.Prima},
			output:    dec05Result{map[door.Part]*int{door.Prima: ptr.Int(357)}},
		},
		"example 1 with an empty line at the end": {
			inputPath: "./testdata/example1_finalempty",
			parts:     door.Parts{door.Prima},
			output:    dec05Result{map[door.Part]*int{door.Prima: ptr.Int(357)}},
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
