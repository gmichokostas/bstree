package bstree

import (
	"testing"
)

func TestBSTree_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		fields BSTree[string, int]
		want   bool
	}{
		{
			name:   "returns true when the tree is empty",
			fields: BSTree[string, int]{Root: nil},
			want:   true,
		},
		{
			name:   "returns false when the tree is not empty",
			fields: BSTree[string, int]{Root: &Node[string, int]{size: 0}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree[string, int]{
				Root: tt.fields.Root,
			}
			if got := tree.IsEmpty(); got != tt.want {
				t.Errorf("BSTree.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Size(t *testing.T) {
	tests := []struct {
		name   string
		fields BSTree[string, int]
		want   int
	}{
		{
			name:   "returns the size of the tree",
			fields: BSTree[string, int]{Root: nil},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree[string, int]{
				Root: tt.fields.Root,
			}
			if got := tree.Size(); got != tt.want {
				t.Errorf("BSTree.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Height(t *testing.T) {
	tests := []struct {
		name   string
		fields BSTree[string, int]
		want   int
	}{
		{
			name:   "returns the height of the tree",
			fields: BSTree[string, int]{Root: &Node[string, int]{size: 1, left: &Node[string, int]{size: 1}}},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree[string, int]{
				Root: tt.fields.Root,
			}
			if got := tree.Height(); got != tt.want {
				t.Errorf("BSTree.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Put(t *testing.T) {
	type args struct {
		key   string
		value int
	}
	type want struct {
		value int
		ok    bool
	}
	tests := []struct {
		name   string
		fields BSTree[string, int]
		args   args
		want   want
	}{
		{
			name:   "puts an item in the tree",
			fields: BSTree[string, int]{Root: nil},
			args:   args{key: "key", value: 1},
			want:   want{value: 1, ok: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree[string, int]{
				Root: tt.fields.Root,
			}
			tree.Put(tt.args.key, tt.args.value)
			if value, ok := tree.Get(tt.args.key); ok != tt.want.ok && value != tt.want.value {
				t.Errorf("BSTree.Find() = %v, want %v", value, tt.want.value)
			}
		})
	}
}

func TestBSTree_Get(t *testing.T) {
	type args struct {
		key   string
		value int
	}
	type want struct {
		value int
		ok    bool
	}
	tests := []struct {
		name   string
		fields BSTree[string, int]
		args   args
		want   want
		want1  bool
	}{
		{
			name:   "returns the value when it finds the key in the tree",
			fields: BSTree[string, int]{Root: &Node[string, int]{size: 1, key: "key", value: 1}},
			args:   args{key: "key", value: 1},
			want:   want{value: 1, ok: true},
		},
		{
			name:   "returns nil when it can't find the item in the tree",
			fields: BSTree[string, int]{Root: &Node[string, int]{size: 1, key: "key", value: 1}},
			args:   args{key: "foo", value: 1},
			want:   want{value: 0, ok: true},
			want1:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree[string, int]{
				Root: tt.fields.Root,
			}
			if got, ok := tree.Get(tt.args.key); ok != true && got != tt.want.value {
				t.Errorf("BSTree.Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_String(t *testing.T) {
	tests := []struct {
		name   string
		fields BSTree[string, int]
		want   string
	}{
		{
			name:   "returns the string representation of the tree",
			fields: BSTree[string, int]{Root: nil},
			want:   "<nil>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree[string, int]{
				Root: tt.fields.Root,
			}
			if got := tree.String(); got != tt.want {
				t.Errorf("BSTree.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Min(t *testing.T) {
	tests := []struct {
		name   string
		fields BSTree[string, int]
		want   string
	}{
		{
			name:   "returns the min item of the tree and true when found",
			fields: BSTree[string, int]{Root: &Node[string, int]{key: "key", value: 2, left: &Node[string, int]{key: "min", value: 1}}},
			want:   "min",
		},
		{
			name:   "returns empty string when there are no items in the tree",
			fields: BSTree[string, int]{Root: nil},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree[string, int]{
				Root: tt.fields.Root,
			}
			got, _ := tree.Min()
			if got != tt.want {
				t.Errorf("BSTree.Min() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestBSTree_DeleteMin(t *testing.T) {
	tests := []struct {
		name   string
		fields BSTree[string, int]
		want   error
	}{
		{
			name:   "Removes the smallest key and associated value from the symbol table.",
			fields: BSTree[string, int]{Root: &Node[string, int]{key: "foo", value: 2, left: &Node[string, int]{key: "min", value: 1}}},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree[string, int]{
				Root: tt.fields.Root,
			}
			if got := tree.DeleteMin(); got != tt.want {
				t.Errorf("BSTree.DeleteMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Delete(t *testing.T) {
	tests := []struct {
		name   string
		fields BSTree[string, int]
		args   string
		want   bool
	}{
		{
			name:   "Removes the specified key and its associated value from this symbol table",
			fields: BSTree[string, int]{Root: &Node[string, int]{key: "foo", value: 1, size: 1}},
			args:   "foo",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree[string, int]{
				Root: tt.fields.Root,
			}
			tree.Delete(tt.args)

			if got, ok := tree.Get(tt.args); ok == tt.want && got == tt.fields.Root.value {
				t.Errorf("BSTree.Get() got = %v, want %v", got, nil)
			}
		})
	}
}
