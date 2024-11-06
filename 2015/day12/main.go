package main

import (
	"adventofcode/utilities"
	"encoding/json"
	"fmt"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	inputLines := utilities.ReadInputFileToLines("2015/day12/input.txt")

	var jsonValue interface{}
	_ = json.Unmarshal([]byte(inputLines[0]), &jsonValue)

	sum := findSum(jsonValue)

	fmt.Println(sum)
}

func findSum(value interface{}) int {
	total := 0

	switch value.(type) {
	case string:
		return 0
	case float64:
		return int(value.(float64))
	case []interface{}:
		for _, v := range value.([]interface{}) {
			total += findSum(v)
		}
	case map[string]interface{}:
		for _, v := range value.(map[string]interface{}) {
			total += findSum(v)
		}
	}

	return total
}

func partTwo() {
	inputLines := utilities.ReadInputFileToLines("2015/day12/input.txt")

	var jsonValue interface{}
	_ = json.Unmarshal([]byte(inputLines[0]), &jsonValue)

	sum := findSumExceptValue(jsonValue, "red")

	fmt.Println(sum)
}

func findSumExceptValue(value interface{}, avoid string) int {
	total := 0

	switch value.(type) {
	case string:
		return 0
	case float64:
		return int(value.(float64))
	case []interface{}:
		for _, v := range value.([]interface{}) {
			total += findSumExceptValue(v, avoid)
		}
	case map[string]interface{}:
		containsAvoid := false
		for _, v := range value.(map[string]interface{}) {
			parsedValue, ok := v.(string)
			if ok && parsedValue == avoid {
				containsAvoid = true
			}
		}

		if !containsAvoid {
			for _, v := range value.(map[string]interface{}) {
				total += findSumExceptValue(v, avoid)
			}
		}
	}

	return total
}
