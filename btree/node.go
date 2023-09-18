package btree

import (
	"container/list"
	"indexers/index"
)

// import "container/list"

// node represents the node in the tree. It contains the following attributes:
//   - A pointer to the parent node.
//   - A pointer to a slice of items.
//   - A bool that tells us whether the node is a leaf.
type node[T index.Key] struct {
	parent *node[T]
	items  []*item[T]
	leaf   bool
}

func newNode[T index.Key](parent *node[T], items *[]item[T], isLeaf bool) *node[T] {
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
	// prev := n.items[0]
	for i := 0; i < len(n.items); i++ {
		curr := n.items[i]
		if k < curr.key {
			if n.isLeaf() {
				if i > 0 {
					prev := n.items[i-1]
					item := newItem[T](prev.after, k, v, curr.before)
				} else {
					item := newItem[T](nil, k, v, curr.before)
				}

				// joining items on new item
				prefix := append(n.items[:i], item)
				n.items = append(prefix, n.items[i:])
			} else {
				curr.before.insert(k, v)
			}
			return
		}
	}

	item := newItem[T](n.items[len(*n.items)-1].after, k, v, nil)
	n.items = append(n.items, item)
}

// find will look for the key k starting at node n
func (n *node[T]) find(k T) (*node[T], *item[T]) {
	// if leaf, do equality check
	// if not leaf
	//	the items are [10, 20, 30, 40, 50] and k is 25
	//	left = 10, right = 50, mid = 30
	//	left = 10, right = 30, mid = 20
	//	left = 20, right = 30, and difference between left and right is 1-index

	left, right := 0, len(n.items)

	for left < right {
		mid := left + (right-left)/2
		midItem := n.items[mid]
		if midItem.key == k {
			if n.isLeaf() {
				return n, midItem
			}
			return midItem.after.find(k)
		}

		if right-left == 1 {
			return n.items[left].after.find(k)
		}

		if n.items[mid].key > k {
			right = mid
		} else {
			left = mid
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
