package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	logA := map[int]int{}
	logB := map[int]int{}
	last := 0

	for scanner.Scan() {
		for i, x := range strings.Split(scanner.Text(), ",") {
			y, _ := strconv.Atoi(x)
			logA[y] = i + 1
			last = y
		}
	}

	for i := len(logA) + 1; i <= 30000000; i++ {
		if val, ok := logB[last]; ok {
			temp := val - logA[last]

			if _, ok1 := logA[temp]; ok1 {
				if val2, ok2 := logB[temp]; ok2 {
					logA[temp] = val2
					logB[temp] = i
				} else {
					logB[temp] = i
				}
			} else {
				logA[temp] = i
			}

			last = temp
		} else {
			if _, ok := logA[0]; ok {
				if val1, ok1 := logB[0]; ok1 {
					logA[0] = val1
					logB[0] = i
				} else {
					logB[0] = i
				}
			} else {
				logA[0] = i
			}

			last = 0
		}
	}

	fmt.Println(last)
}
