package ita

import "iter"

type Chain2[K, V any] iter.Seq2[K, V]

// [Pipe2]
func (c Chain2[K, V]) Pipe(f func(K, V) (K, V)) Chain2[K, V] {
	return Chain2[K, V](Pipe2(c.Seq(), f))
}

// [Filter2]
func (c Chain2[K, V]) Filter(f func(K, V) bool) Chain2[K, V] {
	return Chain2[K, V](Filter2(c.Seq(), f))
}

// [Exec2]
func (c Chain2[K, V]) Exec(f func(K, V)) {
	Exec2(c.Seq(), f)
}

func (c Chain2[K, V]) Seq() iter.Seq2[K, V] {
	return iter.Seq2[K, V](c)
}

// [TakeFirst]
func (c Chain2[K, V]) TakeFirst() Chain[K] {
	return Chain[K](TakeFirst(c.Seq()))
}

// [TakeSecond]
func (c Chain2[K, V]) TakeSecond() Chain[V] {
	return Chain[V](TakeSecond(c.Seq()))
}
