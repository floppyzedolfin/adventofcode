package programv2

import (
	"fmt"
	"strconv"

	"github.com/floppyzedolfin/adventofcode/internal/2020/dec14/mask"
	"github.com/floppyzedolfin/adventofcode/internal/2020/dec14/program"
)

// ProgramV2 solves the second part of the problem
type ProgramV2 struct {
	program.Program
	maskOnes []uint8
	maskExes []uint8
}

// New returns a ProgramV2 solver
func New() *ProgramV2 {
	p := ProgramV2{}
	p.Values = make(map[uint64]uint64)
	p.AffectMemory = p.affectMemory
	p.AffectMask = p.affectMask
	return &p
}

func (p *ProgramV2) affectMask(m string) {
	p.maskOnes = make([]uint8, 0)
	p.maskExes = make([]uint8, 0)
	for i, c := range m {
		// remember, the first value we read will override the 2^35 bit.
		index := mask.LastIndex - uint8(i)
		switch c {
		case '1':
			p.maskOnes = append(p.maskOnes, index)
		case '0':
			// nothing to do - 0's don't affect the address
		default:
			// X, x,...
			p.maskExes = append(p.maskExes, index)
		}
	}
}

func (p *ProgramV2) affectMemory(address, value string) error {
	// first decode the value we want to write in some places
	value64, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return fmt.Errorf("unable to convert value to uint64")
	}

	// second, decode the "initial" address
	initialAddress, err := strconv.ParseUint(address, 10, 64)
	if err != nil {
		return fmt.Errorf("unable to convert address to uint16")
	}

	// finally, write the address wherever needed
	p.writeValueAtAddresses(value64, initialAddress)
	return nil
}

// writeValueAtAddresses writes the value at each eligible address
func (p *ProgramV2) writeValueAtAddresses(value, address uint64) {
	for _, index := range p.maskOnes {
		// overwrite memory address bit
		address |= mask.PowerOf2(index)
	}

	p.writeValueAtExes(value, address, p.maskExes)
}

// writeValueAtExes is recursive. The cost of this func is 2^len(exes)
func (p *ProgramV2) writeValueAtExes(value, address uint64, exes []uint8) {
	// this is the end of the task. Save value at computed address
	if len(exes) == 0 {
		p.Values[address] = value
		return
	}

	pow := mask.PowerOf2(exes[0])
	// consider the first X as a 0
	if address&pow == 0 {
		p.writeValueAtExes(value, address, exes[1:])
	} else {
		// overwrite the 1 with a 0
		p.writeValueAtExes(value, address-pow, exes[1:])
	}

	// consider the first X as a 1
	p.writeValueAtExes(value, address|pow, exes[1:])
}
