package v1

/*
 * Common tree structs and APIs
 */

import (
	"fmt"
	"math"

	queue "queue.example.com/api/v1"
)

type (
	// Tree data structure
	Tree struct {
		root   *node
		length int
	}

	// node type in Tree data type
	node struct {
		value interface{}
		left  *node
		right *node
	}
)

func New() *Tree {
	return &Tree{nil, 0}
}

func printInOrder(node *node) {
	if node != nil {
		printInOrder(node.left)
		fmt.Print(" ", node.value)
		printInOrder(node.right)
	}
}

// print tree in left - node - right
func (t *Tree) PrintInOrder() {
	fmt.Print("Inorder Tree traversal :")
	printInOrder(t.root)
	fmt.Println()
}

func printPostOrder(node *node) {
	if node != nil {
		printPostOrder(node.left)
		printPostOrder(node.right)
		fmt.Print(" ", node.value)
	}
}

// print tree in left - right - node
func (t *Tree) PrintPostOrder() {
	fmt.Print("Postorder Tree traversal :")
	printPostOrder(t.root)
	fmt.Println()
}

func printPreOrder(node *node) {
	if node != nil {
		fmt.Print(" ", node.value)
		printPreOrder(node.left)
		printPreOrder(node.right)
	}
}

// print tree in node - left - right
func (t *Tree) PrintPreOrder() {
	fmt.Print("Preorder Tree traversal :")
	printPreOrder(t.root)
	fmt.Println()
}

/*
Level order traversal of a tree is breadth first traversal for the tree.
Approach:
- Create empty queue and enqueue root
- while queue is not empty, dequeue, print, enqueue left, enqueue right
*/
func (t *Tree) PrintBFS() {

	if t.root == nil {
		return
	}

	fmt.Print("BFS Tree traversal :")
	q := queue.New()
	q.Enqueue(t.root)

	for !q.IsEmpty() {
		n := q.Dequeue().(*node)

		fmt.Print(" ", n.value)

		if n.left != nil {
			q.Enqueue(n.left)
		}
		if n.right != nil {
			q.Enqueue(n.right)
		}

	}
	fmt.Println()
}

// computes the number of nodes in a tree
func (t *Tree) Size() int {
	return t.length
}

func height(n *node) int {
	if n == nil {
		return 0
	}

	return (1 + int(math.Max(float64(height(n.left)), float64(height(n.right)))))
}

// computes the height/ depthof a tree
func (t *Tree) Height() int {
	return height(t.root)
}

func maximumInt(n *node) int {
	if n == nil {
		return -1
	}

	return int(math.Max((float64(GetValueInt(n))), math.Max(float64(maximumInt(n.left)), float64(maximumInt(n.right)))))
}

// Returns the max value stored in the tree
func (t *Tree) Maximum() interface{} {
	switch t.root.value.(type) {
	case int:
		return maximumInt(t.root)
	}
	return nil
}

// Insert in a binary tree
func (t *Tree) Insert(val interface{}) {
	if t.length == 0 {
		t.root = makeNode(val)
		t.length++
		return
	}

	n := makeNode(val)

	q := queue.New()
	q.Enqueue(t.root)
	for !q.IsEmpty() {
		temp := q.Dequeue().(*node)

		if temp.left == nil {
			temp.left = n
			break
		} else {
			q.Enqueue(temp.left)
		}

		if temp.right == nil {
			temp.right = n
			break
		} else {
			q.Enqueue(temp.right)
		}

	}
	t.length++

}

func (t *Tree) PrintRootValue() {
	fmt.Println("Root value of tree is : ", t.root.value)
}

func makeNode(val interface{}) *node {
	return &node{val, nil, nil}
}

func GetValueInt(n *node) int {
	return n.value.(int)
}
