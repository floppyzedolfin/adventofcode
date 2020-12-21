package program

import (
	"fmt"
	"strings"

	"github.com/floppyzedolfin/adventofcode/2020/dec08/instruction"
)

// A Program is a list of instructions
type Program struct {
	instructions    []instruction.Instruction
	accumulator     int
	currentPosition int
}

// runNoLoop goes through instructions and exits when it detects a loop
func (p *Program) runNoLoop() {
	count := 0
	// this counter is a safety-net. If we have no loop, we'll go through each line. If we have a loop, let's make sure we didn't miss something.
	for ; count < len(p.instructions); count++ {
		// let's discard the error, as we don't need it
		if ok, _ := p.executeInstruction(); !ok {
			return
		}
	}
	return
}

// executeInstruction returns false if the head of execution has already been executed, and, otherwise, executes the instruction there.
func (p *Program) executeInstruction() (bool, error) {
	if p.currentPosition == len(p.instructions) {
		// we don't want to execute the "next" line.
		return false, nil
	}
	if p.currentPosition < 0 || p.currentPosition > len(p.instructions) {
		// we definitely do not want to execute that line, as we've successfully reached the end of the program! Yee-hee!
		return false, fmt.Errorf("invalid position %d", p.currentPosition)
	}
	i := &p.instructions[p.currentPosition]
	if i.AlreadySeen {
		return false, fmt.Errorf("instruction %v at position %d has already been executed", *i, p.currentPosition)
	}
	jump := 1
	switch i.Op {
	case instruction.NoOperation:
		// nothing to do here.
	case instruction.Accumulate:
		p.accumulator += i.Arg
	case instruction.Jump:
		jump = i.Arg
	}
	p.currentPosition += jump
	i.AlreadySeen = true
	return true, nil
}

// reset lets someone else re-run the program as if brand-new.
func (p *Program) reset() {
	for i := range p.instructions {
		p.instructions[i].AlreadySeen = false
	}
	p.accumulator = 0
	p.currentPosition = 0
}

// RunTillLoop returns the value of the accumulator right before we would execute an instruction for the second time
func RunTillLoop(p Program) int {
	p.runNoLoop()
	return p.accumulator
}

// FixOnce tries to fix the program by flicking between Jump and NoOp.
func FixOnce(p Program) int {
	nbInstructions := len(p.instructions)
	for i := 0; i < nbInstructions; i++ {
		// switch the i-th instruction if it is a Nop/Jump
		switch p.instructions[i].Op {
		case instruction.Accumulate:
			continue
		case instruction.Jump:
			p.instructions[i].Op = instruction.NoOperation
			p.runNoLoop()
			p.instructions[i].Op = instruction.Jump
		case instruction.NoOperation:
			p.instructions[i].Op = instruction.Jump
			p.runNoLoop()
			p.instructions[i].Op = instruction.NoOperation
		}
		if p.currentPosition == len(p.instructions) {
			return p.accumulator
		}
		// clean up
		p.reset()
	}
	// "hopefully", we shouldn't be here. It would mean the exercise can't be solved.
	return 0
}

// Build builds a program from a set of instructions
func Build(instrs []instruction.Instruction) Program {
	p := Program{instructions: make([]instruction.Instruction, 0, len(instrs))}
	for i, instr := range instrs {
		p.instructions[i] = instr
	}
	return p
}

// ParseLine implements the LineParser interface
func (p *Program) ParseLine(line string) error {
	fields := strings.Split(line, " ")
	if len(fields) != 2 {
		return fmt.Errorf("invalid  line '%s'", line)
	}
	inst, err := instruction.New(fields[0], strings.TrimLeft(fields[1],"+"))
	if err != nil {
		return fmt.Errorf("unable to create instruction: %s", err.Error())
	}
	p.instructions = append(p.instructions, inst)
	return nil
}
