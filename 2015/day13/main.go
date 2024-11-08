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
	inputLines := utilities.ReadInputFileToLines("2015/day13/input.txt")
	party := parseInputLines(inputLines)

	guests := utilities.GetMapKeys(party.guests)

	partyPermutations := utilities.GetPermutations(guests)

	maxHappiness := 0

	for _, partyPermutation := range partyPermutations {
		partyHappiness := 0

		for i := 0; i < len(partyPermutation); i++ {
			first := partyPermutation[i]
			second := partyPermutation[(i+1)%len(partyPermutation)]

			pairKey := first + "-" + second
			reversePairKey := second + "-" + first

			partyHappiness += party.pairs[pairKey]
			partyHappiness += party.pairs[reversePairKey]
		}

		if partyHappiness > maxHappiness {
			maxHappiness = partyHappiness
		}
	}

	fmt.Println(maxHappiness)
}

func partTwo() {
	inputLines := utilities.ReadInputFileToLines("2015/day13/input.txt")
	party := parseInputLines(inputLines)

	guests := utilities.GetMapKeys(party.guests)
	for _, guest := range guests {
		firstPair := "nick-" + guest
		secondPair := guest + "-nick"

		party.pairs[firstPair] = 0
		party.pairs[secondPair] = 0
	}

	guests = append(guests, "nick")

	partyPermutations := utilities.GetPermutations(guests)

	maxHappiness := 0

	for _, partyPermutation := range partyPermutations {
		partyHappiness := 0

		for i := 0; i < len(partyPermutation); i++ {
			first := partyPermutation[i]
			second := partyPermutation[(i+1)%len(partyPermutation)]

			pairKey := first + "-" + second
			reversePairKey := second + "-" + first

			partyHappiness += party.pairs[pairKey]
			partyHappiness += party.pairs[reversePairKey]
		}

		if partyHappiness > maxHappiness {
			maxHappiness = partyHappiness
		}
	}

	fmt.Println(maxHappiness)
}

func parseInputLines(inputLines []string) Party {
	party := Party{
		guests: make(map[string]bool),
		pairs:  make(map[string]int),
	}

	for _, line := range inputLines {
		splitLine := strings.Split(line, " ")
		first := splitLine[0]
		second := strings.TrimSuffix(splitLine[10], ".")

		party.guests[first] = true
		party.guests[second] = true

		sign := splitLine[2]
		value, _ := strconv.Atoi(splitLine[3])

		if sign == "lose" {
			value = value * -1
		}

		pairKey := first + "-" + second

		party.pairs[pairKey] = value
	}

	return party
}

type Party struct {
	guests map[string]bool
	pairs  map[string]int
}
