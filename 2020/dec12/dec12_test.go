package dec12

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec12Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts     door.Parts
		output    dec12Result
		errMsg    string
	}{
		"example1": {
			inputPath: "test_data/example1",
			parts:     door.Parts{door.Prima},
			output:    dec12Result{map[door.Part]*int{door.Prima: ptr.Int(25)}},
		},
		"example1.2": {
			inputPath: "test_data/example1",
			parts:     door.Parts{door.Secunda},
			output:    dec12Result{map[door.Part]*int{door.Secunda: ptr.Int(286)}},
		},
		"nominal": {
			inputPath: "input",
			parts:     door.Parts{door.Prima, door.Secunda},
			output:    dec12Result{map[door.Part]*int{door.Prima: ptr.Int(2847), door.Secunda: ptr.Int(29839)}},
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
