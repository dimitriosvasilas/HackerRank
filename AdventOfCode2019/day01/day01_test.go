package day01

import (
	"io"
	"log"
	"os"
	"testing"

	"../utils"
	"github.com/stretchr/testify/assert"
)

var modules []int

func TestMain(m *testing.M) {
	readF, err := utils.ReadFile("day01.input")
	if err != nil {
		log.Fatal(err)
	}
	modules = make([]int, 0)
	for {
		lines, err := readF.ReadAndParseLines(1, " ", false, int(0))
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		modules = append(modules, lines[0][0].(int))
	}
	returnCode := m.Run()
	os.Exit(returnCode)
}

func TestModuleMassToFuelRequired(t *testing.T) {
	assert.Equal(t, 2, moduleMassToFuelRequired(12), "")
	assert.Equal(t, 2, moduleMassToFuelRequired(14), "")
	assert.Equal(t, 654, moduleMassToFuelRequired(1969), "")
	assert.Equal(t, 33583, moduleMassToFuelRequired(100756), "")
	fuelRequired := 0
	for _, m := range modules {
		fuelRequired += moduleMassToFuelRequired(m)
	}
	assert.Equal(t, 3297626, fuelRequired, "")
}

func TestFuelForFuel(t *testing.T) {
	assert.Equal(t, 2, fuelForFuel(14), "")
	assert.Equal(t, 966, fuelForFuel(1969), "")
	assert.Equal(t, 50346, fuelForFuel(100756), "")
	fuelRequiredComplete := 0
	for _, m := range modules {
		fuelRequiredComplete += fuelForFuel(m)
	}
	assert.Equal(t, 4943578, fuelRequiredComplete, "")
}
