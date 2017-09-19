package bstree

import "fmt"

// Node is the basic data unit of the binary tree
type Node struct {
	size        int
	item        Item
	left, right *Node
}

// Size returns the size of the node
func (node *Node) Size() int {
	if node == nil {
		return 0
	}

	return node.size
}

func (node *Node) String() string {
	return fmt.Sprintf("Item: %v (Left: %v Right: %v)", node.item, node.left, node.right)
}
