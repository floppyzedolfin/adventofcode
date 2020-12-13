package dec10

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/door"
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
		"nominal Parts": {
			inputPath: "input",
			parts:     door.Parts{},
			output:    dec10Result{},
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
