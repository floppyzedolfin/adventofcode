package program

import (
	"fmt"
	"regexp"

	"github.com/floppyzedolfin/adventofcode/pkg/fileparser"
)

// Program holds the list Values of stored in the memory. It needs to know how to affect memory values
type Program struct {
	Values                  map[uint64]uint64
	AffectMemory func(address, value string) error
	AffectMask   func(mask string)
}

// Sum sums the entries stored in a program's memory
func (p Program) Sum() uint64 {
	var sum uint64
	for _, v := range p.Values {
		sum += v
	}
	return sum
}

// ParseFile reads the content of an input initialisation program and populates the program
func (p *Program) ParseFile(inputFile string) error {
	err := fileparser.ParseFile(inputFile, p)
	if err != nil {
		return fmt.Errorf("error while parsing file '%s': %s", inputFile, err.Error())
	}
	return nil
}

// ParseLine implements the LineParser interface
func (p *Program) ParseLine(line string) error {
	// lines can be either memory affectation or Mask overriding
	// let's try memory affectation first - because the input file we have has more of that one
	const memoryAffectationExpression = `^mem\[(\d+)\] = (\d+)$`
	memRE := regexp.MustCompile(memoryAffectationExpression)
	memLine := memRE.FindStringSubmatch(line)
	if len(memLine) == 3 {
		if err := p.AffectMemory(memLine[1], memLine[2]); err != nil {
			return fmt.Errorf("unable to affect a memory value when reading '%s': %s", line, err.Error())
		}
		return nil
	}

	// it wasn't a memory allocation, let's try a Mask overriding
	const maskOverridingExpression = `^mask = ([01X]{36})$`
	maskRE := regexp.MustCompile(maskOverridingExpression)
	maskLine := maskRE.FindStringSubmatch(line)
	if len(maskLine) == 2 {
		// the regexp already took care of nasty situations such as too short a Mask, or unexpected characters such as Radagast
		p.AffectMask(maskLine[1])
		return nil
	}

	return fmt.Errorf("unable to parse line '%s', it is neither a memory allocation nor a Mask definition", line)
}
