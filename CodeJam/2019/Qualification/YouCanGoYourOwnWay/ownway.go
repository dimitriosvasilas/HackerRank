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
	in := strings.Split(line, "\n")

	T, _ := strconv.ParseInt(in[0], 10, 0)
	for i := 0; i < int(T); i++ {
		line, _ := reader.ReadString('\n')
		line, _ = reader.ReadString('\n')
		s := strings.Split(line, "\n")

		path := ""
		for _, c := range s[0] {
			if string(c) == "E" {
				path = path + "S"
			} else {
				path = path + "E"
			}
		}
		fmt.Printf("Case #%v: %v\n", i+1, path)
	}
}
