package dec02

import (
	"bufio"
	"fmt"
	"os"

	"github.com/floppyzedolfin/adventofcode/door"
)

// New builds a solver that can solve the exercise of Dec 02.
func New(inputPath string) door.Solver {
	return dec02Solver{inputPath: inputPath}
}

// Implementation of the solver for dec02
type dec02Solver struct {
	inputPath string
}

// Solve implements the Solver interface
func (s dec02Solver) Solve() (door.Result, error) {
	passwords, err := readLines(s.inputPath)
	if err != nil {
		return nil, fmt.Errorf("readLines: %s", err.Error())
	}
	validPasswords := 0
	for _, p := range passwords {
		// change here to use the validator of the first part.
		if p.isValid2() {
			validPasswords++
		}
	}
	return dec02Result{validPasswords: validPasswords}, nil
}

// Implementation of the result for dec02
type dec02Result struct {
	validPasswords int
}

// String implements the Result interface
func (r dec02Result) String() string {
	return fmt.Sprintf("The number of valid passwords is %d.", r.validPasswords)
}

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
