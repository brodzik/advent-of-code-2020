package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func power(x int, y int, p int) int {
	res := 1
	x = x % p

	if x == 0 {
		return 0
	}

	for y > 0 {
		if y%2 == 1 {
			res = (res * x) % p
		}

		y = y / 2
		x = (x * x) % p
	}

	return res
}

func discreteLogarithm(a int, b int, m int) int {
	n := int(math.Sqrt(float64(m)) + 1)

	an := power(a, n, m)

	value := map[int]int{}

	cur := an
	for i := 1; i <= n; i++ {
		if _, ok := value[cur]; !ok {
			value[cur] = i
		}

		cur = (cur * an) % m
	}

	cur = b
	for i := 0; i <= n; i++ {
		if v, ok := value[cur]; ok {
			ans := v*n - i

			if ans < m {
				return ans
			}
		}

		cur = (cur * a) % m
	}

	return -1
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	card_public_key, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	door_public_key, _ := strconv.Atoi(scanner.Text())

	card_loop := discreteLogarithm(7, card_public_key, 20201227)
	door_loop := discreteLogarithm(7, door_public_key, 20201227)

	fmt.Println(power(door_public_key, card_loop, 20201227))
	fmt.Println(power(card_public_key, door_loop, 20201227))
}
