package utilities

import "strings"

// GetAllIndices returns every index of the substring in the string
func GetAllIndices(str, substr string) []int {
	indices := make([]int, 0)

	index := 0
	offset := 0
	window := len(substr)

	cut := str

	for {
		index = strings.Index(cut, substr)

		if index == -1 {
			break
		}

		indices = append(indices, offset+index)

		offset += index + window

		cut = str[offset:]
	}

	return indices
}
