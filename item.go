package bstree

// Item provides an interface for comparing
type Item interface {
	LessThan(other Item) bool
	MoreThan(other Item) bool
}
