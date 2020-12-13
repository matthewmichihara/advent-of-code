package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// x === b mod a
// N = product of all as
// n = N/a
type Bus struct {
	b int64
	a int64
	n int64
	inv int64
}

func main() {
	bytes, err := ioutil.ReadFile("2020/13/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Fields(string(bytes))

	buses := make([]Bus, 0)
	N := int64(1)
	for b, id := range strings.Split(lines[1], ",") {
		if id == "x" {
			continue
		}
		a, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		N *= int64(a)
		buses = append(buses, Bus{int64(a-b), int64(a), 0, 0})
	}

	fmt.Printf("N=%v\n", N)

	for i, bus := range buses {
		bus.n = N/bus.a
		bus.inv = inverse(bus.n, bus.a)
		buses[i] = bus
	}

	sum := int64(0)
	for _, bus := range buses {
		fmt.Println(bus)

		partial := bus.b * bus.n * bus.inv
		fmt.Println(partial)
		sum += partial

	}

	fmt.Println(sum % N)
}

func inverse(n int64, a int64) int64 {
	for i := int64(1); i < n; i++ {
		if (n * i) % a == 1 {
			return i
		}
	}
	panic("could not find inverse")
}
