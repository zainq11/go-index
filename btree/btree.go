package btree

import i "indexers/index"

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
	if i != nil {
		i.v = append(i.v, v)
		return nil
	}

	if b.isFull(n) {
		// split
		// Create 2 nodes (n1 & n2) with half of the data each
		// Create a parent node p. p.parent = n.parent
		// n1.parent = p  and n2.parent = p
		b.split(n)
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
		return i.v
	}
	return nil
}

func (b *Btree[T]) split(n *node[T]) error {
	l := len(n.items)
	m := l / 2

	parent := n.parent
	median := n.items[m]

	top := newNode[T](parent, []*item[T]{median}, false)

	first := newNode[T](top, n.items[0:m], true)
	second := newNode[T](top, n.items[m+1:l], true)

	median.before = first
	median.after = second

	return nil
}

func (b *Btree[T]) isFull(n *node[T]) bool {
	return len(n.items) == b.Order-1
}
