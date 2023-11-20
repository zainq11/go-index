package btree

import i "indexers/index"

type Item[T i.Key] struct {
	Before *node[T] `json:"before"`
	Key    T        `json:"key"`
	Value  i.Value  `json:"value"`
	After  *node[T] `json:"after"`
}

func newItem[K i.Key](before *node[K], k K, v i.Value, after *node[K]) *Item[K] {
	return &Item[K]{
		Before: before,
		Key:    k,
		Value:  v,
		After:  after,
	}
}

func (i Item[T]) setBefore(n *node[T]) {
	i.Before = n
}

func (i Item[T]) setAfter(n *node[T]) {
	i.After = n
}

func (i Item[T]) less(other Item[T]) bool {
	return i.Key < other.Key
}
