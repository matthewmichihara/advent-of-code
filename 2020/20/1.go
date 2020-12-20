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
	id         int
	t, r, b, l string
}

type Tile2 struct {
	id int
	contents []string
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

		tiles = append(tiles, Tile2 {id, lines[1:]})



		top := tileGrid[0]
		right := ""
		for i := 0; i < len(tileGrid); i++ {
			right += string(tileGrid[i][len(tileGrid[i])-1])
		}
		bottom := tileGrid[len(tileGrid)-1]
		left := ""
		for i := 0; i < len(tileGrid); i++ {
			left += string(tileGrid[i][0])
		}

		tiles = append(tiles, Tile{id, top, right, bottom, left})
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
		fmt.Println(grid[0][0].id *grid[0][size-1].id* grid[size-1][0].id*grid[size-1][size-1].id)
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
				if grid[r-1][c].b != variant.t {
					continue
				}
			}

			// left
			if c != 0 {
				if grid[r][c-1].r != variant.l {
					continue
				}
			}

			newGrid := make([][]Tile, len(grid))
			for i := range grid {
				newGrid[i] = make([]Tile, len(grid[i]))
				copy(newGrid[i], grid[i])
			}
			newGrid[r][c] = variant

			newTiles := make([]Tile,len(tiles))
			copy(newTiles, tiles)
			newTiles = append(newTiles[:tileIndex], newTiles[tileIndex+1:]...)

			if place(newTiles, newGrid, nextR, nextC, size) {
				return true
			}
		}
	}

	return false
}
//
//// original
//abc
//def
//ghi
//
//// original horizontal
//cba
//fed
//ihg
//
//// original vertical
//ghi
//def
//abc
//
//// 90
//gda
//heb
//ifc
//
//// 90 horizontal
//adg
//beh
//cfi
//
//// 90 vertical
//ifc
//heb
//gda
//
//// 180
//ihg
//fed
//cba
//
//// 270
//cfi
//beh
//adg





func variants(tile Tile) []Tile {
	verticalFlip := Tile{tile.id, tile.b, reverse(tile.r), tile.t, reverse(tile.l)}
	horizontalFlip := Tile{tile.id, reverse(tile.t), tile.l, reverse(tile.b), tile.r}
	rotate90 := Tile{tile.id, reverse(tile.l), tile.t, reverse(tile.r), tile.b}
	vertical90Flip := Tile{tile.id, rotate90.b, reverse(rotate90.r), rotate90.t, reverse(rotate90.l)}
	horizontal90Flip := Tile{tile.id, reverse(rotate90.t), rotate90.l, reverse(rotate90.b), rotate90.r}
	rotate180 := Tile{tile.id, reverse(rotate90.l), rotate90.t, reverse(rotate90.r), rotate90.b}
	rotate270 := Tile{tile.id, reverse(rotate180.l), rotate180.t, reverse(rotate180.r), rotate180.b}
	return []Tile{tile, verticalFlip, horizontalFlip, rotate90, vertical90Flip, horizontal90Flip, rotate180, rotate270}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
