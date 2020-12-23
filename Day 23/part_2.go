package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	next := make([]int, 1000001)
	first := -1
	last := -1

	for _, v := range line {
		x := int(v - '0')

		if first == -1 {
			first = x
		} else {
			next[last] = x
		}

		last = x
	}

	for i := len(line) + 1; i <= 1000000; i++ {
		next[last] = i
		last = i
	}

	next[last] = first

	for i := 0; i < 10000000; i++ {
		dest := first - 1

		if dest < 1 {
			dest = 1000000
		}

		pickup1 := next[first]
		pickup2 := next[pickup1]
		pickup3 := next[pickup2]

		next[first] = next[pickup3]
		first = next[first]

		for dest == pickup1 || dest == pickup2 || dest == pickup3 {
			dest--

			if dest < 1 {
				dest = 1000000
			}
		}

		next[pickup3] = next[dest]
		next[dest] = pickup1
	}

	fmt.Println(next[1] * next[next[1]])
}
