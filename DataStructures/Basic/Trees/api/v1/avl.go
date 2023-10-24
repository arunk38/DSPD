package v1

import "fmt"

/*
 * functionalities related to binary search tree
 */

// returns new AVL tree
func NewAVL() *AvlTree {
	return &AvlTree{nil, 0}
}

// Create and a AVL type node with height set to 1
func makeAVLNode(key any, val any) *avlNode {
	return &avlNode{key, val, 1, nil, nil}
}

// return hieght stored in given avl node
func (node *avlNode) GetAVLNodeHeight() int {
	if node == nil {
		return 0
	}

	return node.height
}

func isGreaterThan(a, b any) bool {
	switch a.(type) {
	case int:
		return a.(int) > b.(int)
	}
	return false
}

func isLessThan(a, b any) bool {
	switch a.(type) {
	case int:
		return a.(int) < b.(int)
	}
	return false
}

func (n *avlNode) updateHeight() {
	left, right := 0, 0
	if n.left != nil {
		left = n.left.height
	}
	if n.right != nil {
		right = n.right.height
	}
	if left > right {
		n.height = left + 1
	} else {
		n.height = right + 1
	}
}

func (n *avlNode) balanceFactor() int {
	// calculate balance factor
	left, right := 0, 0
	if n.left != nil {
		left = n.left.height
	}
	if n.right != nil {
		right = n.right.height
	}
	return left - right
}

// return the new root
func (n *avlNode) rebalance() *avlNode {
	bf := n.balanceFactor()
	if bf > 1 {
		// LR
		if n.left != nil && n.left.balanceFactor() < 0 {
			n.left = n.left.leftRotation()
		}
		// LL
		n = n.rightRotation()
	} else if bf < -1 {
		// RL
		if n.right != nil && n.right.balanceFactor() > 0 {
			n.right = n.right.rightRotation()
		}
		// RR
		n = n.leftRotation()
	}
	return n
}

func (n *avlNode) rightRotation() *avlNode {
	n1 := n.left
	n.left = n1.right
	n1.right = n
	n.updateHeight()
	n1.updateHeight()
	return n1
}

func (n *avlNode) leftRotation() *avlNode {
	n1 := n.right
	n.right = n1.left
	n1.left = n
	n.updateHeight()
	n1.updateHeight()
	return n1
}

func (n *avlNode) insertUtil(curr *avlNode) *avlNode {
	if n == nil {
		return curr
	}
	if isLessThan(curr.key, n.key) {
		n.left = n.left.insertUtil(curr)
	} else if isGreaterThan(curr.key, n.key) {
		n.right = n.right.insertUtil(curr)
	} else {
		n.value = curr.value
	}
	n.updateHeight()
	return n.rebalance()
}

// insert in to avl tree and self balance the tree
func (t *AvlTree) InsertAVL(key any, val any) *AvlTree {
	node := makeAVLNode(key, val)
	if t.root == nil {
		node.updateHeight()
		t.root = node
	} else {
		t.root = t.root.insertUtil(node)
	}
	t.length++
	return t
}

func BulkInsertAVL(t *AvlTree, val []any, key []any) *AvlTree {
	for i := range val {
		t.InsertAVL(key[i], val[i])
	}

	return t
}

func getInOrderValuesUtil(n *avlNode) []any {
	var val []any
	if n != nil {
		val = append(getInOrderValuesUtil(n.left), n.value)
		val = append(val, getInOrderValuesUtil(n.right)...)

	}

	return val
}

func (t *AvlTree) GetInOrderValues() []any {
	if t.root == nil {
		return []any{"Empty Tree"}
	}
	return getInOrderValuesUtil(t.root)
}

// Returns height/depth of avl tree
func (t *AvlTree) Height() int {
	if t.root == nil {
		return 0
	}
	return t.root.height
}

// returns total number of nodes in avl tree
func (t *AvlTree) Size() int {
	return t.length
}

func (t *AvlTree) AVLGetRoot() *avlNode {
	return t.root
}

func GetNodeValue(n *avlNode) any {
	if n == nil {
		return "<nil>"
	}
	return n.value
}

func GetNodeKey(n *avlNode) any {
	if n == nil {
		return "<nil>"
	}
	return n.key
}

// Search a given key in avl Tree
func (t *AvlTree) Search(key any) *avlNode {
	rootNode := t.root
	for current := rootNode; current != nil; {
		if isLessThan(key, current.key) {
			current = current.left
		} else if isGreaterThan(key, current.key) {
			current = current.right
		} else {
			return current
		}
	}
	// if not found
	return nil
}

func (n *avlNode) delete(key any) *avlNode {
	if n == nil {
		return nil
	}

	if isLessThan(key, n.key) {
		n.left = n.left.delete(key)
	} else if isGreaterThan(key, n.key) {
		n.right = n.right.delete(key)
	} else {
		if n.right != nil {
			// swap current key with minimum from right (next highest) and delete leaf node
			successor := n.right
			for successor.left != nil {
				successor = successor.left
			}
			n.key = successor.key
			n.value = successor.value
			n.right = n.right.delete(successor.key)
		} else if n.left != nil {
			// node with only left child
			n = n.left
		} else {
			// node has no children
			n = nil
			return n
		}
	}
	if n != nil {
		n.updateHeight()
		n = n.rebalance()
	}
	return n
}

func (t *AvlTree) Delete(key any) *AvlTree {
	t.root = t.root.delete(key)
	return t
}

func printNodeInfo(node *avlNode) {
	fmt.Printf("[  (k, v, h) = (%d, %s, %d) left key = ", node.key.(int), node.value.(string), node.height)

	if node.left != nil {
		fmt.Print(" ", node.left.key.(int))
	} else {
		fmt.Print("<nil> ")
	}
	fmt.Print(" Right Key = ")
	if node.right != nil {
		fmt.Print(" ", node.right.key.(int))
	} else {
		fmt.Print("<nil>")
	}
	fmt.Println("]")
}

func printUtil(node *avlNode) {
	if node == nil {
		return
	}
	printNodeInfo(node)
	printUtil(node.left)
	printUtil(node.right)

}

func (t *AvlTree) PrintAVL() int {
	fmt.Print("\n************	START	************\n\n")
	printUtil(t.root)
	fmt.Print("\n************	xENDx	************\n\n")
	return 0
}
