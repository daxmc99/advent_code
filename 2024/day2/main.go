package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		panic(err)
	}
	splitData := bytes.Split(data, []byte("\n"))

	countSafe := 0
	countSafeWithProblemDampener := 0
	for i := 0; i < len(splitData); i++ {
		split := bytes.Split(splitData[i], []byte(" "))
		level := Level{}
		level.values = make([]int, len(split))
		for j := 0; j < len(split); j++ {
			byteToInt, _ := strconv.Atoi(string(split[j]))
			level.values[j] = byteToInt
		}

		if level.isSafe() {
			countSafe++
		}
		if level.isSafeWithProblemDampener() {
			countSafeWithProblemDampener++
		}
	}
	fmt.Printf("Safe levels: %v\n", countSafe)
	fmt.Printf("Safe levels with problem dampener: %v\n", countSafeWithProblemDampener)

}

type Level struct {
	values                  []int
	safe                    *bool
	safeWithProblemDampener *bool
}

func (l *Level) isSafeWithProblemDampener() bool {
	if l.safeWithProblemDampener != nil {
		return *l.safeWithProblemDampener
	}
	// check if the level is safe
	// A level is safe if
	// isSafe is true
	// OR if removing a single value from the level makes it safe

	if l.isSafe() {
		l.safeWithProblemDampener = new(bool)
		*l.safeWithProblemDampener = true
		return true
	}

	for i := 0; i < len(l.values); i++ {
		// remove the value at index i
		// fmt.Printf("values before: %v\n", l.values)
		valuesCopy := make([]int, len(l.values))
		copy(valuesCopy, l.values)
		values := append(valuesCopy[:i], valuesCopy[i+1:]...)
		// fmt.Printf("values: %v\n", values)
		level := Level{values: values}
		if level.isSafe() {
			l.safeWithProblemDampener = new(bool)
			*l.safeWithProblemDampener = true
			return true
		}
	}

	l.safeWithProblemDampener = new(bool)
	*l.safeWithProblemDampener = false
	return false
}

func (l *Level) isSafe() bool {
	if l.safe != nil {
		return *l.safe
	}
	// check if the level is safe
	// A level is safe if
	//	The levels are either all increasing or all decreasing.
	// AND Any two adjacent levels differ by at least one and at most three

	increasing := true
	decreasing := true

	for i := 0; i < len(l.values)-1; i++ {
		diff := abs(l.values[i+1] - l.values[i])
		if diff < 1 || diff > 3 {
			l.safe = new(bool)
			*l.safe = false
			return false
		}
		if l.values[i] >= l.values[i+1] {
			increasing = false
		}
		if l.values[i] <= l.values[i+1] {
			decreasing = false
		}
	}

	isSafe := increasing || decreasing
	l.safe = &isSafe
	return isSafe
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
