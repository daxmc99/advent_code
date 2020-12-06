package main

import (
	"fmt"
	"strings"
)

const maxForestXLen = 31

func main() {
	var forest [][]byte
	data := strings.Split(input, "\n")

	for _, d := range data {
		// d = "#.....##..", append each row
		if len(d) > maxForestXLen {
			panic("forest X len too large")
		}
		forest = append(forest, []byte(d))
	}
	/*
		Right 1, down 1.
		Right 3, down 1. (This is the slope you already checked.)
		Right 5, down 1.
		Right 7, down 1.
		Right 1, down 2.
	*/
}

const downBy = 1
const rightBy = 3

type pos struct {
	x int
	y int
}

func countTrees(rightBy, downBy int, forest [][]byte) int {
	p := pos{0, 0}
	trees := 0
	for p.y+downBy < len(forest) {
		//check the position downBy & rightBy
		p.x = (p.x + rightBy) % maxForestXLen
		p.y += downBy

		val := forest[p.y][p.x]
		fmt.Printf("at pos: %s", string(val))
		if forest[p.y][p.x] == '#' {
			trees++
		}
	}
	return trees
}
