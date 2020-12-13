package dec09

import (
	"fmt"
	"strconv"

	"github.com/floppyzedolfin/adventofcode/fileparser"
	"github.com/floppyzedolfin/adventofcode/ptr"
)

func xmas1Solver(inputFile string, preambleLength int) (*int, error) {
	x := &xmas{preambleLength: preambleLength, data: make([]int, 0)}
	err := fileparser.ParseFile(inputFile, x)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file '%s': %s", inputFile, err.Error())
	}

	return x.firstInvalid(), nil
}

func xmas2Solver(inputFile string, preambleLength int) (*int, error) {
	x := &xmas{preambleLength: preambleLength, data: make([]int, 0)}
	err := fileparser.ParseFile(inputFile, x)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file '%s': %s", inputFile, err.Error())
	}

	invalidValue := x.firstInvalid()
	if invalidValue == nil {
		return nil, fmt.Errorf("expected a weakness, but this is unbreakable -- no xmas invalid number found")
	}

	return x.computeEncryptionWeakness(*invalidValue), nil
}

type xmas struct {
	preambleLength int
	data           []int
}

func (x *xmas) ParseLine(line string) error {
	v, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("error while reading line '%s': %s", line, err.Error())
	}
	x.data = append(x.data, v)
	return nil
}

// firstInvalid finds the first invalid value of the dataset (if any)
func (x *xmas) firstInvalid() *int {
	for index, value := range x.data {
		if !x.checkAtIndex(index) {
			return &value
		}
	}
	return nil
}

// checkAtIndex checks whether the value at the given position is valid.
func (x *xmas) checkAtIndex(pos int) bool {
	// while in preamble, everything is valid.
	if pos < x.preambleLength {
		return true
	}
	needle := x.data[pos]
	// check the sums
	for i, v1 := range x.data[pos-x.preambleLength : pos-1] {
		for _, v2 := range x.data[pos-x.preambleLength+i+1 : pos] {
			if v1+v2 == needle {
				// found it !
				return true
			}
		}
	}
	return false
}

// computeEncryptionWeakness returns the sum of the smallest and largest values in the set of contiguous numbers that add up to the requested value
func (x *xmas) computeEncryptionWeakness(v int) *int {
	dataLength := len(x.data)-1
	start, end, sum := 0, 1, x.data[0]
	for {
		// this shouldn't happen, as we should, at some point, point at v's position
		if end == dataLength {
			return nil
		}
		// we've gone too far
		if x.data[end] == v {
			return nil
		}
		// remember that, in Go, only the first matching case is executed
		switch {
		case sum == v:
			return ptr.Int(x.minMaxSum(start, end))
		case sum < v:
			// need to read more values
			end++
			sum += x.data[end]
		case sum > v:
			// read too much, need to discard values:
			sum -= x.data[start]
			start++
		}
	}
}

// minMaxSum adds the lowest and the highest values found in the input data between the bounds.
func (x *xmas) minMaxSum(start, end int) int {
	min, max := x.data[start], x.data[start]
	for _,v := range x.data[start:end+1] {
		switch {
		case v < min:
			min = v
		case v > max:
			max = v
		}
	}
	return min+max
}
