package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Name           string
	R1, R2, R3, R4 int
}

func main() {
	bytes, err := ioutil.ReadFile("2020/16/input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(bytes), "\n\n")
	rulesInput, yourTicketInput, nearbyTicketsInput := parts[0], parts[1], parts[2]
	valid := make(map[int]bool)
	rules := make([]Rule, 0)
	r := regexp.MustCompile(`(.+): (\d+)-(\d+) or (\d+)-(\d+)`)
	for _, line := range strings.Split(rulesInput, "\n") {
		matches := r.FindStringSubmatch(line)
		r1, _ := strconv.Atoi(matches[2])
		r2, _ := strconv.Atoi(matches[3])
		r3, _ := strconv.Atoi(matches[4])
		r4, _ := strconv.Atoi(matches[5])

		for i := r1; i <= r2; i++ {
			valid[i] = true
		}

		for i := r3; i <= r4; i++ {
			valid[i] = true
		}

		rules = append(rules, Rule{matches[1], r1, r2, r3, r4})
	}

	candidates := make(map[Rule]map[int]struct{})
	for _, rule := range rules {
		positions := make(map[int]struct{})
		for i := range rules {
			positions[i] = struct{}{}
		}
		candidates[rule] = positions
	}

	for _, line := range strings.Split(nearbyTicketsInput, "\n")[1:] {
		for i, s := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(s)
			if !valid[n] {
				break
			}

			for rule, positions := range candidates {
				if n >= rule.R1 && n <= rule.R2 || n >= rule.R3 && n <= rule.R4 {
					continue
				}
				delete(positions, i)
			}
		}
	}

	// Field name to index
	fields := make(map[string]int)
	for range candidates {
		for rule, positions := range candidates {
			if len(positions) == 1 {
				for i := range positions {
					fields[rule.Name] = i
				}

				// Now remove this index from all other candidate positions.
				for _, positions2 := range candidates {
					delete(positions2, fields[rule.Name])
				}

				break
			}
		}
	}

	departureIndexes := make(map[int]bool)
	for f, i := range fields {
		if strings.HasPrefix(f, "departure") {
			departureIndexes[i] = true
		}
	}

	product := 1
	for _, line := range strings.Split(yourTicketInput, "\n")[1:] {
		for i, s := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(s)
			if departureIndexes[i] {
				product *= n
			}
		}
	}

	fmt.Println(product)
}
