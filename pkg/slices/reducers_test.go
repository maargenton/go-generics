package slices_test

import (
	"testing"

	"github.com/maargenton/go-generics/pkg/slices"
	"github.com/maargenton/go-testpredicate/pkg/require"
)

func TestReduce(t *testing.T) {
	var v = makeRange(4)
	var r = slices.Reduce(v, 0, func(a int, memo int) int {
		return memo + 1
	})
	require.That(t, r).Eq(4)
}

func TestCount(t *testing.T) {
	var v = makeRange(4)
	var r = slices.Count(v, func(a int) bool {
		return a%2 == 0
	})
	require.That(t, r).Eq(2)
}

func TestCountBy(t *testing.T) {
	var v = makeRange(5)
	var r = slices.CountBy(v, func(a int) string {
		if a%2 == 0 {
			return "even"
		}
		return "odd"
	})
	require.That(t, r).Field("even").Eq(3)
	require.That(t, r).Field("odd").Eq(2)
}

func TestGroupBy(t *testing.T) {
	var v = makeRange(5)
	var r = slices.GroupBy(v, func(a int) string {
		if a%2 == 0 {
			return "even"
		}
		return "odd"
	})
	require.That(t, r).Field("even").Eq([]int{0, 2, 4})
	require.That(t, r).Field("odd").Eq([]int{1, 3})
}
