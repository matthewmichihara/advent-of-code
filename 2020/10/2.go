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
	jolts = append(jolts, 0)
	for _, line := range lines {
		jolt, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		jolts = append(jolts, jolt)
	}
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	set := make(map[int]int64)
	fmt.Println(ways(set, jolts, 0))
}

func ways(set map[int]int64, jolts []int, index int) int64 {
	if index == len(jolts)-1 {
		return 1
	}

	if set[index] != 0 {
		return set[index]
	}

	end := index + 4
	if len(jolts) < end {
		end = len(jolts)
	}

	jolt := jolts[index]
	var sum int64 = 0
	for start := index + 1; start < end; start++ {
		if jolts[start]-jolt <= 3 {
			sum += ways(set, jolts, start)
		}
	}

	set[index] = sum
	return sum
}
