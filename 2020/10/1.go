package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("2020/10/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Fields(string(bytes))
	jolts := make([]int, 0)
	for _, line := range lines {
		jolt, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		jolts = append(jolts, jolt)
	}

	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	prev := 0
	one, three := 0, 0
	for _, jolt := range jolts {
		diff := jolt - prev
		switch diff {
		case 1:
			one++
		case 3:
			three++
		}
		prev = jolt
	}

	fmt.Println(one * three)
}
