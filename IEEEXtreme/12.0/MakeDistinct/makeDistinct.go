package main

import (
	"fmt"
	"sort"
)

func readN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := scan(&all[i]); m != 1 {
		panic(err)
	}
	readN(all, i+1, n-1)
}

func scan(a *int) (int, error) {
	return fmt.Scan(a)
}

func main() {
	var N int
	fmt.Scanln(&N)
	arr := make([]int, N)
	readN(arr, 0, N)
	sort.Ints(arr)
	fmt.Println(arr)
	sum := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			j := i
			for j < len(arr) && arr[j] <= arr[j-1] {
				arr[j] = arr[j] + 1
				sum++
				j++
			}
		}
	}
	fmt.Print(sum)
}
