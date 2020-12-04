package dec02

import (
	"bufio"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of passwords - its lines.
func readLines(path string) ([]password, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var passwords []password
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p, err := buildPassword(scanner.Text())
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, p)
	}
	return passwords, scanner.Err()
}
