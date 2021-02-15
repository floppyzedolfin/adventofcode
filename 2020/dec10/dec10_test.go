package dec10

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec10Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts     door.Parts
		output    dec10Result
		errMsg    string
	}{
		"example1": {
			inputPath: "./testdata/example1",
			parts:     door.Parts{door.Prima},
			output:    dec10Result{map[door.Part]*int64{door.Prima: ptr.Int64(35)}},
		},
		"example2": {
			inputPath: "./testdata/example2",
			parts:     door.Parts{door.Prima, door.Secunda},
			output:    dec10Result{map[door.Part]*int64{door.Prima: ptr.Int64(220), door.Secunda: ptr.Int64(19208)}},
		},
		"nominal dataset": {
			inputPath: "input",
			parts:     door.Parts{door.Prima, door.Secunda},
			output:    dec10Result{map[door.Part]*int64{door.Prima: ptr.Int64(2070), door.Secunda: ptr.Int64(24179327893504)}},
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
