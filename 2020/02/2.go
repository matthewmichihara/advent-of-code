package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("02/input.txt")
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

		if string(password[min - 1]) == string(letter) && string(password[max - 1]) == string(letter) {
			continue
		}

		if string(password[min - 1]) == string(letter) || string(password[max - 1]) == string(letter) {
			numValid++
		}
	}

	fmt.Println(numValid)
}
