package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	name string
	arg1 int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	ins := []instruction{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(line[1])
		ins = append(ins, instruction{line[0], val})
	}

	visited := map[int]bool{}

	i := 0
	acc := 0

	for {
		_, v := visited[i]

		if v {
			break
		} else {
			visited[i] = true
		}

		if ins[i].name == "nop" {
			i += 1
		} else if ins[i].name == "acc" {
			acc += ins[i].arg1
			i += 1
		} else if ins[i].name == "jmp" {
			i += ins[i].arg1
		}
	}

	fmt.Println(acc)
}
