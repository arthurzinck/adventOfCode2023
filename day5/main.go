package main

import (
	_ "embed"
	"fmt"
	"github.com/echojc/aocutil"
	"log"
	"strings"
)

//go:embed 2023_5_example.txt
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
		lines, err = input.Strings(2023, 5)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		lines = strings.Split(i, "\n")
	}
	fmt.Println("Results Part 1 : ", part1(lines))
	fmt.Println("Results Part 2 : ", part2(lines))
}

func part1(lines []string) int {
	var total int = 0

	return total
}

func part2(lines []string) int {
	var total int = 0

	return total
}
