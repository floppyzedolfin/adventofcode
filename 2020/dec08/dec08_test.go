package dec08

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec08Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts     door.Parts
		output    dec08Result
		errMsg    string
	}{
		"nominal Parts": {
			inputPath: "input",
			parts:     door.Parts{door.Prima, door.Secunda},
			output:    dec08Result{acc: map[door.Part]*int{door.Prima: ptr.Int(1600), door.Secunda: ptr.Int(1543)}},
		},
		"example 1": {
			inputPath: "./test_data/example1",
			parts:     door.Parts{door.Prima, door.Secunda},
			output:    dec08Result{acc: map[door.Part]*int{door.Prima: ptr.Int(5), door.Secunda: ptr.Int(8)}},
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
