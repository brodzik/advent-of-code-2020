package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search(bags *map[string]map[string]int, k string, v map[string]int, good_bags *map[string]bool) bool {
	_, good := (*good_bags)[k]

	if k == "shiny gold" || good {
		return true
	}

	for k1, _ := range v {
		if search(bags, k1, (*bags)[k1], good_bags) {
			(*good_bags)[k] = true
			return true
		}
	}

	return false
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

	good_bags := map[string]bool{}

	for k0, v0 := range bags {
		search(&bags, k0, v0, &good_bags)
	}

	fmt.Println(len(good_bags))
}
