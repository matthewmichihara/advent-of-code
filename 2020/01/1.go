package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("2020/01/input.txt")
	if err != nil {
		panic(err)
	}

	entries := make([]int, 0)
	for _, line := range strings.Split(string(input), "\n") {
		entry, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		entries = append(entries, entry)
	}

	for i, first := range entries {
		for _, second := range entries[i+1:] {
			if first + second == 2020 {
				fmt.Println(first * second)
				return
			}
		}
	}
}
