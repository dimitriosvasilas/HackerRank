package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func digits(n int) []int {
	digits := make([]int, 0)
	for n != 0 {
		digits = append(digits, n%10)
		n = n / 10
	}
	return digits
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	in := strings.Split(line, "\n")

	T, _ := strconv.ParseInt(in[0], 10, 0)
	for i := 0; i < int(T); i++ {
		line, _ := reader.ReadString('\n')
		in := strings.Split(line, "\n")
		N, _ := strconv.ParseInt(in[0], 10, 0)

		digs := digits(int(N))

		a := 0
		b := 0
		for i, d := range digs {
			if d == 4 {
				a = a + (d-1)*int((math.Pow(10, float64(i))))
				b = b + int((math.Pow(10, float64(i))))
			} else {
				a = a + d*int((math.Pow(10, float64(i))))
			}
		}
		fmt.Printf("Case #%v: %v %v\n", i+1, a, b)
	}
}
