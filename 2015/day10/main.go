package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

var input = "1321131112"

func partOne() {
	currentValue := input

	for i := 0; i < 40; i++ {
		currentValue = lookAndSay(currentValue)
	}

	fmt.Println(len(currentValue))
}

func partTwo() {
	currentValue := input

	for i := 0; i < 50; i++ {
		currentValue = lookAndSay(currentValue)
	}

	fmt.Println(len(currentValue))
}

func lookAndSay(value string) string {
	expanded := strings.Split(value, "")
	newValue := make([]string, 0)

	currentValue := expanded[0]
	valueCounter := 1

	for i := 0; i < len(expanded)-1; i++ {
		nextValue := expanded[i+1]

		if nextValue == currentValue {
			valueCounter++
			continue
		} else {
			newValue = append(newValue, strconv.Itoa(valueCounter))
			newValue = append(newValue, currentValue)

			currentValue = nextValue
			valueCounter = 1
		}
	}

	newValue = append(newValue, strconv.Itoa(valueCounter))
	newValue = append(newValue, currentValue)

	return strings.Join(newValue, "")
}
