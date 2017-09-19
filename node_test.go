package bstree

import "testing"

func TestNode_Size(t *testing.T) {
	type fields struct {
		size  int
		item  Item
		left  *Node
		right *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "returns the size of the node with zero children",
			fields: fields{size: 0},
			want:   0,
		},
		{
			name:   "returns the size of the node with one children",
			fields: fields{size: 1, left: &Node{size: 0}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				size:  tt.fields.size,
				item:  tt.fields.item,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := node.Size(); got != tt.want {
				t.Errorf("Node.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_String(t *testing.T) {
	type fields struct {
		size  int
		item  Item
		left  *Node
		right *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns the string representation of the node",
			fields: fields{size: 0},
			want:   "Item: <nil> (Left: <nil> Right: <nil>)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				size:  tt.fields.size,
				item:  tt.fields.item,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := node.String(); got != tt.want {
				t.Errorf("Node.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
