package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MaxInt = int(^uint(0) >> 1)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	data := map[string]map[string]bool{}
	frequency := map[string]int{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")

		sp := strings.Split(line, " contains ")

		for _, x := range strings.Split(sp[0], " ") {
			if _, ok := frequency[x]; ok {
				frequency[x] += 1
			} else {
				frequency[x] = 1
			}
		}

		for _, x := range strings.Split(sp[1], " ") {
			if _, ok1 := data[x]; !ok1 {
				data[x] = map[string]bool{}

				for _, y := range strings.Split(sp[0], " ") {
					data[x][y] = true
				}
			} else {
				temp := map[string]bool{}

				for _, y := range strings.Split(sp[0], " ") {
					temp[y] = true

					if _, ok2 := data[x][y]; !ok2 {
						delete(data[x], y)
					}
				}

				for y, _ := range data[x] {
					if _, ok2 := temp[y]; !ok2 {
						delete(data[x], y)
					}
				}
			}
		}
	}

	res := map[string]string{}

	for len(data) > 0 {
		k_best := ""
		len_best := MaxInt

		for k, v := range data {
			if len(v) < len_best {
				k_best = k
				len_best = len(v)
			}
		}

		v_best := data[k_best]
		delete(data, k_best)

		for k, v := range data {
			for x, _ := range v_best {
				if _, ok := v[x]; ok {
					delete(data[k], x)
				}
			}
		}

		temp := ""
		for k, _ := range v_best {
			temp = k
			break
		}

		res[temp] = k_best
	}

	sum := 0

	for k, v := range frequency {
		if _, ok := res[k]; !ok {
			sum += v
		}
	}

	fmt.Println(sum)
}
