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

	fmt.Println(len(flipped))
}
