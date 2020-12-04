package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
For example, suppose you have the following list:

1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
Each line gives the password policy and then the password. The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is not;
it contains no instances of b, but needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the limits of their respective policies.`
*/

type pwConfig struct {
	min int
	max int
	req string
	pw  string
}

func main() {
	data := strings.Split(input, "\n")

	numValidConfigs := 0
	numValidTobogganConfigs := 0
	for _, d := range data {
		conf := marshalConfig(d)
		if conf.validateSledRental() {
			numValidConfigs++
		}
		fmt.Println("The number of valid configs are: ", numValidConfigs)
		if conf.validateTobogganRental() {
			numValidTobogganConfigs++
		}
	}

	fmt.Println("The number of valid toboggan configs: ", numValidTobogganConfigs)
}

func marshalConfig(c string) pwConfig {
	config := pwConfig{}
	raw := strings.Split(c, " ")
	minMax := strings.Split(raw[0], "-")
	config.min, _ = strconv.Atoi(minMax[0])
	config.max, _ = strconv.Atoi(minMax[1])
	config.req = strings.Trim(raw[1], ":")
	config.pw = raw[2]
	return config
}

func (c pwConfig) validateSledRental() bool {
	n := strings.Count(c.pw, c.req)
	if n >= c.min && n <= c.max {
		return true
	}
	return false
}

/*
Given the same example list from above:

1-3 a: abcde is valid: position 1 contains a and position 3 does not.
1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.

*/

func (c pwConfig) validateTobogganRental() bool {
	// given indexes assume non-zero first index
	// ie, a given index of "1" would index to position 0
	firstPos := c.pw[c.min-1] == c.req[0]

	secondPos := c.pw[c.max-1] == c.req[0]
	// xor
	if firstPos != secondPos {
		return true
	}

	return false
}
