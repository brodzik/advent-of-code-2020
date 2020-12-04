package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hasKey(passport map[string]string, key string) bool {
	_, ok := passport[key]
	return ok
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	data := ""
	valid := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			passport := make(map[string]string)
			for _, field := range strings.Split(data[:len(data)-1], " ") {
				s := strings.Split(field, ":")
				passport[s[0]] = s[1]
			}

			if hasKey(passport, "byr") && hasKey(passport, "iyr") && hasKey(passport, "eyr") && hasKey(passport, "hgt") && hasKey(passport, "hcl") && hasKey(passport, "ecl") && hasKey(passport, "pid") {
				valid += 1
			}

			data = ""
		} else {
			data += line + " "
		}
	}

	fmt.Println(valid)
}
