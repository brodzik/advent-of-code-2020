package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	count := 0

	data := make(map[byte]int)
	group := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			for _, v := range data {
				if v == group {
					count += 1
				}
			}

			data = make(map[byte]int)
			group = 0
		} else {
			for _, v := range line {
				_, exists := data[byte(v)]

				if exists {
					data[byte(v)] += 1
				} else {
					data[byte(v)] = 1
				}
			}

			group += 1
		}
	}

	for _, v := range data {
		if v == group {
			count += 1
		}
	}

	fmt.Println(count)
}
