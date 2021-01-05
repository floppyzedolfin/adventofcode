package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/floppyzedolfin/adventofcode/2020/dec01"
	"github.com/floppyzedolfin/adventofcode/2020/dec02"
	"github.com/floppyzedolfin/adventofcode/2020/dec03"
	"github.com/floppyzedolfin/adventofcode/2020/dec04"
	"github.com/floppyzedolfin/adventofcode/2020/dec05"
	"github.com/floppyzedolfin/adventofcode/2020/dec06"
	"github.com/floppyzedolfin/adventofcode/2020/dec07"
	"github.com/floppyzedolfin/adventofcode/2020/dec08"
	"github.com/floppyzedolfin/adventofcode/2020/dec09"
	"github.com/floppyzedolfin/adventofcode/2020/dec10"
	"github.com/floppyzedolfin/adventofcode/2020/dec11"
	"github.com/floppyzedolfin/adventofcode/2020/dec12"
	"github.com/floppyzedolfin/adventofcode/door"
)

const (
	december = 12
)

var (
	// flag holders
	currentDoorFlag  int
	currentPartsFlag string
)

func init() {
	// flags
	const (
		lastDoor     = 24
		defaultParts = "1,2"
	)

	flag.IntVar(&currentDoorFlag, "door", lastDoor, "the number of the door you want to open")
	flag.StringVar(&currentPartsFlag, "parts", defaultParts, "parts you want to run (separated with commas)")
}

func main() {
	d, p, err := readArguments()
	if err != nil {
		panic(err.Error())
	}

	var s door.Solver

	switch d {
	case 1:
		s = dec01.New("2020/dec01/input")
	case 2:
		s = dec02.New("2020/dec02/input")
	case 3:
		s = dec03.New("2020/dec03/input")
	case 4:
		s = dec04.New("2020/dec04/input")
	case 5:
		s = dec05.New("2020/dec05/input")
	case 6:
		s = dec06.New("2020/dec06/input")
	case 7:
		s = dec07.New("2020/dec07/input")
	case 8:
		s = dec08.New("2020/dec08/input")
	case 9:
		s = dec09.New("2020/dec09/input")
	case 10:
		s = dec10.New("2020/dec10/input")
	case 11:
		s = dec11.New("2020/dec11/input")
	case 12:
		s = dec12.New("2020/dec12/input")
	default:
		log.Fatalf("can't open door %d yet", d)
	}

	res, err := s.Solve(p)
	if err != nil {
		log.Fatalf("no treat for today: an error has occurred when opening the door %d: %s", d, err.Error())
	}

	fmt.Printf(" ---\nResult of day %d:\n%s ---\n", d, res.String())
}

func readArguments() (int, door.Parts, error) {
	flag.Parse()

	// handle doors (ho ho ho)
	// Argument was optional, use the default value of "today" instead.
	date := time.Now()
	if date.Year() == door.Year && date.Month() == time.December {
		// Set the default door to open. There's no opening of future doors.
		day := date.Day()
		if day < currentDoorFlag {
			currentDoorFlag = day
		}
	}

	// deal with parts
	parts, err := door.ParseParts(currentPartsFlag)
	if err != nil {
		return 0, nil, fmt.Errorf("unable to parse parts: %s", err.Error())
	}

	return currentDoorFlag, parts, nil
}
