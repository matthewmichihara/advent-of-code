package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("2020/16/input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(bytes), "\n\n")
	rulesInput, nearbyTicketsInput := parts[0], parts[2]

	valid := make(map[int]bool)
	r := regexp.MustCompile(`.+: (\d+)-(\d+) or (\d+)-(\d+)`)
	for _, line := range strings.Split(rulesInput, "\n") {
		matches := r.FindStringSubmatch(line)
		r1, _ := strconv.Atoi(matches[1])
		r2, _ := strconv.Atoi(matches[2])
		r3, _ := strconv.Atoi(matches[3])
		r4, _ := strconv.Atoi(matches[4])

		for i := r1; i <= r2; i++{
			valid[i] = true
		}

		for i := r3; i <= r4; i++ {
			valid[i] = true
		}
	}

	ticketScanningErrorRate := 0

	for _, line := range strings.Split(nearbyTicketsInput, "\n") {
		for _, s := range strings.Split(line , ",") {
			n, _ := strconv.Atoi(s)
			if !valid[n] {
				ticketScanningErrorRate += n
				break
			}
		}
	}

	fmt.Println(ticketScanningErrorRate)
}

