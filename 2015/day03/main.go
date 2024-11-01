package main

import (
	"adventofcode/utilities"
	"fmt"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	inputCharacters := utilities.ReadInputFileToCharacters("2015/day03/input.txt")

	x := 0
	y := 0

	presents := make(map[string]int)

	for _, direction := range inputCharacters {
		currentLocation := fmt.Sprintf("%d:%d", x, y)

		presents[currentLocation]++

		switch direction {
		case "^":
			y++
		case ">":
			x++
		case "<":
			x--
		case "v":
			y--
		}
	}

	multiplePresents := 0

	for _, v := range presents {
		if v >= 1 {
			multiplePresents++
		}
	}

	fmt.Println(multiplePresents)
}

func partTwo() {
	inputCharacters := utilities.ReadInputFileToCharacters("2015/day03/input.txt")

	santaX := 0
	santaY := 0
	roboX := 0
	roboY := 0

	presents := make(map[string]int)

	presents["0:0"] = 2

	for index, direction := range inputCharacters {
		santaMoves := index%2 == 0
		santaLoc := fmt.Sprintf("%d:%d", santaX, santaY)
		roboLoc := fmt.Sprintf("%d:%d", roboX, roboY)

		if santaMoves {
			switch direction {
			case "^":
				santaY++
			case ">":
				santaX++
			case "<":
				santaX--
			case "v":
				santaY--
			}

			presents[santaLoc]++
		} else {
			switch direction {
			case "^":
				roboY++
			case ">":
				roboX++
			case "<":
				roboX--
			case "v":
				roboY--
			}

			presents[roboLoc]++
		}
	}

	multiplePresents := 0

	for _, v := range presents {
		if v >= 1 {
			multiplePresents++
		}
	}

	fmt.Println(multiplePresents)
}
