package bezout

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtendedEuclideanAlgorithm(t *testing.T) {
	tt := map[string]struct {
		a big.Int
		b big.Int
		x big.Int
		y big.Int
	}{
		"5x5 + 12*(-2) = 1":             {*big.NewInt(5), *big.NewInt(12), *big.NewInt(5), *big.NewInt(-2)},
		"1723*(-2014) + 4021 * 863 = 1": {*big.NewInt(2014), *big.NewInt(4021), *big.NewInt(-1723), *big.NewInt(863)},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			x, y := extendedEuclideanAlgorithm(tc.a, tc.b)
			assert.Equal(t, tc.x, x)
			assert.Equal(t, tc.y, y)
		})
	}
}

func TestSolve(t *testing.T) {
	tt := map[string]struct {
		c   []BBPair
		res int64
	}{
		"23 is the smallest int X such that X % 3 == 2; X % 5 == 3; and X % 7 == 2": {
			c: []BBPair{
				{
					Divisor:   *big.NewInt(3),
					Remainder: *big.NewInt(2),
				},
				{
					Divisor:   *big.NewInt(5),
					Remainder: *big.NewInt(3),
				},
				{
					Divisor:   *big.NewInt(7),
					Remainder: *big.NewInt(2),
				},
			},
			res: 23,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			r := Solve(tc.c)
			assert.Equal(t, tc.res, r.Int64())
		})
	}

}
