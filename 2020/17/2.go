package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Point4 struct {
	X, Y, Z, W int
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/17/input.txt")
	points := make(map[Point4]bool)
	for y, line := range strings.Split(string(bytes), "\n") {
		for x, state := range line {
			point := Point4{x, y, 0, 0}
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
		nextPoints := make(map[Point4]bool)
		for point, state := range points {
			nextPoints[point] = state
		}

		for point := range nextPoints {
			neighbors := neighbors4(point)
			for _, neighbor := range neighbors {
				if !points[neighbor] {
					points[neighbor] = false
				}
			}
		}

		nextPoints = make(map[Point4]bool)
		for point, active := range points {
			neighbors := neighbors4(point)
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

func neighbors4(p Point4) []Point4 {
	neighbors := make([]Point4, 0)
	for x := p.X - 1; x <= p.X+1; x++ {
		for y := p.Y - 1; y <= p.Y+1; y++ {
			for z := p.Z - 1; z <= p.Z+1; z++ {
				for w := p.W - 1; w <= p.W+1; w++ {
					if x == p.X && y == p.Y && z == p.Z && w == p.W {
						continue
					}
					neighbors = append(neighbors, Point4{x, y, z, w})
				}
			}
		}
	}
	return neighbors
}
