# go-generics

Go utility library building functional-style features on upcoming Go 1.18
generics.

[![Latest](
  https://img.shields.io/github/v/tag/maargenton/go-generics?color=blue&label=latest&logo=go&logoColor=white&sort=semver)](
  https://pkg.go.dev/github.com/maargenton/go-generics)
[![Build](
  https://img.shields.io/github/workflow/status/maargenton/go-generics/build?label=build&logo=github&logoColor=aaaaaa)](
  https://github.com/maargenton/go-generics/actions?query=branch%3Amaster)
[![Codecov](
  https://img.shields.io/codecov/c/github/maargenton/go-generics?label=codecov&logo=codecov&logoColor=aaaaaa&token=fVZ3ZMAgfo)](
  https://codecov.io/gh/maargenton/go-generics)
[![Go Report Card](
  https://goreportcard.com/badge/github.com/maargenton/go-generics)](
  https://goreportcard.com/report/github.com/maargenton/go-generics)


---------------------------


## Usage, pre 1.18 release

1. Install lates Go 1.18 beta

```
go install golang.org/dl/go1.18beta2@latest
go1.18beta2 download
```

## Motivation

I have been using Go professionally for many years. With the long awaited
introduction of generics right around the corner with the upcoming release of Go
1.18, I felt the urge to get right into it and do something useful.

The new version of the language introduces type parameters for types and
functions. It does not include any new generic functions in the standard
library, but provides 3 experimental packages containing a few useful
definitions and functions.

One of the things commonly found in other languages that I really miss in Go is
a set of functional-style primitives to slice and dice and transform slices
(like Ruby Enumerable or Javascript underscore.js). Slices in Go are one of the
few built-in type-parameterized types, making them a prime target for a library
of type-parameterized generic functions.

The target feature set is a collection of expressive functions to iterate
through (`Each`) and transform (`Map`, `FlatMap`) elements of a slice, either
one by one or by consecutive subsets, overlapping (`Cons`) or disjoint
(`Slice`).

## API Design

### Flexibility and performance

`FlatMap` transforms each element of the input into zero, one or more output
elements; it is the most flexible, but also the most expansive of the available
variants. `Map` is an optimized variant for 1-to-1 mapping that is about 10
times faster, `FilterMap` is an optimized variant for 1-to-(0 or 1) mapping that
is 5 times faster, and `Each` is an optimized variant for 1-to-0 mapping that
just iterates through the input without capturing any output.

The main motivation for providing so many variants of the basic `Map` function
is the significant runtime performance difference associated with each use-case,
and the benefits that can be gained by selecting the right one.

### Composition

In Ruby, functional primitives are chainable and allow for the definition of quite sophisticated processing with a compact syntax.

```ruby
array.each_cons(2).flat_map { |a, b| ... }
```

Unfortunately, with Go 1.18 generics, type parameters only apply to types and
functions, not methods. A generic type can have methods, but those methods do
not accept any additional type parameter. This makes a call-chaining style
syntax impossible to implement.

Consequently, it seems that the best we can do is to provide basic functions
(`Each`, `Map`, `FlatMap`) and selected composite functions for the most useful
cases. For example `slice.FlatMapCons(v, 2, ...)` is equivalent to
`slices.FlatMap(slices.Cons(v, 2), ...)` but without using intermediate storage
for the result of `Cons`.

x         | ∅         | Cons | Slice | SliceBetween | SliceBy
--------- | --------- | - | - | - | -
∅         | ∅         | Cons | Slice | SliceBetween | SliceBy
Each      | Each      | EachCons | EachSlice | EachSliceBetween | EachSliceBy
Map       | Map       | MapCons | MapSlice | MapSliceBetween | MapSliceBy
FilterMap | FilterMap | FilterMapCons | FilterMapSlice | FilterMapSliceBetween | FilterMapSliceBy
FlatMap   | FlatMap   | FlatMapCons | FlatMapSlice | FlatMapSliceBetween | FlatMapSliceBy
