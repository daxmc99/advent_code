package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	parseRules(input)
}

type bagRules map[string]map[string]int

func parseRules(input string) bagRules {
	rawInput := strings.Split(input, ".")
	myRules := make(bagRules)
	for _, rule := range rawInput {
		parsed := strings.Split(rule, " bags contain ")
		if len(parsed) != 2 {
			fmt.Println("Found prefix without suffix", parsed)
			continue
		}
		prefix := parsed[0]
		myRules[prefix] = make(map[string]int)

		re := regexp.MustCompile(`(\d) (\w+ \w+)`)
		for _, match := range re.FindAllString(parsed[1], -1) {
			parsedMatch := strings.SplitN(match, " ", 2)
			fmt.Printf("parsedMatch found %v \n", parsedMatch)
			rawCount := parsedMatch[0]
			rawBag := parsedMatch[1]
			fmt.Printf("c is: %c, bag is: %s\n", rawCount, rawBag)
			count, err := strconv.Atoi(rawCount)


			myRules[prefix][rawBag] =
		}
	}
	return myRules
}
