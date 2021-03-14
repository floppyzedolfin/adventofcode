package dec01

import (
	"fmt"
)

// findMatch2 returns either the ints that sum up to the target, or an error.
// findMatch2 has a O(n) complexity, but it can't really be adapted for Part Two.
func findMatch2(data []int, target int) ([]int, error) {
	missing := make(map[int]struct{})
	for _, int1 := range data {
		// check if we were looking for this item
		if _, ok := missing[int1]; ok {
			return []int{int1, target-int1}, nil
		}
		// store the remainder
		missing[target-int1] = struct{}{}
	}
	return nil, fmt.Errorf("unable to find two ints that sum up to %d", target)
}

// findMatch3 is the non-recursive implementation of the Part Two.
// It's in O(n²), as it calls findMatch2 n times.
func findMatch3(data[]int, target int) ([]int, error) {
	if len(data) < 3 {
		return nil, fmt.Errorf("dataset is too small, we need at least 3 items")
	}
	// make sure we don't go too far
	for i, int1 := range data[:len(data)-3] {
		res, err := findMatch2(data[i+1:], target-int1)
		if err != nil {
			// this is not a complete failure. It only means this item cannot find buddies - but there are other ghotis in the sea
			continue
		}
		// notice we only care for the first result, as we suppose we'll have only one match
		return append(res, int1), nil
	}
	return nil, fmt.Errorf("unable to find a match of 3 items that sum up to %d", target)
}

// lazyFind2 has a O(n²) complexity...
func lazyFind2(data []int, target int) ([]int, error) {
	for i := range data[:len(data)-2] {
		for j := range data[i+1:] {
			if i + j == target {
				return []int{i, j}, nil
			}
		}
	}
	return nil, fmt.Errorf("unable to find ints that sum up to %d", target)
}

// lazyFind3 has a O(n³) complexity...
func lazyFind3(data []int, target int) ([]int, error) {
	l := len(data)
	for i := range data[:l-2] {
		for j := range data[i+1:l-1] {
			// let's shift properly. We want to start right after j (hence (j+1)), which started right after i (hence (i+1)).
			for k := range data[(i+1)+(j+1):] {
				if i+j+k == target {
					return []int{i, j, k}, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unable to find ints that sum up to %d", target)
}
