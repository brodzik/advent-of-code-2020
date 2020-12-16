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

	good := [][]int{}

	for _, x := range nearby {
		ok1 := true

		for _, y := range x {
			ok := false

			for _, r := range ranges {
				if (y >= r.a1 && y <= r.a2) || (y >= r.b1 && y <= r.b2) {
					ok = true
					break
				}
			}

			if !ok {
				ok1 = false
			}
		}

		if ok1 {
			good = append(good, x)
		}
	}

	good = append(good, ticket)

	tab := map[string]map[int]int{}

	for _, x := range good {
		for i, y := range x {
			for k, r := range ranges {
				if _, ok := tab[k]; !ok {
					tab[k] = map[int]int{}
				}

				if (y >= r.a1 && y <= r.a2) || (y >= r.b1 && y <= r.b2) {
					if _, ok := tab[k][i]; ok {
						tab[k][i] += 1
					} else {
						tab[k][i] = 1
					}
				}
			}
		}
	}

	for k1, v1 := range tab {
		for k2, v2 := range v1 {
			if v2 == len(good) {
				tab[k1][k2] = 1
			} else {
				delete(tab[k1], k2)
			}
		}
	}

	res := map[string]int{}

	for len(tab) > 0 {
		k_min, v_min := "", map[int]int{}

		for k, v := range tab {
			if len(v) < len(v_min) || len(v_min) == 0 {
				k_min = k
				v_min = v
			}
		}

		idx := 0

		for k, _ := range v_min {
			idx = k
			break
		}

		res[k_min] = idx

		for k, v := range tab {
			if _, ok := v[idx]; ok {
				delete(tab[k], idx)
			}

			if len(tab[k]) <= 0 {
				delete(tab, k)
			}
		}
	}

	temp := 1

	for k, v := range res {
		if strings.Contains(k, "departure") {
			temp *= ticket[v]
		}
	}

	fmt.Println(temp)
}
