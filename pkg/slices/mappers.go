package slices

// Filter returns a copy of v that includes only the elements for which `f`
// returns true.
func Filter[T any](v []T, f func(a T) bool) []T {
	var r []T
	for _, a := range v {
		if f(a) {
			r = append(r, a)
		}
	}
	return r
}

// Map invokes `f` with each element of `v` and collects one result per element.
func Map[T any, U any](v []T, f func(a T) U) []U {
	var r = make([]U, 0, len(v))
	for _, a := range v {
		r = append(r, f(a))
	}
	return r
}

// FlatMap invokes `f` with each element of `v` and collects zero, one or more
// results per element.
func FlatMap[T any, U any](v []T, f func(a T) []U) []U {
	var r []U
	for _, a := range v {
		r = append(r, f(a)...)
	}
	return r
}

// FlatMap invokes `f` with each element of `v` and collects zero or one result
// per element.
func FilterMap[T any, U any](v []T, f func(a T) (U, bool)) []U {
	var r []U
	for _, a := range v {
		if aa, keep := f(a); keep {
			r = append(r, aa)
		}
	}
	return r
}

// ---------------------------------------------------------------------------
// Cons

// MapCons invokes `f` with each `Cons(n)` of `v` and collects one result per
// invocation.
func MapCons[T any, U any](v []T, n int, f func(a []T) U) []U {
	var r = make([]U, 0, len(v))
	EachCons(v, n, func(a []T) {
		r = append(r, f(a))
	})
	return r
}

// FlatMapCons invokes `f` with each `Cons(n)` of `v` and collects zero, one or
// more results per invocation.
func FlatMapCons[T any, U any](v []T, n int, f func(a []T) []U) []U {
	var r []U
	EachCons(v, n, func(a []T) {
		r = append(r, f(a)...)
	})
	return r
}

// FilterMapCons invokes `f` with each `Cons(n)` of `v` and collects zero or one
// result per invocation.
func FilterMapCons[T any, U any](v []T, n int, f func(a []T) (U, bool)) []U {
	var r []U
	EachCons(v, n, func(a []T) {
		if aa, keep := f(a); keep {
			r = append(r, aa)
		}
	})
	return r
}

// Cons
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// Slice

// MapSlice invokes `f` with each `Slice(n)` of `v` and collects one result
// per invocation.
func MapSlice[T any, U any](v []T, n int, f func(a []T) U) []U {
	var r = make([]U, 0, len(v))
	EachSlice(v, n, func(a []T) {
		r = append(r, f(a))
	})
	return r
}

// FlatMapSlice invokes `f` with each `Slice(n)` of `v` and collects zero, one
// or more results per invocation.
func FlatMapSlice[T any, U any](v []T, n int, f func(a []T) []U) []U {
	var r []U
	EachSlice(v, n, func(a []T) {
		r = append(r, f(a)...)
	})
	return r
}

// FilterMapSlice invokes `f` with each `Slice(n)` of `v` and collects zero or
// one result per invocation.
func FilterMapSlice[T any, U any](v []T, n int, f func(a []T) (U, bool)) []U {
	var r []U
	EachSlice(v, n, func(a []T) {
		if aa, keep := f(a); keep {
			r = append(r, aa)
		}
	})
	return r
}

// Slice
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// SliceBetween

// MapSliceBetween slices `v` according to  `slicer`, invokes `f` with each
// slice and collects one result per invocation.
func MapSliceBetween[T any, U any](v []T, slicer func(a, b T) bool, f func(a []T) U) []U {
	var r = make([]U, 0, len(v))
	EachSliceBetween(v, slicer, func(a []T) {
		r = append(r, f(a))
	})
	return r
}

// FlatMapSliceBetween slices `v` according to  `slicer`, invokes `f` with each
// slice and collects zero, one or more results per invocation.
func FlatMapSliceBetween[T any, U any](v []T, slicer func(a, b T) bool, f func(a []T) []U) []U {
	var r []U
	EachSliceBetween(v, slicer, func(a []T) {
		r = append(r, f(a)...)
	})
	return r
}

// FilterMapSliceBetween slices `v` according to  `slicer`, invokes `f` with
// each slice and collects zero or one result per invocation.
func FilterMapSliceBetween[T any, U any](v []T, slicer func(a, b T) bool, f func(a []T) (U, bool)) []U {
	var r []U
	EachSliceBetween(v, slicer, func(a []T) {
		if aa, keep := f(a); keep {
			r = append(r, aa)
		}
	})
	return r
}

// SliceBetween
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// SliceBy

// MapSliceBy slices `v` according to  `slicer`, invokes `f` with each slice and
// collects one result per invocation.
func MapSliceBy[T any, U comparable, V any](v []T, slicer func(a T) U, f func(a []T) V) []V {
	var r = make([]V, 0, len(v))
	EachSliceBy(v, slicer, func(a []T) {
		r = append(r, f(a))
	})
	return r
}

// FlatMapSliceBy slices `v` according to  `slicer`, invokes `f` with each slice
// and collects zero, one or more results per invocation.
func FlatMapSliceBy[T any, U comparable, V any](v []T, slicer func(a T) U, f func(a []T) []V) []V {
	var r []V
	EachSliceBy(v, slicer, func(a []T) {
		r = append(r, f(a)...)
	})
	return r
}

// FilterMapSliceBy slices `v` according to  `slicer`, invokes `f` with each
// slice and collects zero or one result per invocation.
func FilterMapSliceBy[T any, U comparable, V any](v []T, slicer func(a T) U, f func(a []T) (V, bool)) []V {
	var r []V
	EachSliceBy(v, slicer, func(a []T) {
		if aa, keep := f(a); keep {
			r = append(r, aa)
		}
	})
	return r
}

// SliceBy
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// Zip

// MapZip zips the slices of v into one tuple per matching index, invokes `f`
// with each tuple and collects one result per invocation.
func MapZip[T any, U comparable](v [][]T, f func(a []T) U) []U {
	var r = make([]U, 0, len(v))
	EachZip(v, func(a []T) {
		r = append(r, f(a))
	})
	return r
}

// FlatMapZip zips the slices of v into one tuple per matching index, invokes
// `f` with each tuple and collects zero, one or more results per invocation.
func FlatMapZip[T any, U comparable](v [][]T, f func(a []T) []U) []U {
	var r []U
	EachZip(v, func(a []T) {
		r = append(r, f(a)...)
	})
	return r
}

// FilterMapZip zips the slices of v into one tuple per matching index, invokes
// `f` with each tuple and collects zero or one result per invocation.
func FilterMapZip[T any, U comparable](v [][]T, f func(a []T) (U, bool)) []U {
	var r []U
	EachZip(v, func(a []T) {
		if aa, keep := f(a); keep {
			r = append(r, aa)
		}
	})
	return r
}

// Zip
// ---------------------------------------------------------------------------
