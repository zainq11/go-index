package btree

import (
	// i "indexers/index"
	"container/list"
	"reflect"
	"testing"
)

func Test_newNode(t *testing.T) {
	type X = int
	type args struct {
		parent *node[X]
		items  []*item[X]
		isLeaf bool
	}

	parent := &node[X]{
		parent: nil,
		items:  list.New(),
		leaf:   false,
	}
	tests := []struct {
		name          string
		args          args
		wantParent    *node[X]
		wantLeaf      bool
		wantItemsSize int
	}{
		{
			// TODO: Add test cases.
			"creation success",
			args{
				parent: parent,
				items:  list.List.New(),
				isLeaf: false,
			},
			parent,
			false,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newNode(tt.args.parent, tt.args.isLeaf)

			if got != nil {
				if !reflect.DeepEqual(got.parent, tt.wantParent) {
					t.Errorf("parent = %v, want %v", got.parent, tt.wantParent)
				}
				if tt.wantLeaf != got.isLeaf() {
					t.Errorf("leaf = %v, want %v", got.isLeaf(), tt.wantLeaf)
				}
				if tt.wantItemsSize != got.items.Len() {
					t.Errorf("size = %v, want %v", got.items.Len(), tt.wantItemsSize)
				}
			}
		})
	}
}

func Test_node_find(t *testing.T) {
	type X = int
	type args struct {
		k X
	}

	// prepare root with an item
	root := &node[X]{
		parent: nil,
		items:  list.New(),
		leaf:   true,
	}
	it := newItem[X](nil, 1, "value", nil)
	root.items.PushFront(it)

	tests := []struct {
		name     string
		n        *node[X]
		args     args
		destnode *node[X]
		destitem *item[X]
	}{
		{
			"finding item in root",
			root,
			args{
				k: it.key,
			},
			root,
			it,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, i := tt.n.find(tt.args.k)
			if !reflect.DeepEqual(n, tt.destnode) {
				t.Errorf("node[X].find() got = %v, want %v", n, tt.destnode)
			}
			if !reflect.DeepEqual(i, tt.destitem) {
				t.Errorf("node[X].find() got1 = %v, want %v", i, tt.destitem)
			}
		})
	}
}

// func Test_node_insert(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		n    *node[X]
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			n := &node[X]{}
// 			n.insert()
// 		})
// 	}
// }
//
//
// func Test_node_isLeaf(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		n    *node[X]
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			n := &node[X]{}
// 			if got := n.isLeaf(); got != tt.want {
// 				t.Errorf("node[X].isLeaf() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func Test_node_median(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		n    *node[X]
// 		want *item[X]
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			n := &node[X]{}
// 			if got := n.median(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("node[X].median() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
