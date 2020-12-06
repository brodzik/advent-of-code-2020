package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	data := make(map[byte]bool)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			count += len(data)
			data = make(map[byte]bool)
		} else {
			for _, v := range line {
				data[byte(v)] = true
			}
		}
	}

	count += len(data)

	fmt.Println(count)
}
