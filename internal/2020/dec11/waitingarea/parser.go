package waitingarea

import (
	"fmt"

	"github.com/floppyzedolfin/adventofcode/pkg/fileparser"
)

// ReadSeats reads the seats and returns them
func ReadSeats(inputFile string) (*WaitingArea, error) {
	wa := &WaitingArea{seats: make([][]seat, 0)}
	err := fileparser.ParseFile(inputFile, wa)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file '%s': %s", inputFile, err.Error())
	}
	return wa, nil
}

// ParseLine implements the LineParser interface
func (wa *WaitingArea) ParseLine(line string) error {
	currentRow := len(wa.seats)
	row := make([]seat, len(line))
	for i, c := range line {
		row[i] = seat{
			status: c,
			posX:   i,
			posY:   currentRow,
		}
		if c == seatOccupiedChar {
			wa.occupied++
		}
	}
	wa.seats = append(wa.seats, row)
	return nil
}
