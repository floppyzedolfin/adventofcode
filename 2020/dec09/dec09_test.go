package dec09

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec09Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		preambleLength int
		parts     door.Parts
		output    dec09Result
		errMsg    string
	}{
		"nominal Parts": {
			inputPath: "input",
			preambleLength: 25,
			parts:     door.Parts{door.Prima},
			output:    dec09Result{invalidNumber: map[door.Part]*int{door.Prima: ptr.Int(18272118)}},
		},
		"example 1": {
			inputPath: "./test_data/example1",
			preambleLength: 5,
			parts:     door.Parts{door.Prima, door.Secunda},
			output:    dec09Result{invalidNumber: map[door.Part]*int{door.Prima: ptr.Int(127), door.Secunda: ptr.Int(62)}},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			preambleLength = tc.preambleLength
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
