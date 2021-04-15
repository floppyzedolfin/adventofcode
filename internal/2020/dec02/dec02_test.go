package dec02

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestDec02Result_String(t *testing.T) {
	tt := map[string]struct {
		result dec02Result
		output string
	}{
		"part 1": {
			result: dec02Result{validPasswordsPrima: ptr.Int(4)},
			output: "The number of valid passwords for Part 1 is 4.\n",
		},
		"part 2": {
			result: dec02Result{validPasswordsSecunda: ptr.Int(6)},
			output: "The number of valid passwords for Part 2 is 6.\n",
		},
		"parts 1 & 2": {
			result: dec02Result{validPasswordsPrima: ptr.Int(4), validPasswordsSecunda: ptr.Int(6)},
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
		"nominal Prima": {
			inputPath: "./input",
			parts: door.Parts{door.Prima},
			output: dec02Result{validPasswordsPrima: ptr.Int(556)},
		},
		"nominal Secunda": {
			inputPath: "./input",
			parts: door.Parts{door.Secunda},
			output: dec02Result{validPasswordsSecunda: ptr.Int(605)},
		},
		"nominal Prima & Secunda": {
			inputPath: "./input",
			parts: door.Parts{door.Prima, door.Secunda},
			output: dec02Result{validPasswordsPrima: ptr.Int(556), validPasswordsSecunda: ptr.Int(605)},
		},
		"no parts": {
			inputPath: "./input",
			parts: door.Parts{},
			output: dec02Result{},
		},
		"invalid contents": {
			inputPath: "./testdata/invalid_contents",
			parts: door.Parts{door.Prima},
			errMsg: "unable to parse input file './testdata/invalid_contents'",
		},
		"dangerous range": {
			inputPath: "./testdata/dangerous_range",
			parts: door.Parts{door.Prima, door.Secunda},
			output: dec02Result{validPasswordsPrima: ptr.Int(0), validPasswordsSecunda: ptr.Int(0)},
		},
		"invalid range": {
			inputPath: "./testdata/invalid_range",
			parts: door.Parts{door.Prima, door.Secunda},
			errMsg: "unable to parse input file './testdata/invalid_range'",
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
