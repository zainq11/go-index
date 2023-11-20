package btree

import (
	i "indexers/index"
	"testing"
)

func TestBtree_InsertAndSearch(t *testing.T) {
	type arg struct {
		k int
		v i.Value
	}

	tests := []struct {
		name    string
		order   int
		args    []arg
		wantErr bool
	}{
		{
			name:  "success",
			order: 3,
			args: []arg{
				{
					k: 5,
					v: "e",
				},
				{
					k: 1,
					v: "a",
				},
				{
					k: 2,
					v: "b",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBTree(tt.order)
			for _, a := range tt.args {
				if err := b.Insert(a.k, &a.v); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

		})
	}
}
