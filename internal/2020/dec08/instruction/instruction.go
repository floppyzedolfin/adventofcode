package instruction

import (
	"fmt"
	"strconv"
)

// New builds an instruction from strings.
func New(op, arg string) (Instruction, error) {
	var i Instruction
	i.Op = Operation(op)
	if i.Op != Accumulate && i.Op != Jump && i.Op != NoOperation {
		return Instruction{}, fmt.Errorf("invalid operation '%s'", op)
	}
	iArg, err := strconv.Atoi(arg)
	if err != nil {
		return Instruction{}, fmt.Errorf("invalid argument '%s'", arg)
	}
	i.Arg = iArg
	return i, nil
}
