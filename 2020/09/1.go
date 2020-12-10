package main

import (
	"fmt"
	"io/ioutil"
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

	for i := 25; i < len(nums); i++ {
		n := nums[i]
		window := nums[i-25 : i]
		if !sums(window, n) {
			fmt.Println(n)
			return
		}
	}

}

func sums(window []int, n int) bool {
	for j, a := range window {
		for _, b := range window[j+1:] {
			if a+b == n {
				return true
			}
		}
	}
	return false
}
