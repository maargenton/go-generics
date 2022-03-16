package slices_test

import (
	"testing"

	"github.com/maargenton/go-generics/pkg/slices"
	"github.com/maargenton/go-testpredicate/pkg/require"
)

func TestFilter(t *testing.T) {
	var v = makeRange(4)
	var r = slices.Filter(v, func(a int) bool {
		return a%2 == 0
	})
	require.That(t, r).Eq([]int{0, 2})
}

func TestMap(t *testing.T) {
	var v = makeRange(4)
	var r = slices.Map(v, func(a int) float64 {
		return float64(a) * 1.5
	})
	require.That(t, r).Eq([]float64{0, 1.5, 3, 4.5})
}

func TestFlatMap(t *testing.T) {
	var v = makeRange(4)
	var r = slices.FlatMap(v, func(a int) []int {
		return makeRange(a)
	})
	require.That(t, r).Eq([]float64{0, 0, 1, 0, 1, 2})
}

func TestFilterMap(t *testing.T) {
	var v = makeRange(4)
	var r = slices.FilterMap(v, func(a int) (int, bool) {
		return a, a%2 == 0
	})
	require.That(t, r).Eq([]float64{0, 2})
}

func TestMapCons(t *testing.T) {
	var v = makeRange(4)
	var r = slices.MapCons(v, 3, func(a []int) int {
		return len(a)
	})
	require.That(t, r).Eq([]int{3, 3})
}

func TestFlatMapCons(t *testing.T) {
	var v = makeRange(4)
	var r = slices.FlatMapCons(v, 3, func(a []int) []int {
		return a
	})
	require.That(t, r).Eq([]int{0, 1, 2, 1, 2, 3})
}

func TestFilterMapCons(t *testing.T) {
	var v = makeRange(4)
	var r = slices.FilterMapCons(v, 3, func(a []int) (int, bool) {
		return a[2], a[2]%2 == 0
	})
	require.That(t, r).Eq([]int{2})
}

func TestMapSlice(t *testing.T) {
	var v = makeRange(4)
	var r = slices.MapSlice(v, 3, func(a []int) int {
		return len(a)
	})
	require.That(t, r).Eq([]int{3, 1})
}

func TestFlatMapSlice(t *testing.T) {
	var v = makeRange(4)
	var r = slices.FlatMapSlice(v, 3, func(a []int) []int {
		return append(append([]int{}, a...), -1)
	})
	require.That(t, v).Eq(makeRange(4))
	require.That(t, r).Eq([]int{0, 1, 2, -1, 3, -1})
}

func TestFilterMapSlice(t *testing.T) {
	var v = makeRange(4)
	var r = slices.FilterMapSlice(v, 3, func(a []int) (int, bool) {
		return len(a), len(a) < 3
	})
	require.That(t, r).Eq([]int{1})
}

func TestMapSliceBetween(t *testing.T) {
	var v = []int{1, 2, 4, 3, 5}
	var slicer = func(a, b int) bool { return b < a }
	var r = slices.MapSliceBetween(v, slicer, func(a []int) int {
		return len(a)
	})
	require.That(t, r).Eq([]int{3, 2})
}

func TestFlatMapSliceBetween(t *testing.T) {
	var v = []int{1, 2, 4, 3, 5}
	var slicer = func(a, b int) bool { return b < a }
	var r = slices.FlatMapSliceBetween(v, slicer, func(a []int) []int {
		return append(append([]int{}, a...), -1)
	})
	require.That(t, r).Eq([]int{1, 2, 4, -1, 3, 5, -1})
}

func TestFilterMapSliceBetween(t *testing.T) {
	var v = []int{1, 2, 4, 3, 5}
	var slicer = func(a, b int) bool { return b < a }
	var r = slices.FilterMapSliceBetween(v, slicer, func(a []int) (int, bool) {
		return len(a), len(a) < 3
	})
	require.That(t, r).Eq([]int{2})
}

func TestMapSliceBy(t *testing.T) {
	var v = []int{2, 4, 6, 3, 5}
	var slicer = func(a int) bool { return a%2 == 0 }
	var r = slices.MapSliceBy(v, slicer, func(a []int) int {
		return len(a)
	})
	require.That(t, r).Eq([]int{3, 2})
}

func TestFlatMapSliceBy(t *testing.T) {
	var v = []int{2, 4, 6, 3, 5}
	var slicer = func(a int) bool { return a%2 == 0 }
	var r = slices.FlatMapSliceBy(v, slicer, func(a []int) []int {
		return append(append([]int{}, a...), -1)
	})
	require.That(t, r).Eq([]int{2, 4, 6, -1, 3, 5, -1})
}

func TestFilterMapSliceBy(t *testing.T) {
	var v = []int{2, 4, 6, 3, 5}
	var slicer = func(a int) bool { return a%2 == 0 }
	var r = slices.FilterMapSliceBy(v, slicer, func(a []int) (int, bool) {
		return len(a), len(a) < 3
	})
	require.That(t, r).Eq([]int{2})
}

func TestMapZip(t *testing.T) {
	var a = []int{1, 2, 3, 4}
	var b = []int{5, 6, 7}
	var r = slices.MapZip(slices.Make(a, b), func(a []int) int {
		return a[0] + a[1]
	})
	require.That(t, r).Eq([]int{6, 8, 10})
}

func TestFlatMapZip(t *testing.T) {
	var a = []int{1, 2, 3, 4}
	var b = []int{5, 6, 7}
	var r = slices.FlatMapZip(slices.Make(a, b), func(a []int) []int {
		return append(append([]int{}, a...), -1)
	})
	require.That(t, r).Eq([]int{1, 5, -1, 2, 6, -1, 3, 7, -1})
}

func TestFilterMapZip(t *testing.T) {
	var a = []int{1, 2, 3, 4}
	var b = []int{5, 6, 7}
	var r = slices.FilterMapZip(slices.Make(a, b), func(a []int) (int, bool) {
		return a[0] + a[1], a[0]%2 != 0
	})
	require.That(t, r).Eq([]int{6, 10})
}
