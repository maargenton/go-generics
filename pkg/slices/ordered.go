package slices

import (
	// "constraints"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Sort returns a sorted copy of `v` according to the comparison function
// `less`. The original slice is not modified.
func Sort[T any](v []T, less func(a, b T) bool) []T {
	var r = append([]T{}, v...)
	slices.SortFunc(r, less)
	return r
}

// StableSort returns a sorted copy of `v` according to the comparison function
// `less`, preserving the original order of elements that compare equal. The
// original slice is not modified.
func StableSort[T any](v []T, less func(a, b T) bool) []T {
	var r = append([]T{}, v...)
	slices.SortStableFunc(r, less)
	return r
}

// Uniq returns a copy of `v` where only the first occurrence of a value is
// preserved. The input must be a slice of comparable values, but does not have
// to be sorted.
func Uniq[T comparable](v []T) []T {
	var m = make(map[T]struct{}, len(v))
	var r []T
	for _, a := range v {
		if _, ok := m[a]; !ok {
			r = append(r, a)
			m[a] = struct{}{}
		}
	}
	return r
}

// Min returns the minimum value of `v`. The type of `v` must be a slice with a
// value type that defines a strict ordered relationship. If `v` is empty, the
// function return the zero value of the underlying value type.
func Min[T constraints.Ordered](v []T) T {
	var min T
	for i, a := range v {
		if i == 0 || a < min {
			min = a
		}
	}
	return min
}

// Max returns the maximum value of `v`. The type of `v` must be a slice with a
// value type that defines a strict ordered relationship. If `v` is empty, the
// function return the zero value of the underlying value type.
func Max[T constraints.Ordered](v []T) T {
	var max T
	for i, a := range v {
		if i == 0 || a > max {
			max = a
		}
	}
	return max
}

// MinMax returns both the minimum abd maximum value of `v`. The type of `v`
// must be a slice with a value type that defines a strict ordered relationship.
// If `v` is empty, the function return a pair of zero values for the underlying
// value type.
func MinMax[T constraints.Ordered](v []T) (T, T) {
	var min, max T
	for i, a := range v {
		if i == 0 {
			min = a
			max = a
		}
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	return min, max
}

// ---

// SortBy returns a sorted copy of `v` to the natural sort order of the result
// of invoking `f` on each element. The original slice is not modified.
func SortBy[T any, U constraints.Ordered](v []T, f func(a T) U) []T {
	var r = append([]T{}, v...)
	var comp = func(a, b T) bool {
		var aa = f(a)
		var bb = f(b)
		return aa < bb
	}
	slices.SortFunc(r, comp)
	return r
}

// StableSortBy returns a sorted copy of `v` according to the natural sort order
// of the result of invoking `f` on each element, preserving the original order
// of elements that compare equal. The original slice is not modified.
func StableSortBy[T any, U constraints.Ordered](v []T, f func(a T) U) []T {
	var r = append([]T{}, v...)
	var comp = func(a, b T) bool {
		var aa = f(a)
		var bb = f(b)
		return aa < bb
	}
	slices.SortStableFunc(r, comp)
	return r
}

// UniqBy returns a copy of `v` where only the first occurrence of a value is
// preserved. Values are considered equal if the results of invoking `f` on them
// are equal. The input does not have to be sorted.
func UniqBy[T any, U comparable](v []T, f func(a T) U) []T {
	var m = make(map[U]struct{}, len(v))
	var r []T
	for _, a := range v {
		var aa = f(a)
		if _, ok := m[aa]; !ok {
			r = append(r, a)
			m[aa] = struct{}{}
		}
	}
	return r
}

// MinBy returns the first element of `v` for which the result of invoking `f`
// yields the smallest value. If `v` is empty, the zero value of the underlying
// value type is returned.
func MinBy[T any, U constraints.Ordered](v []T, f func(a T) U) T {
	var min T
	var minv U
	for i, a := range v {
		var aa = f(a)
		if i == 0 {
			min = a
			minv = aa
		}
		if aa < minv {
			min = a
			minv = aa
		}
	}
	return min
}

// MaxBy returns the first element of `v` for which the result of invoking `f`
// yields the largest value. If `v` is empty, the zero value of the underlying
// value type is returned.
func MaxBy[T any, U constraints.Ordered](v []T, f func(a T) U) T {
	var max T
	var maxv U
	for i, a := range v {
		var aa = f(a)
		if i == 0 {
			max = a
			maxv = aa
		}
		if aa > maxv {
			max = a
			maxv = aa
		}
	}
	return max
}

// MinMaxBy returns the first elements of `v` for which the result of invoking
// `f` yields the smallest and the largest values. If `v` is empty, zero values
// of the underlying value type are returned.
func MinMaxBy[T any, U constraints.Ordered](v []T, f func(a T) U) (T, T) {
	var min, max T
	var minv, maxv U
	for i, a := range v {
		var aa = f(a)
		if i == 0 {
			min = a
			max = a
			minv = aa
			maxv = aa
		}
		if aa < minv {
			min = a
			minv = aa
		}
		if aa > maxv {
			max = a
			maxv = aa
		}
	}
	return min, max
}
