package dec01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/floppyzedolfin/adventofcode/door"
)

// New builds a solver that can solve the exercise of Dec 01.
func New(inputPath string) door.Solver {
	return dec01Solver{inputPath: inputPath}
}

// Implementation of the solver for dec01
type dec01Solver struct {
	inputPath string
}

// Solve implements the Solver interface
func (s dec01Solver) Solve() (door.Result, error) {
	lines, err := readLines(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read lines: %s", err.Error())
	}

	matches, err := findMatch3(lines, 2020)
	if err != nil {
		return nil, fmt.Errorf("unable to find match: %s", err.Error())
	}
	return dec01Result{productPart2:  product(matches)}, nil
}

// Implementation of the result for dec01
type dec01Result struct {
	productPart1 int
	productPart2 int
}

// String implements the Result interface
func (r dec01Result) String() string {
	if r.productPart1 == 0 && r.productPart2 == 0 {
		return "No match found."
	}
	output := strings.Builder{}
	if r.productPart1 != 0 {
		output.WriteString(fmt.Sprintf("The product for Part 1 is %d.\n", r.productPart1))
	}
	if r.productPart2 != 0 {
		output.WriteString(fmt.Sprintf("The product for Part 2 is %d.\n", r.productPart2))
	}
	return output.String()
}

// readLines reads a whole file into memory
// and returns a slice of ints - its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, l)
	}
	return lines, scanner.Err()
}

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

// product performs the product of ints in a slice.
func product(data []int) int {
	p := 1
	for _, d := range data {
		p *= d
	}
	return p
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
