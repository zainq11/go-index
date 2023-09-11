package btree

import (
	"container/list"
	i "indexers/index"
)

// Btree top level construct for the Btree implementation
type Btree[T i.Key] struct {
	Order              int
	minChildrenForRoot int
	root               *node[T]
}

func NewBTree[T i.Key](order int) *Btree[T] {
	// Create the Btree instance with order
	return &Btree[T]{
		Order:              order,
		minChildrenForRoot: 2,
		root:               &node[T]{},
	}
}

func (b *Btree[T]) Insert(k T, v *i.Value) error {
	n, i := b.root.find(k)

	// key found, replace the value
	if i != nil {
		i.value = v
		return nil
	}

	// node found, insert the value into node
	if n != nil {
		n.insert(k, v)
	}

	if b.isFull(n) {
		// split
		// Create 2 nodes (n1 & n2) with half of the data each
		// Create a parent node p. p.parent = n.parent
		// n1.parent = p  and n2.parent = p
		b.split(n)
	}

	return nil
}

func (b *Btree[T]) Delete(k T) bool {
	//TODO implement me
	panic("implement me")
}

func (b *Btree[T]) Search(k T) i.Value {
	_, i := b.root.find(k)
	if i != nil {
		return i.value
	}
	return nil
}

func (b *Btree[T]) split(n *node[T]) error {
	l := n.items.Len()
	m := l / 2

	curr := n.items.Front()
	
	for i := 0; i  < m; i++ {
		curr = curr.Next()	
	}

	median := curr.Value.(*item[T])

	items := list.New()
	items.PushFront(median)

	// split the list at median point
	left, right := n.splitAt(median)
	n.promote(median, left, right)

	return nil
}

func (b *Btree[T]) isFull(n *node[T]) bool {
	return n.items.Len() == b.Order-1
}
