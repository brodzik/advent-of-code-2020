package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	cards1 := []int{}
	cards2 := []int{}
	load1 := true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "Player 1:" {
			load1 = true
		} else if line == "Player 2:" {
			load1 = false
		} else if line != "" {
			temp, _ := strconv.Atoi(scanner.Text())

			if load1 {
				cards1 = append(cards1, temp)
			} else {
				cards2 = append(cards2, temp)
			}
		}
	}

	for len(cards1) > 0 && len(cards2) > 0 {
		if cards1[0] > cards2[0] {
			cards1 = append(cards1, cards1[0], cards2[0])
		} else {
			cards2 = append(cards2, cards2[0], cards1[0])
		}

		cards1 = cards1[1:]
		cards2 = cards2[1:]
	}

	sum := 0

	if len(cards1) > len(cards2) {
		for i, v := range cards1 {
			sum += v * (len(cards1) - i)
		}
	} else {
		for i, v := range cards2 {
			sum += v * (len(cards2) - i)
		}
	}

	fmt.Println(sum)
}
