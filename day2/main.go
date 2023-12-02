package main

import (
	_ "embed"
	"fmt"
	"github.com/echojc/aocutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed 2023_2_example.txt
var i string
var test bool

func main() {
	test = false
	var lines []string
	if !test {
		input, err := aocutil.NewInputFromFile("../session_id")
		if err != nil {
			log.Fatal(err)
		}
		lines, err = input.Strings(2023, 2)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		lines = strings.Split(i, "\n")
	}
	fmt.Println("Results Part 1 : ", part1(lines))
	fmt.Println("Results Part 2 : ", part2(lines))
}

func numOfCube(colour string, line string) int {
	var numCubes int = 0
	firstReplacements := map[string]string{
		`red`:   `\d+ red`,
		`green`: `\d+ green`,
		`blue`:  `\d+ blue`,
	}

	secondReplacements := map[string]string{
		`red`:   ` red`,
		`green`: ` green`,
		`blue`:  ` blue`,
	}

	fistRegex := regexp.MustCompile(firstReplacements[colour])
	var CubeStrings []string = fistRegex.FindAllString(line, -1)

	secondRegex := regexp.MustCompile(secondReplacements[colour])
	var CubeNum string

	for _, CubeString := range CubeStrings {
		CubeNum = secondRegex.ReplaceAllString(CubeString, "")
		cubeInt, err := strconv.Atoi(CubeNum)
		if err != nil {
			fmt.Println("Error during conversion", err)
			return 100
		}
		if numCubes < cubeInt {
			numCubes = cubeInt
		}

	}
	return numCubes
}

func gameIdentifier(line string) int {
	var gameId int

	firstRegex := regexp.MustCompile(`Game \d+:`)
	var firstMatches []string = firstRegex.FindAllString(line, -1)

	for _, firstMatch := range firstMatches {

		secondRegex := regexp.MustCompile(`\d+`)
		var secondMatches string = secondRegex.FindString(firstMatch)

		matchInt, err := strconv.Atoi(secondMatches)
		if err != nil {
			fmt.Println("Error during conversion", err)
			return 0
		}
		gameId = matchInt
	}
	return gameId
}

func part1(lines []string) int {

	var total int

	for _, line := range lines {
		var gameId int = gameIdentifier(line)
		var redCubes int = numOfCube("red", line)
		var greenCubes int = numOfCube("green", line)
		var blueCubes int = numOfCube("blue", line)
		//fmt.Println(gameId, redCubes, greenCubes, blueCubes)
		if redCubes <= 12 && greenCubes <= 13 && blueCubes <= 14 {
			//fmt.Println("adds Game number : ", gameId)
			total = total + gameId
		}
	}
	return total
}

func part2(lines []string) int {

	var total int

	for _, line := range lines {

		//var gameId int = gameIdentifier(line)
		var redCubes int = numOfCube("red", line)
		var greenCubes int = numOfCube("green", line)
		var blueCubes int = numOfCube("blue", line)

		total = total + (redCubes * greenCubes * blueCubes)

	}
	return total
}
