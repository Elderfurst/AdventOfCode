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
	input := utilities.ReadInputFileToGrid("2024/day04/input.txt")

	count := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			character := input[i][j]

			if character == "X" {
				// check north
				if i >= 3 {
					word := input[i][j] + input[i-1][j] + input[i-2][j] + input[i-3][j]

					if word == "XMAS" {
						count++
					}
				}
				// check south
				if i < len(input)-3 {
					word := input[i][j] + input[i+1][j] + input[i+2][j] + input[i+3][j]

					if word == "XMAS" {
						count++
					}
				}
				// check west
				if j >= 3 {
					word := input[i][j] + input[i][j-1] + input[i][j-2] + input[i][j-3]

					if word == "XMAS" {
						count++
					}
				}
				// check east
				if j < len(input[i])-3 {
					word := input[i][j] + input[i][j+1] + input[i][j+2] + input[i][j+3]

					if word == "XMAS" {
						count++
					}
				}
				// check northeast
				if i >= 3 && j < len(input[i])-3 {
					word := input[i][j] + input[i-1][j+1] + input[i-2][j+2] + input[i-3][j+3]

					if word == "XMAS" {
						count++
					}
				}
				// check southeast
				if i < len(input)-3 && j < len(input[i])-3 {
					word := input[i][j] + input[i+1][j+1] + input[i+2][j+2] + input[i+3][j+3]

					if word == "XMAS" {
						count++
					}
				}
				// check northwest
				if i >= 3 && j >= 3 {
					word := input[i][j] + input[i-1][j-1] + input[i-2][j-2] + input[i-3][j-3]

					if word == "XMAS" {
						count++
					}
				}
				// check southwest
				if i < len(input)-3 && j >= 3 {
					word := input[i][j] + input[i+1][j-1] + input[i+2][j-2] + input[i+3][j-3]

					if word == "XMAS" {
						count++
					}
				}
			}
		}
	}

	fmt.Println(count)
}

func partTwo() {
	input := utilities.ReadInputFileToGrid("2024/day04/input.txt")

	count := 0

	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			character := input[i][j]

			if character == "A" {
				first := input[i-1][j-1] + input[i][j] + input[i+1][j+1]
				second := input[i+1][j-1] + input[i][j] + input[i-1][j+1]

				if (first == "MAS" || utilities.Reverse(first) == "MAS") &&
					(second == "MAS" || utilities.Reverse(second) == "MAS") {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
