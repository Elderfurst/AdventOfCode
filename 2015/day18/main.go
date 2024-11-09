package main

import (
	"adventofcode/utilities"
	"fmt"
	"maps"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	input := utilities.ReadInputFileToLines("2015/day18/input.txt")

	lights := make(map[coordinate]int)

	for i, row := range input {
		for j, column := range strings.Split(row, "") {
			key := newCoordinate(j, i)

			switch column {
			case "#":
				lights[key] = 1
			case ".":
				lights[key] = 0
			}
		}
	}

	for step := 0; step < 100; step++ {
		nextStep := maps.Clone(lights)

		for key, value := range lights {
			nw := newCoordinate(key.x-1, key.y+1)
			n := newCoordinate(key.x, key.y+1)
			ne := newCoordinate(key.x+1, key.y+1)
			e := newCoordinate(key.x+1, key.y)
			se := newCoordinate(key.x+1, key.y-1)
			s := newCoordinate(key.x, key.y-1)
			sw := newCoordinate(key.x-1, key.y-1)
			w := newCoordinate(key.x-1, key.y)

			nwv := lights[nw]
			nv := lights[n]
			nev := lights[ne]
			ev := lights[e]
			sev := lights[se]
			sv := lights[s]
			swv := lights[sw]
			wv := lights[w]

			neighbors := nwv + nv + nev + ev + sev + sv + swv + wv

			switch value {
			case 1:
				if neighbors != 2 && neighbors != 3 {
					nextStep[key] = 0
				}
			case 0:
				if neighbors == 3 {
					nextStep[key] = 1
				}
			}
		}

		lights = maps.Clone(nextStep)
	}

	on := 0

	for _, value := range lights {
		if value > 0 {
			on++
		}
	}

	fmt.Println(on)
}

func partTwo() {
	input := utilities.ReadInputFileToLines("2015/day18/input.txt")

	lights := make(map[coordinate]int)

	for i, row := range input {
		for j, column := range strings.Split(row, "") {
			key := newCoordinate(j, i)

			switch column {
			case "#":
				lights[key] = 1
			case ".":
				lights[key] = 0
			}
		}
	}

	topLeft := newCoordinate(0, 0)
	topRight := newCoordinate(99, 0)
	bottomLeft := newCoordinate(0, 99)
	bottomRight := newCoordinate(99, 99)

	lights[topLeft] = 1
	lights[topRight] = 1
	lights[bottomLeft] = 1
	lights[bottomRight] = 1

	for step := 0; step < 100; step++ {
		nextStep := maps.Clone(lights)

		for key, value := range lights {
			if key == topLeft || key == topRight || key == bottomLeft || key == bottomRight {
				continue
			}

			nw := newCoordinate(key.x-1, key.y+1)
			n := newCoordinate(key.x, key.y+1)
			ne := newCoordinate(key.x+1, key.y+1)
			e := newCoordinate(key.x+1, key.y)
			se := newCoordinate(key.x+1, key.y-1)
			s := newCoordinate(key.x, key.y-1)
			sw := newCoordinate(key.x-1, key.y-1)
			w := newCoordinate(key.x-1, key.y)

			nwv := lights[nw]
			nv := lights[n]
			nev := lights[ne]
			ev := lights[e]
			sev := lights[se]
			sv := lights[s]
			swv := lights[sw]
			wv := lights[w]

			neighbors := nwv + nv + nev + ev + sev + sv + swv + wv

			switch value {
			case 1:
				if neighbors != 2 && neighbors != 3 {
					nextStep[key] = 0
				}
			case 0:
				if neighbors == 3 {
					nextStep[key] = 1
				}
			}
		}

		lights = maps.Clone(nextStep)
	}

	on := 0

	for _, value := range lights {
		if value > 0 {
			on++
		}
	}

	fmt.Println(on)
}

type coordinate struct {
	x int
	y int
}

func newCoordinate(x, y int) coordinate {
	return coordinate{
		x: x,
		y: y,
	}
}
