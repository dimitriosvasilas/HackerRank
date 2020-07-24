package main

import "fmt"

func readInput() int {
	var in int
	fmt.Scanln(&in)
	return in
}

func write(seq []bool, query string) {
	fmt.Printf(query)
	for _, b := range seq {
		fmt.Printf(" ")
		if b {
			fmt.Printf("1")
		} else {
			fmt.Printf("0")
		}
	}
	fmt.Printf("\n")
}

func changeBit(seq []bool, pos int) {
	seq[pos] = !seq[pos]
}
func main() {
	T := readInput()
	for i := 0; i < T; i++ {
		length := readInput()
		seq := make([]bool, length)

		write(seq, "Q")
		prevResponse := readInput()
		if prevResponse <= length/2 {
			for i := range seq {
				seq[i] = true
			}
			prevResponse = length - prevResponse
		}
		currentResponse := 0
		for i := range seq {
			changeBit(seq, i)
			write(seq, "Q")
			currentResponse = readInput()
			if currentResponse == length {
				write(seq, "A")
				break
			}
			if currentResponse <= prevResponse {
				changeBit(seq, i)
			}
			prevResponse = currentResponse
		}
	}
}
