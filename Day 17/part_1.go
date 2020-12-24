package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
	z int
}

func count_adjacent_active(p Point, active map[Point]bool) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				t := Point{p.x + i, p.y + j, p.z + k}
				_, ok := active[t]

				if t != p && ok {
					count++
				}
			}
		}
	}

	return count
}

func get_adjacent_inactive(p Point, active map[Point]bool) []Point {
	temp := []Point{}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				t := Point{p.x + i, p.y + j, p.z + k}
				_, ok := active[t]

				if t != p && !ok {
					temp = append(temp, t)
				}
			}
		}
	}

	return temp
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	active := map[Point]bool{}
	x := 0

	for scanner.Scan() {
		for y, v := range scanner.Text() {
			if byte(v) == '#' {
				active[Point{x, y, 0}] = true
			}
		}

		x++
	}

	for i := 0; i < 6; i++ {
		temp := map[Point]bool{}

		for p, _ := range active {
			adj_active := count_adjacent_active(p, active)

			if adj_active == 2 || adj_active == 3 {
				// set active
				temp[p] = true
			} else {
				// set inactive
			}

			for _, a := range get_adjacent_inactive(p, active) {
				if count_adjacent_active(a, active) == 3 {
					// set active
					temp[a] = true
				}
			}
		}

		active = temp
	}

	fmt.Println(len(active))
}
