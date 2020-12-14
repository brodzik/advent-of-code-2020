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

			for i, v := range mask {
				if byte(v) == '0' {
					val = unsetBit(val, len(mask)-i-1)
				} else if byte(v) == '1' {
					val = setBit(val, len(mask)-i-1)
				}
			}

			data[key] = val
		}
	}

	sum := 0

	for _, v := range data {
		sum += v
	}

	fmt.Println(sum)
}
