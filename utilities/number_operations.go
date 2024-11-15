package utilities

import "cmp"

// Min returns the smaller of the two submitted values
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of the two submitted values
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
