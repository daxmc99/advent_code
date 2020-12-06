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
	fmt.Println("Ans 1:", sum(data))
	fmt.Println("Ans 2:", sum3(data))
}

func sum(data []string) int {
	// O(n^2) bad :(
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

func sum3(data []string) int {
	// don't have to be too smart if the lang if fast
	for _, i := range data {
		n1, _ := strconv.Atoi(i)
		for _, j := range data {
			n2, _ := strconv.Atoi(j)
			for _, k := range data {
				n3, _ := strconv.Atoi(k)

				if n1+n2+n3 == targetSum {
					return n1 * n2 * n3
				}
			}
		}
	}
	return -1
}
