package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	x, y := 0, 0
	dx, dy := 10, 1

	for scanner.Scan() {
		line := scanner.Text()
		action := byte(line[0])
		arg, _ := strconv.Atoi(line[1:])

		if action == 'N' {
			dy += arg
		} else if action == 'S' {
			dy -= arg
		} else if action == 'E' {
			dx += arg
		} else if action == 'W' {
			dx -= arg
		} else if action == 'L' {
			for i := 0; i < arg; i += 90 {
				dx, dy = -dy, dx
			}
		} else if action == 'R' {
			for i := 0; i < arg; i += 90 {
				dx, dy = dy, -dx
			}
		} else if action == 'F' {
			x += arg * dx
			y += arg * dy
		}
	}

	fmt.Println(Abs(x) + Abs(y))
}
