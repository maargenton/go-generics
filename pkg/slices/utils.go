package slices

// Make is an helper function that creates a slice from individual elements
// using type inference to determine the type of the resulting slice. All input
// elements must have matching type.
func Make[T any](v ...T) []T {
	return v
}

// Private helpers

func imin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
