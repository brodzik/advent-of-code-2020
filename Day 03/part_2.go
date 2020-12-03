package main

import (
	"bufio"
	"fmt"
	"os"
)

func countTrees(data []string, right int, down int) int {
	x_max := len(data[0])
	y_max := len(data)

	x := right
	y := down

	trees := 0

	for y < y_max {
		if data[y][x] == '#' {
			trees += 1
		}
		x = (x + right) % x_max
		y += down
	}

	return trees
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	data := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	fmt.Println(countTrees(data, 1, 1) * countTrees(data, 3, 1) * countTrees(data, 5, 1) * countTrees(data, 7, 1) * countTrees(data, 1, 2))
}
