package v1

/*
 * functionalities related to binary search tree
 */

import "math"

func searchBSTInt(root *node, val int) (bool, interface{}) {
	if root == nil {
		return false, nil
	} else if GetValueInt(root) == val {
		return true, root
	} else if val <= GetValueInt(root) {
		return searchBSTInt(root.left, val)
	} else {
		return searchBSTInt(root.right, val)
	}
}

func (t *Tree) SearchBST(val interface{}) (bool, interface{}) {
	switch v := val.(type) {
	case int:
		return searchBSTInt(t.root, v)
	}

	return false, nil
}

func insertBSTInt(root *node, data *node) *node {

	if root == nil {
		return data
	} else if GetValueInt(data) <= GetValueInt(root) {
		root.left = insertBSTInt(root.left, data)
	} else {
		root.right = insertBSTInt(root.right, data)
	}

	return root
}

// Insert in a Binary Search Tree
func (t *Tree) InsertBST(val interface{}) {
	switch val.(type) {
	case int:
		t.root = insertBSTInt(t.root, makeNode(val))
	}
	t.length++
}

func isBSTInt(n *node, min int, max int) bool {
	if n == nil {
		return true
	}
	x := GetValueInt(n)

	// check if max value in left subtree is smaller
	// than the node and min value in right subtree greater than
	// the node.
	if x < min || x > max {
		return false
	}

	return isBSTInt(n.left, min, x-1) && isBSTInt(n.right, x+1, max)
}

// check if the tree is a BST or not
func (t *Tree) IsBST() bool {
	if t.root == nil {
		return true
	}

	switch t.root.value.(type) {
	case int:
		return isBSTInt(t.root, math.MinInt32, math.MaxInt32)
	}

	return false
}

func minNode(n *node) *node {
	if n == nil {
		return nil
	}

	for n.left != nil {
		n = n.left
	}

	return n
}

func deleteNodeInt(n *node, val int) *node {
	if n == nil {
		// nothing to do
		return nil
	}

	if val < GetValueInt(n) {
		// node in left side of tree
		n.left = deleteNodeInt(n.left, val)
	} else if val > GetValueInt(n) {
		// node in right side of tree
		n.right = deleteNodeInt(n.right, val)
	} else {
		// If no child exists: return nil
		if n.left == nil && n.right == nil {
			return nil
		}
		// If left child exists: return left child.
		if n.left != nil && n.right == nil {
			return n.left
		}
		// If right child exists: return right child.
		if n.left == nil && n.right != nil {
			return n.right
		}
		// If both child exists:
		// get the min node from right child
		temp := minNode(n.right)

		n.value = temp.value
		n.right = deleteNodeInt(n.right, GetValueInt(temp))
	}

	return n
}

// - Leaf node, delete directly
// - single child, copy child to node and delete child
// - has both children
//   - Find inOrder successor
//   - Copy successor contents to node, delete successor
func (t *Tree) DeleteNodeBST(val interface{}) {
	switch t.root.value.(type) {
	case int:
		t.root = deleteNodeInt(t.root, val.(int))
	}
}
