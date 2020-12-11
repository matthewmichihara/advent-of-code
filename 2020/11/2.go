package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("2020/11/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Fields(string(bytes))
	floor := make([][]string, len(lines))
	for r, line := range lines {
		floor[r] = strings.Split(line, "")
	}

	for {
		changed := false
		occupied := 0
		nextFloor := make([][]string, len(floor))
		for r, row := range floor {
			nextFloor[r] = make([]string, len(row))
			for c := range row {
				n := next2(floor, r, c)
				nextFloor[r][c] = n

				if floor[r][c] != n {
					changed = true
				}
				if n == "#" {
					occupied++
				}
			}
		}
		if !changed {
			fmt.Println(occupied)
			return
		}
		floor = nextFloor
	}
}

func next2(floor [][]string, r int, c int) string {
	pos := floor[r][c]
	switch pos {
	case ".":
		return "."
	case "L":
		for _, v := range adjacent2(floor, r, c) {
			if v == "#" {
				return "L"
			}
		}
		return "#"
	case "#":
		count := 0
		for _, v := range adjacent2(floor, r, c) {
			if v == "#" {
				count++
			}
			if count >= 5 {
				return "L"
			}
		}
		return "#"
	default:
		panic("Invalid state")
	}
}

func adjacent2(floor [][]string, r int, c int) []string {
	adj := make([]string, 0)

	// Up
	for i := r - 1; i >= 0; i-- {
		if floor[i][c] != "." {
			adj = append(adj, floor[i][c])
			break
		}
	}

	// Down
	for i := r + 1; i < len(floor); i++ {
		if floor[i][c] != "." {
			adj = append(adj, floor[i][c])
			break
		}
	}

	// Left
	for i := c - 1; i >= 0; i-- {
		if floor[r][i] != "." {
			adj = append(adj, floor[r][i])
			break
		}
	}

	// Right
	for i := c + 1; i < len(floor[r]); i++ {
		if floor[r][i] != "." {
			adj = append(adj, floor[r][i])
			break
		}
	}

	// Top left
	for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if floor[i][j] != "." {
			adj = append(adj, floor[i][j])
			break
		}
	}

	// Top right
	for i, j := r-1, c+1; i >= 0 && j < len(floor[i]); i, j = i-1, j+1 {
		if floor[i][j] != "." {
			adj = append(adj, floor[i][j])
			break
		}
	}

	// Bottom left
	for i, j := r+1, c-1; i < len(floor) && j >= 0; i, j = i+1, j-1 {
		if floor[i][j] != "." {
			adj = append(adj, floor[i][j])
			break
		}
	}

	// Bottom right
	for i, j := r+1, c+1; i < len(floor) && j < len(floor[i]); i, j = i+1, j+1 {
		if floor[i][j] != "." {
			adj = append(adj, floor[i][j])
			break
		}
	}

	return adj
}
