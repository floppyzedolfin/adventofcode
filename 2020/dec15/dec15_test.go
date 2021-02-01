package dec15

import (
	"testing"

	"github.com/floppyzedolfin/adventofcode/door"
	"github.com/floppyzedolfin/adventofcode/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Most of the tests here are commented, because they take an awful lot of time.
func TestDec15Result_Solve(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		parts     door.Parts
		output    dec15Result
		errMsg    string
	}{
		// "0,3,6" : {
		// 	inputPath: "test_data/0-3-6",
		// 	parts:     door.Parts{door.Prima, door.Secunda},
		// 	output:    dec15Result{map[door.Part]*int{door.Prima: ptr.Int(436), door.Secunda: ptr.Int(175594)}},
		// },
		"input": {
			inputPath: "input",
			parts:     door.Parts{door.Prima, door.Secunda},
			output:    dec15Result{map[door.Part]*int{door.Prima: ptr.Int(468), door.Secunda: ptr.Int(1801753)}},
		},
		// "example 1": {
		// 	inputPath: "test_data/example1",
		// 	parts:     door.Parts{door.Prima, door.Secunda},
		// 	output:    dec15Result{map[door.Part]*int{door.Prima: ptr.Int(1), door.Secunda: ptr.Int(2578)}},
		// },
		// "example 2": {
		// 	inputPath: "test_data/example2",
		// 	parts:     door.Parts{door.Prima, door.Secunda},
		// 	output:    dec15Result{map[door.Part]*int{door.Prima: ptr.Int(10), door.Secunda: ptr.Int(3544142)}},
		// },
		// "example 3": {
		// 	inputPath: "test_data/example3",
		// 	parts:     door.Parts{door.Prima, door.Secunda},
		// 	output:    dec15Result{map[door.Part]*int{door.Prima: ptr.Int(27), door.Secunda: ptr.Int(261214)}},
		// },
		// "example 4": {
		// 	inputPath: "test_data/example4",
		// 	parts:     door.Parts{door.Prima, door.Secunda},
		// 	output:    dec15Result{map[door.Part]*int{door.Prima: ptr.Int(78), door.Secunda: ptr.Int(6895259)}},
		// },
		// "example 5": {
		// 	inputPath: "test_data/example5",
		// 	parts:     door.Parts{door.Prima, door.Secunda},
		// 	output:    dec15Result{map[door.Part]*int{door.Prima: ptr.Int(438), door.Secunda: ptr.Int(18)}},
		// },
		// "example 6": {
		// 	inputPath: "test_data/example6",
		// 	parts:     door.Parts{door.Prima, door.Secunda},
		// 	output:    dec15Result{map[door.Part]*int{door.Prima: ptr.Int(1836), door.Secunda: ptr.Int(362)}},
		// },
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
