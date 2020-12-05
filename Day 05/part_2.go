package main

import (
	"bufio"
	"fmt"
	"os"
)

const COL = 8
const ROW = 128

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var seats [COL * ROW]bool

	min_seat_id := COL * ROW
	max_seat_id := 0

	for scanner.Scan() {
		x_low, x_high := 0, COL-1
		y_low, y_high := 0, ROW-1

		for _, c := range scanner.Text() {
			if c == 'F' {
				y_high -= (y_high-y_low)/2 + 1
			} else if c == 'B' {
				y_low += (y_high-y_low)/2 + 1
			} else if c == 'R' {
				x_low += (x_high-x_low)/2 + 1
			} else if c == 'L' {
				x_high -= (x_high-x_low)/2 + 1
			}
		}

		seat_id := y_low*COL + x_low
		seats[seat_id] = true

		if seat_id < min_seat_id {
			min_seat_id = seat_id
		}

		if seat_id > max_seat_id {
			max_seat_id = seat_id
		}
	}

	for i := min_seat_id; i < max_seat_id; i++ {
		if !seats[i] {
			fmt.Println(i)
		}
	}
}
