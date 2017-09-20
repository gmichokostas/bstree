package bstree

import (
	"reflect"
	"testing"
)

type Int int

func (item Int) LessThan(other Item) bool {
	if item < other.(Int) {
		return true
	}

	return false
}

func (item Int) MoreThan(other Item) bool {
	if item > other.(Int) {
		return true
	}

	return false
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *BSTree
	}{
		{
			name: "returns a pointer to a new tree",
			want: &BSTree{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_IsEmpty(t *testing.T) {
	type fields struct {
		Root *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "returns true when the tree is empty",
			fields: fields{Root: nil},
			want:   true,
		},
		{
			name:   "returns false when the tree is not empty",
			fields: fields{Root: &Node{size: 0}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree{
				Root: tt.fields.Root,
			}
			if got := tree.IsEmpty(); got != tt.want {
				t.Errorf("BSTree.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Size(t *testing.T) {
	type fields struct {
		Root *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "returns the size of the tree",
			fields: fields{Root: nil},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree{
				Root: tt.fields.Root,
			}
			if got := tree.Size(); got != tt.want {
				t.Errorf("BSTree.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Height(t *testing.T) {
	type fields struct {
		Root *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "returns the height of the tree",
			fields: fields{Root: &Node{size: 1, left: &Node{size: 1}}},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree{
				Root: tt.fields.Root,
			}
			if got := tree.Height(); got != tt.want {
				t.Errorf("BSTree.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Put(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		item Item
	}
	type want struct {
		item Int
		ok   bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name:   "puts an item in the tree",
			fields: fields{Root: nil},
			args:   args{item: Int(1)},
			want:   want{item: Int(1), ok: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree{
				Root: tt.fields.Root,
			}
			tree.Put(tt.args.item)
			if item, ok := tree.Find(tt.args.item); ok != tt.want.ok && item != tt.want.item {
				t.Errorf("BSTree.Find() = %v, want %v", ok, tt.want.ok)
			}
		})
	}
}

func TestBSTree_Find(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		item Item
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Item
		want1  bool
	}{
		{
			name:   "returns the item and true when it find the item in the tree",
			fields: fields{Root: &Node{size: 1, item: Int(1)}},
			args:   args{item: Int(1)},
			want:   Int(1),
			want1:  true,
		},
		{
			name:   "returns the nil and false when it can't find the item in the tree",
			fields: fields{Root: &Node{size: 1, item: Int(1)}},
			args:   args{item: Int(2)},
			want:   nil,
			want1:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree{
				Root: tt.fields.Root,
			}
			got, got1 := tree.Find(tt.args.item)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BSTree.Find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BSTree.Find() got = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBSTree_String(t *testing.T) {
	type fields struct {
		Root *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns the string representation of the tree",
			fields: fields{Root: nil},
			want:   "<nil>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree{
				Root: tt.fields.Root,
			}
			if got := tree.String(); got != tt.want {
				t.Errorf("BSTree.String() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestBSTree_Min(t *testing.T) {
	type fields struct {
		Root *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   Item
		want1  bool
	}{
		{
			name:   "returns the min item of the tree and true when found",
			fields: fields{Root: &Node{item: Int(2), left: &Node{item: Int(1)}}},
			want:   Int(1),
			want1:  true,
		},
		{
			name:   "returns nil and false there are no items in the tree",
			fields: fields{Root: nil},
			want:   nil,
			want1:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree{
				Root: tt.fields.Root,
			}
			got, got1 := tree.Min()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BSTree.Min() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BSTree.Min() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name  string
		args  args
		want  *Node
		want1 bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := min(tt.args.node)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("min() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("min() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBSTree_DeleteMin(t *testing.T) {
	type fields struct {
		Root *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "deletes the min item of the tree and returns true",
			fields: fields{Root: &Node{item: Int(2), left: &Node{item: Int(1)}}},
			want:   true,
		},
		{
			name:   "returns false when there are no item in the tree to delete",
			fields: fields{Root: nil},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree{
				Root: tt.fields.Root,
			}
			if got := tree.DeleteMin(); got != tt.want {
				t.Errorf("BSTree.DeleteMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Delete(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		item Item
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Item
		want1  bool
	}{
		{
			name:   "deletes an item from the tree",
			fields: fields{Root: &Node{item: Int(1), size: 1}},
			args:   args{item: Int(1)},
			want:   nil,
			want1:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BSTree{
				Root: tt.fields.Root,
			}
			tree.Delete(tt.args.item)

			got, got1 := tree.Find(tt.args.item)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BSTree.Find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BSTree.Find() got = %v, want %v", got1, tt.want1)
			}
		})
	}
}
