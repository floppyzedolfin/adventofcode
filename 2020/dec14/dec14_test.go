package dec14

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec14Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts     door.Parts
		output    dec14Result
		errMsg    string
	}{
		"example1 part1": {
			inputPath: "testdata/example1",
			parts:     door.Parts{door.Prima},
			output:    dec14Result{map[door.Part]*uint64{door.Prima: ptr.Uint64(165)}},
		},
		"nominal": {
			inputPath: "input",
			parts:     door.Parts{door.Prima, door.Secunda},
			output:    dec14Result{map[door.Part]*uint64{door.Prima: ptr.Uint64(17028179706934), door.Secunda: ptr.Uint64(3683236147222)}},
		},
		"example2 part2": {
			inputPath: "testdata/example2",
			parts:     door.Parts{door.Secunda},
			output:    dec14Result{map[door.Part]*uint64{door.Secunda: ptr.Uint64(208)}},
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
