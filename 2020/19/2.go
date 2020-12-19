package main

import (
	fmt "fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	val     string
	options [][]int
}

func main() {
	bytes, err := ioutil.ReadFile("2020/19/input.txt")
	if err != nil {
		panic(err)
	}

	inputParts := strings.Split(string(bytes), "\n\n")
	ruleInput, messageInput := inputParts[0], inputParts[1]

	r := regexp.MustCompile(`"(.+)"`)
	rules := make(map[int]Rule)
	for _, line := range strings.Split(ruleInput, "\n") {
		lineParts := strings.Split(line, ": ")
		var rule Rule
		options := make([][]int, 0)
		ruleId, _ := strconv.Atoi(lineParts[0])
		matches := r.FindStringSubmatch(lineParts[1])

		if matches != nil {
			val := matches[1]
			rule = Rule{val, options}
		} else {
			for _, optionParts := range strings.Split(lineParts[1], "|") {
				parts := make([]int, 0)
				for _, part := range strings.Fields(optionParts) {
					n, _ := strconv.Atoi(part)
					parts = append(parts, n)
				}
				options = append(options, parts)
			}
			rule = Rule{"", options}
		}
		rules[ruleId] = rule
	}
	memo := make(map[int][]string)

	validMessages := make(map[string]struct{})
	for _, message := range gen(rules, 0, memo) {
		validMessages[message] = struct{}{}
	}

	count := 0
	messages := strings.Split(messageInput, "\n")
	for _, message := range messages {
		_, exists := validMessages[message]
		if exists {
			count++
		}
	}

	// part 1
	fmt.Println(count)

	memo = make(map[int][]string)
	valid42Messages := gen(rules, 42, memo)

	memo = make(map[int][]string)
	valid31Messages := gen(rules, 31, memo)

	regexFirst := strings.Join(valid42Messages, `|`)
	regexSecond := strings.Join(valid31Messages, `|`)
	re := regexp.MustCompile(`^((?:` + regexFirst + `)+)((?:` + regexSecond + `)+)$`)

	count2 := 0
	for _, message := range messages {
		matches := re.FindStringSubmatch(message)
		if matches == nil {
			continue
		} else if len(matches[1]) > len(matches[2]) {
			count2++
		} else {
			continue
		}
	}

	// part 2
	fmt.Println(count2)
}

func gen(rules map[int]Rule, ruleId int, memo map[int][]string) []string {
	_, exists := memo[ruleId]
	if exists {
		return memo[ruleId]
	}

	rs := make([]string, 0)
	rule := rules[ruleId]
	if rule.val != "" {
		rs = append(rs, rule.val)
		memo[ruleId] = rs
		return rs
	}

	for _, option := range rule.options {
		ors := []string{""}
		for _, optionRuleId := range option {
			nextOrs := make([]string, 0)
			for _, or := range ors {
				for _, c := range gen(rules, optionRuleId, memo) {
					nextOrs = append(nextOrs, or+c)
				}
			}
			ors = nextOrs
		}

		for _, or := range ors {
			rs = append(rs, or)
		}
	}

	memo[ruleId] = rs
	return rs
}
