package main

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := strings.Split("589174263", "")
	cups := list.New()
	cupsMap := make(map[int]*list.Element)
	minCup, maxCup := math.MaxInt32, math.MinInt32
	for _, s := range input {
		cup, _ := strconv.Atoi(s)

		if cup < minCup {
			minCup = cup
		}
		if cup > maxCup {
			maxCup = cup
		}

		cupsMap[cup] = cups.PushBack(cup)
	}

	for cup := 10; cup <= 1000000; cup++ {
		if cup < minCup {
			minCup = cup
		}
		if cup > maxCup {
			maxCup = cup
		}

		cupsMap[cup] = cups.PushBack(cup)

		//if cup%1000 == 0 {
		//	fmt.Printf("Processed %v\n", cup)
		//}
	}

	for i := 0; i < 10000000; i++ {
		//if i%1000 == 0 {
		//	fmt.Printf("Iteration %v\n", i)
		//}

		curr := cups.Front()
		dest := curr.Value.(int) - 1

		for {
			if dest < minCup {
				dest = maxCup
			}

			found := false
			pick := curr
			for j := 0; j < 3; j++ {
				pick = pick.Next()
				if dest == pick.Value.(int) {
					found = true
				}
			}
			if !found {
				break
			}
			dest--
		}

		destEl := cupsMap[dest]
		// Move pick after dest
		for j := 0; j < 3; j++ {
			pick := curr.Next()
			cups.MoveAfter(pick, destEl)
			destEl = pick
		}

		cups.MoveToBack(curr)
	}

	first := cupsMap[1].Next().Value.(int)
	second := cupsMap[1].Next().Next().Value.(int)
	fmt.Println(first * second)
	fmt.Printf("Took %v\n", time.Since(start))
}

func p(l *list.List) {
	s := ""
	for e := l.Front(); e != nil; e = e.Next() {
		n := e.Value.(int)
		x := strconv.Itoa(n)
		s += x + " "
	}

	fmt.Println(s)
}
