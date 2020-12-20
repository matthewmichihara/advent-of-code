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
	solution := place(tiles, grid, 0, 0, size)

	// Part 1
	fmt.Printf("1: %v\n",
		solution[0][0].id*
			solution[0][size-1].id*
			solution[size-1][0].id*
			solution[size-1][size-1].id,
	)

	trimmedSolution := make([][]Tile, len(solution))
	for r, row := range solution {
		trimmedSolution[r] = make([]Tile, len(row))
		for c, tile := range row {
			trimmedSolution[r][c] = tile.trim()
		}
	}
	trimmedImage := makeImage(trimmedSolution)

	numMonsters := 0
	fakeTile := Tile{0, trimmedImage}
	for _, tile := range variants(fakeTile) {
		// Take every monster grid area, and see if a monster is there
		for r, row := range tile.contents[:len(tile.contents)-3] {
			for c := range row[:len(row)-20] {
				area := make([][]rune, 3)
				for i := 0; i < 3; i++ {
					area[i] = make([]rune, 20)
					copy(area[i], tile.contents[r+i][c:c+20])
				}
				if hasMonster(area) {
					numMonsters++
				}
			}
		}
		if numMonsters > 0 {
			break
		}
	}

	monsterSize := 15
	numWaves := 0
	for _, row := range trimmedImage {
		for _, item := range row {
			if item == '#' {
				numWaves++
			}
		}
	}

	fmt.Printf("2: %v\n", numWaves-(numMonsters*monsterSize))
}

func hasMonster(area [][]rune) bool {
	monster := [][]rune{
		[]rune("..................O."),
		[]rune("O....OO....OO....OOO"),
		[]rune(".O..O..O..O..O..O..."),
	}

	for r, row := range area {
		for c, item := range row {
			if monster[r][c] == 'O' && item != '#' {
				return false
			}
		}
	}

	return true
}

func makeImage(grid [][]Tile) [][]rune {
	tileLength := len(grid[0][0].contents)
	image := make([][]rune, 0)

	for _, tileRow := range grid {
		for rowIndex := 0; rowIndex < tileLength; rowIndex++ {
			imageRow := make([]rune, 0)
			for _, tile := range tileRow {
				row := tile.contents[rowIndex]
				imageRow = append(imageRow, row...)
			}
			image = append(image, imageRow)
		}
	}

	return image
}

func place(tiles []Tile, grid [][]Tile, r int, c int, size int) [][]Tile {
	if len(tiles) == 0 {
		return grid
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

			solution := place(newTiles, newGrid, nextR, nextC, size)
			if len(solution) != 0 {
				return solution
			}
		}
	}

	return [][]Tile{}
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

func (tile Tile) trim() Tile {
	newContents := make([][]rune, len(tile.contents)-2)
	for i, row := range tile.contents[1 : len(tile.contents)-1] {
		newRow := make([]rune, len(row)-2)
		copy(newRow, row[1:len(row)-1])
		newContents[i] = newRow
	}
	return Tile{tile.id, newContents}
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
