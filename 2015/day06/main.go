package main

import (
	"adventofcode/utilities"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputLines := utilities.ReadInputFileToLines("input.txt")

	instructions := parseInputToInstructions(inputLines)

	lights := make([][]bool, 1000)
	for i := range lights {
		lights[i] = make([]bool, 1000)
	}

	dimmers := make([][]int, 1000)
	for i := range dimmers {
		dimmers[i] = make([]int, 1000)
	}

	for _, instruction := range instructions {
		for i := instruction.startY; i <= instruction.endY; i++ {
			for j := instruction.startX; j <= instruction.endX; j++ {
				switch instruction.action {
				case "turnon":
					lights[i][j] = true
					dimmers[i][j]++
				case "turnoff":
					lights[i][j] = false
					if dimmers[i][j] > 0 {
						dimmers[i][j]--
					}
				case "toggle":
					lights[i][j] = !lights[i][j]
					dimmers[i][j] += 2
				}
			}
		}
	}

	lightsOn := 0
	totalBrightness := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if lights[i][j] {
				lightsOn++
			}
			totalBrightness += dimmers[i][j]
		}
	}

	fmt.Println(lightsOn)
	fmt.Println(totalBrightness)
}

func parseInputToInstructions(lines []string) []instruction {
	instructions := make([]instruction, 0)

	for _, line := range lines {
		line = strings.Replace(line, "turn on", "turnon", 1)
		line = strings.Replace(line, "turn off", "turnoff", 1)

		splitLine := strings.Split(line, " ")

		action := splitLine[0]

		startPoint := splitLine[1]
		parsedStart := strings.Split(startPoint, ",")
		startX, _ := strconv.Atoi(parsedStart[0])
		startY, _ := strconv.Atoi(parsedStart[1])

		endPoint := splitLine[3]
		parsedEnd := strings.Split(endPoint, ",")
		endX, _ := strconv.Atoi(parsedEnd[0])
		endY, _ := strconv.Atoi(parsedEnd[1])

		instruction := instruction{
			action: action,
			startX: startX,
			startY: startY,
			endX:   endX,
			endY:   endY,
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}

type instruction struct {
	action string
	startX int
	startY int
	endX   int
	endY   int
}
