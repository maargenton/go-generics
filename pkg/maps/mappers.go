package maps

// Map invokes `f` on each key-value pair of `m` and collects the returned keys
// and values into a new map.
func Map[T comparable, U any, R comparable, S any](
	m map[T]U, f func(k T, v U) (R, S)) (
	r map[R]S) {

	r = make(map[R]S)
	for k, v := range m {
		rk, rv := f(k, v)
		r[rk] = rv
	}
	return r
}

// FlatMap invokes `f` on each key-value pair of `m` and collects the returned
// keys and values into a new map. In this variant, `f` return a map of results
// with 0, 1 or more key-value pairs.
func FlatMap[T comparable, U any, R comparable, S any](
	m map[T]U, f func(k T, v U) map[R]S) (
	r map[R]S) {

	r = make(map[R]S)
	for k, v := range m {
		for rk, rv := range f(k, v) {
			r[rk] = rv
		}
	}
	return r
}
