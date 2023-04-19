package btree

import i "indexers/index"

type node[T i.Key] struct {
	parent *node[T]
	items  []*item[T]
	leaf   bool
}

func newNode[T i.Key](parent *node[T], items []*item[T], isLeaf bool) *node[T] {
	return &node[T]{
		parent: parent,
		items:  items,
		leaf:   isLeaf,
	}
}

func (n *node[T]) insert() {
	// To insert an item to a node
	// - find the location, in between
	// - update before and next pointers of item, before and next

}

func (n *node[T]) find(k T) (*node[T], *item[T]) {
	for i, item := range n.items {
		if item.k == k {
			return n, item
		}

		if item.k > k {
			if n.isLeaf() {
				return n, nil
			}

			return item.before.find(k)
		}

		if i == len(n.items)-1 {
			if n.isLeaf() {
				return n, nil
			}

			return item.after.find(k)
		}
	}

	return n, nil
}

func (n *node[T]) isLeaf() bool {
	// TODO: implement this as an expression
	return n.leaf
}

func (n *node[T]) median() *item[T] {
	return n.items[len(n.items)/2]
}
