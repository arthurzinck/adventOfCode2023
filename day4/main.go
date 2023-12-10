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

//go:embed 2023_4_example.txt
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
		lines, err = input.Strings(2023, 4)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		lines = strings.Split(i, "\n")
	}
	fmt.Println("Results Part 1 : ", part1(lines))
	fmt.Println("Results Part 2 : ", part2(lines))
}

func toNum(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		//fmt.Println(err)
	}
	return output
}

func findNumberPositions(line string) [][]int {
	regex := regexp.MustCompile(`\d+`)
	var positions [][]int = regex.FindAllStringIndex(line, 50)
	return positions
}

func remove(s []string, index int) []string {

	return append(s[:index], s[index+1:]...)
}

func numberOfWiningNumber(line string) int {
	var numberOfWiningNumber []int
	regex := regexp.MustCompile(`: .* \|`)
	var matches []string = regex.FindAllString(line, -1)
	var matchedNumber []string = strings.Split(matches[0], " ")

	matchedNumber = remove(matchedNumber, 0)
	matchedNumber = remove(matchedNumber, len(matchedNumber)-1)

	for _, number := range matchedNumber {
		if toNum(number) != 0 {
			numberOfWiningNumber = append(numberOfWiningNumber, toNum(number))
		}
	}
	//fmt.Println("WinningNumber : ", numberOfWiningNumber, len(numberOfWiningNumber))
	return len(numberOfWiningNumber)
}

func winingNumbers(line string) []int {
	var winingNumbers []int
	for counter, numberPositions := range findNumberPositions(line) {
		//fmt.Println(line[numberPositions[0]:numberPositions[1]])
		if counter > numberOfWiningNumber(line) {
			break
		} else if counter == 0 {
			continue
		} else {
			winingNumbers = append(winingNumbers, toNum(line[numberPositions[0]:numberPositions[1]]))
		}
	}
	//fmt.Println("winingNumbers:", winingNumbers)
	return winingNumbers
}

func playedNumbers(line string) []int {
	var playedNumbers []int
	for counter, numberPositions := range findNumberPositions(line) {
		//fmt.Println(line[numberPositions[0]:numberPositions[1]])
		if counter > numberOfWiningNumber(line) {
			playedNumbers = append(playedNumbers, toNum(line[numberPositions[0]:numberPositions[1]]))
		}
	}
	//fmt.Println("playedNumbers:", playedNumbers)
	return playedNumbers
}

func numberMatching(line string) []int {
	var matches []int

	for _, playedNumber := range playedNumbers(line) {
		for _, winingNumber := range winingNumbers(line) {
			//fmt.Println("WinNumber: ", winingNumber, "playedNumber:", playedNumber)
			if winingNumber == playedNumber {
				matches = append(matches, playedNumber)
			}
		}
	}
	fmt.Println("Matches:", matches)
	return matches
}

func cardValue(matchedNumbers []int) int {
	var total int
	for counter := 1; counter < len(matchedNumbers)+1; counter++ {

		if counter > 1 {
			total = total + total
			//fmt.Println("Total:", total, "Counter :", counter)
		} else {
			//fmt.Println()
			total = total + counter
			//fmt.Println("Total:", total, "Counter :", counter)
		}
	}
	//fmt.Println("Total :", total)
	return total
}

func part1(lines []string) int {
	var total int
	for _, line := range lines {
		total = total + cardValue(numberMatching(line))
	}
	return total
}

func match(lineNumber int, line string) {

}

func numberOfCards(lines []string) [250]int {
	var numberOfCards [250]int
	for lineNumber, line := range lines {
		matches := numberMatching(line)
		fmt.Println("LineNumber", lineNumber, "NumberOfMatch", len(matches))
		for i := 0; i <= len(matches); i++ {
			if i > 0 {
				if numberOfCards[lineNumber] > 1 {
					numberOfCards[lineNumber+i] = numberOfCards[lineNumber+i] + 1*numberOfCards[lineNumber]
				} else {
					numberOfCards[lineNumber+i] = numberOfCards[lineNumber+i] + 1
				}
			} else {
				numberOfCards[lineNumber+i] = numberOfCards[lineNumber+i] + 1

			}

		}
		fmt.Println("Card ", lineNumber, "numberOfCard", numberOfCards[lineNumber])
	}
	return numberOfCards
}

func part2(lines []string) int {
	var total int
	numberOfCards := numberOfCards(lines)

	fmt.Println(numberOfCards)
	for _, numberOfcard := range numberOfCards {

		total = total + numberOfcard
	}

	return total
}
