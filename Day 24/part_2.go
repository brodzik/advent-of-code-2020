package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func count_black_adjacent(p Point, flipped map[Point]bool) int {
	count := 0

	if _, ok := flipped[Point{p.x + 1, p.y + 1}]; ok {
		count++
	}

	if _, ok := flipped[Point{p.x - 1, p.y + 1}]; ok {
		count++
	}

	if _, ok := flipped[Point{p.x - 1, p.y - 1}]; ok {
		count++
	}

	if _, ok := flipped[Point{p.x + 1, p.y - 1}]; ok {
		count++
	}

	if _, ok := flipped[Point{p.x + 2, p.y}]; ok {
		count++
	}

	if _, ok := flipped[Point{p.x - 2, p.y}]; ok {
		count++
	}

	return count
}

func get_white_adjacent(p Point, flipped map[Point]bool) []Point {
	temp := []Point{}

	if _, ok := flipped[Point{p.x + 1, p.y + 1}]; !ok {
		temp = append(temp, Point{p.x + 1, p.y + 1})
	}

	if _, ok := flipped[Point{p.x - 1, p.y + 1}]; !ok {
		temp = append(temp, Point{p.x - 1, p.y + 1})
	}

	if _, ok := flipped[Point{p.x - 1, p.y - 1}]; !ok {
		temp = append(temp, Point{p.x - 1, p.y - 1})
	}

	if _, ok := flipped[Point{p.x + 1, p.y - 1}]; !ok {
		temp = append(temp, Point{p.x + 1, p.y - 1})
	}

	if _, ok := flipped[Point{p.x + 2, p.y}]; !ok {
		temp = append(temp, Point{p.x + 2, p.y})
	}

	if _, ok := flipped[Point{p.x - 2, p.y}]; !ok {
		temp = append(temp, Point{p.x - 2, p.y})
	}

	return temp
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	flipped := map[Point]bool{}

	for scanner.Scan() {
		line := scanner.Text()
		x, y := 0, 0

		for i := 0; i < len(line); {
			if byte(line[i]) == 'n' {
				if byte(line[i+1]) == 'e' {
					x += 1
					y += 1
				} else if byte(line[i+1]) == 'w' {
					x += -1
					y += 1
				}

				i += 2
			} else if byte(line[i]) == 'e' {
				x += 2
				i++
			} else if byte(line[i]) == 's' {
				if byte(line[i+1]) == 'e' {
					x += 1
					y += -1
				} else if byte(line[i+1]) == 'w' {
					x += -1
					y += -1
				}

				i += 2
			} else if byte(line[i]) == 'w' {
				x -= 2
				i++
			}
		}

		p := Point{x, y}

		if _, ok := flipped[p]; ok {
			delete(flipped, p)
		} else {
			flipped[p] = true
		}
	}

	for i := 0; i < 100; i++ {
		temp := map[Point]bool{}

		for k, _ := range flipped {
			black_adj := count_black_adjacent(k, flipped)

			if black_adj == 0 || black_adj > 2 {
				// flip to white
			} else {
				// flip to black
				temp[k] = true
			}

			for _, w := range get_white_adjacent(k, flipped) {
				if count_black_adjacent(w, flipped) == 2 {
					// flip to black
					temp[w] = true
				}
			}
		}

		flipped = temp
	}

	fmt.Println(len(flipped))
}
