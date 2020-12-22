package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func hash(cards1 []int, cards2 []int) string {
	s := ""

	for _, v := range cards1 {
		s += strconv.Itoa(v) + "_"
	}

	s += "___"

	for _, v := range cards2 {
		s += strconv.Itoa(v) + "_"
	}

	return s
}

func recursive_combat(cards1 []int, cards2 []int) int {
	log := map[string]bool{}

	for true {
		if len(cards1) == 0 {
			return 2
		} else if len(cards2) == 0 {
			return 1
		}

		h := hash(cards1, cards2)

		if _, ok := log[h]; ok {
			return 1
		} else {
			log[h] = true
		}

		card1 := cards1[0]
		cards1 = cards1[1:]

		card2 := cards2[0]
		cards2 = cards2[1:]

		winner := 1

		if len(cards1) >= card1 && len(cards2) >= card2 {
			t1 := make([]int, card1)
			copy(t1, cards1[:card1])

			t2 := make([]int, card2)
			copy(t2, cards2[:card2])

			winner = recursive_combat(t1, t2)
		} else {
			if card1 > card2 {
				winner = 1
			} else {
				winner = 2
			}
		}

		if winner == 1 {
			cards1 = append(cards1, card1, card2)
		} else if winner == 2 {
			cards2 = append(cards2, card2, card1)
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

	return 1
}

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

	recursive_combat(cards1, cards2)
}
