package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("2020/08/input.txt")
	if err != nil {
		panic(err)
	}

	instructions := make([]string, 0)
	for _, line := range strings.Split(string(input), "\n") {
		instructions = append(instructions, line)
	}

	accumulator := 0
	currentInstruction := 0
	visitedInstructions := make(map[int]bool)
	for {
		line := instructions[currentInstruction]
		fields := strings.Fields(line)
		op := fields[0]
		arg := fields[1]

		sign := arg[0]
		num, err := strconv.Atoi(arg[1:])
		if err != nil {
			panic(err)
		}

		switch op {
		case "acc":
			switch sign {
			case '+':
				accumulator += num
			case '-':
				accumulator -= num
			}
			currentInstruction++
		case "jmp":
			switch sign {
			case '+':
				currentInstruction += num
			case '-':
				currentInstruction -= num
			}
		case "nop":
			currentInstruction++
		}

		if visitedInstructions[currentInstruction] {
			fmt.Println(accumulator)
			return
		}

		visitedInstructions[currentInstruction] = true
	}
}
