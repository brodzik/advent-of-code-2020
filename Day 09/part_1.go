package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const OFFSET = 25

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	data := []int{}

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		data = append(data, num)
	}

	for i := OFFSET; i < len(data); i++ {
		sum := data[i]
		set := map[int]bool{}
		good := false

		for j := i - OFFSET; j < i; j++ {
			x := data[j]

			if set[sum-x] {
				good = true
				break
			}

			set[x] = true
		}

		if !good {
			fmt.Println(sum)
			break
		}
	}
}
