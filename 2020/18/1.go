package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type State struct {
	total int
	op    string
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/18/input.txt")

	sum := 0
	for _, equation := range strings.Split(string(bytes), "\n") {
		stack := make([]*State, 0)
		state := &State{0, ""}
		stack = append(stack, state)
		state = stack[len(stack)-1]

		for _, c := range strings.ReplaceAll(equation, " ", "") {
			switch c {
			case '(':
				state = &State{0, ""}
				stack = append(stack, state)
			case ')':
				stack = stack[:len(stack)-1]
				newState := stack[len(stack)-1]
				switch newState.op {
				case "+":
					newState.total += state.total
				case "*":
					newState.total *= state.total
				case "":
					newState.total = state.total
				default:
					panic("illegal state")
				}
				state = newState
			case '+':
				state.op = "+"
			case '*':
				state.op = "*"
			default:
				num, _ := strconv.Atoi(string(c))
				switch state.op {
				case "+":
					state.total += num
				case "*":
					state.total *= num
				case "":
					state.total = num
				default:
					panic("illegal state")
				}
			}
		}
		sum += stack[0].total
	}

	fmt.Println(sum)
}
