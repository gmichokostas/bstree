package bstree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Node is the basic data unit of the binary tree
type Node[K constraints.Ordered, V any] struct {
	size        int
	key         K
	value       V
	left, right *Node[K, V]
}

// Size returns the number of nodes in subtree
func (node *Node[K, V]) Size() int {
	if node == nil {
		return 0
	}

	return node.size
}

func (node *Node[K, V]) String() string {
	return fmt.Sprintf("Key: %v Value: %v (Left: %v Right: %v)", node.key, node.value, node.left, node.right)
}
