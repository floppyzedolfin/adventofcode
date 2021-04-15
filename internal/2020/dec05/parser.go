package dec05

import (
	"fmt"
	"regexp"

	"github.com/floppyzedolfin/adventofcode/pkg/fileparser"
)

func readSeats(inputFile string) ([]int, error) {
	var s seats
	err := fileparser.ParseFile(inputFile, &s)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the file '%s': %s", inputFile, err.Error())
	}

	return s.list, nil
}

// local structure needed to parse the file
type seats struct {
	list []int
}

// ParseLine implements the LineParser interface. We add a seat for each line
func (s *seats) ParseLine(line string) error {
	const seatExpression = `^[FB]{7}[LR]{3}`
	re := regexp.MustCompile(seatExpression)
	if !re.MatchString(line) {
		return fmt.Errorf("invalid line '%s', doesn't match the regexp", line)
	}

	// this should be a const, but Go doesn't support that...
	var positions = map[byte]int{
		'F': 0, // Front of the plane - closer to int 1
		'B': 1, // Back of the plane - closer to int 2^n-1
		'R': 1, // uppeR (???) half - closer to 7
		'L': 0, // Lower half - closer to 0
	}

	seatPos := 0
	// the seats are written in big-endian syntax, hence the reading from left to right.
	for i := 0; i < logPlaneCapacity; i++ {
		seatPos *= 2
		seatPos += positions[line[i]]
	}
	s.list = append(s.list, seatPos)

	return nil
}
