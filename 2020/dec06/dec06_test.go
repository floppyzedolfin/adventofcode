package dec06

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/common"
	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDec06Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts     door.Parts
		output    dec06Result
		errMsg    string
	}{
		"nominal Parts 1&2": {
			inputPath: "input",
			parts:     door.Parts{door.Secunda, door.Prima},
			output:    dec06Result{map[door.Part]*int{door.Prima: common.IntPointer(6630), door.Secunda: common.IntPointer(3437)}},
		},
		"example 1": {
			inputPath: "./test_data/example1",
			parts:     door.Parts{door.Prima},
			output:    dec06Result{map[door.Part]*int{door.Prima: common.IntPointer(11)}},
		},
		"example 2": {
			inputPath: "./test_data/example2",
			parts:     door.Parts{door.Secunda},
			output:    dec06Result{map[door.Part]*int{door.Secunda: common.IntPointer(6)}},
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
