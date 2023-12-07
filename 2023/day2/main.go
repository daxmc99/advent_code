package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 2: Cube Conundrum ---
or example, the record of a few games might look like this:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
In game 1, three sets of cubes are revealed from the bag (and then put back again).
The first set is 3 blue cubes and 4 red cubes; the second set is 1 red cube, 2
green cubes, and 6 blue cubes; the third set is only 2 green cubes.

The Elf would first like to know which games would have been possible if the
bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

In the example above, games 1, 2, and 5 would have been possible if the bag had
been loaded with that configuration. However, game 3 would have been impossible
because at one point the Elf showed you 20 red cubes at once; similarly, game 4
would also have been impossible because the Elf showed you 15 blue cubes at once.
If you add up the IDs of the games that would have been possible, you get 8.

Determine which games would have been possible if the bag had been loaded with
only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
*/

const (
	redCubes   = 12
	greenCubes = 13
	blueCubes  = 14
)

var cubeValues = map[string]int{"red": redCubes, "green": greenCubes, "blue": blueCubes}

func main() {
	file, err := os.Open("./day2.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	sum := 0
	// read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b := scanner.Text()
		val, err := possibleGame(b)
		if err != nil {
			fmt.Fprintln(os.Stderr, "couldn't parse line: ", err, string(b))
		}
		fmt.Println("val: ", val)
		sum += val
		fmt.Println("sum: ", sum)
	}
}

func possibleGame(line string) (int, error) {
	game := strings.Split(line, ":")
	if len(game) < 2 || len(game) > 3 {
		return 0, fmt.Errorf("can't parse line with len %d", len(game))
	}
	idx := strings.Split(game[0], " ")[1]
	gameId, err := strconv.Atoi(idx)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse game id")
	}

	cubeGames := strings.Split(game[1], ";")
	for _, cubeGame := range cubeGames {
		p, err := processGame(cubeGame)
		if err != nil {
			return 0, fmt.Errorf("couldn't process game for game:%s ", line)
		}
		if !p { // game not possible
			fmt.Printf("game not possible: %d\n", gameId)
			return 0, nil
		}
	}
	return gameId, nil
}

// all the games are independent of each other
func processGame(game string) (bool, error) {
	cubes := strings.Split(game, ",")
	// [3 blue, 4 red]
	if len(cubes) < 1 {
		return false, fmt.Errorf("no cubes found")
	}
	for _, cube := range cubes {
		cube = strings.TrimSpace(cube)
		cubeParts := strings.Split(cube, " ")
		if len(cubeParts) != 2 {
			return false, fmt.Errorf("couldn't parse cube")
		}
		cubeValue, ok := cubeValues[cubeParts[1]]
		if !ok {
			return false, fmt.Errorf("couldn't find cube value for %s", cubeParts[1])
		}
		cubeCount, err := strconv.Atoi(cubeParts[0])
		if err != nil {
			return false, fmt.Errorf("couldn't parse cube count")
		}
		if cubeCount > cubeValue {
			return false, nil
		}
	}
	return true, nil
}
