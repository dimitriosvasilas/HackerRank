package main

import (
	"fmt"
	"strconv"
)

type point struct {
	x int
	y int
}

type lineSegment struct {
	start       point
	end         point
	orientation bool
	distance    int
}

type intersection struct {
	point    point
	segment0 lineSegment
	segment1 lineSegment
}

func equalPoints(a, b point) bool {
	return a.x == b.x && a.y == b.y
}

func equalSegments(a, b lineSegment) bool {
	return equalPoints(a.start, b.start) && equalPoints(a.end, b.end)
}

func (l *lineSegment) setOrientation() {
	if l.start.x == l.end.x {
		l.orientation = true
	} else {
		l.orientation = false
	}
}

func distance(a, b point) int {
	if a.x == b.x {
		if b.y > a.y {
			return abs(b.y - a.y)
		} else {
			return abs(a.y - b.y)
		}
	} else {
		if b.x > a.x {
			return abs(b.x - a.x)
		} else {
			return abs(a.x - b.x)
		}
	}
}
func (l *lineSegment) setDistance() {
	l.distance = distance(l.start, l.end)
}

func directionsToSegments(directions []string) []lineSegment {
	segments := make([]lineSegment, 0)
	start := point{x: 0, y: 0}
	for _, d := range directions {
		nSegment := buildSegment(start, d)
		nSegment.setOrientation()
		nSegment.setDistance()
		segments = append(segments, nSegment)
		start = nSegment.end
	}
	return segments
}

func buildSegment(start point, direction string) lineSegment {
	result := lineSegment{start: start}
	end := start
	i, _ := strconv.ParseInt(direction[1:], 10, 64)
	offset := int(i)
	switch direction[0] {
	case 'U':
		end.x = start.x - offset
	case 'D':
		end.x = start.x + offset
	case 'L':
		end.y = start.y - offset
	case 'R':
		end.y = start.y + offset
	}
	result.end = end
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDist(p point) int {
	return abs(p.x) + abs(p.y)
}

func intersectionPoints(a, b lineSegment) []intersection {
	intersections := make([]intersection, 0)
	if equalPoints(a.start, point{x: 0, y: 0}) && equalPoints(b.start, point{x: 0, y: 0}) {
		return intersections
	}
	if a.orientation == b.orientation {
		return intersections
	}
	var horizontal lineSegment
	var vertical lineSegment
	if a.orientation {
		horizontal = a
		vertical = b
	} else {
		horizontal = b
		vertical = a
	}
	if horizontal.start.y > horizontal.end.y {
		temp := horizontal.start
		horizontal.start = horizontal.end
		horizontal.end = temp

	}
	if vertical.start.x > vertical.end.x {
		temp := vertical.start
		vertical.start = vertical.end
		vertical.end = temp
	}
	if horizontal.start.y <= vertical.start.y && horizontal.end.y >= vertical.start.y {
		if vertical.start.x <= horizontal.start.x && vertical.end.x >= horizontal.start.x {
			intersections = append(intersections, intersection{
				point:    point{x: horizontal.start.x, y: vertical.start.y},
				segment0: a,
				segment1: b,
			})
		}
	}
	return intersections
}

func sameAdjacent(str string) bool {
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			return true
		}
	}
	return false
}

func sameAdjacentTwo(str string) bool {
	countAdj := 1
	prev := str[0]
	for i := 1; i < len(str); i++ {
		// fmt.Print("prev:", prev, "str[i]", str[i])
		if prev == str[i] {
			countAdj++
			// fmt.Println("countAdj++", countAdj)
		} else {
			// fmt.Println("countAdj not++", countAdj)
			if countAdj == 2 {
				return true
			}
			countAdj = 1
		}
		prev = str[i]
	}
	// fmt.Println("end: countAdj", countAdj)
	return countAdj == 2
}

func neverDecrease(str string) bool {
	// fmt.Println(str)
	s := '_'
	for _, c := range str {
		if s == '_' {
			s = c
		}
		// fmt.Println("s: ", s, "c:", c)
		if c < s {
			return false
		}
		s = c
	}
	return true
}

func main() {
	count := 0
	for i := 172851; i <= 675869; i++ {
		str := strconv.Itoa(i)
		if sameAdjacent(str) && neverDecrease(str) {
			count++
		}
	}
	fmt.Println(count)

	count = 0
	for i := 172851; i <= 675869; i++ {
		str := strconv.Itoa(i)
		if sameAdjacentTwo(str) && neverDecrease(str) {
			count++
		}
	}
	fmt.Println(sameAdjacentTwo("4441"))
	fmt.Println(sameAdjacentTwo("1444"))
	fmt.Println(sameAdjacentTwo("144433"))
	fmt.Println(count)
}
