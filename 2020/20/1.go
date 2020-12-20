package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Tile struct {
	id       int
	contents [][]rune
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/20/input.txt")
	r := regexp.MustCompile(`^Tile (\d+):$`)
	tiles := make([]Tile, 0)
	for _, chunk := range strings.Split(string(bytes), "\n\n") {
		lines := strings.Split(strings.TrimSpace(chunk), "\n")
		idLine := lines[0]
		matches := r.FindStringSubmatch(idLine)
		id, _ := strconv.Atoi(matches[1])
		tileGrid := lines[1:]
		contents := make([][]rune, len(tileGrid))
		for i, line := range tileGrid {
			contents[i] = []rune(line)
		}
		tiles = append(tiles, Tile{id, contents})
	}

	size := int(math.Sqrt(float64(len(tiles))))
	grid := make([][]Tile, size)
	for i := 0; i < size; i++ {
		row := make([]Tile, size)
		grid[i] = row
	}
	place(tiles, grid, 0, 0, size)
}

func place(tiles []Tile, grid [][]Tile, r int, c int, size int) bool {
	if len(tiles) == 0 {
		fmt.Println(grid[0][0].id * grid[0][size-1].id * grid[size-1][0].id * grid[size-1][size-1].id)
		return true
	}

	nextR, nextC := r, c+1
	if nextC == size {
		nextR++
		nextC = 0
	}

	for tileIndex, tile := range tiles {
		for _, variant := range variants(tile) {
			// check if can place tile
			// top
			if r != 0 {
				if grid[r-1][c].bottom() != variant.top() {
					continue
				}
			}

			// left
			if c != 0 {
				if grid[r][c-1].right() != variant.left() {
					continue
				}
			}

			newGrid := make([][]Tile, len(grid))
			for i := range grid {
				newGrid[i] = make([]Tile, len(grid[i]))
				copy(newGrid[i], grid[i])
			}
			newGrid[r][c] = variant

			newTiles := make([]Tile, len(tiles))
			copy(newTiles, tiles)
			newTiles = append(newTiles[:tileIndex], newTiles[tileIndex+1:]...)

			if place(newTiles, newGrid, nextR, nextC, size) {
				return true
			}
		}
	}

	return false
}

func (tile Tile) top() string {
	return string(tile.contents[0])
}

func (tile Tile) bottom() string {
	return string(tile.contents[len(tile.contents)-1])
}

func (tile Tile) left() string {
	ret := make([]rune, len(tile.contents))
	for i, row := range tile.contents {
		ret[i] = row[0]
	}
	return string(ret)
}

func (tile Tile) right() string {
	ret := make([]rune, len(tile.contents))
	for i, row := range tile.contents {
		ret[i] = row[len(row)-1]
	}
	return string(ret)
}

func variants(tile Tile) []Tile {
	verticalFlip := Tile{tile.id, vertical(tile.contents)}
	horizontalFlip := Tile{tile.id, horizontal(tile.contents)}
	rotate90 := Tile{tile.id, rot90(tile.contents)}
	verticalRotate90 := Tile{tile.id, vertical(rotate90.contents)}
	horizontalRotate90 := Tile{tile.id, horizontal(rotate90.contents)}
	rotate180 := Tile{tile.id, rot90(rotate90.contents)}
	rotate270 := Tile{tile.id, rot90(rotate180.contents)}
	return []Tile{
		tile,
		verticalFlip,
		horizontalFlip,
		rotate90,
		verticalRotate90,
		horizontalRotate90,
		rotate180,
		rotate270,
	}
}

func rot90(contents [][]rune) [][]rune {
	newContents := make([][]rune, len(contents))
	for r, row := range contents {
		newContents[r] = make([]rune, len(row))
		copy(newContents[r], row)
	}

	for r, row := range contents {
		for c := range row {
			newContents[r][c] = contents[len(contents)-c-1][r]
		}
	}

	return newContents
}

func vertical(contents [][]rune) [][]rune {
	newContents := make([][]rune, len(contents))
	for i := range contents {
		newContents[i] = contents[len(contents)-1-i]
	}
	return newContents
}

func horizontal(contents [][]rune) [][]rune {
	newContents := make([][]rune, len(contents))
	for i, row := range contents {
		newContents[i] = reverse(row)
	}
	return newContents
}

func reverse(r []rune) []rune {
	runes := make([]rune, len(r))
	copy(runes, r)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return runes
}
