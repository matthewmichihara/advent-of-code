package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("2020/02/input.txt")
	if err != nil {
		panic(err)
	}

	numValid := 0
	for _, line := range strings.Split(string(input), "\n") {
		var min, max int
		var letter byte
		var password string

		_, err := fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &letter, &password)
		if err != nil {
			panic(err)
		}

		count := strings.Count(password, string(letter))
		if count >= min && count <= max {
			numValid++
		}
	}

	fmt.Println(numValid)
}
