package dec13

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec13Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts     door.Parts
		output    dec13Result
		errMsg    string
	}{
		"example1": {
			inputPath: "test_data/example1",
			parts:     door.Parts{door.Prima},
			output:    dec13Result{map[door.Part]*int64{door.Prima: ptr.Int64(295)}},
		},
		"example1.2": {
			inputPath: "test_data/example1",
			parts:     door.Parts{door.Secunda},
			output:    dec13Result{map[door.Part]*int64{door.Secunda: ptr.Int64(1068781)}},
		},
		"example2": {
			inputPath: "test_data/example2",
			parts:     door.Parts{door.Secunda},
			output:    dec13Result{map[door.Part]*int64{door.Secunda: ptr.Int64(3417)}},
		},
		"example3": {
			inputPath: "test_data/example3",
			parts:     door.Parts{door.Secunda},
			output:    dec13Result{map[door.Part]*int64{door.Secunda: ptr.Int64(754018)}},
		},
		"example4": {
			inputPath: "test_data/example4",
			parts:     door.Parts{door.Secunda},
			output:    dec13Result{map[door.Part]*int64{door.Secunda: ptr.Int64(779210)}},
		},
		"example5": {
			inputPath: "test_data/example5",
			parts:     door.Parts{door.Secunda},
			output:    dec13Result{map[door.Part]*int64{door.Secunda: ptr.Int64(1261476)}},
		},
		"example6": {
			inputPath: "test_data/example6",
			parts:     door.Parts{door.Secunda},
			output:    dec13Result{map[door.Part]*int64{door.Secunda: ptr.Int64(1202161486)}},
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
