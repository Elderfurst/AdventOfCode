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
	input := utilities.ReadInputFileToLines("2015/day16/input.txt")
	familyTree := parseInputLines(input)

	trueSue := sue{
		number:     0,
		properties: make(map[string]int),
	}

	trueSue.properties["children"] = 2
	trueSue.properties["cats"] = 7
	trueSue.properties["samoyeds"] = 2
	trueSue.properties["pomeranians"] = 3
	trueSue.properties["akitas"] = 0
	trueSue.properties["vizslas"] = 0
	trueSue.properties["goldfish"] = 5
	trueSue.properties["cars"] = 2
	trueSue.properties["perfumes"] = 1

	trueSueNumber := 0

	for _, imposter := range familyTree {
		match := true

		for property, count := range imposter.properties {
			if value, exists := trueSue.properties[property]; exists {
				if value != count {
					match = false
					break
				}
			} else {
				match = false
				break
			}
		}

		if match {
			trueSueNumber = imposter.number
			break
		}
	}

	fmt.Println(trueSueNumber)
}

func partTwo() {
	input := utilities.ReadInputFileToLines("2015/day16/input.txt")
	familyTree := parseInputLines(input)

	trueSue := sue{
		number:     0,
		properties: make(map[string]int),
	}

	trueSue.properties["children"] = 2
	trueSue.properties["cats"] = 7
	trueSue.properties["samoyeds"] = 2
	trueSue.properties["pomeranians"] = 3
	trueSue.properties["akitas"] = 0
	trueSue.properties["vizslas"] = 0
	trueSue.properties["goldfish"] = 5
	trueSue.properties["cars"] = 2
	trueSue.properties["perfumes"] = 1

	trueSueNumber := 0

	for _, imposter := range familyTree {
		match := true

		for property, count := range imposter.properties {
			if value, exists := trueSue.properties[property]; exists {
				switch property {
				case "cats":
					if count <= value {
						match = false
					}
				case "trees":
					if count <= value {
						match = false
					}
				case "pomeranians":
					if count >= value {
						match = false
					}
				case "goldfish":
					if count >= value {
						match = false
					}
				default:
					if value != count {
						match = false
					}
				}
			} else {
				match = false
				break
			}
		}

		if match {
			trueSueNumber = imposter.number
			break
		}
	}

	fmt.Println(trueSueNumber)
}

func parseInputLines(inputLines []string) []sue {
	familyTree := make([]sue, 0)

	for _, line := range inputLines {
		splitLine := strings.Split(line, " ")
		number, _ := strconv.Atoi(strings.TrimSuffix(splitLine[1], ":"))

		aunt := sue{
			number:     number,
			properties: make(map[string]int),
		}

		for i := 2; i <= len(splitLine)-2; i += 2 {
			property := strings.TrimSuffix(splitLine[i], ":")
			value, _ := strconv.Atoi(strings.TrimSuffix(splitLine[i+1], ","))

			aunt.properties[property] = value
		}

		familyTree = append(familyTree, aunt)
	}

	return familyTree
}

type sue struct {
	number     int
	properties map[string]int
}
