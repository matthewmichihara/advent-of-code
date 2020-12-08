package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("2019/go/03/input.txt")
	if err != nil {
		panic(err)
	}

	paths := strings.Split(string(bytes), "\n")
	first := strings.Split(paths[0], ",")
	second := strings.Split(paths[1], ",")
	firstPoints := getPointMap(first)
	secondPoints := getPointMap(second)

	minDist := math.MaxInt32
	for point := range firstPoints {
		if secondPoints[point] {
			dist := distanceToOrigin(point)
			if dist < minDist {
				minDist = dist
			}
		}
	}

	fmt.Println(minDist)
}

func distanceToOrigin(point image.Point) int {
	return int(math.Abs(float64(point.X)) + math.Abs(float64(point.Y)))
}

func getPointMap(moves []string) (pointMap map[image.Point]bool) {
	pointMap = make(map[image.Point]bool)
	currPoint := image.Point{X: 0, Y: 0}
	for _, move := range moves {
		dir := rune(move[0])
		dist, err := strconv.Atoi(move[1:])
		if err != nil {
			panic(err)
		}

		for i := 0; i < dist; i++ {
			switch dir {
			case 'R':
				currPoint = image.Point{X: currPoint.X + 1, Y: currPoint.Y}
			case 'L':
				currPoint = image.Point{X: currPoint.X - 1, Y: currPoint.Y}
			case 'U':
				currPoint = image.Point{X: currPoint.X, Y: currPoint.Y + 1}
			case 'D':
				currPoint = image.Point{X: currPoint.X, Y: currPoint.Y - 1}
			}
			pointMap[currPoint] = true
		}
	}
	return
}
