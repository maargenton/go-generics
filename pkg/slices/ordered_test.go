package slices_test

import (
	"testing"

	"github.com/maargenton/go-generics/pkg/slices"
	"github.com/maargenton/go-testpredicate/pkg/require"
)

func TestSort(t *testing.T) {
	var less = func(a, b int) bool { return a < b }
	require.That(t, slices.Sort([]int{3, 1, 2}, less)).Eq([]int{1, 2, 3})

	var greater = func(a, b int) bool { return a > b }
	require.That(t, slices.Sort([]int{3, 1, 2}, greater)).Eq([]int{3, 2, 1})

	var lessf = func(a, b float64) bool { return int(a) < int(b) }
	var v = []float64{3.3, 1.1, 2.2, 3.1}
	var sorted = []float64{1.1, 2.2, 3.3, 3.1}
	require.That(t, slices.Sort(v, lessf)).Eq(sorted)
}

func TestStableSort(t *testing.T) {
	var v = []float64{
		1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9,
		2.0, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9,
	}
	var less = func(a, b float64) bool { return int(a) > int(b) }
	var sorted = []float64{ // by integral part only
		2.0, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9,
		1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9,
	}

	require.That(t, slices.StableSort(v, less)).Eq(sorted)
}

func TestUniq(t *testing.T) {
	var v = []int{1, 2, 3, 1, 2, 4}
	require.That(t, slices.Uniq(v)).Eq([]int{1, 2, 3, 4})
}

func TestMin(t *testing.T) {
	require.That(t, slices.Min([]int{3, 1, 2})).Eq(1)
	require.That(t, slices.Min([]int{3})).Eq(3)
	require.That(t, slices.Min([]int{})).Eq(0)
}

func TestMax(t *testing.T) {
	require.That(t, slices.Max([]int{2, 3, 4, 1})).Eq(4)
	require.That(t, slices.Max([]int{3})).Eq(3)
	require.That(t, slices.Max([]int{})).Eq(0)
}

func TestMinMax(t *testing.T) {
	var min, max = slices.MinMax([]int{2, 3, 4, 1})
	require.That(t, min).Eq(1)
	require.That(t, max).Eq(4)

	min, max = slices.MinMax([]int{3})
	require.That(t, min).Eq(3)
	require.That(t, max).Eq(3)

	min, max = slices.MinMax([]int{})
	require.That(t, min).Eq(0)
	require.That(t, max).Eq(0)
}

func TestSorBy(t *testing.T) {
	var v = []int{3, 1, 2}
	var r = slices.SortBy(v, func(a int) int { return -a })
	require.That(t, r).Eq([]int{3, 2, 1})
}

func TestStableSortBy(t *testing.T) {
	var v = []float64{
		1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9,
		2.0, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9,
	}
	var by = func(a float64) int { return int(-a) }
	var sorted = []float64{ // by integral part only
		2.0, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9,
		1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9,
	}

	require.That(t, slices.StableSortBy(v, by)).Eq(sorted)
}

func TestUniqBy(t *testing.T) {
	var v = []float64{1.1, 2.1, 3.1, 1.2, 2.2, 4.2}
	var by = func(a float64) int { return int(a) }
	var uniq = []float64{1.1, 2.1, 3.1, 4.2}
	require.That(t, slices.UniqBy(v, by)).Eq(uniq)
}

func TestMinBy(t *testing.T) {
	var v = []float64{2.1, 3.1, 1.1, 3.2, 1.2, 2.2}
	var by = func(a float64) int { return int(a) }
	require.That(t, slices.MinBy(v, by)).Eq(1.1)
}

func TestMinByEmpty(t *testing.T) {
	var by = func(a float64) int { return int(a) }
	require.That(t, slices.MinBy([]float64{}, by)).Eq(0.0)
}

func TestMaxBy(t *testing.T) {
	var v = []float64{2.1, 3.1, 1.1, 3.2, 1.2, 2.2}
	var by = func(a float64) int { return int(a) }
	require.That(t, slices.MaxBy(v, by)).Eq(3.1)
}

func TestMaxByEmpty(t *testing.T) {
	var by = func(a float64) int { return int(a) }
	require.That(t, slices.MaxBy([]float64{}, by)).Eq(0.0)
}

func TestMinMaxBy(t *testing.T) {
	var v = []float64{2.1, 3.1, 1.1, 3.2, 1.2, 2.2}
	var by = func(a float64) int { return int(a) }

	var min, max = slices.MinMaxBy(v, by)
	require.That(t, min).Eq(1.1)
	require.That(t, max).Eq(3.1)
}

func TestMinMaxByEmpty(t *testing.T) {
	var by = func(a float64) int { return int(a) }

	var min, max = slices.MinMaxBy([]float64{}, by)
	require.That(t, min).Eq(0.0)
	require.That(t, max).Eq(0.0)
}
