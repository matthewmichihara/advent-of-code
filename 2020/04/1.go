package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	input, err := ioutil.ReadFile("2020/04/input.txt")
	if err != nil {
		panic(err)
	}

	numValid := 0
	for _, passport := range strings.Split(string(input), "\n\n") {
		fields := strings.Fields(passport)

		fieldMap := make(map[string]bool)
		for _, field := range fields {
			split := strings.Split(field, ":")
			fieldMap[split[0]] = true
		}

		missingField := false
		for _, requiredField := range requiredFields {
			if !fieldMap[requiredField] {
				missingField = true
				break
			}
		}

		if !missingField {
			numValid++
		}
	}

	fmt.Println(numValid)
}
