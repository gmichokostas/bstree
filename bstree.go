package bstree

import (
	"fmt"
)

// BSTree represents a binary tree
type BSTree struct {
	Root *Node
}

// New returns a pointer to a new binary tree
func New() *BSTree {
	return &BSTree{}
}

// IsEmpty checks if the BSTree is empty
func (tree *BSTree) IsEmpty() bool {
	return tree.Root == nil
}

// Size returns the size of the tree
func (tree *BSTree) Size() int {
	return tree.Root.Size()
}

// Height returns the height of the tree
func (tree *BSTree) Height() int {
	return height(tree.Root)
}

func height(node *Node) int {
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
func (tree *BSTree) Put(item Item) {
	tree.Root = put(tree.Root, item)
}

func put(node *Node, item Item) *Node {
	if node == nil {
		return &Node{item: item, size: 1}
	}

	if item.LessThan(node.item) {
		node.left = put(node.left, item)
	} else if item.MoreThan(node.item) {
		node.right = put(node.right, item)
	} else {
		node.item = item
	}

	node.size = 1 + node.left.Size() + node.right.Size()
	return node
}

// Find search for an item in the tree
func (tree *BSTree) Find(item Item) (Item, bool) {
	return find(tree.Root, item)
}

func find(node *Node, item Item) (Item, bool) {
	if node == nil {
		return nil, false
	}

	if item.LessThan(node.item) {
		return find(node.left, item)
	} else if item.MoreThan(node.item) {
		return find(node.right, item)
	} else {
		return node.item, true
	}
}

// Min returns the min item of the tree
func (tree *BSTree) Min() (Item, bool) {
	if tree.IsEmpty() {
		return nil, false
	}

	node, found := min(tree.Root)

	return node.item, found
}

func min(node *Node) (*Node, bool) {
	if (node.left) == nil {
		return node, true
	}

	return min(node.left)
}

// InOrderPrint traversal of the tree
func (tree *BSTree) InOrderPrint() {
	inOrder(tree.Root)
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.left)
		fmt.Printf("%v ", node.item)
		inOrder(node.right)
	}
}

func (tree *BSTree) String() string {
	tree.InOrderPrint()

	return fmt.Sprintf("%v", tree.Root)
}
