package bstree

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

// BSTree represents a binary tree
type BSTree[K constraints.Ordered, V any] struct {
	Root *Node[K, V]
}

// IsEmpty checks if the BSTree is empty
func (tree *BSTree[K, V]) IsEmpty() bool {
	return tree.Root == nil
}

// Size returns the size of the tree
func (tree *BSTree[K, V]) Size() int {
	return tree.Root.Size()
}

// Height returns the height of the tree
func (tree *BSTree[K, V]) Height() int {
	return height(tree.Root)
}

func height[K constraints.Ordered, V any](node *Node[K, V]) int {
	if node == nil {
		return 0
	}

	leftHeight := height(node.left)
	rightHeight := height(node.right)

	if leftHeight > rightHeight {
		leftHeight++
		return leftHeight
	}

	rightHeight++
	return rightHeight
}

// Put the item to the node
func (tree *BSTree[K, V]) Put(key K, val V) {
	tree.Root = put(tree.Root, key, val)

}

func put[K constraints.Ordered, V any](node *Node[K, V], key K, val V) *Node[K, V] {
	if node == nil {
		return &Node[K, V]{key: key, value: val, size: 1}
	}

	if key < node.key {
		node.left = put(node.left, key, val)
	} else if key > node.key {
		node.right = put(node.right, key, val)
	} else {
		node.value = val
	}

	node.size = 1 + node.left.Size() + node.right.Size()
	return node
}

// Find search for an item in the tree
func (tree *BSTree[K, V]) Get(key K) (V, bool) {
	return get(tree.Root, key)
}

func get[K constraints.Ordered, V any](node *Node[K, V], key K) (V, bool) {
	if node == nil {
		return *new(V), false
	}

	if key < node.key {
		return get(node.left, key)
	} else if key > node.key {
		return get(node.right, key)
	} else {
		return node.value, true
	}
}

// Min returns the smallest key in the table
func (tree *BSTree[K, V]) Min() (K, error) {
	if tree.IsEmpty() {
		return *new(K), errors.New("calls min() on empty tree")
	}

	node := min(tree.Root)

	return node.key, nil
}

func min[K constraints.Ordered, V any](node *Node[K, V]) *Node[K, V] {
	if (node.left) == nil {
		return node
	}

	return min(node.left)
}

// DeleteMin deletes the smallest key and associated value from the table.
func (tree *BSTree[K, V]) DeleteMin() error {
	if tree.IsEmpty() {
		return errors.New("Symbol table is empty")
	}

	tree.Root = deleteMin(tree.Root)
	return nil
}

func deleteMin[K constraints.Ordered, V any](node *Node[K, V]) *Node[K, V] {
	if node.left == nil {
		return node.right
	}

	node.left = deleteMin(node.left)
	node.size = node.left.Size() + node.right.Size() + 1

	return node
}

// Removes the specified key and its associated value from this symbol table
func (tree *BSTree[K, V]) Delete(key K) {
	tree.Root = delete(tree.Root, key)
}

func delete[K constraints.Ordered, V any](node *Node[K, V], key K) *Node[K, V] {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = delete(node.left, key)
	} else if key > node.key {
		node.right = delete(node.right, key)
	} else {

		if node.right == nil {
			return node.left
		}

		if node.left == nil {
			return node.right
		}

		t := node
		node = min(t.right)
		node.right = deleteMin(t.right)
		node.left = t.left
	}

	node.size = node.left.Size() + node.right.Size() + 1

	return node
}

// InOrderPrint traversal of the tree
func (tree *BSTree[K, V]) InOrderPrint() {
	inOrder(tree.Root)
}

func inOrder[K constraints.Ordered, V any](node *Node[K, V]) {
	if node != nil {
		inOrder(node.left)
		fmt.Printf("%v ", node.value)
		inOrder(node.right)
	}
}

func (tree *BSTree[K, V]) String() string {
	return fmt.Sprintf("%v", tree.Root)
}
