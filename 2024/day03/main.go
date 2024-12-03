package main

import (
	"adventofcode/utilities"
	"fmt"
	"math"
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

	regex, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	do, _ := regexp.Compile("do\\(\\)")
	dont, _ := regexp.Compile("don't\\(\\)")

	dos := do.FindAllStringIndex(input, -1)
	donts := dont.FindAllStringIndex(input, -1)

	commands := regex.FindAllStringIndex(input, -1)

	total := 0

	for _, command := range commands {
		start := command[0]
		end := command[1]

		value := input[start:end]

		if shouldOperate(dos, donts, start) {
			commaSplit := strings.Split(value, ",")

			leftSplit := strings.Split(commaSplit[0], "(")
			rightSplit := strings.Split(commaSplit[1], ")")

			leftNumber, _ := strconv.Atoi(leftSplit[1])
			rightNumber, _ := strconv.Atoi(rightSplit[0])

			total += leftNumber * rightNumber
		}
	}

	fmt.Println(total)
}

func shouldOperate(dos, donts [][]int, index int) bool {
	process := true

	difference := math.MaxInt

	for _, do := range dos {
		start := do[0]

		if start < index {
			tempDifference := index - start

			if tempDifference < difference {
				process = true
				difference = tempDifference
			}
		} else {
			break
		}
	}

	for _, dont := range donts {
		start := dont[0]

		if start < index {
			tempDifference := index - start

			if tempDifference < difference {
				process = false
				difference = tempDifference
			}
		} else {
			break
		}
	}

	return process
}
