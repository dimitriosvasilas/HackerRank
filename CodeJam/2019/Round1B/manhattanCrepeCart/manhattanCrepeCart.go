package main

import (
	"fmt"
)

func solve(P, Q int, X, Y []int, D []rune) (int, int) {
	// Main idea
	// Create an 2d array representing the area
	// Initialize all cells' values to 0
	// For each person, increment the values of the cells they are moving towards to
	// Find the the cell with the maximum value
	// If more than one exist, use the tiebraker rules
	// O(P*Q^2)
	// Insight: The two dimensions are indepentend
	// Solve as two 1D probems

	southNorth := make([]int, Q+1)
	westEast := make([]int, Q+1)

	for p := 0; p < P; p++ {

		switch D[p] {
		case 'N':
			for i := Y[p] + 1; i <= Q; i++ {
				southNorth[i]++
			}
		case 'S':
			for i := 0; i < Y[p]; i++ {
				southNorth[i]++
			}
		case 'W':
			for i := 0; i < X[p]; i++ {
				westEast[i]++
			}
		default:
			for i := X[p] + 1; i <= Q; i++ {
				westEast[i]++
			}
		}
	}
	maxSN := 0
	maxWE := 0
	argMaxSN := 0
	argMaxWE := 0
	for q := 0; q <= Q; q++ {
		if southNorth[q] > maxSN {
			maxSN = southNorth[q]
			argMaxSN = q
		}
		if westEast[q] > maxWE {
			maxWE = westEast[q]
			argMaxWE = q
		}
	}
	return argMaxWE, argMaxSN
}

func main() {
	var T int
	n, err := fmt.Scanf("%d\n", &T)
	if n != 1 || err != nil {
		fmt.Printf("read %v items, err: %v\n", n, err)
		return
	}

	for t := 0; t < T; t++ {
		var P int
		var Q int
		n, err = fmt.Scanf("%d %d\n", &P, &Q)
		if n != 2 || err != nil {
			fmt.Printf("read %v items, err: %v\n", n, err)
			return
		}

		X := make([]int, P)
		Y := make([]int, P)
		D := make([]rune, P)

		for p := 0; p < P; p++ {
			n, err = fmt.Scanf("%d %d %c\n", &X[p], &Y[p], &D[p])
			if n != 3 || err != nil {
				fmt.Printf("read %v items, err: %v\n", n, err)
				return
			}
		}

		x, y := solve(P, Q, X, Y, D)

		fmt.Printf("Case #%v: %v %v\n", t+1, x, y)
	}
}
