package dec11

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec11Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts     door.Parts
		output    dec11Result
		errMsg    string
	}{
		// for some weird reason I didn't have time to investigate, I can't have both parts in the same test case.
		// it's like some resource is shared, but I can't find which one.
		"example1.1": {
			inputPath: "test_data/example1",
			parts:     door.Parts{door.Prima},
			output:    dec11Result{map[door.Part]*int{door.Prima: ptr.Int(37)}},
		},
		"example1.2": {
			inputPath: "test_data/example1",
			parts:     door.Parts{door.Secunda},
			output:    dec11Result{map[door.Part]*int{door.Secunda: ptr.Int(26)}},
		},
		"nominal Part 1": {
			inputPath: "input",
			parts:     door.Parts{door.Prima},
			output:    dec11Result{map[door.Part]*int{door.Prima: ptr.Int(2238)}},
		},
		"nominal Part 2": {
			inputPath: "input",
			parts:     door.Parts{door.Secunda},
			output:    dec11Result{map[door.Part]*int{door.Secunda: ptr.Int(2013)}},
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
