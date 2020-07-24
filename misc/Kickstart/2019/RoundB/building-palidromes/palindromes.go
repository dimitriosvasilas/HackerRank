package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	in := line[:len(line)-1]
	T, _ := strconv.ParseInt(in, 10, 0)
	//fmt.Println(T)
	for t := 0; t < int(T); t++ {
		yesCounter := 0
		line, _ := reader.ReadString('\n')
		lineEOL := line[:len(line)-1]
		in := strings.Split(lineEOL, " ")
		//N, _ := strconv.ParseInt(in[0], 10, 0)
		Q, _ := strconv.ParseInt(in[1], 10, 0)
		//fmt.Println(N, Q)

		line, _ = reader.ReadString('\n')
		s := line[:len(line)]
		//fmt.Printf("Case #%v: %v %v %v \n", t+1, N, Q, s)

		for q := 0; q < int(Q); q++ {
			line, _ := reader.ReadString('\n')
			lineEOL := line[:len(line)-1]
			in := strings.Split(lineEOL, " ")
			L, _ := strconv.ParseInt(in[0], 10, 0)
			R, _ := strconv.ParseInt(in[1], 10, 0)
			//fmt.Printf("Case #%v: %v %v %v .. Question %v %v\n", t+1, N, Q, s, L, R)

			subS := s[(L - 1):R]
			//fmt.Println(subS)

			fr := make(map[rune]int, 0)
			for _, c := range subS {
				//fmt.Println(c)
				if _, ok := fr[c]; !ok {
					fr[c] = 1
				} else {
					fr[c]++
				}

			}
			//fmt.Println(fr)
			foundOdd := false
			palindrome := true
			for _, v := range fr {
				if v%2 == 1 {
					if !foundOdd {
						foundOdd = true
					} else {
						palindrome = false
						break
					}
				}
			}
			//fmt.Printf("Case #%v: %v %v %v .. Question %v %v .. Sub %v .. Palindrome %v\n", t+1, N, Q, s, L, R, subS, palindrome)
			if palindrome {
				yesCounter++
			}
		}
		fmt.Printf("Case #%v: %v\n", t+1, yesCounter)
	}

}
