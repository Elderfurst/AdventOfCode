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
	inputLines := utilities.ReadInputFileToLines("input.txt")

	vowels, _ := regexp.Compile("[aeiou].*[aeiou].*[aeiou]")

	niceWords := 0

	for _, word := range inputLines {
		if vowels.MatchString(word) &&
			doubleLetters(word) &&
			!strings.Contains(word, "ab") &&
			!strings.Contains(word, "cd") &&
			!strings.Contains(word, "pq") &&
			!strings.Contains(word, "xy") {
			niceWords++
		}
	}

	fmt.Println(niceWords)
}

func partTwo() {
	inputLines := utilities.ReadInputFileToLines("input.txt")

	niceWords := 0

	for _, word := range inputLines {
		if repeatDoubleLetters(word) &&
			repeatLetterSpaced(word) {
			niceWords++
		}
	}

	fmt.Println(niceWords)
}

func doubleLetters(word string) bool {
	for i := 0; i < len(word)-1; i++ {
		if word[i] == word[i+1] {
			return true
		}
	}

	return false
}

func repeatDoubleLetters(word string) bool {
	for i := 0; i < len(word)-2; i++ {
		currentSet := word[i : i+2]

		for j := i + 2; j < len(word)-1; j++ {
			checkSet := word[j : j+2]

			if currentSet == checkSet {
				return true
			}
		}
	}

	return false
}

func repeatLetterSpaced(word string) bool {
	for i := 0; i < len(word)-2; i++ {
		currentLetter := string(word[i])
		checkLetter := string(word[i+2])

		if currentLetter == checkLetter {
			return true
		}
	}

	return false
}
