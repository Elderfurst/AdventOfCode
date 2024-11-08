package main

import (
	"adventofcode/utilities"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	test()
	partOne()
}

func partOne() {
	input := utilities.ReadInputFileToLines("2015/day15/input.txt")
	pantry := parseInputLines(input)

	bestScore := 0
	bestCalorieScore := 0

	frosting := pantry[0]
	candy := pantry[1]
	butterscotch := pantry[2]
	sugar := pantry[3]

	for i := 1; i <= 100; i++ {
		for j := 1; j+i <= 100; j++ {
			for k := 1; k+j+i <= 100; k++ {
				for l := 1; l+k+j+i <= 100; l++ {
					cScore := (frosting.capacity * i) + (candy.capacity * j) + (butterscotch.capacity * k) + (sugar.capacity * l)
					dScore := (frosting.durability * i) + (candy.durability * j) + (butterscotch.durability * k) + (sugar.durability * l)
					fScore := (frosting.flavor * i) + (candy.flavor * j) + (butterscotch.flavor * k) + (sugar.flavor * l)
					tScore := (frosting.texture * i) + (candy.texture * j) + (butterscotch.texture * k) + (sugar.texture * l)

					if cScore < 0 || dScore < 0 || fScore < 0 || tScore < 0 {
						continue
					}

					totalScore := cScore * dScore * fScore * tScore

					calories := (frosting.calories * i) + (candy.calories * j) + (butterscotch.calories * k) + (sugar.calories * l)

					if calories == 500 && totalScore > bestCalorieScore {
						bestCalorieScore = totalScore
					}

					if totalScore > bestScore {
						bestScore = totalScore
					}
				}
			}
		}
	}

	fmt.Println(bestScore)
	fmt.Println(bestCalorieScore)
}

func test() {
	butterscotch := ingredient{
		name:       "Butterscotch",
		capacity:   -1,
		durability: -2,
		flavor:     6,
		texture:    3,
		calories:   8,
	}

	cinnamon := ingredient{
		name:       "Cinnamon",
		capacity:   2,
		durability: 3,
		flavor:     -2,
		texture:    -1,
		calories:   3,
	}

	bestScore := 0
	bestCalorieScore := 0

	for i := 1; i <= 100; i++ {
		for j := 1; j+i <= 100; j++ {
			cScore := (butterscotch.capacity * i) + (cinnamon.capacity * j)
			dScore := (butterscotch.durability * i) + (cinnamon.durability * j)
			fScore := (butterscotch.flavor * i) + (cinnamon.flavor * j)
			tScore := (butterscotch.texture * i) + (cinnamon.texture * j)

			if cScore < 0 || dScore < 0 || fScore < 0 || tScore < 0 {
				continue
			}

			totalScore := cScore * dScore * fScore * tScore

			calories := (butterscotch.calories * i) + (cinnamon.calories * j)

			if calories == 500 && totalScore > bestCalorieScore {
				bestCalorieScore = totalScore
			}

			if totalScore > bestScore {
				bestScore = totalScore
			}
		}
	}

	fmt.Println(bestScore)
	fmt.Println(bestCalorieScore)
}

func parseInputLines(inputLines []string) []ingredient {
	pantry := make([]ingredient, 0)

	for _, line := range inputLines {
		splitLine := strings.Split(line, " ")
		name := strings.TrimSuffix(splitLine[0], ":")
		capacity, _ := strconv.Atoi(strings.TrimSuffix(splitLine[2], ","))
		durability, _ := strconv.Atoi(strings.TrimSuffix(splitLine[4], ","))
		flavor, _ := strconv.Atoi(strings.TrimSuffix(splitLine[6], ","))
		texture, _ := strconv.Atoi(strings.TrimSuffix(splitLine[8], ","))
		calories, _ := strconv.Atoi(splitLine[10])

		i := ingredient{
			name:       name,
			capacity:   capacity,
			durability: durability,
			flavor:     flavor,
			texture:    texture,
			calories:   calories,
		}

		pantry = append(pantry, i)
	}

	return pantry
}

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}
