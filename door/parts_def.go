package door

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Part1 = 1
	Part2 = 2
)

// Parts lists the parts we want to solve.
type Parts []int

// ParseParts parses a string and turns it into a Parts
func ParseParts(v string) (Parts, error) {
	items := strings.Split(v, ",")
	parts := make(Parts, len(items))
	for i, item := range items {
		part, err := strconv.Atoi(item)
		if err != nil {
			return nil, fmt.Errorf("unable to parse part %s", item)
		}
		parts[i] = part
	}
	return parts, nil
}

// Contains lets us know whether a value was added to a Parts
func (p *Parts)Contains(value int) bool {
	for _, part := range *p {
		if part == value {
			return true
		}
	}
	return false
}