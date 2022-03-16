package slices_test

import (
	"testing"

	"github.com/maargenton/go-generics/pkg/slices"
)

func BenchmarkMap100(b *testing.B) {
	var n = 100
	var v = make([]int, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, i)
	}

	for n := 0; n < b.N; n++ {
		slices.Map(v, func(v int) float32 {
			return 1.25 * float32(v)
		})
	}
}

func BenchmarkMap10K(b *testing.B) {
	var n = 10000
	var v = make([]int, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, i)
	}

	for n := 0; n < b.N; n++ {
		slices.Map(v, func(v int) float32 {
			return 1.25 * float32(v)
		})
	}
}

func BenchmarkFlatMap100(b *testing.B) {
	var n = 100
	var v = make([]int, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, i)
	}

	for n := 0; n < b.N; n++ {
		slices.FlatMap(v, func(v int) []float32 {
			return []float32{1.25 * float32(v)}
		})
	}
}

func BenchmarkFlatMap10K(b *testing.B) {
	var n = 10000
	var v = make([]int, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, i)
	}

	for n := 0; n < b.N; n++ {
		slices.FlatMap(v, func(v int) []float32 {
			return []float32{1.25 * float32(v)}
		})
	}
}

func BenchmarkFilterMap100(b *testing.B) {
	var n = 100
	var v = make([]int, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, i)
	}

	for n := 0; n < b.N; n++ {
		slices.FilterMap(v, func(v int) (float32, bool) {
			return 1.25 * float32(v), true
		})
	}
}

func BenchmarkFilterMap10K(b *testing.B) {
	var n = 10000
	var v = make([]int, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, i)
	}

	for n := 0; n < b.N; n++ {
		slices.FilterMap(v, func(v int) (float32, bool) {
			return 1.25 * float32(v), true
		})
	}
}
