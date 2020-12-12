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
	dir := 90

	for scanner.Scan() {
		line := scanner.Text()
		action := byte(line[0])
		arg, _ := strconv.Atoi(line[1:])

		if action == 'N' {
			y += arg
		} else if action == 'S' {
			y -= arg
		} else if action == 'E' {
			x += arg
		} else if action == 'W' {
			x -= arg
		} else if action == 'L' {
			dir = (dir + 360 - arg) % 360
		} else if action == 'R' {
			dir = (dir + arg) % 360
		} else if action == 'F' {
			if dir == 0 {
				y += arg
			} else if dir == 90 {
				x += arg
			} else if dir == 180 {
				y -= arg
			} else if dir == 270 {
				x -= arg
			}
		}
	}

	fmt.Println(Abs(x) + Abs(y))
}
