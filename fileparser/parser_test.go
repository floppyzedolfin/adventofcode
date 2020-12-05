package fileparser

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	tt := map[string]struct {
		inputPath string
		output    integers
		errMsg    string
	}{
		"norminal": {
			inputPath: "./test_data/integers",
			output:    integers{[]int{1, 2, 4}},
		},
		"failing": {
			inputPath: "./test_data/bad_integers",
			errMsg:    "unable to parse line, expected an int, read 'foo'",
		},
		"no such file": {
			inputPath: "./test_data/inexisting_file",
			errMsg:    "unable to read file './test_data/inexisting_file'",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			var i integers
			err := ParseFile(tc.inputPath, &i)
			if tc.errMsg != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.output, i)
			}
		})
	}
}

// this struct is only here as an example and for testing purposes
type integers struct {
	ints []int
}

// ParseLine implements the LineParser interface
func (i *integers) ParseLine(line string) error {
	value, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("unable to parse line, expected an int, read '%s'", line)
	}
	// store value
	i.ints = append(i.ints, value)
	return nil
}
