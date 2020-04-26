package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stone struct {
	secsToEat  int64
	energy     int64
	lossPerSec int64
}

func remove(s []stone, i int) []stone {
	tmp := make([]stone, len(s))
	copy(tmp, s)
	tmp[len(tmp)-1], tmp[i] = tmp[i], tmp[len(tmp)-1]
	return tmp[:len(tmp)-1]
}

func max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func positiveExists(s []stone) bool {
	for _, s := range s {
		if s.energy > 0 {
			return true
		}
		s.energy = 0
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	in := line[:len(line)-1]
	T, _ := strconv.ParseInt(in, 10, 0)
	for t := 0; t < int(T); t++ {

		line, _ := reader.ReadString('\n')
		lineEOL := line[:len(line)-1]
		N, _ := strconv.ParseInt(lineEOL, 10, 0)

		stones := make([]stone, N)
		for n := 0; n < int(N); n++ {
			line, _ := reader.ReadString('\n')
			lineEOL := line[:len(line)-1]
			in := strings.Split(lineEOL, " ")
			S, _ := strconv.ParseInt(in[0], 10, 0)
			E, _ := strconv.ParseInt(in[1], 10, 0)
			L, _ := strconv.ParseInt(in[2], 10, 0)
			stones[n] = stone{secsToEat: S, energy: E, lossPerSec: L}
		}
		eaten := make([]stone, 0)
		energy := int64(0)
		for positiveExists(stones) {
			//fmt.Println(stones)
			min := int64(100000000)
			minIndex := -1
			selected := stone{}
			for i, s := range stones {
				sumLoss := int64(0)
				sumEnergy := int64(0)
				temp := remove(stones, i)

				//				fmt.Printf("Selected %v, rest %v\n", s, temp)
				for _, ss := range temp {
					sumLoss += max(ss.energy-s.secsToEat*ss.lossPerSec, 0)
					sumEnergy += ss.energy
				}
				//fmt.Printf("sumLoss %v,\n", sumLoss)
				benefit := sumEnergy - sumLoss
				if benefit < min {
					min = benefit
					minIndex = i
					selected = s
				}
			}
			eaten = append(eaten, selected)
			stones = remove(stones, minIndex)
			for i, ss := range stones {
				stones[i].energy -= selected.secsToEat * ss.lossPerSec
			}
			energy += selected.energy
			//fmt.Printf("Eaten %v\n", selected)
			//fmt.Printf("Energy %v\n", energy)

		}
		fmt.Printf("Case #%v: %v\n", t+1, energy)
	}
}
