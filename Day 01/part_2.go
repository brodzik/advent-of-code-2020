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

	var data []int

	for scanner.Scan() {
		line := scanner.Text()
		x, _ := strconv.Atoi(line)
		data = append(data, x)
	}

	for i, x := range data {
		set := make(map[int]bool)

		for _, y := range data[i+1:] {
			if set[sum-x-y] {
				fmt.Println(x * y * (sum - x - y))
				return
			}

			set[y] = true
		}
	}
}
