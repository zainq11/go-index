package btree

import (
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
		root:               newNode(nil, []*Item[T]{}, true),
	}
}

func (b *Btree[T]) Insert(k T, v *i.Value) error {
	n, i := b.root.find(k)

	// key found, replace the value
	if i != nil {
		i.Value = v
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
		return i.Value
	}
	return nil
}

// split method  
func (b *Btree[T]) split(n *node[T]) {
	m := len(n.Items) / 2

	// split the list at median point
	median := n.Items[m]
	left, right := n.splitAt(m)

	// Promote the median node to the parent
	n.promote(median, left, right)
}

func (b *Btree[T]) isFull(n *node[T]) bool {
	return len(n.Items) == b.Order-1
}
