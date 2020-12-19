package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// this is just disgusting
func main() {
	// Paren pass
	// split by mult
	// eval parts
	// multiply the rest

	bytes, _ := ioutil.ReadFile("2020/18/input.txt")
	sum := 0
	for _, line := range strings.Split(string(bytes), "\n") {
		equation := strings.ReplaceAll(line, " ", "")
		n, _ := strconv.Atoi(eval(equation))
		sum += n
	}
	fmt.Println(sum)
}

func eval(equation string) string {
	if len(equation) == 1 {
		return equation
	}

	// Remove paren
	newEquation1 := ""
	for i := 0; i < len(equation); i++ {
		c := equation[i]
		switch c {
		case '(':
			opens := 0
			for j := i + 1; j < len(equation); j++ {
				d := equation[j]
				switch d {
				case '(':
					opens++
				case ')':
					if opens != 0 {
						opens--
						continue
					}

					numString := eval(equation[i+1 : j])
					newEquation1 += numString
					i = j
					break
				}
			}
		default:
			newEquation1 = newEquation1 + string(c)
		}
	}
	equation = newEquation1

	// After this block, the equation should just have *s
	// 1 + 3 * 5 + 4 * 4 * 4
	// 4     * 9     * 4 * 4
	if strings.Contains(equation, "*") {
		// Split by mult: 1 * 2 + 3 * 4 * 4
		additions := strings.Split(equation, "*")
		prod := 1
		for _, addition := range additions {
			n, _ := strconv.Atoi(eval(addition))
			prod *= n
		}
		equation = strconv.Itoa(prod)
	}

	// add
	total := 0
	if strings.Contains(equation, "+") {
		for _, num := range strings.Split(equation, "+") {
			n, _ := strconv.Atoi(num)
			total += n
		}
		s := strconv.Itoa(total)
		return s
	}

	return equation
}
