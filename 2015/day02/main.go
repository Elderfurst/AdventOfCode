package main

import (
	"adventofcode/utilities"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputLines := utilities.ReadInputFileToLines("2015/day02/input.txt")

	totalSquareFeet := 0
	totalRibbon := 0

	for index, line := range inputLines {
		_ = index

		splitLine := strings.Split(line, "x")

		l, err := strconv.Atoi(splitLine[0])
		if err != nil {
			panic(err)
		}
		w, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic(err)
		}
		h, err := strconv.Atoi(splitLine[2])
		if err != nil {
			panic(err)
		}

		side1 := l * w
		side2 := w * h
		side3 := h * l

		presentDimensions := (2 * side1) + (2 * side2) + (2 * side3)

		if side1 <= side2 && side1 <= side3 {
			presentDimensions += side1
			totalRibbon += l + l + w + w
		} else if side2 <= side1 && side2 <= side3 {
			presentDimensions += side2
			totalRibbon += w + w + h + h
		} else if side3 <= side1 && side3 <= side2 {
			presentDimensions += side3
			totalRibbon += h + h + l + l
		}

		totalSquareFeet += presentDimensions
		totalRibbon += l * w * h
	}

	fmt.Println(totalSquareFeet)
	fmt.Println(totalRibbon)
}
