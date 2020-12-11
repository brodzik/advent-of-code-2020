package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid struct {
	data   [][]byte
	x_max  int
	y_max  int
	hashes map[string]bool
}

func (g *Grid) AddRow(line string) {
	row := []byte{}

	for _, v := range line {
		row = append(row, byte(v))
	}

	g.data = append(g.data, row)

	g.x_max = len(row)
	g.y_max += 1
}

func (g *Grid) CountOccupied() int {
	count := 0

	for y := 0; y < g.y_max; y++ {
		for x := 0; x < g.x_max; x++ {
			if g.data[y][x] == '#' {
				count += 1
			}
		}
	}

	return count
}

func (g *Grid) CountAdjacentOccupied(x int, y int) int {
	count := 0

	if x-1 >= 0 && y-1 >= 0 && g.data[y-1][x-1] == '#' {
		count += 1
	}

	if y-1 >= 0 && g.data[y-1][x] == '#' {
		count += 1
	}

	if x+1 < g.x_max && y-1 >= 0 && g.data[y-1][x+1] == '#' {
		count += 1
	}

	if x-1 >= 0 && g.data[y][x-1] == '#' {
		count += 1
	}

	if x+1 < g.x_max && g.data[y][x+1] == '#' {
		count += 1
	}

	if x-1 >= 0 && y+1 < g.y_max && g.data[y+1][x-1] == '#' {
		count += 1
	}

	if y+1 < g.y_max && g.data[y+1][x] == '#' {
		count += 1
	}

	if x+1 < g.x_max && y+1 < g.y_max && g.data[y+1][x+1] == '#' {
		count += 1
	}

	return count
}

func (g *Grid) Step() bool {
	temp := make([][]byte, g.y_max)
	for y := 0; y < g.y_max; y++ {
		temp[y] = make([]byte, g.x_max)
		copy(temp[y], g.data[y])
	}

	for y := 0; y < g.y_max; y++ {
		for x := 0; x < g.x_max; x++ {
			if g.data[y][x] == 'L' && g.CountAdjacentOccupied(x, y) == 0 {
				temp[y][x] = '#'
			} else if g.data[y][x] == '#' && g.CountAdjacentOccupied(x, y) >= 4 {
				temp[y][x] = 'L'
			}
		}
	}

	for y := 0; y < g.y_max; y++ {
		copy(g.data[y], temp[y])
	}

	hash := g.GetHash()

	if g.hashes[hash] {
		return false
	} else {
		g.hashes[hash] = true
		return true
	}
}

func (g *Grid) GetHash() string {
	hash := ""

	for y := 0; y < g.y_max; y++ {
		hash += string(g.data[y])
	}

	return hash
}

func (g *Grid) Print() {
	for y := 0; y < g.y_max; y++ {
		fmt.Println(string(g.data[y]))
	}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	grid := Grid{[][]byte{}, 0, 0, map[string]bool{}}

	for scanner.Scan() {
		grid.AddRow(scanner.Text())
	}

	for grid.Step() {
	}

	fmt.Println(grid.CountOccupied())
}
