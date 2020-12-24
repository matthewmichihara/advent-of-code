package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split("589174263", "")
	cups := make([]int, len(input))

	minCup, maxCup := math.MaxInt32, math.MinInt32
	for i, s := range input {
		cup, _ := strconv.Atoi(s)

		if cup < minCup {
			minCup = cup
		}
		if cup > maxCup {
			maxCup = cup
		}

		cups[i] = cup
	}

	for i := 0; i < 100; i++ {
		fmt.Println(cups)

		curr, pick := cups[0], cups[1:4]
		dupe := make([]int, len(cups))
		copy(dupe, cups)
		cups = append(dupe[:1], dupe[4:]...)
		fmt.Printf("curr: %v, pick: %v, cups: %v\n", curr, pick, cups)

		dest := curr - 1
		for {
			found := false
			for _, p := range pick {
				if dest == p {
					found = true
				}
			}
			if !found {
				break
			}
			dest--
		}

		if dest < minCup {
			max := math.MinInt32
			for _, cup := range cups {
				if cup > max {
					max = cup
				}
			}
			dest = max
		}

		destIndex := -1
		for i, cup := range cups {
			if cup == dest {
				destIndex = i
				break
			}
		}

		fmt.Printf("dest: %v\n", dest)
		next := make([]int, 0)
		for _, cup := range cups[:destIndex+1] {
			next = append(next, cup)
		}
		for _, cup := range pick {
			next = append(next, cup)
		}
		for _, cup := range cups[destIndex+1:] {
			next = append(next, cup)
		}

		cups = append(next[1:], next[0])
	}

	fmt.Println(cups)

	oneIndex := -1
	for i, cup := range cups {
		if cup == 1 {
			oneIndex = i
			break
		}
	}

	cups = append(cups[oneIndex+1:], cups[:oneIndex]...)
	for _, cup := range cups {
		fmt.Print(cup)
	}
	fmt.Println()
}
