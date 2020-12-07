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

	originalProgram := make([]int, 0)
	for _, s := range strings.Split(string(bytes), ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		originalProgram = append(originalProgram, i)
	}

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			program := make([]int, len(originalProgram))
			copy(program, originalProgram)
			addr := 0
			program[1] = noun
			program[2] = verb
			finished := false
			for !finished {
				switch program[addr] {
				case 1:
					firstAddr := program[addr+1]
					secondAddr := program[addr+2]
					thirdAddr := program[addr+3]
					program[thirdAddr] = program[firstAddr] + program[secondAddr]
					addr += 4
				case 2:
					firstAddr := program[addr+1]
					secondAddr := program[addr+2]
					thirdAddr := program[addr+3]
					program[thirdAddr] = program[firstAddr] * program[secondAddr]
					addr += 4
				case 99:
					if program[0] == 19690720 {
						fmt.Println(100*noun + verb)
						return
					}
					finished = true
				}
			}
		}
	}
}
