package main

import (
	"adventofcode/utilities"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	inputLines := utilities.ReadInputFileToLines("2015/day08/input.txt")

	totalCharacters := 0
	totalEscapedCharacters := 0

	for _, line := range inputLines {
		totalCharacters += len(line)

		// remove first and last quotation
		noQuotes := strings.Trim(line, "\"")

		// replace all instances of an escaped quote with just the letter Q
		noEscapedQuotes := strings.ReplaceAll(noQuotes, "\\\"", "Q")

		// replace all instances of an escaped slash with the letter S
		noSlashes := strings.ReplaceAll(noEscapedQuotes, "\\\\", "S")

		// replace all hex values with H
		regex, _ := regexp.Compile("\\\\x[a-z0-9]{2}")

		noHex := regex.ReplaceAllString(noSlashes, "H")

		totalEscapedCharacters += len(noHex)
	}

	fmt.Println(totalCharacters - totalEscapedCharacters)
}

func partTwo() {
	inputLines := utilities.ReadInputFileToLines("2015/day08/input.txt")

	totalCharacters := 0
	totalEscapedCharacters := 0

	for _, line := range inputLines {
		totalCharacters += len(line)

		encodeSlashes := strings.ReplaceAll(line, "\\", "\\\\")

		encodeQuotes := strings.ReplaceAll(encodeSlashes, "\"", "\\\"")

		totalEscapedCharacters += len(encodeQuotes) + 2
	}

	fmt.Println(totalEscapedCharacters - totalCharacters)
}
