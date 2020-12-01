package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

)

// readLines reads a whole file into memory
// and returns a slice of its lines.
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

func findMatch(data []int, target int) (int, error) {
	missing := make(map[int]struct{})
	for _, line1 := range data {
		// check if we were looking for this item
		if _, ok := missing[line1]; ok {
			return line1*(target-line1), nil
		}
		// store the remainder
		missing[target-line1] = struct{}{}
	}
	return -1, fmt.Errorf("not found")
}

func main() {
	lines, err := readLines("./input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	res, err := findMatch(lines, 2020)
	if err !=nil {
		log.Fatalf("findMatch: %s", err)
	}
	fmt.Printf("res: %d\n", res)
}
