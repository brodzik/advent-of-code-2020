package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const TARGET = 18272118
const MaxInt = int(^uint(0) >> 1)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	data := []int{}

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		data = append(data, num)
	}

	sum := 0
	visited := map[int]int{}
	a, b := 0, 0

	for i := 0; i < len(data); i++ {
		sum += data[i]

		if sum == TARGET && i > 0 {
			a = 0
			b = i
			break
		}

		j, exists := visited[sum-TARGET]

		if exists && i-j > 0 {
			a = j
			b = i
			break
		}

		visited[sum] = i + 1
	}

	min := MaxInt
	max := 0

	for i := a; i < b; i++ {
		if data[i] < min {
			min = data[i]
		}

		if data[i] > max {
			max = data[i]
		}
	}

	fmt.Println(min + max)
}
