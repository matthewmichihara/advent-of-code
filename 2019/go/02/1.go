package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("2019/go/02/input.txt")
	if err != nil {
		panic(err)
	}

	program := make([]int, 0)
	for _, s := range strings.Split(string(bytes), ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		program = append(program, i)
	}

	program[1] = 12
	program[2] = 2

	pos := 0
	for {
		switch program[pos] {
		case 1:
			firstPos := program[pos+1]
			secondPos := program[pos+2]
			thirdPos := program[pos+3]
			program[thirdPos] = program[firstPos] + program[secondPos]
			pos += 4
		case 2:
			firstPos := program[pos+1]
			secondPos := program[pos+2]
			thirdPos := program[pos+3]
			program[thirdPos] = program[firstPos] * program[secondPos]
			pos += 4
		case 99:
			fmt.Println(program[0])
			return
		}
	}
}
