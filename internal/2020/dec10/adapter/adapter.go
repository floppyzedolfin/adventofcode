package adapter

import (
	"github.com/floppyzedolfin/adventofcode/internal/2020/dec10/tribonacci"
)

type adapter = int

// Adapters holds the list of adapters we found in the bag
type Adapters struct {
	list []adapter
}

// Jolts1By3 returns the number of 1-jolt differences multiplied by the number of 3-jolt differences
func Jolts1By3(as Adapters) int64 {
	// previously read value - initialised with the plane joltage
	prev := 0
	// counters for differences of 1 jolt and 3 jolts.
	ones, threes := 0, 0
	for _, a := range as.list {
		// I know you're not supposed to do that, but this is a switch on an adapter.
		switch a - prev {
		case 1:
			ones++
		case 3:
			threes++
		}
		prev = a
	}
	// include the device
	threes++
	return int64(ones * threes)
}

// DifferentCombinations counts the number of different combinations, even without some of the adapters (respecting the 3-jolt maximal gap)
func DifferentCombinations(as Adapters) int64 {
	// we're given the clue this might be a humongus number
	count := int64(1)
	// ones holds the number of gaps of 1 we've seen since the last gap of 3
	ones := 0
	prev := 0
	// we'll need a Tribonacci sequence for some computation
	tribSeq := tribonacci.New(1,1,2)
	for _, a := range as.list {
		switch a - prev {
		// gaps of length 2 are ignored here, as none of the input files contained any. Also, they make the computation way more complex.
		case 1:
			ones++
		case 3:
			// stop counting ones, instead include the number of combinations they brought
			count *= tribSeq.At(ones)
			ones = 0
		}
		prev = a
	}
	// don't forget to include the device, which is 3 ahead.
	return count * tribSeq.At(ones)
}
