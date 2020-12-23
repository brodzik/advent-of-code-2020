package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func remove(X []int, idx int) []int {
	return append(X[:idx], X[idx+1:]...)
}

func insert(X []int, idx int, val int) []int {
	if len(X) == idx {
		return append(X, val)
	}

	X = append(X[:idx+1], X[idx:]...)
	X[idx] = val

	return X
}

func find(X []int, val int) int {
	for i, v := range X {
		if v == val {
			return i
		}
	}

	return -1
}

func contains(X []int, val int) bool {
	return find(X, val) != -1
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	x_max := 0
	X := []int{}

	for _, v := range scanner.Text() {
		x := int(v - 48)
		X = append(X, x)

		if x > x_max {
			x_max = x
		}
	}

	current := X[0]

	for i := 0; i < 100; i++ {
		idx := find(X, current)
		picked := []int{X[(idx+1)%len(X)], X[(idx+2)%len(X)], X[(idx+3)%len(X)]}

		X = remove(X, find(X, picked[0]))
		X = remove(X, find(X, picked[1]))
		X = remove(X, find(X, picked[2]))

		dest := current - 1
		if dest == 0 {
			dest = x_max
		}

		for contains(picked, dest) {
			dest = (dest - 1 + x_max) % x_max
			if dest == 0 {
				dest = x_max
			}
		}

		dest_idx := find(X, dest)

		X = insert(X, dest_idx+1, picked[0])
		X = insert(X, dest_idx+2, picked[1])
		X = insert(X, dest_idx+3, picked[2])

		current = X[(find(X, current)+1)%len(X)]
	}

	idx := find(X, 1)
	res := ""

	for i := 0; i < len(X)-1; i++ {
		res += strconv.Itoa(X[(idx+1+i)%len(X)])
	}

	fmt.Println(res)
}
