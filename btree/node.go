package btree

import i "indexers/index"

type Node[T i.Key] struct {
	Parent *Node[T]
	Items  []*Item[T]
	isLeaf bool
}

func NewNode[T i.Key](parent *Node[T], items []*Item[T], isLeaf bool) *Node[T] {
	return &Node[T]{
		Parent: parent,
		Items:  items,
		isLeaf: isLeaf,
	}
}

func (n *Node[T]) insert() {
	// To insert an item to a node
	// - find the location, in between
	// - update before and next pointers of item, before and next

}

func (n *Node[T]) find(k T) (*Node[T], *Item[T]) {
	for i, item := range n.Items {
		if item.K == k {
			return n, item
		}

		if item.K > k {
			if n.IsLeaf() {
				return n, nil
			}

			return item.Before.find(k)
		}

		if i == len(n.Items)-1 {
			if n.IsLeaf() {
				return n, nil
			}

			return item.After.find(k)
		}
	}

	return n, nil
}

func (n *Node[T]) IsLeaf() bool {
	return n.isLeaf
}

func (n *Node[T]) median() *Item[T] {
	return n.Items[len(n.Items)/2]
}
