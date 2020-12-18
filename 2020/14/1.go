package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("2020/14/input.txt")
	if err != nil {
		panic(err)
	}

	mem := make(map[int]int64)

	var setMask, clearMask int64

	for _, line := range strings.Split(string(bytes), "\n") {
		parts := strings.Split(line, " = ")
		first, second := parts[0], parts[1]
		if first == "mask" {
			setMask = 0
			clearMask = -1

			for _, bit := range second {
				switch bit {
				case 'X':
					setMask <<= 1
					clearMask = (clearMask << 1) | 1
				case '0':
					setMask <<= 1
					clearMask <<= 1
				case '1':
					setMask = (setMask << 1) | 1
					clearMask = (clearMask << 1) | 1
				default:
					panic("invalid state")
				}
			}

			continue
		}

		var addr int
		_, err := fmt.Sscanf(first, "mem[%d]", &addr)
		if err != nil {
			panic(err)
		}

		val, err := strconv.Atoi(second)
		if err != nil {
			panic(err)
		}

		res := (int64(val) | setMask) & clearMask
		mem[addr] = res
	}

	var sum int64
	for _, val := range mem {
		sum += val
	}

	fmt.Println(sum)
}
