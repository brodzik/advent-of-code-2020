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
		a, _ := strconv.Atoi(limits[0])
		b, _ := strconv.Atoi(limits[1])

		token := strings.Split(s[1], ":")[0][0]

		text := s[2]

		if (text[a-1] == token && text[b-1] != token) || (text[a-1] != token && text[b-1] == token) {
			valid += 1
		}
	}

	fmt.Println(valid)
}
