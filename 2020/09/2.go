package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("2020/09/input.txt")
	if err != nil {
		panic(err)
	}

	nums := make([]int, 0)
	for _, s := range strings.Fields(string(bytes)) {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			window := nums[i:j]
			sum := 0
			min := math.MaxInt32
			max := 0
			for _, n := range window {
				sum += n
				if n < min {
					min = n
				}
				if n > max {
					max = n
				}
			}
			if sum == 466456641 {
				fmt.Println(min + max)
				return
			}
		}
	}
}
