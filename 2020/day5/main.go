package main

import (
	"fmt"
	"math"
	"strings"
)

/*
	Find your seat on a airline that uses binary space partitioning
	FBFBBFFRLR, where F means front
	Step 1: Find the highest seat id
*/

const (
	MaxRow = 127
	MaxCol = 7
)

func seatID(row int, col int) int {
	return row*8 + col
}

func main() {
	maxID := 0
	seatList := strings.Split(input, "\n")
	for _, seat := range seatList {
		r, c := findSeat(seat)
		id := seatID(r, c)
		if id > maxID {
			maxID = id
		}
	}

	fmt.Println("Max seat id is: ", maxID)
}

func findSeat(seat string) (row int, col int) {

	row = findRow(seat[:8])

	col = findCol(seat[8:])
	return
}

func findRow(s string) int {

	min := 0.0
	max := float64(MaxRow)
	halve := 0.0
	for i := range s {
		// taking either upper or lower halve  of this range
		halve = max - min
		d := math.Round(halve / 2)

		if s[i] == 'F' {
			max -= d
		} else if s[i] == 'B' {
			min += d
		} else {
			panic("findRow: unexpected char at index")
		}
	}
	if min == max {
		return int(min)
	}
	return -1
}

func findCol(s string) int {
	min := 0.0
	max := float64(MaxCol)
	halve := 0.0

	for i := range s {
		halve = max - min
		d := math.Round(halve / 2)

		if s[i] == 'L' {
			max -= d
		} else if s[i] == 'R' {
			min += d
		} else {
			panic("findCol: unexpected char at index")
		}
	}
	if min == max {
		return int(min)
	}
	return -1
}
