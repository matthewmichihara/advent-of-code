package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/22/input.txt")

	decks := [][]int{{}, {}}
	for i, chunk := range strings.Split(string(bytes), "\n\n") {
		for _, line := range strings.Split(chunk, "\n")[1:] {
			card, _ := strconv.Atoi(line)
			decks[i] = append(decks[i], card)
		}
	}

	a, b := decks[0], decks[1]
	_, score := combat(a, b)
	fmt.Println(score)
}

func score(n []int) int {
	sum := 0
	for i, card := range n {
		sum += (len(n) - i) * card
	}
	return sum
}

func combat(ain []int, bin []int) (string, int) {
	a, b := make([]int, len(ain)), make([]int, len(bin))
	copy(a, ain)
	copy(b, bin)
	seen := make(map[string]struct{})

	for len(a) > 0 && len(b) > 0 {
		key := fmt.Sprintf("a:%v-b:%v", score(a), score(b))
		if _, exists := seen[key]; exists {
			return "a", score(a)
		}
		seen[key] = struct{}{}

		winner := ""
		if a[0] <= len(a[1:]) && b[0] <= len(b[1:]) {
			subWinner, _ := combat(a[1:a[0]+1], b[1:b[0]+1])
			winner = subWinner
		} else if a[0] > b[0] {
			winner = "a"
		} else if b[0] > a[0] {
			winner = "b"
		} else {
			panic("Invalid state.")
		}

		switch winner {
		case "a":
			a = append(a[1:], a[0], b[0])
			b = b[1:]
		case "b":
			b = append(b[1:], b[0], a[0])
			a = a[1:]
		default:
			panic("Invalid state.")
		}
	}

	if len(a) == 0 {
		return "b", score(b)
	} else if len(b) == 0 {
		return "a", score(a)
	} else {
		panic("Invalid state.")
	}
}
