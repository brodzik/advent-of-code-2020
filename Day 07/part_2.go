package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search(bags *map[string]map[string]int, k string, v int, subbags map[string]int) int {
	sum := 0

	for k1, v1 := range subbags {
		sum += v*v1 + search(bags, k1, v*v1, (*bags)[k1])
	}

	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	bags := map[string]map[string]int{}

	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, " bags contain ")

		key := temp[0]
		values := map[string]int{}

		if temp[1] != "no other bags." {
			for _, v := range strings.Split(temp[1], ", ") {
				w := strings.Split(v, " ")
				x, _ := strconv.Atoi(w[0])
				values[w[1]+" "+w[2]] = x
			}
		}

		bags[key] = values
	}

	fmt.Println(search(&bags, "shiny gold", 1, bags["shiny gold"]))
}
