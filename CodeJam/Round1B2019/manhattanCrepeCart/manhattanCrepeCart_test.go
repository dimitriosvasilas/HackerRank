package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	P := 1
	Q := 10
	X := []int{5}
	Y := []int{5}
	D := []rune{'N'}

	x, y := solve(P, Q, X, Y, D)
	assert.Equal(t, x, 0, "")
	assert.Equal(t, y, 6, "")

	P = 4
	Q = 10
	X = []int{2, 2, 1, 3}
	Y = []int{4, 6, 5, 5}
	D = []rune{'N', 'S', 'E', 'W'}

	x, y = solve(P, Q, X, Y, D)
	assert.Equal(t, x, 2, "")
	assert.Equal(t, y, 5, "")

	P = 8
	Q = 10
	X = []int{0, 0, 0, 0, 0, 0, 0, 1}
	Y = []int{2, 3, 3, 4, 5, 5, 8, 5}
	D = []rune{'S', 'N', 'N', 'N', 'S', 'S', 'S', 'W'}

	x, y = solve(P, Q, X, Y, D)
	assert.Equal(t, x, 0, "")
	assert.Equal(t, y, 4, "")
}
