package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
It seems like the goal of the program is just to multiply some numbers.
It does that with instructions like mul(X,Y), where X and Y are each 1-3 digit
numbers. For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024.
Similarly, mul(123,4) would multiply 123 by 4.

However, because the program's memory has been corrupted, there are also many
invalid characters that should be ignored, even if they look like part of a mul
instruction. Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.

For example, consider the following section of corrupted memory:

xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
Only the four highlighted sections are real mul instructions. Adding up the
result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).
*/

func main() {

	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	// custom split function, which should only return text like mul(X,Y)
	// it should return each digit
	disabledCount := 0
	enabledCount := 0

	enabled := true
	const do string = "do()"
	const dont string = "don't()"
	const mul string = "mul("

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		for i := 0; i < len(data); i++ {
			// The do() instruction enables
			// the don't() instruction disables
			if data[i] == 'd' {
				if string(data[i:i+4]) == do {
					enabled = true
					enabledCount++
					i += 3
				} else if string(data[i:i+7]) == dont {
					enabled = false
					disabledCount++
					i += 6
				}
			}

			if data[i] == 'm' && string(data[i:i+4]) == mul {
				hasSeenDigit := false
				hasSeenComma := false
				for j := i + 4; j < len(data); j++ {
					if data[j] == ')' {
						if hasSeenDigit && hasSeenComma {
							if !enabled {
								fmt.Printf("Disabled mul instruction: %s\n", data[i:j+1])
								break
							}
							fmt.Printf("Found mul instruction: %s\n", data[i:j+1])
							return j + 1, data[i+4 : j], nil
						} else {
							break
						}
					}

					if data[j] == ',' {
						hasSeenComma = true
						continue
					}
					if data[j] >= '0' && data[j] <= '9' {
						hasSeenDigit = true
					} else { // non-digit character
						fmt.Printf("Invalid mul instruction: %s\n", data[i:j])
						i = j
						break
					}
				}
			}
		}
		return len(data), nil, nil
	}

	result := 0
	scanner.Split(split)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("Valid token: %s\n", text)
		vals := strings.Split(text, ",")
		x, _ := strconv.Atoi(vals[0])
		y, _ := strconv.Atoi(vals[1])
		result += x * y
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %vs\n", err)
	}
	fmt.Println("Result:", result)
	fmt.Println("Enabled count:", enabledCount)
	fmt.Println("Disabled count:", disabledCount)
}
