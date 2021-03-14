package dec04

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec04Result_String(t *testing.T) {
	tt := map[string]struct {
		result dec04Result
		output string
	}{
		"part 1": {
			result: dec04Result{validPassports: map[door.Part]*int{door.Prima: ptr.Int(4)}},
			output: "The number of valid passports for Part 1 is 4.\n",
		},
		"part 2": {
			result: dec04Result{validPassports: map[door.Part]*int{door.Secunda: ptr.Int(6)}},
			output: "The number of valid passports for Part 2 is 6.\n",
		},
		"parts 1 & 2": {
			result: dec04Result{validPassports: map[door.Part]*int{door.Prima: ptr.Int(4), door.Secunda: ptr.Int(6)}},
			output: "The number of valid passports for Part 1 is 4.\nThe number of valid passports for Part 2 is 6.\n",
		},
		"empty result": {
			result: dec04Result{},
			output: "No job done by the elves today.",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.output, tc.result.String())
		})
	}
}

func TestDec04Result_Solve(t *testing.T) {
	tt := map[string] struct {
		inputPath string
		parts door.Parts
		output dec04Result
		errMsg string
	} {
		"nominal Part 1": {
			inputPath: "./input",
			parts: door.Parts{door.Prima},
			output: dec04Result{ map[door.Part]*int{door.Prima: ptr.Int(182)}},
		},
		"nominal Part 2": {
			inputPath: "./input",
			parts: door.Parts{door.Secunda},
			output: dec04Result{ map[door.Part]*int{door.Secunda: ptr.Int(109)}},
		},
		"nominal Parts 1 & 2": {
			inputPath: "./input",
			parts: door.Parts{door.Prima, door.Secunda},
			output: dec04Result{ map[door.Part]*int{door.Prima: ptr.Int(182), door.Secunda: ptr.Int(109)}},
		},
		"nominal no parts": {
			inputPath: "./input",
			parts: door.Parts{},
			output: dec04Result{map[door.Part]*int{}},
		},
		"example 1" : {
			inputPath: "./testdata/example1",
			parts: door.Parts{door.Prima},
			output: dec04Result{map[door.Part]*int{door.Prima: ptr.Int(2)}},
		},
		"example 1 with an empty line at the end" : {
			inputPath: "./testdata/example1_finalempty",
			parts: door.Parts{door.Prima},
			output: dec04Result{map[door.Part]*int{door.Prima: ptr.Int(2)}},
		},
		"example 2" : {
			inputPath: "./testdata/example2",
			parts: door.Parts{door.Secunda},
			output: dec04Result{map[door.Part]*int{door.Secunda: ptr.Int(0)}},
		},
		"example 3" : {
			inputPath: "./testdata/example3",
			parts: door.Parts{door.Secunda},
			output: dec04Result{map[door.Part]*int{door.Secunda: ptr.Int(4)}},
		},

	}

	for name, tc := range tt {
		t.Run(name, func (t *testing.T){
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
