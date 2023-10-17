package v1

/*
 * functionalities related to binary search tree
 */

import "math"

func searchBST(root *node, val int) bool {
	if root == nil {
		return false
	} else if root.value.(int) == val {
		return true
	} else if val <= root.value.(int) {
		return searchBST(root.left, val)
	} else {
		return searchBST(root.right, val)
	}
}

func (t *Tree) SearchBST(val interface{}) bool {
	return searchBST(t.root, val.(int))
}

func insert(root *node, data *node) *node {

	if root == nil {
		return data
	} else if data.value.(int) <= root.value.(int) {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}

	return root
}

func isBST(n *node, min int, max int) bool {
	if n == nil {
		return true
	}
	x := n.value.(int)

	// check if max value in left subtree is smaller
	// than the node and min value in right subtree greater than
	// the node.
	if x < min || x > max {
		return false
	}	

	return isBST(n.left, min, x-1 ) && isBST(n.right, x+1, max)
}

// check if the tree is a BST or not
func (t *Tree) IsBST() bool {
	return isBST(t.root, math.MinInt32, math.MaxInt32)
}

// Insert in a Binary Search Tree
func (t *Tree) InsertBST(val interface{}) {
	t.root = insert(t.root, makeNode(val))
	t.length++
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

func deleteNode(n *node, val int) *node {
	if n == nil {
		// nothing to do
		return nil
	}

	if val < n.value.(int) {
		// node in left side of tree
		n.left = deleteNode(n.left, val)
	} else if val > n.value.(int) {
		// node in right side of tree
		n.right = deleteNode(n.right, val)
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
		n.right = deleteNode(n.right, temp.value.(int))
	}

	return n
}

//
// - Leaf node, delete directly
// - single child, copy child to node and delete child
// - has both children
// 		* Find inOrder successor
//		* Copy successor contents to node, delete successor 
//
func (t *Tree) DeleteNodeBST(val interface{}) {
	t.root = deleteNode(t.root, val.(int))
}