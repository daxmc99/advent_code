package main

import (
	"fmt"
	"strings"
)

type group struct {
	numPeople      int
	uniqueCount    int
	duplicateCount int
	people         []person
}

type person struct {
	answers string
}

func main() {
	rawGroups := strings.Split(input, "\n\n")
	var groups []group
	for _, g := range rawGroups {
		answers := strings.Split(g, "\n")
		numPersons := len(answers)

		var people []person
		for _, response := range answers {
			people = append(people, person{response})
		}

		groups = append(groups, group{numPersons, -1, -1, people})
	}
	fmt.Println(validateAnswers(groups[1]))
	uniq, dupl := sumCounts(groups)
	fmt.Printf("uniq: %d, dupl: %d", uniq, dupl)
}

func sumCounts(squad []group) (int, int) {
	sumUnique := 0
	sumDupl := 0
	for _, g := range squad {
		uniqueAnswers := validateAnswers(g)
		duplicateAnswers := allYes(g)
		g.uniqueCount = uniqueAnswers
		g.duplicateCount = duplicateAnswers
		sumUnique += uniqueAnswers
		sumDupl += duplicateAnswers

	}
	return sumUnique, sumDupl
}

// answers must not be duplicated
func validateAnswers(g group) int {
	answerSet := make(map[rune]bool)
	for _, person := range g.people {
		for _, char := range person.answers {
			answerSet[char] = true
		}
	}

	return len(answerSet)
}

// only count answers where everyone said yes
func allYes(g group) int {

	// special case
	if len(g.people) == 1 {
		return len(g.people[0].answers)
	}
	allYes := 0
	answerSet := make(map[rune]int)
	for _, person := range g.people {
		for _, char := range person.answers {
			// check for existence
			_, ok := answerSet[char]
			if !ok {
				answerSet[char] = 1
			} else {
				answerSet[char]++
			}
		}
	}

	for char := range answerSet {
		if answerSet[char] == len(g.people) {
			allYes++
		}
	}
	return allYes
}
