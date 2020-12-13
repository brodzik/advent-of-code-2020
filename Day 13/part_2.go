package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inverse(a int, m int) int {
	m0 := m
	x0, x1 := 0, 1

	if m == 1 {
		return 0
	}

	for a > 1 {
		q := a / m
		m, a = a%m, m
		x0, x1 = x1-q*x0, x0
	}

	if x1 < 0 {
		x1 += m0
	}

	return x1
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	M := 1
	a := []int{}
	n := []int{}

	scanner.Scan()
	scanner.Scan()
	for i, v := range strings.Split(scanner.Text(), ",") {
		if v != "x" {
			v, _ := strconv.Atoi(v)

			M *= v
			a = append(a, (v-i)%v)
			n = append(n, v)
		}
	}

	res := 0
	for i := 0; i < len(a); i++ {
		pp := M / n[i]
		res += a[i] * inverse(pp, n[i]) * pp
	}

	res %= M

	fmt.Println(res)
}
