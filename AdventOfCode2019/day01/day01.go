package day01

import (
	"math"
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
