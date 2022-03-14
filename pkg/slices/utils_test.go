package slices_test

import (
	"testing"

	"github.com/maargenton/go-generics/pkg/slices"
	"github.com/maargenton/go-testpredicate/pkg/require"
)

func makeRange(n int) []int {
	var v = make([]int, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, i)
	}
	return v
}

func TestMake(t *testing.T) {
	var v = slices.Make(1, 2, 3)
	require.That(t, v).Eq([]int{1, 2, 3})
}
