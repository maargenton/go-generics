package slices_test

import (
	"testing"

	"github.com/maargenton/go-generics/pkg/slices"
	"github.com/maargenton/go-testpredicate/pkg/require"
)

func TestCons(t *testing.T) {
	var v = makeRange(4)
	require.That(t, slices.Cons(v, 3)).Eq([][]int{{0, 1, 2}, {1, 2, 3}})
	require.That(t, slices.Cons(v, 4)).Eq([][]int{{0, 1, 2, 3}})
	require.That(t, slices.Cons(v, 5)).Eq([][]int{})
}

func TestSlice(t *testing.T) {
	var v = makeRange(8)
	require.That(t, slices.Slice(v, 3)).Eq([][]int{{0, 1, 2}, {3, 4, 5}, {6, 7}})
	require.That(t, slices.Slice(v, 5)).Eq([][]int{{0, 1, 2, 3, 4}, {5, 6, 7}})
}

func TestSliceBetween(t *testing.T) {
	var slicer = func(a, b int) bool {
		return b < a
	}
	var v = []int{1, 3, 5, 2, 4, 6}
	require.That(t, slices.SliceBetween(v, slicer)).Eq(
		[][]int{{1, 3, 5}, {2, 4, 6}})
	require.That(t, slices.SliceBetween(makeRange(1), slicer)).Eq([][]int{{0}})
}

func TestSliceBy(t *testing.T) {
	var slicer = func(v int) int {
		return v / 3
	}
	require.That(t, slices.SliceBy(makeRange(10), slicer)).Eq(
		[][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9}})
	require.That(t, slices.SliceBy(makeRange(1), slicer)).Eq([][]int{{0}})
}

func TestZip(t *testing.T) {
	var a = []int{1, 2, 3}
	var b = []int{4, 5, 6, 7}
	require.That(t, slices.Zip(a, b)).Eq([][]int{{1, 4}, {2, 5}, {3, 6}})
	require.That(t, slices.Zip(a, b, nil)).Eq([][]int{})
}
