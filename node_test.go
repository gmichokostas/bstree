package bstree

import (
	"testing"
)

func TestNode_Size(t *testing.T) {
	tests := []struct {
		name   string
		fields Node[string, int]
		want   int
	}{
		{
			name:   "returns the size of the node with zero children",
			fields: Node[string, int]{size: 0},
			want:   0,
		},
		{
			name:   "returns the size of the node with one children",
			fields: Node[string, int]{size: 1, left: &Node[string, int]{size: 0}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node[string, int]{
				size:  tt.fields.size,
				key:   tt.fields.key,
				value: tt.fields.value,
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
	tests := []struct {
		name   string
		fields Node[string, int]
		want   string
	}{
		{
			name:   "returns the string representation of the node",
			fields: Node[string, int]{size: 0},
			want:   "Key:  Value: 0 (Left: <nil> Right: <nil>)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node[string, int]{
				size:  tt.fields.size,
				key:   tt.fields.key,
				value: tt.fields.value,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := node.String(); got != tt.want {
				t.Errorf("Node.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
