package day02

import (
	"log"
	"os"
	"testing"

	"../utils"
	"github.com/stretchr/testify/assert"
)

var instructions []int

func TestMain(m *testing.M) {
	readF, err := utils.ReadFile("day02.input")
	if err != nil {
		log.Fatal(err)
	}
	input, err := readF.ReadAndParseLines(1, ",", true, int(0))
	if err != nil {
		log.Fatal(err)
	}
	instructions = make([]int, len(input[0]))
	for i, elem := range input[0] {
		instructions[i] = elem.(int)
	}
	returnCode := m.Run()
	os.Exit(returnCode)
}

func TestRunProgram(t *testing.T) {
	assert.Equal(t, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, runProgram([]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}), "")
	assert.Equal(t, []int{2, 0, 0, 0, 99}, runProgram([]int{1, 0, 0, 0, 99}), "")
	assert.Equal(t, []int{2, 3, 0, 6, 99}, runProgram([]int{2, 3, 0, 3, 99}), "")
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801}, runProgram([]int{2, 4, 4, 5, 99, 0}), "")
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, runProgram([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}), "")

}

func TestPart1(t *testing.T) {
	assert.Equal(t, 5534943, part1(instructions), "")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 7603, part2(instructions), "")
}

// func TestFuelForFuel(t *testing.T) {
// 	assert.Equal(t, 2, fuelForFuel(14), "")
// 	assert.Equal(t, 966, fuelForFuel(1969), "")
// 	assert.Equal(t, 50346, fuelForFuel(100756), "")
// 	fuelRequiredComplete := 0
// 	for _, m := range modules {
// 		fuelRequiredComplete += fuelForFuel(m)
// 	}
// 	assert.Equal(t, 4943578, fuelRequiredComplete, "")
// }
