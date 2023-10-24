package v1

// Tree data structure
type Tree struct {
	root   *node
	length int
}

// node type in Tree data type
type node struct {
	value any
	left  *node
	right *node
}

// Tree data structure
type AvlTree struct {
	root   *avlNode
	length int
}

// node type for avl trees
type avlNode struct {
	key    any
	value  any
	height int
	left   *avlNode
	right  *avlNode
}
