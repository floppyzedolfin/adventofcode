package instruction

// Operation describes the action to take for a line
type Operation string

const (
	Accumulate  Operation = "acc"
	Jump        Operation = "jmp"
	NoOperation Operation = "nop"
)

// Instruction holds what is necessary to understand a single executable line
type Instruction struct {
	Op  Operation
	Arg int
	AlreadySeen bool
}
