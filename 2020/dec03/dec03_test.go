package dec03

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec03Result_String(t *testing.T) {
	tt := map[string]struct {
		result dec03Result
		output string
	}{
		"part 1": {
			result: dec03Result{treeCountPrima: ptr.Int(4)},
			output: "The number of trees hit for Part 1 is 4.\n",
		},
		"part 2": {
			result: dec03Result{treeCountSecunda: ptr.Int(6)},
			output: "The number of trees hit for Part 2 is 6.\n",
		},
		"parts 1 & 2": {
			result: dec03Result{treeCountPrima: ptr.Int(4), treeCountSecunda: ptr.Int(6)},
			output: "The number of trees hit for Part 1 is 4.\nThe number of trees hit for Part 2 is 6.\n",
		},
		"empty result": {
			result: dec03Result{},
			output: "No job done by the elves today.",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.output, tc.result.String())
		})
	}
}

func TestDec03Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts     door.Parts
		output    dec03Result
		errMsg    string
	}{
		"nominal Prima": {
			inputPath: "./input",
			parts:     door.Parts{door.Prima},
			output:    dec03Result{treeCountPrima: ptr.Int(187)},
		},
		"nominal Secunda": {
			inputPath: "./input",
			parts:     door.Parts{door.Secunda},
			output:    dec03Result{treeCountSecunda: ptr.Int(4723283400)},
		},
		"nominal Prima & Secunda": {
			inputPath: "./input",
			parts:     door.Parts{door.Prima, door.Secunda},
			output:    dec03Result{treeCountPrima: ptr.Int(187), treeCountSecunda: ptr.Int(4723283400)},
		},
		"no parts": {
			inputPath: "./input",
			parts:     door.Parts{},
			output:    dec03Result{},
		},
		"invalid contents": {
			// if the input is garbage, we simply don't find trees
			inputPath: "./test_data/invalid_contents",
			parts:     door.Parts{door.Prima},
			output:    dec03Result{treeCountPrima: ptr.Int(0)},
		},
		"no track x": {
			inputPath: "./test_data/no_track_x",
			parts:     door.Parts{door.Prima},
			output:    dec03Result{treeCountPrima: ptr.Int(0)},
		},
		"no track y": {
			inputPath: "./test_data/no_track_y",
			parts:     door.Parts{door.Secunda},
			output:    dec03Result{treeCountSecunda: ptr.Int(0)},
		},
		"missing input": {
			inputPath: "./test_data/missing_file",
			parts:     door.Parts{door.Prima},
			errMsg:    "unable to parse input file './test_data/missing_file'",
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
