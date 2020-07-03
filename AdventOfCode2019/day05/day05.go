package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"../utils"
)

type program struct {
	instructions   []int
	programCounter int
	finished       bool
}

type operand struct {
	value int
	mode  int
}

func (p *program) run() {
	buf := bufio.NewReader(os.Stdin)
	p.programCounter = 0
	p.finished = false
	for !p.finished {
		fmt.Println("instruction", p.instructions[p.programCounter])
		opCode, op0, op1, op2 := p.getValues(p.instructions[p.programCounter])
		fmt.Println("op", opCode)
		switch opCode {
		case 1:
			// fmt.Println(p.instructions[p.programCounter+1], p.instructions[p.programCounter+2], p.instructions[p.programCounter+3])
			p.add(op0, op1, op2)
			p.programCounter += 4
		case 2:
			p.multiply(op0, op1, op2)
			p.programCounter += 4
		case 3:
			p.input(5, op0)
			p.programCounter += 2
		case 4:
			p.output(op0)
			p.programCounter += 2
		case 5:
			// fmt.Println(p.instructions[p.programCounter+1], p.instructions[p.programCounter+2])
			p.jumpIfTrue(op0, op1)
		case 6:
			p.jumpIfFalse(op0, op1)
		case 7:
			p.lessThan(op0, op1, op2)
			p.programCounter += 4
		case 8:
			p.equals(op0, op1, op2)
			p.programCounter += 4
		case 99:
			p.halt()
		default:
			fmt.Println("Invalid opcode")
			os.Exit(1)
		}
		fmt.Println("p.programCounter", p.programCounter)
		_, _ = buf.ReadBytes('\n')
	}
}

func (p *program) getValues(op int) (int, int, int, int) {
	opCode := op % 100
	if opCode == 99 {
		return opCode, 0, 0, 0
	}
	if opCode == 3 {
		return opCode, p.instructions[p.programCounter+1], 0, 0
	}
	if opCode == 4 {
		if math.Floor(float64((op%1000)/100)) == 1 {
			return opCode, p.instructions[p.programCounter+1], 0, 0
		} else {
			return opCode, p.instructions[p.instructions[p.programCounter+1]], 0, 0
		}
	}
	if opCode == 1 || opCode == 2 || (opCode >= 5 && opCode <= 8) {
		var arg1, arg2, arg3 int
		if math.Floor(float64((op%1000)/100)) == 1 {
			arg1 = p.instructions[p.programCounter+1]
		} else {
			arg1 = p.instructions[p.instructions[p.programCounter+1]]
		}
		if math.Floor(float64((op%10000)/1000)) == 1 {
			arg2 = p.instructions[p.programCounter+2]
		} else {
			arg2 = p.instructions[p.instructions[p.programCounter+2]]
		}
		if opCode != 5 && opCode != 6 {
			arg3 = p.instructions[p.programCounter+3]
		}
		return opCode, arg1, arg2, arg3
	}
	// modes := make([]int, 3)
	// paramModes := opcode / 100
	// modes[0] = paramModes % 10
	// paramModes /= 10
	// modes[1] = paramModes % 10
	// paramModes /= 10
	// modes[2] = paramModes % 10
	// return 0, 0, opcode % 100, modes
	return 0, 0, 0, 0
}

func (p *program) getOperand(op operand) int {
	if op.mode == 0 {
		return p.instructions[op.value]
	}
	return op.value
}

func (p *program) halt() {
	p.finished = true
}

func (p *program) jumpIfTrue(op0, op1 int) {
	if op0 != 0 {
		p.programCounter = op1
	} else {
		p.programCounter += 3
	}
}

func (p *program) jumpIfFalse(op0, op1 int) {
	if op0 == 0 {
		p.programCounter = op1
	} else {
		p.programCounter += 3
	}
}

func (p *program) lessThan(op0, op1, resPos int) {
	if op0 < op1 {
		p.instructions[resPos] = 1
	} else {
		p.instructions[resPos] = 0
	}
}

func (p *program) equals(op0, op1, resPos int) {
	if op0 == op1 {
		p.instructions[resPos] = 1
	} else {
		p.instructions[resPos] = 0
	}
}

func (p *program) add(op0, op1, resPos int) {
	fmt.Println("adding", op0, op1, "and saving", op0+op1, "to position", resPos)
	p.instructions[resPos] = op0 + op1
}

func (p *program) multiply(op0, op1, resPos int) {
	p.instructions[resPos] = op0 * op1
}

func (p *program) input(op, pos int) {
	fmt.Println("saving", op, "to position", pos)
	p.instructions[pos] = op
}

func (p *program) output(pos int) {
	fmt.Println(pos)
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

// func part2(instructions []int) int {
// 	instructionsTemp := make([]int, len(instructions))
// 	for i := 0; i < 99; i++ {
// 		for j := 0; j < 99; j++ {
// 			copy(instructionsTemp, instructions)
// 			p := &program{instructions: instructionsTemp}
// 			p.fixProgram(i, j)
// 			p.run()
// 			if p.getReturnValue() == 19690720 {
// 				return 100*i + j
// 			}
// 		}
// 	}
// 	return -1
// }

func main() {
	readF, err := utils.ReadFile("in")
	if err != nil {
		log.Fatal(err)
	}
	input, err := readF.ReadAndParseLines(1, ",", true, int(0))
	if err != nil {
		log.Fatal(err)
	}
	instructions := make([]int, len(input[0]))
	for i, elem := range input[0] {
		instructions[i] = elem.(int)
	}
	p := &program{instructions: instructions}
	p.run()
}
