package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("2020/13/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Fields(string(bytes))
	est, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}

	bestTime := math.MaxInt32
	bestId := 0
	for _, id := range strings.Split(lines[1], ",") {
		if id == "x" {
			continue
		}
		time, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		c := time - (est % time)
		if c < bestTime {
			bestTime = c
			bestId, err = strconv.Atoi(id)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println(bestTime * bestId)
}
