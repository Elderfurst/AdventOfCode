package main

import "fmt"

func main() {
	first := solvePassword("vzbxkghb")
	fmt.Println(first)

	second := solvePassword(first)
	fmt.Println(second)
}

func solvePassword(input string) string {
	currentValue := []rune(input)

	currentValue = incrementAttempt(currentValue)

	for {
		if checkValidity(currentValue) {
			break
		}

		currentValue = incrementAttempt(currentValue)
	}

	return string(currentValue)
}

func checkValidity(attempt []rune) bool {
	avoid := []rune{
		105, 108, 111,
	}

	if containsAnyRune(attempt, avoid) {
		return false
	}

	if !containsIncreasingWindow(attempt) {
		return false
	}

	if !containsDoubles(attempt) {
		return false
	}

	return true
}

func incrementAttempt(attempt []rune) []rune {
	minValue := rune(97)
	maxValue := rune(122)

	bumpNext := true

	for i := len(attempt) - 1; i >= 0; i-- {
		if bumpNext {
			increase := attempt[i] + 1

			if increase > maxValue {
				increase = minValue
			} else {
				bumpNext = false
			}

			attempt[i] = increase
		} else {
			break
		}
	}

	return attempt
}

func containsAnyRune(attempt []rune, avoid []rune) bool {
	for _, bad := range avoid {
		for _, v := range attempt {
			if v == bad {
				return true
			}
		}
	}

	return false
}

func containsIncreasingWindow(attempt []rune) bool {
	for i := 0; i < len(attempt)-2; i++ {
		if attempt[i]+1 == attempt[i+1] && attempt[i+1]+1 == attempt[i+2] {
			return true
		}
	}

	return false
}

func containsDoubles(attempt []rune) bool {
	first := rune(-1)
	second := rune(-1)

	for i := 0; i < len(attempt)-1; i++ {
		current := attempt[i]
		next := attempt[i+1]

		if current == next {
			if first < 0 {
				first = current
			} else if first > 0 && current != first {
				second = current
			}
		}
	}

	if first > 0 && second > 0 {
		return true
	} else {
		return false
	}
}
