package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getBit(x int, n int) int {
	return (x >> n) & 1
}

func setBit(x int, n int) int {
	return x | 1<<n
}

func unsetBit(x int, n int) int {
	return x & ^(1 << n)
}

func check(data *map[int]int, key int, val int, mask string, i int) {
	if i < len(mask) {
		if byte(mask[i]) == '0' {
			check(data, key, val, mask, i+1)
		} else if byte(mask[i]) == '1' {
			check(data, setBit(key, len(mask)-i-1), val, mask, i+1)
		} else if byte(mask[i]) == 'X' {
			check(data, unsetBit(key, len(mask)-i-1), val, mask, i+1)
			check(data, setBit(key, len(mask)-i-1), val, mask, i+1)
		}
	} else {
		(*data)[key] = val
	}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	data := map[int]int{}
	mask := ""

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " = ")

		if line[0] == "mask" {
			mask = line[1]
		} else {
			temp := strings.ReplaceAll(line[0], "mem[", "")

			key, _ := strconv.Atoi(temp[:len(temp)-1])
			val, _ := strconv.Atoi(line[1])

			check(&data, key, val, mask, 0)
		}
	}

	sum := 0

	for _, v := range data {
		sum += v
	}

	fmt.Println(sum)
}
