package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("2019/go/01/input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, line := range strings.Fields(string(bytes)) {
		mass, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		fuel := (mass / 3) - 2
		for fuel > 0 {
			sum += fuel
			fuel = (fuel / 3) - 2
		}
	}

	fmt.Println(sum)
}
