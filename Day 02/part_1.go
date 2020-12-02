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

	valid := 0

	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, " ")

		limits := strings.Split(s[0], "-")
		min, _ := strconv.Atoi(limits[0])
		max, _ := strconv.Atoi(limits[1])

		token := strings.Split(s[1], ":")[0][0]

		text := s[2]

		count := 0

		for _, v := range text {
			if byte(v) == token {
				count += 1
			}
		}

		if count >= min && count <= max {
			valid += 1
		}
	}

	fmt.Println(valid)
}
