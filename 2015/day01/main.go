package main

import (
	"adventofcode/utilities"
	"fmt"
)

func main() {
	inputCharacters := utilities.ReadInputFileToCharacters("2015/day01/input.txt")

	currentFloor := 0
	firstBasement := true

	for index, character := range inputCharacters {
		switch character {
		case "(":
			currentFloor++
		case ")":
			currentFloor--
		}

		if currentFloor < 0 && firstBasement == true {
			fmt.Println(index + 1)
			firstBasement = false
		}
	}

	fmt.Println(currentFloor)
}
