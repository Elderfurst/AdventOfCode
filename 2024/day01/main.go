package main

import (
	"adventofcode/utilities"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	input := utilities.ReadInputFileToLines("2024/day01/input.txt")

	locations := parseInputLines(input)

	slices.Sort(locations.first)
	slices.Sort(locations.second)

	totalDistance := 0

	for i := 0; i < len(locations.first); i++ {
		first := locations.first[i]
		second := locations.second[i]

		totalDistance += int(math.Abs(float64(first - second)))
	}

	fmt.Println(totalDistance)
}

func partTwo() {
	input := utilities.ReadInputFileToLines("2024/day01/input.txt")

	locations := parseInputLines(input)

	similarity := 0

	for _, first := range locations.first {
		similarity += first * utilities.GetCount(first, locations.second)
	}

	fmt.Println(similarity)
}

func parseInputLines(inputLines []string) locations {
	locations := locations{
		first:  make([]int, 0),
		second: make([]int, 0),
	}

	for _, line := range inputLines {
		splitLine := strings.Split(line, "   ")
		first, _ := strconv.Atoi(splitLine[0])
		second, _ := strconv.Atoi(splitLine[1])

		locations.first = append(locations.first, first)
		locations.second = append(locations.second, second)
	}

	return locations
}

type locations struct {
	first  []int
	second []int
}
