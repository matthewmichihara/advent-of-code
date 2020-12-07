package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("2020/04/input.txt")
	if err != nil {
		panic(err)
	}

	numValid := 0
	for _, passport := range strings.Split(string(input), "\n\n") {
		fields := strings.Fields(passport)

		fieldMap := make(map[string]string)
		for _, field := range fields {
			split := strings.Split(field, ":")
			fieldMap[split[0]] = split[1]
		}

		birthYear, err := strconv.Atoi(fieldMap["byr"])
		if err != nil {
			continue
		}
		if birthYear < 1920 || birthYear > 2002 {
			continue
		}

		issueYear, err := strconv.Atoi(fieldMap["iyr"])
		if err != nil {
			continue
		}
		if issueYear < 2010 || issueYear > 2020 {
			continue
		}

		expirationYear, err := strconv.Atoi(fieldMap["eyr"])
		if err != nil {
			continue
		}
		if expirationYear < 2020 || expirationYear > 2030 {
			continue
		}

		height := fieldMap["hgt"]
		heightRegexp := regexp.MustCompile(`^(\d+)(in|cm)$`)
		match := heightRegexp.FindStringSubmatch(height)
		if len(match) != 3 {
			continue
		}

		digit, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}
		if match[2] == "cm" {
			if digit < 150 || digit > 193 {
				continue
			}
		} else if match[2] == "in" {
			if digit < 59 || digit > 76 {
				continue
			}
		} else {
			continue
		}

		hairColor := fieldMap["hcl"]
		hairColorRegExp := regexp.MustCompile(`^#[a-z0-9]{6}$`)
		if !hairColorRegExp.MatchString(hairColor) {
			continue
		}

		eyeColor := fieldMap["ecl"]
		eyeColorRegexp := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
		if !eyeColorRegexp.MatchString(eyeColor) {
			continue
		}

		passportId := fieldMap["pid"]
		passportIdRegexp := regexp.MustCompile(`^[0-9]{9}$`)
		if !passportIdRegexp.MatchString(passportId) {
			continue
		}

		numValid++
	}

	fmt.Println(numValid)
}
