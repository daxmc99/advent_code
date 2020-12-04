package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Find the two values that sum to 2020, multiply them and return the result
const targetSum = 2020

func main() {
	data := strings.Split(input, "\n")
	fmt.Println(sum(data))
}

func sum(data []string) int {
	// O(n^n) bad :(
	for _, j := range data {
		n, err := strconv.Atoi(j)
		if err != nil {
			fmt.Printf("Encountered err: %v", err)
			return -1
		}
		for _, k := range data {
			m, err := strconv.Atoi(k)
			if err != nil {
				fmt.Printf("Encountered err: %v", err)
				return -1
			}
			if n+m == targetSum {
				return n * m
			}
		}
	}
	return -1
}
