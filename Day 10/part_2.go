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

	v_max := 0
	adapters := []int{}

	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())

		if v > v_max {
			v_max = v
		}

		adapters = append(adapters, v)
	}

	sort.Ints(adapters)

	temp := map[int]int{}
	temp[0] = 1

	for _, v := range adapters {
		s := 0

		for i := 1; i <= 3; i++ {
			if value, ok := temp[v-i]; ok {
				s += value
			}
		}

		temp[v] = s
	}

	fmt.Println(temp[v_max])
}
