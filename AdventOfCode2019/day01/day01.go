package main

import (
	"fmt"
	"io"
	"log"
	"math"

	"../utils"
)

func moduleMassToFuelRequired(module int) int {
	return module/3 - 2
}

func fuelForFuel(module int) int {
	totalFuel := 0
	fuel := int(math.Floor(float64(module)/3.0)) - 2
	for fuel > 0 {
		totalFuel += fuel
		module = fuel
		fuel = int(math.Floor(float64(module)/3.0)) - 2
	}
	return totalFuel
}

func main() {
	readF, err := utils.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	fuelRequired := 0
	fuelRequiredComplete := 0
	for {
		lines, err := readF.ReadAndParseLines(1, " ", false, int(0))
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		moduleMass := lines[0][0].(int)
		fuelRequired += moduleMassToFuelRequired(moduleMass)
		fuelRequiredComplete += fuelForFuel(moduleMass)
	}
	fmt.Println(fuelRequired)
	fmt.Println(fuelRequiredComplete)
	readF.Close()
}

// 3297626
// 4943578
