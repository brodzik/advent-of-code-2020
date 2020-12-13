package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxInt = int(^uint(0) >> 1)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	time, _ := strconv.Atoi(scanner.Text())

	best_diff := MaxInt
	best_bus := 0

	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), ",") {
		if v != "x" {
			v, _ := strconv.Atoi(v)
			diff := v - (time % v)

			if diff < best_diff {
				best_diff = diff
				best_bus = v
			}
		}
	}

	fmt.Println(best_bus * best_diff)
}
