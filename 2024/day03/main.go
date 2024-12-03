package main

import (
	"adventofcode/utilities"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	input := utilities.ReadInputFileToString("2024/day03/input.txt")

	regex, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")

	commands := regex.FindAllString(input, -1)

	total := 0

	for _, command := range commands {
		commaSplit := strings.Split(command, ",")

		leftSplit := strings.Split(commaSplit[0], "(")
		rightSplit := strings.Split(commaSplit[1], ")")

		leftNumber, _ := strconv.Atoi(leftSplit[1])
		rightNumber, _ := strconv.Atoi(rightSplit[0])

		total += leftNumber * rightNumber
	}

	fmt.Println(total)
}

func partTwo() {
	input := utilities.ReadInputFileToString("2024/day03/input.txt")

	regex, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)|do\\(\\)|don't\\(\\)")

	commands := regex.FindAllStringIndex(input, -1)

	process := true

	total := 0

	for _, command := range commands {
		start := command[0]
		end := command[1]

		value := input[start:end]

		if value == "do()" {
			process = true

		} else if value == "don't()" {
			process = false
		} else {
			if process {
				commaSplit := strings.Split(value, ",")

				leftSplit := strings.Split(commaSplit[0], "(")
				rightSplit := strings.Split(commaSplit[1], ")")

				leftNumber, _ := strconv.Atoi(leftSplit[1])
				rightNumber, _ := strconv.Atoi(rightSplit[0])

				total += leftNumber * rightNumber
			}
		}
	}

	fmt.Println(total)
}
