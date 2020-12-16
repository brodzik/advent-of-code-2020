package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	a1 int
	a2 int
	b1 int
	b2 int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	ranges := map[string]Range{}
	ticket := []int{}
	nearby := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "your ticket:" {
			scanner.Scan()
			for _, v := range strings.Split(scanner.Text(), ",") {
				temp, _ := strconv.Atoi(v)
				ticket = append(ticket, temp)
			}
		} else if line == "nearby tickets:" {
			for scanner.Scan() {
				t := []int{}

				for _, v := range strings.Split(scanner.Text(), ",") {
					temp, _ := strconv.Atoi(v)
					t = append(t, temp)
				}

				nearby = append(nearby, t)
			}
		} else {
			if line != "" {
				sp := strings.Split(line, ": ")
				key := sp[0]
				vals := strings.Split(sp[1], " or ")

				r1 := strings.Split(vals[0], "-")
				a1, _ := strconv.Atoi(r1[0])
				a2, _ := strconv.Atoi(r1[1])

				r2 := strings.Split(vals[1], "-")
				b1, _ := strconv.Atoi(r2[0])
				b2, _ := strconv.Atoi(r2[1])

				ranges[key] = Range{a1, a2, b1, b2}
			}
		}
	}

	sum := 0

	for _, x := range nearby {
		for _, y := range x {
			ok := false

			for _, r := range ranges {
				if (y >= r.a1 && y <= r.a2) || (y >= r.b1 && y <= r.b2) {
					ok = true
					break
				}
			}

			if !ok {
				sum += y
			}
		}
	}

	fmt.Println(sum)
}
