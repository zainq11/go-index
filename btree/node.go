package btree

import (
	"indexers/index"
)

// A pointer to the parent node.
// A pointer to a slice of Items.
// A bool that tells us whether the node is a leaf.
type node[T index.Key] struct {
	Parent *node[T]
	Items  []*Item[T] `json:"Items"`
	Leaf   bool
}

func newNode[T index.Key](parent *node[T], Items []*Item[T], isLeaf bool) *node[T] {
	return &node[T]{
		Parent: parent,
		Items:  Items,
		Leaf:   isLeaf,
	}
}

func (n *node[T]) insert(k T, v index.Value) {
	n.insertItem(newItem[T](nil, k, v, nil))
}

func (n *node[T]) insertItem(it *Item[T]) {
	for i := 0; i < len(n.Items); i++ {

		curr := n.Items[i]
		if it.Key >= curr.Key {
			continue
		}

		if i == 0 {
			if it.After != nil {
				curr.Before = it.After
			} else {
				it.After = curr.Before
			}

			n.Items = append([]*Item[T]{it}, n.Items...)
		} else {
			prev := n.Items[i-1]

			if it.Before == nil {
				it.Before = prev.After
			} else {
				prev.After = it.Before
			}

			if it.After == nil {
				it.After = curr.Before
			} else {
				curr.Before = it.After
			}

			pre := append(n.Items[:i], it)
			n.Items = append(pre, n.Items[i:]...)
		}
		return
	}

	// Reached the end of items
	if len(n.Items) > 0 {
		last := n.Items[len(n.Items)-1]
		if it.Before == nil {
			it.Before = last.After
		} else {
			last.After = it.Before
		}
	}

	n.Items = append(n.Items, it)
}

// find will look for the Key k starting at node n
func (n *node[T]) find(k T) (*node[T], *Item[T]) {
	// if leaf, do equality check
	// if not leaf
	//	the Items are [10, 20, 30, 40, 50] and k is 25
	//	left = 10, right = 50, mid = 30
	//	left = 10, right = 30, mid = 20
	//	left = 20, right = 30, and difference between left and right is 1-index

	if len(n.Items) == 0 {
		return n, nil
	}

	left, right := 0, len(n.Items)-1

	// Binary search Item recursively
	for left < right {
		mid := left + (right-left)/2
		midItem := n.Items[mid]
		if midItem.Key == k {
			// there is no further node to lookup key
			if n.isLeaf() {
				return n, midItem
			}

			if midItem.After == nil {
				return n, nil
			}

			return midItem.After.find(k)
		}

		if right-left == 1 {
			leftItem := n.Items[left]
			if leftItem.After == nil {
				return n, nil
			}
			return n.Items[left].After.find(k)
		}

		if midItem.Key > k {
			right = mid
		} else {
			left = mid
		}
	}
	return n, nil
}

func (n *node[T]) isLeaf() bool {
	// TODO: implement this as an expression
	return n.Leaf
}

// splitAt splits node n into left and right. Point of splitting is 'at'
func (n *node[T]) splitAt(i int) (*node[T], *node[T]) {
	// init left and right node's slices
	// var l []*Item[T]
	// var r []*Item[T]

	// split n.Items into left and right sub-lists
	// l := l1
	// for curr := n.Items.Front(); curr != nil; curr = curr.Next() {
	// 	l.PushBack(curr.Value)
	// 	if curr.Value.(*Item[T]) == at {
	// 		l = l2
	// 		continue
	// 	}
	// }

	// return left and right child nodes
	return newNode[T](n.Parent, n.Items[:i], n.isLeaf()), newNode[T](n.Parent, n.Items[i:], n.isLeaf())
}

// promote turns node n into the parent of left and right
func (n *node[T]) promote(Item *Item[T], left *node[T], right *node[T]) {
	Item.Before = left
	Item.After = right

	if n.Parent == nil {
		n.Leaf = false
		left.Parent = n
		right.Parent = n
	} else {
		left.Parent = n.Parent
		right.Parent = n.Parent
		n.Parent.insertItem(Item)
	}
}
