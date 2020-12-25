package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/24/input.txt")
	allDirs := make([][]string, 0)
	for _, line := range strings.Split(string(bytes), "\n") {
		dirs := make([]string, 0)
		for i := 0; i < len(line); i++ {
			c := string(line[i])
			if c == "s" || c == "n" {
				c = line[i : i+2]
				i += 1
			}
			dirs = append(dirs, c)
		}
		allDirs = append(allDirs, dirs)
	}

	maxDirsLength := 0
	for _, dirs := range allDirs {
		if len(dirs) > maxDirsLength {
			maxDirsLength = len(dirs)
		}
	}
	// We're going to start in the middle.
	maxDirsLength = maxDirsLength * 2 + 200
	start := image.Point{X: maxDirsLength / 2, Y: maxDirsLength / 2}
	grid := make([][]bool, maxDirsLength)
	for i := range grid {
		grid[i] = make([]bool, maxDirsLength)
	}

	for _, dirs := range allDirs {
		tile := start
		for _, dir := range dirs {
			tile = move(tile, dir)
		}
		grid[tile.X][tile.Y] = !grid[tile.X][tile.Y]
	}

	// Count
	black := 0
	for _, row := range grid {
		for _, tile := range row {
			if tile {
				black++
			}
		}
	}

	fmt.Printf("Part 1: %v\n", black)

	for i := 0; i < 100; i++ {
		nextGrid := make([][]bool, len(grid))
		for x := range grid {
			nextGrid[x] = make([]bool, len(grid[x]))
			copy(nextGrid[x], grid[x])
			for y, tile := range grid[x] {
				adjs := adj(image.Point{X: x, Y: y})
				numBlack := 0
				for _, adj := range adjs {
					if adj.X >= 0 && adj.Y >= 0 && adj.X < len(grid) && adj.Y < len(grid[x]) && grid[adj.X][adj.Y] {
						numBlack++
					}
				}
				if tile {
					if numBlack == 0 || numBlack > 2 {
						nextGrid[x][y] = false
					}
				} else {
					if numBlack == 2 {
						nextGrid[x][y] = true
					}
				}
			}
		}
		grid = nextGrid
	}

	// Count
	black = 0
	for _, row := range grid {
		for _, tile := range row {
			if tile {
				black++
			}
		}
	}
	fmt.Printf("Part 2: %v\n", black)
}

func adj(tile image.Point) []image.Point {
	return []image.Point{
		move(tile, "nw"),
		move(tile, "ne"),
		move(tile, "sw"),
		move(tile, "se"),
		move(tile, "w"),
		move(tile, "e"),
	}
}

func move(from image.Point, dir string) image.Point {
	var x, y int
	if from.X%2 == 0 {
		switch dir {
		case "nw":
			x, y = from.X-1, from.Y
		case "ne":
			x, y = from.X-1, from.Y+1
		case "sw":
			x, y = from.X+1, from.Y
		case "se":
			x, y = from.X+1, from.Y+1
		}
	} else {
		switch dir {
		case "nw":
			x, y = from.X-1, from.Y-1
		case "ne":
			x, y = from.X-1, from.Y
		case "sw":
			x, y = from.X+1, from.Y-1
		case "se":
			x, y = from.X+1, from.Y
		}
	}

	switch dir {
	case "w":
		x, y = from.X, from.Y-1
	case "e":
		x, y = from.X, from.Y+1
	}
	return image.Point{X: x, Y: y}
}
