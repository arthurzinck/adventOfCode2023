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

//go:embed 2023_1_example.txt
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
		lines, err = input.Strings(2023, 1)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		lines = strings.Split(i, "\n")
	}
	fmt.Println("Results Part 1 : ", part1(lines))
	fmt.Println("Results Part 2 :", part2(lines))
}

func part1(lines []string) int {
	// fmt.Println(lines)
	var total int
	re := regexp.MustCompile("[^0-9]")
	for _, line := range lines {

		var numbers string = re.ReplaceAllString(line, "")

		if len(numbers) > 0 {

			var triage string = fmt.Sprintf("%s%s", numbers[0:1], numbers[len(numbers)-1:])

			castInt, err := strconv.Atoi(triage)

			if err != nil {
				fmt.Println("Error during conversion")
				return 1
			}

			total = total + castInt
		}
	}
	return total
}

func replaceWordsWithNumbers(input string) string {
	replacements := map[string]string{
		`nine`:  "n9e",
		`eight`: "e8t",
		`seven`: "s7n",
		`six`:   "s6x",
		`five`:  "f5e",
		`four`:  "f4r",
		`three`: "t3e",
		`two`:   "t2o",
		`one`:   "o1e",
		`zero`:  "z0o",
	}

	for word, number := range replacements {
		re := regexp.MustCompile(word)
		input = re.ReplaceAllString(input, number)
	}
	return input
}

func part2(lines []string) int {
	//fmt.Println(lines)

	var total int

	for _, line := range lines {
		var numbersWithoutLetters string = replaceWordsWithNumbers(line)
		re10 := regexp.MustCompile(`[^0-9]`)
		var numbersWithNumOnly string = re10.ReplaceAllString(numbersWithoutLetters, "")

		if len(numbersWithNumOnly) != 0 {
			var triage string = fmt.Sprintf("%s%s", numbersWithNumOnly[0:1], numbersWithNumOnly[len(numbersWithNumOnly)-1:])
			castInt, err := strconv.Atoi(triage)

			if err != nil {
				fmt.Println("Error during conversion")
				return 1
			}
			total = total + castInt
		} else {
			fmt.Println("err")
		}

	}
	return total

}
