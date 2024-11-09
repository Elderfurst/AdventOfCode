package main

import (
	"adventofcode/utilities"
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := utilities.ReadInputFileToLines("2015/day17/input.txt")

	containers := make([]int, 0)

	for _, line := range input {
		parsed, _ := strconv.Atoi(line)
		containers = append(containers, parsed)
	}

	subsets := utilities.GetSubsets(containers)

	exact := 0

	containerCount := make(map[int]int)

	for _, subset := range subsets {
		totalVolume := 0

		for _, container := range subset {
			totalVolume += container
		}

		if totalVolume == 150 {
			containerCount[len(subset)]++

			exact++
		}
	}

	maxCount := math.MaxInt
	combinations := 0

	for key, value := range containerCount {
		if key < maxCount {
			maxCount = key
			combinations = value
		}
	}

	fmt.Println(exact)
	fmt.Println(combinations)
}
