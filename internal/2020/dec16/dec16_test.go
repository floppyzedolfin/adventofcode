package dec16

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec16Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts     door.Parts
		output    dec16Result
		errMsg    string
	}{
		"example Part1": {
			inputPath: "testdata/example1",
			parts:     door.Parts{door.Prima},
			output:    dec16Result{map[door.Part]*uint64{door.Prima: ptr.Uint64(71)}},
		},
		"example Part2": {
			inputPath: "testdata/example2",
			parts:     door.Parts{door.Secunda},
			output:    dec16Result{map[door.Part]*uint64{door.Secunda: ptr.Uint64(132)}},
		},
		"input": {
			inputPath: "input",
			parts:     door.Parts{door.Prima, door.Secunda},
			output:    dec16Result{map[door.Part]*uint64{door.Prima: ptr.Uint64(20048), door.Secunda: ptr.Uint64(4810284647569)}},
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
