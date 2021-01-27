package programv1

import (
	"fmt"
	"strconv"

	"github.com/floppyzedolfin/adventofcode/2020/dec14/mask"
	"github.com/floppyzedolfin/adventofcode/2020/dec14/program"
)

// ProgramV1 solves the first part of the problem
type ProgramV1 struct {
	program.Program
	mask map[uint8]rune // each position has either '0', '1' or 'X'
}

// New returns a ProgramV1 solver
func New() *ProgramV1 {
	var p ProgramV1
	p.Values = make(map[uint64]uint64)
	p.mask = make(map[uint8]rune)
	p.AffectMemory = p.affectMemory
	p.AffectMask = p.affectMask
	return &p
}

// parseMask saves a mask to the program
func (p ProgramV1) affectMask(m string) {
	for i, c := range m {
		// remember, the first value we read will override the 2^35 bit.
		p.mask[mask.LastIndex-uint8(i)] = c
	}
}

func (p *ProgramV1) affectMemory(address, value string) error {
	// first, decode the value to an uint64
	value64, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return fmt.Errorf("unable to convert value to uint64")
	}

	// second, write the value at the correct address
	memAddress, err := strconv.ParseUint(address, 10, 64)
	if err != nil {
		return fmt.Errorf("unable to convert address to uint16")
	}

	// finally, apply the mask and save it
	p.Values[memAddress] = p.applyMask(value64)
	return nil
}

// applyMask applies p's mask to a value
func (p *ProgramV1) applyMask(value uint64) uint64 {
	for i, m := range p.mask {
		p := mask.PowerOf2(i)
		switch m {
		case '0':
			// the mask removes the value if present
			if value&p != 0 {
				value -= p
			}
		case '1':
			// the mask adds the value if absent
			if value&p == 0 {
				value += p
			}
		default:
			// nothing to be done
		}
	}
	return value
}
