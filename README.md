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
2. Open VSCode on the workspace and update the go tools. The workspace contains
   local settings to use the installed beta rather than the default version. In
   the command palette, select "Go: Install/Update Tools".

## Motivation

With the long awaited introduction of generics right around the corner with the
upcoming release of Go 1.18, it is time to get right into it and build something
useful.

The new version of the language introduces type parameters for types and
functions. It does not include any new generic functions in the standard
library, but provides 3 experimental packages containing a few useful
definitions and functions.

One of the things commonly found in other languages that is still missing in Go
is a set of functional-style primitives to slice and dice and transform slices
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

#### Iteration and transformation functions

- `Each( []T, func(T) )`: Invoke the given function with each element of the
  input. The function does not return anything.
- `Map( []T, func(T)U )`: Invoke the given function with each element of the
  input. The function returns a single value per element that gets collected in
  the result.
- `FilterMap( []T, func(T)(U,bool) )`: Invoke the given function with each
  element of the input. The function returns an optional value per element that
  gets collected in the result. The function return value is discarded if the
  second return value is false.
- `FlatMap( []T, func(T)[]U )`: Invoke the given function with each element of
  the input. The function returns a slice of values per element that gets
  collected in the result.

#### Traversal modes

- `∅`: When no specific traversal mode is specified, the input is enumerated,
  one element at a  time.
- `Cons(n)`: Iterate over each contiguous overlapping n-tuple of the input. An
  input shorter than `n` result in no iteration.
- `Slice(n)`: Iterate over each contiguous disjoint n-tuple of the input. All
  iterations are of size `n` except the last one which can be shorter.
- `SliceBy( func(T)U )`: Iterate over each contiguous disjoint variable-size
  tuples of the input for which all elements result is the same value when
  invoking the given function.
- `SliceBetween( func(T,T)bool )`: The given function is invoked with each
  cons(2) of the input and the input is split at the point between the two
  elements when the function returns true. The iteration happens over each
  resulting split, which are contiguous disjoint variable-size tuples of the
  input.
- `Zip( ...[]T )`: Iterate over tuples formed by taking the same-index element
  in each input.

#### Composite function names

Composition is achieved by mixing an iteration/transform method and a traversal
mode. For example `FlatMapZip( []T{a,b,c}, f)` iterates over 3-tuple composed of
matching elements of a, b, and c, invokes f and appends the resulting elements
to the final results.

Some functions like `FlatMapSliceBetween()` expect two separate functions, one
for slicing, one for mapping. For readability, it might be good practice to
define one or both as local variables rather than inline.
