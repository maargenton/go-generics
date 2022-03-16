package maps_test

import (
	"strings"
	"testing"

	"github.com/maargenton/go-generics/pkg/maps"
	"github.com/maargenton/go-testpredicate/pkg/require"
)

func TestMap(t *testing.T) {

	var v = map[string]int{
		"foo": 1,
		"bar": 2,
	}
	var r = maps.Map(v, func(k string, v int) (string, int) {
		return strings.ToUpper(k), v * 2

	})
	require.That(t, r).MapKeys().IsEqualSet([]string{"FOO", "BAR"})
	require.That(t, r).Field("FOO").Eq(2)
	require.That(t, r).Field("BAR").Eq(4)
}

func TestFlatMap(t *testing.T) {

	var v = map[string]int{
		"foo": 1,
		"bar": 2,
	}
	var r = maps.FlatMap(v, func(k string, v int) map[string]int {
		return map[string]int{
			k:                  v,
			strings.ToUpper(k): v * 2,
		}
	})
	require.That(t, r).MapKeys().IsEqualSet([]string{"foo", "bar", "FOO", "BAR"})
	require.That(t, r).Field("foo").Eq(1)
	require.That(t, r).Field("bar").Eq(2)
	require.That(t, r).Field("FOO").Eq(2)
	require.That(t, r).Field("BAR").Eq(4)
}
