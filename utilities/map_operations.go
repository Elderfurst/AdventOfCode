package utilities

// GetMapKeys returns a slice of the keys from a given map
func GetMapKeys[T comparable, V any](m map[T]V) []T {
	keys := make([]T, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}
