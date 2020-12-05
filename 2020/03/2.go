package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("03/input.txt")
	if err != nil {
		panic(err)
	}

	trees := make([][]bool, 0)
	for _, line := range strings.Split(string(input), "\n") {
		row := make([]bool, 0)
		for _, tile := range strings.Split(line, "") {
			if tile == "." {
				row = append(row, false)
			} else if tile == "#" {
				row = append(row, true)
			} else {
				panic("Unknown tile: " + tile)
			}
		}
		trees = append(trees, row)
	}

	maxCols := len(trees[0])

	type Slope struct {
		Row int
		Col int
	}

	slopes := []Slope{
		{Row: 1, Col: 1},
		{Row: 1, Col: 3},
		{Row: 1, Col: 5},
		{Row: 1, Col: 7},
		{Row: 2, Col: 1},
	}

	product := 1
	for _, slope := range slopes {
		numTrees := 0
		currRow := 0
		currCol := 0

		for currRow < len(trees) {
			if trees[currRow][currCol] {
				numTrees++
			}
			currRow += slope.Row
			currCol = (currCol + slope.Col) % maxCols
		}
		product *= numTrees
	}
	fmt.Println(product)
}
