package main

import (
	"fmt"
	"image"
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
				n := next1(floor, r, c)
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

func next1(floor [][]string, r int, c int) string {
	pos := floor[r][c]
	switch pos {
	case ".":
		return "."
	case "L":
		for _, v := range adjacent1(floor, r, c) {
			if v == "#" {
				return "L"
			}
		}
		return "#"
	case "#":
		count := 0
		for _, v := range adjacent1(floor, r, c) {
			if v == "#" {
				count++
			}
			if count >= 4 {
				return "L"
			}
		}
		return "#"
	default:
		panic("Invalid state")
	}
}

func adjacent1(floor [][]string, r int, c int) []string {
	points := []image.Point{
		{r - 1, c - 1},
		{r - 1, c},
		{r - 1, c + 1},
		{r, c - 1},
		{r, c + 1},
		{r + 1, c - 1},
		{r + 1, c},
		{r + 1, c + 1},
	}

	adj := make([]string, 0)
	for _, p := range points {
		if p.X < 0 || p.X >= len(floor) {
			continue
		}
		if p.Y < 0 || p.Y >= len(floor[0]) {
			continue
		}
		adj = append(adj, floor[p.X][p.Y])
	}

	return adj
}
