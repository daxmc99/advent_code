package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		panic(err)
	}
	// split the data into columns

	column1 := make([]int, 0)
	column2 := make([]int, 0)

	// 35039   67568
	// 61770   80134
	dataSplit := bytes.Split(data, []byte("\n"))
	// [ [ 35039 67568 ] [ 61770 80134 ] ]
	for i := 0; i < len(dataSplit); i++ {
		slic := bytes.Split(dataSplit[i], []byte(" "))
		byteToInt, _ := strconv.Atoi(string(slic[0]))
		column1 = append(column1, byteToInt)
		byteToInt, _ = strconv.Atoi(string(slic[3]))
		column2 = append(column2, byteToInt)
	}
	// sort them
	sort.IntSlice(column1).Sort()
	sort.IntSlice(column2).Sort()
	fmt.Printf("column1: %v\n", column1)
	fmt.Printf("column2: %v\n", column2)

	occurrenceMap := make(map[int]int)
	distance := 0
	for i := 0; i < len(column1); i++ {
		// first check to see if coulmn2[i] is in the map
		val, ok := occurrenceMap[column2[i]]
		if !ok {
			occurrenceMap[column2[i]] = 1
		} else {
			occurrenceMap[column2[i]] = val + 1
		}

		distance += Abs(column2[i] - column1[i])
	}

	fmt.Printf("distance: %v\n", distance)

	simScore := 0
	// get similariy score
	for i := 0; i < len(column1); i++ {
		if val, ok := occurrenceMap[column1[i]]; ok {
			// 3 occurences of 35039 means, 3 *
			simScore += val * column1[i]
		}
	}
	fmt.Printf("simScore: %v\n", simScore)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
