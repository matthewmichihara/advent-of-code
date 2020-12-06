package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("06/input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, group := range strings.Split(string(input), "\n\n") {
		yesAnswers := make(map[string]int)
		groupSlice := strings.Split(group, "\n")
		for _, answers := range groupSlice {
			for _, answer := range strings.Split(answers, "") {
				yesAnswers[answer] += 1
			}
		}
		for _, v := range yesAnswers {
			if v == len(groupSlice) {
				sum += 1
			}
		}
	}

	fmt.Println(sum)
}
