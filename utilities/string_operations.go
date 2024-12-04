package utilities

import (
	"strings"
	"unicode/utf8"
)

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

// Reverse returns a reversed version of the provided string
func Reverse(str string) string {
	size := len(str)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(str[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
