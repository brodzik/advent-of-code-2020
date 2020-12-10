package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	adapters := []int{}

	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, v)
	}

	sort.Ints(adapters)

	prev_v := 0

	diff1 := 0
	diff3 := 1

	for _, v := range adapters {
		diff := v - prev_v

		if diff == 1 {
			diff1 += 1
		} else if diff == 3 {
			diff3 += 1
		}

		prev_v = v
	}

	fmt.Println(diff1 * diff3)
}
