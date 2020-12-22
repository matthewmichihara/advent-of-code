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
	for len(a) > 0 && len(b) > 0 {
		if a[0] > b[0] {
			a = append(a[1:], a[0], b[0])
			b = b[1:]
		} else if a[0] < b[0] {
			b = append(b[1:], b[0], a[0])
			a = a[1:]
		} else {
			panic("Invalid state.")
		}
	}

	var winner []int
	if len(a) == 0 {
		winner = b
	} else {
		winner = a
	}

	sum := 0
	for i, card := range winner {
		sum += (len(winner) - i) * card
	}

	fmt.Println(sum)
}
