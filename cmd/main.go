package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/floppyzedolfin/adventofcode/2020/dec01"
	"github.com/floppyzedolfin/adventofcode/2020/dec02"
	"github.com/floppyzedolfin/adventofcode/door"
)

const (
	currentYear = 2020
	december = 12
)

var (
	lastDoor = 24
)

func init() {
	date := time.Now()
	if date.Year() == currentYear  && date.Month() == time.December {
		// Set the default door to open. There's no opening of future doors.
		day := date.Day()
		if day < lastDoor {
			lastDoor = day
		}
	}
}

func main() {
	d := getDoorOfRun()
	var s door.Solver

	switch d {
	case 1:
		s = dec01.New("2020/dec01/input")
	case 2:
		s = dec02.New("2020/dec02/input")
	default:
		log.Fatalf("can't open door %d yet", d)
	}

	res, err := s.Solve()
	if err != nil  {
		log.Fatalf("no treat for today: an error has occurred when opening the door %d: %s", d, err.Error())
	}

	fmt.Printf(" ---\nResult of day %d:\n%s\n ---\n", d, res.String())
}

func getDoorOfRun() int {
	// Extract the door to open from the command line.
	d := flag.Int("door", lastDoor, "the number of the door you want to open")
	flag.Parse()
	if d == nil {
		// Argument was optional, use the default value of "today" instead.
		d = &lastDoor
	}
	return *d
}
