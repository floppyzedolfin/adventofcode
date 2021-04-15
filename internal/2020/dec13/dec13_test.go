package dec13

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/pkg/door"
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
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
			inputPath: "testdata/example1",
			parts:     door.Parts{door.Prima},
			output:    dec13Result{map[door.Part]*int64{door.Prima: ptr.Int64(295)}},
		},
		"example1.2": {
			inputPath: "testdata/example1",
			parts:     door.Parts{door.Secunda},
			output:    dec13Result{map[door.Part]*int64{door.Secunda: ptr.Int64(1068781)}},
		},
		"example2": {
			inputPath: "testdata/example2",
			parts:     door.Parts{door.Secunda},
			output:    dec13Result{map[door.Part]*int64{door.Secunda: ptr.Int64(3417)}},
		},
		"example3": {
			inputPath: "testdata/example3",
			parts:     door.Parts{door.Secunda},
			output:    dec13Result{map[door.Part]*int64{door.Secunda: ptr.Int64(754018)}},
		},
		"example4": {
			inputPath: "testdata/example4",
			parts:     door.Parts{door.Secunda},
			output:    dec13Result{map[door.Part]*int64{door.Secunda: ptr.Int64(779210)}},
		},
		"example5": {
			inputPath: "testdata/example5",
			parts:     door.Parts{door.Secunda},
			output:    dec13Result{map[door.Part]*int64{door.Secunda: ptr.Int64(1261476)}},
		},
		"example6": {
			inputPath: "testdata/example6",
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
