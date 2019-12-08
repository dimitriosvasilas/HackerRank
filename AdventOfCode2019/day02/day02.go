package day02

type program struct {
	instructions   []int
	programCounter int
	finished       bool
}

func (p *program) run() {
	p.programCounter = 0
	p.finished = false
	for !p.finished {
		switch p.instructions[p.programCounter] {
		case 1:
			p.add(p.instructions[p.programCounter+1], p.instructions[p.programCounter+2], p.instructions[p.programCounter+3])
		case 2:
			p.multiply(p.instructions[p.programCounter+1], p.instructions[p.programCounter+2], p.instructions[p.programCounter+3])
		case 99:
			p.halt()
		}
		p.programCounter += 4
	}
}

func (p *program) halt() {
	p.finished = true
}

func (p *program) add(op0Pos, op1Pos, resPos int) {
	p.instructions[resPos] = p.instructions[op0Pos] + p.instructions[op1Pos]
}

func (p *program) multiply(op0Pos, op1Pos, resPos int) {
	p.instructions[resPos] = p.instructions[op0Pos] * p.instructions[op1Pos]
}

func (p program) getReturnValue() int {
	return p.instructions[0]
}

func (p *program) fixProgram(noun, verb int) {
	p.instructions[1] = noun
	p.instructions[2] = verb
}

func runProgram(instructions []int) []int {
	p := &program{instructions: instructions}
	p.run()
	return p.instructions
}

func part1(instructions []int) int {
	instructionsTemp := make([]int, len(instructions))
	copy(instructionsTemp, instructions)
	p := &program{instructions: instructionsTemp}
	p.fixProgram(12, 2)
	p.run()
	return p.getReturnValue()
}

func part2(instructions []int) int {
	instructionsTemp := make([]int, len(instructions))
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			copy(instructionsTemp, instructions)
			p := &program{instructions: instructionsTemp}
			p.fixProgram(i, j)
			p.run()
			if p.getReturnValue() == 19690720 {
				return 100*i + j
			}
		}
	}
	return -1
}
