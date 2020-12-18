package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Mask struct {
	set int64
	clr int64
}

func main() {
	bytes, err := ioutil.ReadFile("2020/14/input.txt")
	if err != nil {
		panic(err)
	}

	mem := make(map[int64]int64)
	var mask string

	for _, line := range strings.Split(string(bytes), "\n") {
		parts := strings.Split(line, " = ")
		first, second := parts[0], parts[1]
		if first == "mask" {
			mask = second
			continue
		}

		var addr int64
		_, err := fmt.Sscanf(first, "mem[%d]", &addr)
		if err != nil {
			panic(err)
		}

		val, err := strconv.Atoi(second)
		if err != nil {
			panic(err)
		}

		for _, res := range getAddrs(addr, mask) {
			mem[res] = int64(val)
		}
	}

	var sum int64
	for _, val := range mem {
		sum += val
	}

	fmt.Println(sum)
}

func getAddrs(addr int64, maskString string) []int64 {
	addrs := make([]int64, 0)
	addrs = append(addrs, addr)

	for i := len(maskString)-1; i>= 0; i-- {
		bit := maskString[i]
		shift := len(maskString) - 1 - i
		switch bit {
		case '1':
			for j := range addrs {
				addrs[j] |= 1 << shift
			}
		case 'X':
			ones := make([]int64, len(addrs))
			copy(ones, addrs)

			// add zeros
			for j := range addrs {
				addrs[j] &= (1 << shift) ^ -1
			}

			// add ones
			for j := range ones {
				ones[j] |= 1 << shift
			}

			addrs = append(addrs, ones...)
		}
	}
	return addrs
}
