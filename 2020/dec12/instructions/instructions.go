package instructions

// Instruction holds a single move of the vessel
type Instruction struct {
	Action Action
	Value  int
}

// Instructions holds the list of instructions the vessel will accomplish
type Instructions struct {
	Commands []Instruction
}
