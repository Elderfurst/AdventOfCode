package main

import (
	"adventofcode/utilities"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	input := utilities.ReadInputFileToLines("2024/day02/input.txt")

	reports := parseInputLines(input)

	safe := 0

	for _, report := range reports {
		if checkSafety(report.levels) {
			safe++
		}
	}

	fmt.Println(safe)
}

func partTwo() {
	input := utilities.ReadInputFileToLines("2024/day02/input.txt")

	reports := parseInputLines(input)

	safe := 0

	for _, report := range reports {
		if checkSafety(report.levels) {
			safe++
			continue
		}

		safeReport := false

		for i := 0; i < len(report.levels); i++ {
			dampener := utilities.RemoveIndex(report.levels, i)

			if checkSafety(dampener) {
				safeReport = true
				break
			}
		}

		if safeReport {
			safe++
		}
	}

	fmt.Println(safe)
}

func checkSafety(levels []int) bool {
	safe := false

	ascending := sort.SliceIsSorted(levels, func(i, j int) bool {
		return levels[i] < levels[j]
	})

	descending := sort.SliceIsSorted(levels, func(i, j int) bool {
		return levels[i] > levels[j]
	})

	safeLevels := true

	for i := 0; i < len(levels)-1; i++ {
		current := levels[i]
		next := levels[i+1]

		difference := int(math.Abs(float64(current - next)))

		if difference < 1 || difference > 3 {
			safeLevels = false
			break
		}
	}

	if (ascending || descending) && safeLevels {
		safe = true
	}

	return safe
}

func parseInputLines(inputLines []string) []report {
	reports := make([]report, 0)

	for _, line := range inputLines {
		report := report{
			levels: make([]int, 0),
		}
		splitLine := strings.Split(line, " ")

		for _, element := range splitLine {
			level, _ := strconv.Atoi(element)

			report.levels = append(report.levels, level)
		}

		reports = append(reports, report)
	}

	return reports
}

type report struct {
	levels []int
}
