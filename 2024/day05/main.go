package main

import (
	"adventofcode/utilities"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	input := utilities.ReadInputFileToLines("2024/day05/input.txt")

	instructions := parseInstructions(input)

	answer := 0

	for _, update := range instructions.updates {
		if isValid(instructions.orderings, update) {
			middle := len(update) / 2

			answer += update[middle]
		}
	}

	fmt.Println(answer)
}

func partTwo() {
	input := utilities.ReadInputFileToLines("2024/day05/input.txt")

	instructions := parseInstructions(input)

	answer := 0

	for _, update := range instructions.updates {
		valid := isValid(instructions.orderings, update)

		if valid {
			continue
		}

		adjusted := update

		for !valid {
			brokenRules := brokenRules(instructions.orderings, adjusted)

			brokenRule := brokenRules[0]

			beforeIndex := slices.Index(adjusted, brokenRule.before)
			afterIndex := slices.Index(adjusted, brokenRule.after)

			adjusted = utilities.SwapIndices(adjusted, beforeIndex, afterIndex)

			valid = isValid(instructions.orderings, adjusted)
		}

		middle := len(adjusted) / 2

		answer += adjusted[middle]
	}

	fmt.Println(answer)
}

func brokenRules(orderings []order, update []int) []order {
	broken := make([]order, 0)

	for _, ordering := range orderings {
		beforeIndex := slices.Index(update, ordering.before)
		afterIndex := slices.Index(update, ordering.after)

		if beforeIndex >= 0 && afterIndex >= 0 && afterIndex < beforeIndex {
			broken = append(broken, ordering)
		}
	}

	return broken
}

func isValid(orderings []order, update []int) bool {
	valid := true

	for _, ordering := range orderings {
		beforeIndex := slices.Index(update, ordering.before)
		afterIndex := slices.Index(update, ordering.after)

		if beforeIndex >= 0 && afterIndex >= 0 && afterIndex < beforeIndex {
			valid = false
			break
		}
	}

	return valid
}

func parseInstructions(input []string) instructions {
	instructions := instructions{
		orderings: make([]order, 0),
		updates:   make([][]int, 0),
	}

	ordering := true

	for _, line := range input {
		if line == "" {
			ordering = false
			continue
		}

		if ordering {
			split := strings.Split(line, "|")

			first, _ := strconv.Atoi(split[0])
			second, _ := strconv.Atoi(split[1])

			order := order{
				before: first,
				after:  second,
			}

			instructions.orderings = append(instructions.orderings, order)
		} else {
			split := strings.Split(line, ",")

			update := make([]int, 0)

			for _, i := range split {
				parsed, _ := strconv.Atoi(i)

				update = append(update, parsed)
			}

			instructions.updates = append(instructions.updates, update)
		}
	}

	return instructions
}

type instructions struct {
	orderings []order
	updates   [][]int
}

type order struct {
	before int
	after  int
}
