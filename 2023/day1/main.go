package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

/*
	Advent of Code 2023 Day 1

	The newly-improved calibration document consists of lines of text; each line
	originally contained a specific calibration value that the Elves now need to
	recover. On each line, the calibration value can be found by combining the
	first digit and the last digit (in that order) to form a single two-digit number.

	For example:

	1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
	In this example, the calibration values of these four lines are 12, 38, 15, and 77.
	Adding these together produces 142.

	Consider your entire calibration document. What is the sum of all of the
	calibration values?
*/

func main() {

	file, err := os.Open("./day1.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	sum := 0
	// read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b := scanner.Bytes()
		val, err := parseLine(b)
		if err != nil {
			fmt.Fprintln(os.Stderr, "couldn't parse line: ", err, string(b))
		}
		fmt.Println("val: ", val)
		sum += val
		fmt.Println("sum: ", sum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("sum: ", sum)

	return
}

func parseLine(line []byte) (int, error) {
	if len(line) <= 0 {
		return 0, fmt.Errorf("empty line")
	}
	var first, last string

	for _, char := range line {
		if unicode.IsDigit(rune(char)) {
			if first == "" {
				first = string(char)
			}
			last = string(char)
		}
	}

	if first != "" && last != "" {
		final := fmt.Sprintf("%s%s", first, last)
		return strconv.Atoi(final)
	} else {
		return 0, fmt.Errorf("no digits found")
	}

}
