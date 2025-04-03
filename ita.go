// Package ita provides adaptors for iter.Seq and iter.Seq2.
package ita

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// NumAt enables iter.Seq functionality for objects providing only a count function and an element-at-index function.
//
//   - function num returns the total number of elements.
//   - function at takes an index and returns the corresponding element.
//
// NumAt returns an iterator from 0 to num() - 1, yielding each element at the corresponding index.
func NumAt[N constraints.Integer, T any](num func() N, at func(N) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		n := int64(num())
		for i := range n {
			if !yield(at(N(i))) {
				return
			}
		}
	}
}

// NumAtIndex is a variation of NumAt that returns an iterator with indexes.
func NumAtIndex[N constraints.Integer, T any](num func() N, at func(N) T) iter.Seq2[N, T] {
	return func(yield func(N, T) bool) {
		n := int64(num())
		for i := range n {
			ni := N(i)
			if !yield(ni, at(ni)) {
				return
			}
		}
	}
}

// KeysValues enables iter.Seq functionality for objects providing only a keys function and an element-by-key function.
func KeysValues[K comparable, V any](keys func() []K, value func(K) V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range keys() {
			if !yield(k, value(k)) {
				return
			}
		}
	}
}

// KeysValuesSeq is a variation of KeysValues that accepts a sequence of keys.
func KeysValuesSeq[K comparable, V any](keys iter.Seq[K], value func(K) V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k := range keys {
			if !yield(k, value(k)) {
				return
			}
		}
	}
}

// Pipe applies a transformation function f to seq and pipes it to another operation.
func Pipe[T any](seq iter.Seq[T], f func(T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func Pipe2[K, V any](seq iter.Seq2[K, V], f func(K, V) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

// Filter filters seq and pipes it to another operation.
func Filter[T any](seq iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if !f(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

func Filter2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !f(k, v) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

// Exec applies a function f to each element in seq.
func Exec[T any](seq iter.Seq[T], f func(T)) {
	for v := range seq {
		f(v)
	}
}

func Exec2[K, V any](seq iter.Seq2[K, V], f func(K, V)) {
	for k, v := range seq {
		f(k, v)
	}
}

// TakeFirst returns iter.Seq that iterates indices or keys from iter.Seq2.
func TakeFirst[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range seq {
			if !yield(k) {
				return
			}
		}
	}
}

// TakeSecond returns iter.Seq that iterates values from iter.Seq2.
func TakeSecond[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				return
			}
		}
	}
}
