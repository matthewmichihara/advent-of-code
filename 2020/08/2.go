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

	for _, program := range generatePrograms(instructions) {
		accumulator, loop := getAccumulator(program)
		if !loop {
			fmt.Println(accumulator)
		}
	}
}

func getAccumulator(program []string) (accumulator int, loop bool) {
	accumulator = 0
	loop = false
	currentInstruction := 0
	visitedInstructions := make(map[int]bool)
	for {
		if currentInstruction == len(program) {
			return
		}

		line := program[currentInstruction]
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
			loop = true
			return
		}

		visitedInstructions[currentInstruction] = true
	}
}

func generatePrograms(program []string) (programs [][]string) {
	programs = make([][]string, 0)
	for i, instruction := range program {
		switch instruction[:3] {
		case "jmp":
			newProgram := make([]string, len(program))
			copy(newProgram, program)
			newLine := strings.Replace(newProgram[i], "jmp", "nop", 1)
			newProgram[i] = newLine
			programs = append(programs, newProgram)
			continue
		case "nop":
			newProgram := make([]string, len(program))
			copy(newProgram, program)
			newLine := strings.Replace(newProgram[i], "nop", "jmp", 1)
			newProgram[i] = newLine
			programs = append(programs, newProgram)
			continue
		}
	}
	return
}
