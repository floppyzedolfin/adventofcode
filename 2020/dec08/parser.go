package dec08

import (
	"fmt"

	"github.com/floppyzedolfin/adventofcode/2020/dec08/program"
	"github.com/floppyzedolfin/adventofcode/fileparser"
)

func readInstructions(inputFile string) (program.Program, error) {
	p := program.Build(nil)
	err := fileparser.ParseFile(inputFile, &p)
	if err != nil {
		return program.Program{}, fmt.Errorf("unable to parse file '%s': %s", inputFile, err.Error())
	}
	return p, nil
}
