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

	lefts := map[string]string{
		"N": "W",
		"W": "S",
		"S": "E",
		"E": "N",
	}

	rights := map[string]string{
		"N": "E",
		"E": "S",
		"S": "W",
		"W": "N",
	}

	pos := image.Point{}
	dir := "E"
	for _, instruction := range strings.Fields(string(bytes)) {
		move := string(instruction[0])
		num, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic(err)
		}

		switch move {
		case "N":
			pos.Y += num
		case "S":
			pos.Y -= num
		case "W":
			pos.X -= num
		case "E":
			pos.X += num
		case "F":
			switch dir {
			case "N":
				pos.Y += num
			case "S":
				pos.Y -= num
			case "W":
				pos.X -= num
			case "E":
				pos.X += num
			default:
				panic("Invalid state")
			}
		case "L":
			for num > 0 {
				dir = lefts[dir]
				num -= 90
			}
		case "R":
			for num > 0 {
				dir = rights[dir]
				num -= 90
			}
		default:
			panic("Invalid state")
		}
	}

	fmt.Println(int(math.Abs(float64(pos.X)) + math.Abs(float64(pos.Y))))
}
