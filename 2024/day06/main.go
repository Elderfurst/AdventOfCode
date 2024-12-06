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
	input := utilities.ReadInputFileToCoordinateHash("2024/day06/input.txt")

	fmt.Println(solve(input))
}

func partTwo() {
	input := utilities.ReadInputFileToCoordinateHash("2024/day06/input.txt")

	blockers := 0

	for i := 0; i < 130; i++ {
		for j := 0; j < 130; j++ {
			current := coordinate{i, j}

			value := input[key(current)]

			if value == "." {
				input[key(current)] = "#"

				result := solve(input)

				if result < 0 {
					blockers++
				}

				input[key(current)] = "."
			}
		}
	}

	fmt.Println(blockers)
}

func solve(input map[string]string) int {
	guard := findGuard(input)

	visited := make(map[string]bool)

	direction := "up"

	loopCount := 0

	for {
		loopCount++

		if loopCount > 15000 {
			break
		}

		visited[key(guard)] = true

		var next coordinate

		switch direction {
		case "up":
			next = coordinate{guard.x, guard.y - 1}
		case "down":
			next = coordinate{guard.x, guard.y + 1}
		case "left":
			next = coordinate{guard.x - 1, guard.y}
		case "right":
			next = coordinate{guard.x + 1, guard.y}
		}

		value, found := input[key(next)]

		if !found {
			break
		}

		if value == "#" {
			switch direction {
			case "up":
				direction = "right"
			case "down":
				direction = "left"
			case "left":
				direction = "up"
			case "right":
				direction = "down"
			}
		} else {
			guard = next
		}
	}

	if loopCount > 15000 {
		return -1
	}

	return len(visited)
}

func findGuard(input map[string]string) coordinate {
	for key, value := range input {
		if value == "^" {
			split := strings.Split(key, ",")

			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])

			return coordinate{x, y}
		}
	}

	return coordinate{-1, -1}
}

func key(coordinate coordinate) string {
	return fmt.Sprintf("%d,%d", coordinate.x, coordinate.y)
}

type coordinate struct {
	x, y int
}
