package utilities

import "slices"

// GetPermutations returns all permutations of the provided slice
func GetPermutations[T any](values []T) [][]T {
	var inner func(int, []T)
	var result [][]T

	inner = func(count int, slice []T) {
		if count == 1 {
			temp := make([]T, len(slice))
			copy(temp, slice)
			result = append(result, temp)
		} else {
			for i := 0; i < count; i++ {
				inner(count-1, slice)
				if count%2 == 1 {
					temp := slice[i]
					slice[i] = slice[count-1]
					slice[count-1] = temp
				} else {
					temp := slice[0]
					slice[0] = slice[count-1]
					slice[count-1] = temp
				}
			}
		}
	}

	inner(len(values), values)
	return result
}

// FillSpacesWithValues returns all permutations of values in the number of spaces provided
func FillSpacesWithValues[T any](values []T, spaces int) [][]T {
	var inner func(int)
	var result [][]T
	var subset []T

	inner = func(iterator int) {
		if iterator == spaces {
			temp := make([]T, len(subset))
			copy(temp, subset)
			result = append(result, temp)
			return
		}

		for _, value := range values {
			subset = append(subset, value)
			inner(iterator + 1)
			subset = subset[:len(subset)-1]
		}
	}

	inner(0)
	return result
}

// GetSubsets returns all possible subsets of the provided slice
func GetSubsets[T any](values []T) [][]T {
	var inner func(int)
	var result [][]T
	var subset []T

	inner = func(iterator int) {
		if iterator == len(values) {
			temp := make([]T, len(subset))
			copy(temp, subset)
			result = append(result, temp)
			return
		}

		subset = append(subset, values[iterator])
		inner(iterator + 1)
		subset = subset[:len(subset)-1]
		inner(iterator + 1)
	}

	inner(0)
	return result
}

// GetCount returns the number of times needle appears in haystack
func GetCount[S []T, T comparable](needle T, haystack S) int {
	var count int

	for {
		i := slices.Index(haystack, needle)
		if i == -1 {
			break
		}
		count++
		haystack = haystack[i+1:]
	}

	return count
}

// RemoveIndex returns a new slice equal to the passed slice minus the element specified at index
func RemoveIndex[T any](slice []T, index int) []T {
	temp := make([]T, 0)
	temp = append(temp, slice[:index]...)
	return append(temp, slice[index+1:]...)
}

// SwapIndices returns a new slice with the values at the two indices swapped
func SwapIndices[T any](slice []T, i, j int) []T {
	temp := make([]T, 0)

	var first int
	var second int

	if i < j {
		first = i
		second = j
	} else {
		first = j
		second = i
	}

	firstValue := slice[first]
	secondValue := slice[second]

	temp = append(temp, slice[:first]...)
	temp = append(temp, secondValue)
	temp = append(temp, slice[first+1:second]...)
	temp = append(temp, firstValue)
	return append(temp, slice[second+1:]...)
}
