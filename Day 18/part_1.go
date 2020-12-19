package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func eval(s string) int {
	if !strings.Contains(s, "(") {
		return eval_simple(s)
	}

	temp1 := []int{}
	temp2 := map[int]int{}

	for i, x := range s {
		if byte(x) == '(' {
			temp1 = append(temp1, i)
		} else if byte(x) == ')' {
			n := len(temp1) - 1
			temp1, temp2[temp1[n]] = temp1[:n], i
		}
	}

	a, b := 0, 0

	for a, b = range temp2 {
		break
	}

	return eval(s[:a] + strconv.Itoa(eval(s[a+1:b])) + s[b+1:])
}

func eval_simple(s string) int {
	sp := strings.Split(s, " ")
	res, _ := strconv.Atoi(sp[0])

	for i := 1; i < len(sp); i += 2 {
		temp, _ := strconv.Atoi(sp[i+1])

		if sp[i] == "+" {
			res += temp
		} else if sp[i] == "*" {
			res *= temp
		}
	}

	return res
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += eval(line)
	}

	fmt.Println(sum)
}
