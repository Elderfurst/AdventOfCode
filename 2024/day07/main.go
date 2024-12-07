package main

import (
	"adventofcode/utilities"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	input := utilities.ReadInputFileToLines("2024/day07/input.txt")

	equations := parseInput(input)

	operators := []string{"+", "x"}

	answer := 0

	for _, equation := range equations {
		operatorPermutations := utilities.FillSpacesWithValues(operators, len(equation.components)-1)

		for _, permutation := range operatorPermutations {
			solution := equation.components[0]

			for i := 1; i < len(equation.components); i++ {
				switch permutation[i-1] {
				case "+":
					solution += equation.components[i]
				case "x":
					solution *= equation.components[i]
				}
			}

			if solution == equation.result {
				answer += equation.result
				break
			}
		}
	}

	fmt.Println(answer)
}

func partTwo() {
	input := utilities.ReadInputFileToLines("2024/day07/input.txt")

	equations := parseInput(input)

	operators := []string{"+", "x", "||"}

	answer := 0

	for _, equation := range equations {
		operatorPermutations := utilities.FillSpacesWithValues(operators, len(equation.components)-1)

		for _, permutation := range operatorPermutations {
			solution := equation.components[0]

			for i := 1; i < len(equation.components); i++ {
				switch permutation[i-1] {
				case "+":
					solution += equation.components[i]
				case "x":
					solution *= equation.components[i]
				case "||":
					left := strconv.Itoa(solution)
					right := strconv.Itoa(equation.components[i])

					combination := left + right

					solution, _ = strconv.Atoi(combination)
				}
			}

			if solution == equation.result {
				answer += equation.result
				break
			}
		}
	}

	fmt.Println(answer)
}

func parseInput(input []string) []equation {
	equations := make([]equation, 0)

	for _, line := range input {
		split := strings.Split(line, ": ")

		result, _ := strconv.Atoi(split[0])

		rawComponents := strings.Split(split[1], " ")

		components := make([]int, 0)

		for _, raw := range rawComponents {
			component, _ := strconv.Atoi(raw)

			components = append(components, component)
		}

		equations = append(equations, equation{
			result:     result,
			components: components,
		})
	}

	return equations
}

type equation struct {
	result     int
	components []int
}
