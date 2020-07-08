package main

import (
	"fmt"
	"os"
)

func query(n int) {
	os.Stdout.WriteString(fmt.Sprintf("%v\n", n))
}

func result(r1, r2, r3, r4, r5, r6 int) {
	os.Stdout.WriteString(fmt.Sprintf("%v %v %v %v %v %v\n", r1, r2, r3, r4, r5, r6))
}

func main() {
	f, err := os.Create("./dbg")
	defer f.Close()
	var T int
	var W int
	n, err := fmt.Scanf("%d %d\n", &T, &W)
	if n != 2 || err != nil {
		fmt.Printf("read %v items, err: %v\n", n, err)
		return
	}
	f.WriteString(fmt.Sprintf("T: %v, W: %v\n", T, W))
	for t := 0; t < T; t++ {
		for w := 0; w < W; w++ {
			query(w + 1)

			var R int
			n, err := fmt.Scanf("%d\n", &R)
			if n != 1 || err != nil {
				fmt.Printf("read %v items, err: %v\n", n, err)
				return
			}
			//	fmt.Printf("R: %v\n", R)
			if R == -1 {
				os.Exit(0)
			}
		}
		result(1, 2, 3, 4, 5, 6)
		var outcome int
		n, err = fmt.Scanf("%d\n", &outcome)
		if n != 1 || err != nil {
			fmt.Printf("read %v items, err: %v\n", n, err)
			return
		}
		//fmt.Printf("outcome: %v\n", outcome)
		if outcome == -1 {
			os.Exit(0)
		}
	}
}
