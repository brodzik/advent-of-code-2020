package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const sum = 2020

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	set := make(map[int]bool)

	for scanner.Scan() {
		line := scanner.Text()
		x, _ := strconv.Atoi(line)

		if set[sum-x] {
			fmt.Println(x * (sum - x))
			return
		}

		set[x] = true
	}
}
