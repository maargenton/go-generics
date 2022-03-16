package slices

// Reduce invokes `f` with each element of `v` and the updated memo from the
// previous invocation.
func Reduce[T, U any](v []T, memo U, f func(a T, memo U) U) U {
	for _, a := range v {
		memo = f(a, memo)
	}
	return memo
}

// Count invokes `f` with each element of `v` and counts the true results.
func Count[T any](s []T, f func(v T) bool) int {
	var count = 0
	for _, v := range s {
		if f(v) {
			count++
		}
	}
	return count
}

// CountBy invokes `f` with each element of `v` and counts the number of
// elements for each value returned by `f`.
func CountBy[T any, U comparable](s []T, f func(v T) U) map[U]int {
	var r = make(map[U]int)
	for _, v := range s {
		var vv = f(v)
		r[vv]++
	}
	return r
}

// GroupBy returns a map that groups all the elements of `v` by the value
// returned by `f`.
func GroupBy[T any, U comparable](v []T, f func(v T) U) map[U][]T {
	var r = make(map[U][]T)
	for _, a := range v {
		var aa = f(a)
		r[aa] = append(r[aa], a)
	}
	return r
}
