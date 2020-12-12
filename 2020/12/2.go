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
	bytes, err := ioutil.ReadFile("2020/12/input.txt")
	if err != nil {
		panic(err)
	}

	pos := image.Point{}
	way := image.Point{X: 10, Y: 1}
	for _, instruction := range strings.Fields(string(bytes)) {
		move := string(instruction[0])
		num, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic(err)
		}

		switch move {
		case "N":
			way.Y += num
		case "S":
			way.Y -= num
		case "W":
			way.X -= num
		case "E":
			way.X += num
		case "F":
			pos.X += num * way.X
			pos.Y += num * way.Y
			num--
		case "L":
			for num > 0 {
				way = left(way)
				num -= 90
			}
		case "R":
			for num > 0 {
				way = right(way)
				num -= 90
			}
		default:
			panic("Invalid state")
		}
	}

	fmt.Println(int(math.Abs(float64(pos.X)) + math.Abs(float64(pos.Y))))
}

func left(pos image.Point) image.Point {
	return image.Point{X: -pos.Y, Y: pos.X}
}

func right(pos image.Point) image.Point {
	return image.Point{X: pos.Y, Y: -pos.X}
}
