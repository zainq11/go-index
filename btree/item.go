package btree

import i "indexers/index"

type item[T i.Key] struct {
	before *node[T]
	key    T
	value  i.Value
	after  *node[T]
}

func newItem[K i.Key](before *node[K], k K, v i.Value, after *node[K]) *item[K] {
	return &item[K]{
		before: before,
		key:    k,
		value:  v,
		after:  after,
	}
}

func (i item[T]) setBefore(n *node[T]) {
	i.before = n
}

func (i item[T]) setAfter(n *node[T]) {
	i.after = n
}

func (i item[T]) less(other item[T]) bool {
	return i.key < other.key
}
