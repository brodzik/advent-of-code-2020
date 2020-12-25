package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	data [][]int
}

func (t *Tile) Print() {
	for x := 0; x < len(t.data); x++ {
		line := ""

		for y := 0; y < len(t.data); y++ {
			if t.data[x][y] == 1 {
				line += "#"
			} else {
				line += "."
			}
		}

		fmt.Println(line)
	}

	fmt.Println("")
}

func (t *Tile) Rotate() {
	N := len(t.data)

	for x := 0; x < N/2; x++ {
		for y := x; y < N-x-1; y++ {
			temp := t.data[x][y]

			t.data[x][y] = t.data[y][N-1-x]
			t.data[y][N-1-x] = t.data[N-1-x][N-1-y]
			t.data[N-1-x][N-1-y] = t.data[N-1-y][x]
			t.data[N-1-y][x] = temp
		}
	}
}

func (t *Tile) FlipVertical() {
	N := len(t.data)

	for x := 0; x < N/2; x++ {
		for y := 0; y < N; y++ {
			temp := t.data[x][y]

			t.data[x][y] = t.data[N-1-x][y]
			t.data[N-1-x][y] = temp
		}
	}
}

func (t *Tile) FlipHorizontal() {
	N := len(t.data)

	for x := 0; x < N; x++ {
		for y := 0; y < N/2; y++ {
			temp := t.data[x][y]

			t.data[x][y] = t.data[x][N-1-y]
			t.data[x][N-1-y] = temp
		}
	}
}

func (t *Tile) GetEdges() []string {
	N := len(t.data)
	res := make([]string, 4)

	for i := 0; i < N; i++ {
		res[0] = res[0] + strconv.Itoa(t.data[0][i])
		res[1] = res[1] + strconv.Itoa(t.data[i][0])
		res[2] = res[2] + strconv.Itoa(t.data[N-1][i])
		res[3] = res[3] + strconv.Itoa(t.data[i][N-1])
	}

	return res
}

func Reverse(s string) string {
	res := ""

	for _, v := range s {
		res = string(v) + res
	}

	return res
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	tiles := map[int]Tile{}

	temp_id := 0
	temp_data := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Tile") {
			sp := strings.Split(line, " ")
			t, _ := strconv.Atoi(sp[1][:len(sp[1])-1])

			temp_id = t
		} else if line == "" {
			tiles[temp_id] = Tile{temp_data}
			temp_data = [][]int{}
		} else {
			t := []int{}

			for _, v := range line {
				if byte(v) == '#' {
					t = append(t, 1)
				} else {
					t = append(t, 0)
				}
			}

			temp_data = append(temp_data, t)
		}
	}

	tiles[temp_id] = Tile{temp_data}

	edges := map[string][]int{}

	for k, v := range tiles {
		for _, e := range v.GetEdges() {
			rev := Reverse(e)

			if _, ok := edges[e]; ok {
				edges[e] = append(edges[e], k)
			} else if _, ok := edges[rev]; ok {
				edges[rev] = append(edges[rev], k)
			} else {
				edges[e] = []int{k}
			}
		}
	}

	graph := map[int][]int{}

	for _, v := range edges {
		if len(v) == 2 {
			if _, ok := graph[v[0]]; ok {
				graph[v[0]] = append(graph[v[0]], v[1])
			} else {
				graph[v[0]] = []int{v[1]}
			}

			if _, ok := graph[v[1]]; ok {
				graph[v[1]] = append(graph[v[1]], v[0])
			} else {
				graph[v[1]] = []int{v[0]}
			}

		}
	}

	res := 1

	for k, v := range graph {
		if len(v) == 2 {
			res *= k
		}
	}

	fmt.Println(res)
}
