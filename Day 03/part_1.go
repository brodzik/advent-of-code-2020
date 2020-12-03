package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	data := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	x_max := len(data[0])
	y_max := len(data)

	x := 3
	y := 1

	trees := 0

	for y < y_max {
		if data[y][x] == '#' {
			trees += 1
		}
		x = (x + 3) % x_max
		y += 1
	}

	fmt.Println(trees)
}
