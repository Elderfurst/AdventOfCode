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
	inputLines := utilities.ReadInputFileToLines("input.txt")
	commands := parseInputLines(inputLines)

	completedCommands := 0
	signals := make(map[string]uint16)

	for completedCommands < len(commands) {
		for i := 0; i < len(commands); i++ {
			currentCommand := commands[i]
			if !currentCommand.complete {
				switch currentCommand.action {
				case "VALUE":
					value, found := getValue(currentCommand.inputs[0], signals)
					if found {
						signals[currentCommand.output] = value
						currentCommand.complete = true
						completedCommands++
					}
				case "AND":
					first, foundFirst := getValue(currentCommand.inputs[0], signals)
					if foundFirst {
						second, foundSecond := getValue(currentCommand.inputs[1], signals)
						if foundSecond {
							signals[currentCommand.output] = first & second
							currentCommand.complete = true
							completedCommands++
						}
					}
				case "OR":
					first, foundFirst := getValue(currentCommand.inputs[0], signals)
					if foundFirst {
						second, foundSecond := getValue(currentCommand.inputs[1], signals)
						if foundSecond {
							signals[currentCommand.output] = first | second
							currentCommand.complete = true
							completedCommands++
						}
					}
				case "LSHIFT":
					first, found := getValue(currentCommand.inputs[0], signals)
					if found {
						second, _ := getValue(currentCommand.inputs[1], signals)

						signals[currentCommand.output] = first << second
						currentCommand.complete = true
						completedCommands++
					}
				case "RSHIFT":
					first, found := getValue(currentCommand.inputs[0], signals)
					if found {
						second, _ := getValue(currentCommand.inputs[1], signals)

						signals[currentCommand.output] = first >> second
						currentCommand.complete = true
						completedCommands++
					}
				case "NOT":
					input, found := getValue(currentCommand.inputs[0], signals)
					if found {
						signals[currentCommand.output] = ^input
						currentCommand.complete = true
						completedCommands++
					}
				}
			}
		}
	}

	fmt.Println(signals["a"])
}

func partTwo() {
	inputLines := utilities.ReadInputFileToLines("input.txt")
	commands := parseInputLines(inputLines)

	completedCommands := 0
	signals := make(map[string]uint16)

	for completedCommands < len(commands) {
		for i := 0; i < len(commands); i++ {
			currentCommand := commands[i]
			if !currentCommand.complete {
				switch currentCommand.action {
				case "VALUE":
					value, found := getValue(currentCommand.inputs[0], signals)
					if found {
						if currentCommand.output == "b" {
							value = 956
						}

						signals[currentCommand.output] = value
						currentCommand.complete = true
						completedCommands++
					}
				case "AND":
					first, foundFirst := getValue(currentCommand.inputs[0], signals)
					if foundFirst {
						second, foundSecond := getValue(currentCommand.inputs[1], signals)
						if foundSecond {
							signals[currentCommand.output] = first & second
							currentCommand.complete = true
							completedCommands++
						}
					}
				case "OR":
					first, foundFirst := getValue(currentCommand.inputs[0], signals)
					if foundFirst {
						second, foundSecond := getValue(currentCommand.inputs[1], signals)
						if foundSecond {
							signals[currentCommand.output] = first | second
							currentCommand.complete = true
							completedCommands++
						}
					}
				case "LSHIFT":
					first, found := getValue(currentCommand.inputs[0], signals)
					if found {
						second, _ := getValue(currentCommand.inputs[1], signals)

						signals[currentCommand.output] = first << second
						currentCommand.complete = true
						completedCommands++
					}
				case "RSHIFT":
					first, found := getValue(currentCommand.inputs[0], signals)
					if found {
						second, _ := getValue(currentCommand.inputs[1], signals)

						signals[currentCommand.output] = first >> second
						currentCommand.complete = true
						completedCommands++
					}
				case "NOT":
					input, found := getValue(currentCommand.inputs[0], signals)
					if found {
						signals[currentCommand.output] = ^input
						currentCommand.complete = true
						completedCommands++
					}
				}
			}
		}
	}

	fmt.Println(signals["a"])
}

func getValue(input string, signals map[string]uint16) (uint16, bool) {
	parsedInput, err := strconv.ParseUint(input, 10, 16)
	if err == nil {
		return uint16(parsedInput), true
	} else {
		existingValue, found := signals[input]

		return existingValue, found
	}
}

func parseInputLines(inputLines []string) []*command {
	commands := make([]*command, 0)

	for _, line := range inputLines {
		newCommand := &command{
			inputs:   make([]string, 0),
			complete: false,
		}

		splitLine := strings.Split(line, "->")

		rawInputs := strings.TrimSpace(splitLine[0])
		splitInputs := strings.Split(rawInputs, " ")

		if len(splitInputs) == 3 {
			newCommand.inputs = append(newCommand.inputs, splitInputs[0], splitInputs[2])
			newCommand.action = splitInputs[1]
		} else if len(splitInputs) == 2 {
			newCommand.action = splitInputs[0]
			newCommand.inputs = append(newCommand.inputs, splitInputs[1])
		} else {
			newCommand.inputs = append(newCommand.inputs, splitInputs[0])
			newCommand.action = "VALUE"
		}

		rawOutput := strings.TrimSpace(splitLine[1])

		newCommand.output = rawOutput

		commands = append(commands, newCommand)
	}

	return commands
}

type command struct {
	inputs   []string
	action   string
	output   string
	complete bool
}
