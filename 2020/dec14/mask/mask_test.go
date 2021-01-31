package mask

import (
	"testing"
)

// This benchmark file is used to explain why I get powers of 2 using a func rather than using a LUT (original intention)
var randomNumber = uint8(4) // chosen by fair dice roll, guaranteed to be random.

// BenchmarkPowersOf2Array benchmarks the array
func BenchmarkPowersOf2Array(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = powersOf2[randomNumber]
	}
}

// BenchmarkPowerOf2Function benchmarks the func
func BenchmarkPowerOf2Function(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PowerOf2(randomNumber)
	}
}

// powersOf2 is a lookup table to avoid repeating the computation of powers of 2.
var powersOf2 = map[uint8]uint64{
	0:  1 << 0,
	1:  1 << 1,
	2:  1 << 2,
	3:  1 << 3,
	4:  1 << 4,
	5:  1 << 5,
	6:  1 << 6,
	7:  1 << 7,
	8:  1 << 8,
	9:  1 << 9,
	10: 1 << 10,
	11: 1 << 11,
	12: 1 << 12,
	13: 1 << 13,
	14: 1 << 14,
	15: 1 << 15,
	16: 1 << 16,
	17: 1 << 17,
	18: 1 << 18,
	19: 1 << 19,
	20: 1 << 20,
	21: 1 << 21,
	22: 1 << 22,
	23: 1 << 23,
	24: 1 << 24,
	25: 1 << 25,
	26: 1 << 26,
	27: 1 << 27,
	28: 1 << 28,
	29: 1 << 29,
	30: 1 << 30,
	31: 1 << 31,
	32: 1 << 32,
	33: 1 << 33,
	34: 1 << 34,
	35: 1 << 35,
}
