package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("2020/07/input.txt")
	if err != nil {
		panic(err)
	}

	bags := make(map[string]map[string]int)
	for _, line := range strings.Split(string(input), "\n") {
		name, children := parseLine(line)
		bags[name] = children
	}

	numBags := 0
	for bag := range bags {
		if canHold(bags, bag, "shiny gold") {
			numBags++
		}
	}

	fmt.Printf("1: %v\n", numBags)
	fmt.Printf("2: %v\n", holds(bags, "shiny gold"))
}

func parseLine(line string) (name string, children map[string]int) {
	line = line[:len(line)-1] // Remove trailing .
	split := strings.Split(line, " contain ")
	name = bagType(split[0])
	children = make(map[string]int)
	if split[1] == "no other bags" {
		return
	}

	innersSplit := strings.Split(split[1], ", ")
	for _, inner := range innersSplit {
		childBagType, count := bagTypeWithCount(inner)
		children[childBagType] = count
	}
	return
}

func bagType(s string) string {
	fields := strings.Fields(s)
	return strings.Join(fields[:len(fields)-1], " ")
}

func bagTypeWithCount(s string) (bagType string, count int) {
	fields := strings.Fields(s)
	count, err := strconv.Atoi(fields[0])
	if err != nil {
		panic(err)
	}
	bagType = strings.Join(fields[1:len(fields)-1], " ")
	return
}

func canHold(bags map[string]map[string]int, outerBag string, innerBag string) bool {
	children := bags[outerBag]
	for child := range children {
		if child == innerBag {
			return true
		}

		if canHold(bags, child, innerBag) {
			return true
		}
	}
	return false
}

func holds(bags map[string]map[string]int, bag string) int {
	children := bags[bag]
	numBags := 0
	for child, childCount := range children {
		numBags += childCount + (holds(bags, child) * childCount)
	}
	return numBags
}
