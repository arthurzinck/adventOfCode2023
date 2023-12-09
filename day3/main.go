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

//go:embed 2023_3_example.txt
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
		lines, err = input.Strings(2023, 3)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		lines = strings.Split(i, "\n")
	}
	fmt.Println("Results Part 1 : ", part1(lines))
	fmt.Println("Results Part 2 : ", part2(lines))
}

func findNumberPositions(line string) [][]int {
	regex := regexp.MustCompile(`\d+`)
	var positions [][]int = regex.FindAllStringIndex(line, 50)
	return positions
}

func getEndLineNumber(lenLines int, lineNumber int) int {
	if lineNumber == lenLines-1 {
		return lineNumber
	} else {
		return lineNumber + 1
	}
}

func getStartLineNumber(lineNumber int) int {
	if lineNumber == 0 {
		return lineNumber
	} else {
		return lineNumber - 1
	}
}

func getStartPositionNumber(startPosition int) int {
	if startPosition == 0 {
		return startPosition
	} else {
		return startPosition - 1
	}
}

func getEndPositionNumber(lenLine int, endPosition int) int {
	if endPosition == lenLine {
		return endPosition
	} else {
		return endPosition + 1
	}
}

func adjacentSymbol(lines []string, lineNumber int, startPosition int, endPosition int) bool {

	var charsToTest = []string{"~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "\\", "|", "/", "<", ">", "?"}

	var startLine int = getStartLineNumber(lineNumber)
	var endLine int = getEndLineNumber(len(lines), lineNumber)
	startPosition = getStartPositionNumber(startPosition)
	endPosition = getEndPositionNumber(len(lines), endPosition)

	for i := startLine; i <= endLine; i++ {
		for j := startPosition; j < endPosition; j++ {
			for _, toTest := range charsToTest {
				if j == len(lines[lineNumber]) {
					if strings.Contains(lines[i][j:], toTest) {
						return true
					}
				} else {
					if strings.Contains(lines[i][j:j+1], toTest) {
						return true
					}
				}
			}
		}
	}
	return false
}
func toNum(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return output
}

func part1(lines []string) int {
	var total int
	var taken []int
	var notTaken []int
	for lineNumber, line := range lines {
		for _, coordinates := range findNumberPositions(line) {
			if adjacentSymbol(lines, lineNumber, coordinates[0], coordinates[1]) {
				total = total + toNum(lines[lineNumber][coordinates[0]:coordinates[1]])
				taken = append(taken, toNum(lines[lineNumber][coordinates[0]:coordinates[1]]))
			} else {
				notTaken = append(notTaken, toNum(lines[lineNumber][coordinates[0]:coordinates[1]]))
			}
		}
	}
	//fmt.Println("TAKEN :", taken)
	//fmt.Println("NOT TAKEN :", notTaken)
	return total
}

func findStarPosition(line string) [][]int {
	regex := regexp.MustCompile(`\*`)
	var positions [][]int = regex.FindAllStringIndex(line, 50)

	return positions
}

func intersect(a, b, c, d int) bool {
	return a <= d && c <= b
}

func checkAround(lines []string, lineNumber int, coordinates []int) []string {
	var startLine int = getStartLineNumber(lineNumber)
	var endLine int = getEndLineNumber(len(lines), lineNumber)

	var intersects []string
	for i := startLine; i <= endLine; i++ {
		for _, numberCoordinates := range findNumberPositions(lines[i]) {
			if intersect(coordinates[0], coordinates[1], numberCoordinates[0], numberCoordinates[1]) {
				intersects = append(intersects, lines[i][numberCoordinates[0]:numberCoordinates[1]])
			}
		}
	}
	if len(intersects) >= 2 {
		return intersects
	}
	return []string{"0", "0"}
}

func part2(lines []string) int {
	var total int
	var gears []string

	for lineNumber, line := range lines {
		for _, coordinates := range findStarPosition(line) {
			gears = checkAround(lines, lineNumber, coordinates)
			total = total + (toNum(gears[0]) * toNum(gears[1]))
		}
	}
	return total
}
