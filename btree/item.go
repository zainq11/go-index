package btree

import i "indexers/index"

type Item[T i.Key] struct {
	Before *Node[T]
	V      []*i.Value
	K      T
	After  *Node[T]
}

func NewItem[K i.Key](k K, v i.Value) *Item[K] {
	return &Item[K]{}
}

func (i Item[T]) Less(other Item[T]) bool {
	return i.K < other.K
}
