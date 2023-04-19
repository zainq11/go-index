package btree

import i "indexers/index"

type item[T i.Key] struct {
	before *node[T]
	v      []*i.Value
	k      T
	after  *node[T]
}

func newItem[K i.Key](k K, v i.Value) *item[K] {
	return &item[K]{}
}

func (i item[T]) less(other item[T]) bool {
	return i.k < other.k
}
