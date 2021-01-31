package dec14

import (
	"fmt"

	"github.com/floppyzedolfin/adventofcode/2020/dec14/programv1"
	"github.com/floppyzedolfin/adventofcode/2020/dec14/programv2"
	"github.com/floppyzedolfin/adventofcode/ptr"
)

// sumValuesV1 computes the sum of the values written in the memory
func sumValuesV1(inputFile string) (*uint64, error) {
	p1 := programv1.New()
	err := p1.ParseFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("unable to parse input file '%s': %s", inputFile, err.Error())
	}
	return ptr.Uint64(p1.Sum()), nil
}

// sumValuesV2 computes the sum of the values of a program2 input
func sumValuesV2(inputFile string) (*uint64, error) {
	p2 := programv2.New()
	err := p2.ParseFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("unable to parse input file '%s': %s", inputFile, err.Error())
	}
	return ptr.Uint64(p2.Sum()), nil
}
