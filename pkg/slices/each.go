package slices

// Cons returns a slice of slices consisting of successive overlapping n-tuple
// of elements. All resulting slices have a length of `n`. The result is empty
// if the input is shorter than `n`.
func Cons[T any](v []T, n int) [][]T {
	var r = make([][]T, 0, len(v)-n+1)
	EachCons(v, n, func(v []T) {
		r = append(r, v)
	})
	return r
}

// EachCons invokes `f` with each element returned by Cons()
func EachCons[T any](v []T, n int, f func(v []T)) {
	var l = len(v) - n + 1
	for i := 0; i < l; i++ {
		f(v[i : i+n])
	}
}

// Slice returns a slice of slices consisting of successive non-overlapping
// n-tuple of elements. All resulting slices have a length of `n`, except the
// last one which may be shorter.
func Slice[T any](v []T, n int) [][]T {
	var r = make([][]T, 0, len(v)-n+1)
	EachSlice(v, n, func(v []T) {
		r = append(r, v)
	})
	return r
}

// EachSlice invokes `f` with each element returned by Slice()
func EachSlice[T any](v []T, n int, f func(v []T)) {
	var l = len(v)
	for i := 0; i < l; i += n {
		f(v[i:imin(l, i+n)])
	}
}

// SliceBetween invokes the slicer function with each consecutive elements a and
// b and splits the input between a and b if the slicer returns true. The result
// is a slice of slices containing all the resulting splits.
func SliceBetween[T any](v []T, slicer func(a, b T) bool) [][]T {
	var r [][]T
	EachSliceBetween(v, slicer, func(v []T) {
		r = append(r, v)
	})
	return r
}

// EachSliceBetween invokes `f` with each element returned by SliceBetween()
func EachSliceBetween[T any](v []T, slicer func(a, b T) bool, f func(v []T)) {
	var l = len(v)
	var s = 0
	for e := 1; e < l; e++ {
		if slicer(v[e-1], v[e]) {
			f(v[s:e])
			s = e
		}
	}
	f(v[s:l])
}

// SliceBy slices the input into contiguous slices for which the function
// 'slicer' returns the same value.
func SliceBy[T any, U comparable](v []T, slicer func(a T) U) [][]T {
	var r [][]T
	EachSliceBy(v, slicer, func(v []T) {
		r = append(r, v)
	})
	return r
}

// EachSliceBy invokes `f` with each element returned by SliceBy()
func EachSliceBy[T any, U comparable](v []T, slicer func(a T) U, f func(v []T)) {
	var p, n U
	var s = 0
	for e, vv := range v {
		n = slicer(vv)
		if p != n {
			if s != e {
				f(v[s:e])
			}
			s = e
			p = n
		}
	}
	if s != len(v) {
		f(v[s:])
	}
}

// Zip takes an list of slices and return a slice of slices where each resulting
// element is a tuple composed of the elements of each input at a given index.
// The length od the output matches the length of the shortest input.
func Zip[T any](v ...[]T) [][]T {
	var r [][]T
	EachZip(v, func(v []T) {
		r = append(r, v)
	})
	return r
}

// EachZip invokes `f` with each element returned by Zip(). The arguments must
// be passed as slice rather than variadic because of the trailing function
// argument.
func EachZip[T any](v [][]T, f func(v []T)) {
	var n = len(v)
	var l = 0
	for i, vv := range v {
		if i == 0 || len(vv) <= l {
			l = len(vv)
		}
	}
	for i := 0; i < l; i++ {
		var rr = make([]T, 0, n)
		for _, vv := range v {
			rr = append(rr, vv[i])
		}
		f(rr)
	}
}
