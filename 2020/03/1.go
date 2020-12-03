package main

import (
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
	rowChange := 1
	colChange := 3
	currRow := 0
	currCol := 0
	numTrees := 0

	for currRow < len(trees) {
		if trees[currRow][currCol] {
			numTrees++
		}
		currRow += rowChange
		currCol = (currCol + colChange) % maxCols
	}

	println(numTrees)
}
