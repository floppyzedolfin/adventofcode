package door

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts_Contains(t *testing.T) {
	tt := map[string]struct {
		parts  Parts
		value  int
		result bool
	}{
		"value is in parts": {
			parts:  Parts{1, 2},
			value:  2,
			result: true,
		},
		"value is not in parts" : {
		parts: Parts{1},
		value: 2,
		result: false,
	},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T){
			res := tc.parts.Contains(tc.value)
			assert.Equal(t, tc.result, res)
		})
	}
}
