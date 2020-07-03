package main

import (
	"fmt"
	"log"
	"strconv"

	"../utils"
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

func main() {
	readF, err := utils.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	wire0 := make([]string, 0)
	wire1 := make([]string, 0)
	input, err := readF.ReadAndParseLines(1, ",", true, "")
	for _, dir := range input[0] {
		wire0 = append(wire0, dir.(string))
	}
	input, err = readF.ReadAndParseLines(1, ",", true, "")
	for _, dir := range input[0] {
		wire1 = append(wire1, dir.(string))
	}

	// wire0 = []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	// wire1 = []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	segmentsA := directionsToSegments(wire0)
	segmentsB := directionsToSegments(wire1)
	intersections := make([]intersection, 0)
	for _, pA := range segmentsA {
		for _, pB := range segmentsB {
			intersections = append(intersections, intersectionPoints(pA, pB)...)
		}
	}

	min := -1
	for _, p := range intersections {
		if p.point.x != 0 && p.point.y != 0 {
			distance := manhattanDist(p.point)
			if min == -1 || distance < min {
				min = distance
			}
		}
	}

	fmt.Println(min)

	min = -1
	for _, intersection := range intersections {
		steps := 0
		for _, segment := range segmentsA {
			if equalSegments(segment, intersection.segment0) || equalSegments(segment, intersection.segment1) {
				steps += distance(segment.start, intersection.point)
				break
			}
			steps += segment.distance
		}
		for _, segment := range segmentsB {
			if equalSegments(segment, intersection.segment0) || equalSegments(segment, intersection.segment1) {
				steps += distance(segment.start, intersection.point)
				break
			}
			steps += segment.distance
		}
		if min == -1 || steps < min {
			min = steps
		}
	}
	fmt.Println(min)
}
