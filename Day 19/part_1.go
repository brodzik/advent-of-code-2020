package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func match1(rule []string, rules map[string][][]string, s string) bool {
	for i, r := range rule {
		if r[0] == '"' {
			if len(s) > 0 && r[1] == s[0] {
				s = s[1:]
			} else {
				return false
			}
		} else {
			temp := [][]string{}

			for _, x := range rules[r] {
				for _, z := range rule[i+1:] {
					x = append(x, z)
				}

				temp = append(temp, x)
			}

			return match(temp, rules, s)
		}
	}

	return len(s) == 0
}

func match(rule [][]string, rules map[string][][]string, s string) bool {
	for _, r := range rule {
		if match1(r, rules, s) {
			return true
		}
	}

	return false
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	rules := map[string][][]string{}
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, ":") {
			temp1 := strings.Split(line, ": ")

			subrules := [][]string{}

			for _, x := range strings.Split(temp1[1], " | ") {
				subrule := []string{}

				for _, y := range strings.Split(x, " ") {
					subrule = append(subrule, y)
				}

				subrules = append(subrules, subrule)
			}

			rules[temp1[0]] = subrules
		} else {
			if line != "" {
				if match(rules["0"], rules, line) {
					sum += 1
				}
			}
		}
	}

	fmt.Println(sum)
}
