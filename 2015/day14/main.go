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

var totalTime = 2503

func partOne() {
	input := utilities.ReadInputFileToLines("2015/day14/input.txt")
	stable := parseInputLines(input)

	maxDistance := 0

	for _, deer := range stable {
		cycle := deer.duration + deer.rest

		cycleCount := totalTime / cycle

		distance := deer.speed * deer.duration * cycleCount

		partialCycle := totalTime % cycle

		if partialCycle >= deer.duration {
			distance += deer.speed * deer.duration
		} else {
			distance += deer.speed * partialCycle
		}

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	fmt.Println(maxDistance)
}

func partTwo() {
	input := utilities.ReadInputFileToLines("2015/day14/input.txt")
	stable := parseInputLines(input)

	distance := make(map[string]int)
	points := make(map[string]int)

	for i := 0; i < totalTime; i++ {
		for _, deer := range stable {
			cycle := deer.duration + deer.rest

			resting := (i % cycle) >= deer.duration

			if !resting {
				distance[deer.name] += deer.speed
			}
		}

		maxDistance := 0

		for _, deer := range stable {
			if distance[deer.name] > maxDistance {
				maxDistance = distance[deer.name]
			}
		}
		for _, deer := range stable {
			if distance[deer.name] == maxDistance {
				points[deer.name]++
			}
		}
	}

	maxPoints := 0

	for _, deer := range stable {
		if points[deer.name] > maxPoints {
			maxPoints = points[deer.name]
		}
	}

	fmt.Println(maxPoints)
}

func parseInputLines(inputLines []string) []reindeer {
	stable := make([]reindeer, 0)

	for _, line := range inputLines {
		splitLine := strings.Split(line, " ")
		name := splitLine[0]
		speed, _ := strconv.Atoi(splitLine[3])
		duration, _ := strconv.Atoi(splitLine[6])
		rest, _ := strconv.Atoi(splitLine[13])

		newDeer := reindeer{
			name:     name,
			speed:    speed,
			duration: duration,
			rest:     rest,
		}

		stable = append(stable, newDeer)
	}

	return stable
}

type reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
}
