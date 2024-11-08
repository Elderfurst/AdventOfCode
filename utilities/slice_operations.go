package utilities

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
