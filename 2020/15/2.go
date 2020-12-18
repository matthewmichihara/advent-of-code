package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "16,12,1,0,15,7,11"
	nums := make(map[int]int)
	for i, s := range strings.Split(input, ",") {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		nums[num] = i + 1
	}

	next := 0
	for i := len(nums) + 1; i < 30000000; i++ {
		oldNext := next
		if nums[next] == 0 {
			next = 0
		} else {
			next = i - nums[next]
		}
		nums[oldNext] = i
	}

	fmt.Println(next)
}
