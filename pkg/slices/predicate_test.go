package slices_test

import (
	"testing"

	"github.com/maargenton/go-generics/pkg/slices"
	"github.com/maargenton/go-testpredicate/pkg/require"
)

// Any
// All
// None
// One

func TestAny(t *testing.T) {
	var v = makeRange(4)
	require.That(t, slices.Any(v, func(a int) bool { return a < 2 })).IsTrue()
	require.That(t, slices.Any(v, func(a int) bool { return a > 3 })).IsFalse()
}

func TestAll(t *testing.T) {
	var v = makeRange(4)
	require.That(t, slices.All(v, func(a int) bool { return a <= 3 })).IsTrue()
	require.That(t, slices.All(v, func(a int) bool { return a < 3 })).IsFalse()
}

func TestNone(t *testing.T) {
	var v = makeRange(4)
	require.That(t, slices.None(v, func(a int) bool { return a > 3 })).IsTrue()
	require.That(t, slices.None(v, func(a int) bool { return a < 3 })).IsFalse()
}

func TestOne(t *testing.T) {
	var v = makeRange(4)
	require.That(t, slices.One(v, func(a int) bool { return a < 1 })).IsTrue()
	require.That(t, slices.One(v, func(a int) bool { return a < 2 })).IsFalse()
	require.That(t, slices.One(v, func(a int) bool { return a > 3 })).IsFalse()
}
