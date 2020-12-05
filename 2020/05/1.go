package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	seats, err := ioutil.ReadFile("05/input.txt")
	if err != nil {
		panic(err)
	}

	maxSeatId := 0
	for _, seat := range strings.Split(string(seats), "\n") {
		minRow, maxRow := 0, 128
		minCol, maxCol := 0, 8
		for _, direction := range strings.Split(seat, "") {
			switch direction {
			case "F":
				maxRow -= (maxRow - minRow) / 2
			case "B":
				minRow += (maxRow - minRow) / 2
			case "L":
				maxCol -= (maxCol - minCol) / 2
			case "R":
				minCol += (maxCol - minCol) / 2
			}
		}

		seatId := minRow * 8 + minCol
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}

	fmt.Println(maxSeatId)
}
