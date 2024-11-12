package main

import (
	"adventofcode/utilities"
	"fmt"
	"sort"
	"strings"
)

var molecule = "ORnPBPMgArCaCaCaSiThCaCaSiThCaCaPBSiRnFArRnFArCaCaSiThCaCaSiThCaCaCaCaCaCaSiRnFYFArSiRnMgArCaSiRnPTiTiBFYPBFArSiRnCaSiRnTiRnFArSiAlArPTiBPTiRnCaSiAlArCaPTiTiBPMgYFArPTiRnFArSiRnCaCaFArRnCaFArCaSiRnSiRnMgArFYCaSiRnMgArCaCaSiThPRnFArPBCaSiRnMgArCaCaSiThCaSiRnTiMgArFArSiThSiThCaCaSiRnMgArCaCaSiRnFArTiBPTiRnCaSiAlArCaPTiRnFArPBPBCaCaSiThCaPBSiThPRnFArSiThCaSiThCaSiThCaPTiBSiRnFYFArCaCaPRnFArPBCaCaPBSiRnTiRnFArCaPRnFArSiRnCaCaCaSiThCaRnCaFArYCaSiRnFArBCaCaCaSiThFArPBFArCaSiRnFArRnCaCaCaFArSiRnFArTiRnPMgArF"

func main() {
	partOne()
	partTwo()
}

func partOne() {
	input := utilities.ReadInputFileToLines("2015/day19/input.txt")

	instructions := getInstructions(input)

	distinctResults := make(map[string]bool)

	for _, instruction := range instructions {
		indices := utilities.GetAllIndices(molecule, instruction.from)

		window := len(instruction.from)

		for _, index := range indices {
			transformation := molecule[:index] + instruction.to + molecule[index+window:]

			distinctResults[transformation] = true
		}
	}

	fmt.Println(len(distinctResults))
}

// Essentially for part 2 we want to start with our molecule and work our way backwards
// instead of starting with 'e' and working forwards. This will let us greedily iterate
// over the 1000 shortest options each iteration instead of having to explore every single
// option if we were expanding instead of contracting
func partTwo() {
	input := utilities.ReadInputFileToLines("2015/day19/input.txt")

	instructions := getInstructions(input)

	fmt.Println(stepsUntilE(molecule, instructions))
}

func stepsUntilE(start string, instructions []instruction) int {
	set := []string{start}

	iterations := 0

	for {
		distinctResults := make(map[string]bool)

		for _, attempt := range set {
			for _, instruction := range instructions {
				indices := utilities.GetAllIndices(attempt, instruction.to)

				window := len(instruction.to)

				for i := len(indices) - 1; i >= 0; i-- {
					index := indices[i]

					transformation := attempt[:index] + instruction.from + attempt[index+window:]

					distinctResults[transformation] = true
				}
			}
		}

		allOptions := utilities.GetMapKeys(distinctResults)

		sort.Slice(allOptions, func(i, j int) bool {
			return len(allOptions[i]) < len(allOptions[j])
		})

		if len(allOptions) > 1000 {
			set = allOptions[:1000]
		} else {
			set = allOptions
		}

		iterations++

		for _, attempt := range set {
			if attempt == "e" {
				return iterations
			}
		}
	}
}

func getInstructions(input []string) []instruction {
	instructions := make([]instruction, 0)

	for _, line := range input {
		splitLine := strings.Split(line, " ")

		newInstruction := instruction{
			from: splitLine[0],
			to:   splitLine[2],
		}

		instructions = append(instructions, newInstruction)
	}

	return instructions
}

type instruction struct {
	from string
	to   string
}
