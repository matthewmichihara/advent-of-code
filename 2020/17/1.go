package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Point struct {
	X, Y, Z int
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/17/input.txt")
	points := make(map[Point]bool)
	for y, line := range strings.Split(string(bytes), "\n") {
		for x, state := range line {
			point := Point{x, y, 0}
			switch state {
			case '#':
				points[point] = true
			case '.':
				points[point] = false
			default:
				panic("invalid state")
			}
		}
	}

	for i := 0; i < 6; i++ {
		// Make sure space is big enough to include inactive points at the border.
		nextPoints := make(map[Point]bool)
		for point, state := range points {
			nextPoints[point] = state
		}

		for point := range nextPoints {
			neighbors := neighbors(point)
			for _, neighbor := range neighbors {
				if !points[neighbor] {
					points[neighbor] = false
				}
			}
		}

		nextPoints = make(map[Point]bool)
		for point, active := range points {
			neighbors := neighbors(point)
			activeCount := 0
			for _, neighbor := range neighbors {
				if points[neighbor] {
					activeCount++
				}
			}
			if active {
				if activeCount == 2 || activeCount == 3 {
					nextPoints[point] = true
				} else {
					nextPoints[point] = false
				}
			} else {
				if activeCount == 3 {
					nextPoints[point] = true
				} else {
					nextPoints[point] = false
				}
			}
		}

		points = nextPoints
	}

	count := 0
	for point := range points {
		if points[point] {
			count++
		}
	}
	fmt.Println(count)
}

func neighbors(p Point) []Point {
	neighbors := make([]Point, 0)
	for x := p.X - 1; x <= p.X+1; x++ {
		for y := p.Y - 1; y <= p.Y+1; y++ {
			for z := p.Z - 1; z <= p.Z+1; z++ {
				if x == p.X && y == p.Y && z == p.Z {
					continue
				}
				neighbors = append(neighbors, Point{x, y, z})
			}
		}
	}
	return neighbors
}
