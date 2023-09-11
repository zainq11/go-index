package btree

import "indexers/index"
import "container/list"

// node represents the node in the tree. It contains the following attributes:
//   - A pointer to the parent node.
//   - A pointer to a slice of items.
//   - A bool that tells us whether the node is a leaf.
type node[T index.Key] struct {
	parent *node[T]
	items  *list.List
	leaf   bool
}

func newNode[T index.Key](parent *node[T], items *list.List, isLeaf bool) *node[T] {
	return &node[T]{
		parent: parent,
		items:  items,
		leaf:   isLeaf,
	}
}

func (n *node[T]) insert(k T, v *index.Value) {
	// To insert an item to a node
	// - find the location, in between
	// - update before and next pointers of item, before and next
	for curr := n.items.Front(); curr != nil; curr = curr.Next() {
		// do something with e.Value
		currItem := curr.Value.(*item[T])
		// Find the first node with value < key
		// We need to decide between
		//	- if leaf, Insert item in current node
		//	- else Choose next node
		if k < currItem.key {
			if n.isLeaf() {
				var prevItem *item[T]
				if curr.Prev() != nil {
					prevItem = curr.Prev().Value.(*item[T])
				}
				item := newItem[T](prevItem.after, k, v, currItem.before)
				n.items.InsertBefore(item, curr)

			}

		} else {
			currItem.before.insert(k, v)
		}
	}
}

// find will look for the key k starting at node n
func (n *node[T]) find(k T) (*node[T], *item[T]) {
	for curr := n.items.Front(); curr != nil; curr = curr.Next() {
		item := curr.Value.(*item[T])

		if item.key == k {
			return n, item
		}

		if item.key > k {
			if n.isLeaf() {
				return n, nil
			}
			return item.before.find(k)
		}

		if curr == n.items.Back() {
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

// splitAt splits node n into left and right. Point of splitting is 'at'
func (n *node[T]) splitAt(at *item[T]) (*node[T], *node[T]) {
	// init left and right node's lists
	l1 := list.New()
	l2 := list.New()

	// split n.items into left and right sub-lists
	l := l1
	for curr := n.items.Front(); curr != nil; curr = curr.Next() {
		l.PushBack(curr.Value)
		if curr.Value.(*item[T]) == at {
			l = l2
			continue
		}
	}

	// return left and right child nodes
	return newNode[T](n, l1, true), newNode[T](n, l2, true)
}

// promote turns node n into the parent of left and right
func (n *node[T]) promote(item *item[T], left *node[T], right *node[T]) {
	n.leaf = false
	item.before = left
	item.after = right
	n.items = list.New()
	n.items.PushBack(item)
}
