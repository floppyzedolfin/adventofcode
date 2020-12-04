package dec01

import (
	"bufio"
	"os"
	"strconv"
)

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
