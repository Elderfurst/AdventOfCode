package main

import (
	"adventofcode/utilities"
	"fmt"
	"math"
	"slices"
)

var input = 36000000

func main() {
	partOne()
	partTwo()
}

func partOne() {
	house := 0
	gifts := 0

	for gifts < input {
		house++

		gifts = getGiftCount(house)
	}

	fmt.Println(house)
}

func partTwo() {
	elf := 0
	presents := make(map[int]int)
	houses := make(map[int]bool)

	// idk loop long enough to check enough houses I guess?
	for len(houses) < 1000000 {
		elf++

		for i := 1; i <= 50; i++ {
			presents[i*elf] += elf * 11

			if presents[i*elf] >= input {
				houses[i*elf] = true
			}
		}
	}

	houseNumbers := utilities.GetMapKeys(houses)
	slices.Sort(houseNumbers)

	fmt.Println(houseNumbers[0])
}
func getGiftCount(house int) int {
	divisors := make([]int, 0)

	converted := int(math.Sqrt(float64(house)))

	for i := 1; i <= converted; i++ {
		if house%i == 0 {
			// Only add duplicate divisors once
			if house/i == i {
				divisors = append(divisors, i)
			} else {
				divisors = append(divisors, i, house/i)
			}
		}
	}

	gifts := 0

	for _, divisor := range divisors {
		gifts += divisor * 10
	}

	return gifts
}
