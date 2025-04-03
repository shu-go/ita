package ita

import "iter"

// Chain is a type that allows chaining of operations on a sequence of elements.
//
// [Pipe], [Filter], and [Exec] are more performant (speed and memory) with Chain.
//
// Example:
//
//	ita.Chain[string](slices.Values([]string{"a", "b", "c"})).
//	  Pipe(strings.ToUpper).
//	  Exec(func(s string) { fmt.Println(s) })
type Chain[T any] iter.Seq[T]

// Pipe applies a transformation function f to Chain and pipes it to another operation.
func (c Chain[T]) Pipe(f func(T) T) Chain[T] {
	return Chain[T](Pipe(c.Seq(), f))
}

// [Filter]
func (c Chain[T]) Filter(f func(T) bool) Chain[T] {
	return Chain[T](Filter(c.Seq(), f))
}

// [Exec]
func (c Chain[T]) Exec(f func(T)) {
	Exec(c.Seq(), f)
}

// Seq returns itself as iter.Seq.
func (c Chain[T]) Seq() iter.Seq[T] {
	return iter.Seq[T](c)
}
