package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("2020/06/input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, group := range strings.Split(string(input), "\n\n") {
		yesAnswers := make(map[string]bool)
		for _, answers := range strings.Split(group, "\n") {
			for _, answer := range strings.Split(answers, "") {
				yesAnswers[answer] = true
			}
		}
		sum += len(yesAnswers)
	}

	fmt.Println(sum)
}
