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

	for j := 0; j < len(ins); j++ {
		temp := make([]instruction, len(ins))
		copy(temp, ins)

		if temp[j].name == "nop" {
			temp[j].name = "jmp"
		} else if temp[j].name == "jmp" {
			temp[j].name = "nop"
		} else {
			continue
		}

		visited := map[int]bool{}
		failed := false

		i := 0
		acc := 0

		for i < len(temp) {
			_, v := visited[i]

			if v {
				failed = true
				break
			} else {
				visited[i] = true
			}

			if temp[i].name == "nop" {
				i += 1
			} else if temp[i].name == "acc" {
				acc += temp[i].arg1
				i += 1
			} else if temp[i].name == "jmp" {
				i += temp[i].arg1
			}
		}

		if !failed {
			fmt.Println(acc)
			break
		}
	}
}
