package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

/*
	Find your seat on a airline that uses binary space partitioning
	FBFBBFFRLR, where F means front
	Step 1: Find the highest seat id
*/

const (
	maxrow = 127
	maxcol = 7
)

func seatID(row int, col int) int {
	return row*8 + col
}

func main() {
	maxID := 0
	seatList := strings.Split(input, "\n")

	idList := make([]int, len(seatList))
	for _, seat := range seatList {
		r, c := findSeat(seat)
		id := seatID(r, c)
		if id == 0 {
			continue
		}
		idList = append(idList, id)
		if id > maxID {
			maxID = id
		}
	}

	sort.Ints(idList)
	start := 0
	for i := range idList {
		if i == len(idList)-1 {
			break //stop overruns
		}
		// set start to the value after the current bad entry
		if idList[i] == idList[i+1] || idList[i] == 0 {
			start = i + 1
		}
	}
	cleanIDList := idList[start:]
	count := cleanIDList[0]
	for i := range cleanIDList {

		if count != cleanIDList[i] {
			fmt.Printf("next expected value was %d, but was actually %d\n", count, cleanIDList[i])
			fmt.Printf("Missing id value is %d", count)
			break
		}
		count++
	}

	fmt.Println("Max seat id is: ", maxID)
}

func findSeat(seat string) (row int, col int) {

	row = findRow(seat[:7])

	col = findCol(seat[7:])
	return
}

func findRow(s string) int {

	min := 0.0
	max := float64(maxrow)
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
	max := float64(maxcol)
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
