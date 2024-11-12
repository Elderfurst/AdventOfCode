package main

import (
	"adventofcode/utilities"
	"fmt"
	"strings"
)

var molecule = "ORnPBPMgArCaCaCaSiThCaCaSiThCaCaPBSiRnFArRnFArCaCaSiThCaCaSiThCaCaCaCaCaCaSiRnFYFArSiRnMgArCaSiRnPTiTiBFYPBFArSiRnCaSiRnTiRnFArSiAlArPTiBPTiRnCaSiAlArCaPTiTiBPMgYFArPTiRnFArSiRnCaCaFArRnCaFArCaSiRnSiRnMgArFYCaSiRnMgArCaCaSiThPRnFArPBCaSiRnMgArCaCaSiThCaSiRnTiMgArFArSiThSiThCaCaSiRnMgArCaCaSiRnFArTiBPTiRnCaSiAlArCaPTiRnFArPBPBCaCaSiThCaPBSiThPRnFArSiThCaSiThCaSiThCaPTiBSiRnFYFArCaCaPRnFArPBCaCaPBSiRnTiRnFArCaPRnFArSiRnCaCaCaSiThCaRnCaFArYCaSiRnFArBCaCaCaSiThFArPBFArCaSiRnFArRnCaCaCaFArSiRnFArTiRnPMgArF"

func main() {
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
