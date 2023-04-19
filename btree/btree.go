package btree

import i "indexers/index"

// Btree top level construct for the Btree implementation
type Btree[T i.Key] struct {
	Order              int
	minChildrenForRoot int
	root               *Node[T]
}

func NewBTree[T i.Key](order int) *Btree[T] {
	// Create the Btree instance with order
	return &Btree[T]{
		Order:              order,
		minChildrenForRoot: 2,
		root:               &Node[T]{},
	}
}

func (b *Btree[T]) Insert(k T, v *i.Value) error {
	n, i := b.root.find(k)
	if i != nil {
		i.V = append(i.V, v)
		return nil
	}

	if b.IsFull(n) {
		// Split
		// Create 2 nodes (n1 & n2) with half of the data each
		// Create a parent node p. p.Parent = n.Parent
		// n1.parent = p  and n2.parent = p
		b.Split(n)
		return nil
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
		return i.V
	}
	return nil
}

func (b *Btree[T]) Split(n *Node[T]) error {
	l := len(n.Items)
	m := l / 2

	parent := n.Parent
	median := n.Items[m]

	top := NewNode[T](parent, []*Item[T]{median}, false)

	first := NewNode[T](top, n.Items[0:m], true)
	second := NewNode[T](top, n.Items[m+1:l], true)

	median.Before = first
	median.After = second

	return nil
}

func (b *Btree[T]) IsFull(n *Node[T]) bool {
	return len(n.Items) == b.Order-1
}
