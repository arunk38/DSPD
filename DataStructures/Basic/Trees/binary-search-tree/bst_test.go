package main

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	tree "trees.example.com/api/v1"

	"testing"
)

// Lowest Common Ancestor in BST (LCA)
// - If both key1 and key2 is smaller than root’s val, then lca exist in left subtree.
// - If both key1 and key2 is greater than root’s val, then lca exist in right subtree.
// - Else this root is LCA.
func findLCAInt(root interface{}, key1 int, key2 int) int {

	for root != nil {
		val := tree.GetValueInt(root)
		// If both key1 and key2 is smaller than root's val, then lca exist in left subtree.
		if (key1 < val) && (key2 < val) {
			root = tree.GetLeftNode(root)
		} else if (key1 > val) && (key2 > val) {
			// If both key1 and key2 is greater than root's val, then lca exist in right subtree.
			root = tree.GetRightNode(root)
		} else {
			// Else this root is LCA.
			break
		}
	}

	return tree.GetValueInt(root)
}

// Testing LCA
func TestLCA(t *testing.T) {
	data := []int{20, 8, 22, 4, 12, 10, 14}

	fmt.Println("\n\n\t\t***** Lowest Common Ancestor *****")
	fmt.Println("Test tree data: ", data)
	fmt.Println()

	bst := tree.New()
	for _, i := range data {
		bst.InsertBST(i)
	}

	type args struct { // function args definition
		a interface{}
		b int
		c int
	}

	tests := []struct { // test case definition
		name string
		args args
		want int
	}{ // test cases
		{ // test case
			name: "LCA TC 1",
			args: args{bst.GetRoot(), 10, 14},
			want: 12,
		},
		{ // test case
			name: "LCA TC 2",
			args: args{bst.GetRoot(), 10, 22},
			want: 20,
		},
		{ // test case
			name: "LCA TC 3",
			args: args{bst.GetRoot(), 14, 8},
			want: 8,
		},
	}

	for _, tt := range tests { // test cases executions
		t.Run(tt.name, func(t *testing.T) {
			got := findLCAInt(tt.args.a, tt.args.b, tt.args.c) // function call
			assert.Equal(t, tt.want, got)                      // test case assertions
		})
	}
}

func testSearchBST(t *testing.T, tree *tree.Tree) {
	type args struct { // function args definition
		a int
	}
	tests := []struct { // test case definition
		name string
		args args
		want bool
	}{ // test cases
		{ // test case
			name: "Search data that is not present",
			args: args{5},
			want: false,
		},
		{ // test case
			name: "Search data that is present",
			args: args{6},
			want: true,
		},
	}
	for _, tt := range tests { // test cases executions
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tree.SearchBST(tt.args.a) // function call
			assert.Equal(t, tt.want, got)       // test case assertions
		})
	}
}

// Testing basic insert, delete, search in a BST
func TestBasic(t *testing.T) {
	bst := tree.New()

	data := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	fmt.Println("\n\n\t\t***** Basic BST *****")
	fmt.Println("Tree data: ", data)
	for _, i := range data {
		bst.InsertBST(i)
	}
	bst.PrintRootValue()

	bst.PrintInOrder()

	testSearchBST(t, bst)

	fmt.Println("Is my tree BST? : ", bst.IsBST())

	fmt.Println("Deleting 10 from BST")
	bst.DeleteNodeBST(10)
	bst.PrintInOrder()

	fmt.Println("Deleting 81 from BST")
	bst.DeleteNodeBST(81)
	bst.PrintInOrder()

	fmt.Println("Deleting 8 from BST")
	bst.DeleteNodeBST(8)
	bst.PrintInOrder()
	bst.PrintRootValue()

	data = []int{2, 1, 3, 4, 6, 5}
	bst1 := tree.New()
	for _, i := range data {
		bst1.InsertBST(i)
	}
	bst1.PrintInOrder()
}

// Inorder Successor of an input node can also be defined as the node with the smallest key greater than the key of input node.
// Approach:
//   - next node in Inorder traversal of the Binary Tree.
//   - NULL for the last node in Inoorder traversal.
func inorderSuccessor(root interface{}, node interface{}) int {

	// If node's right exist, simply return the min_node from right.
	rightNode := tree.GetRightNode(node)
	if rightNode != nil {
		return tree.GetValueInt(tree.MinNode(rightNode))
	}

	var successor interface{}
	successor = nil

	for root != nil {
		if tree.GetValueInt(node) < tree.GetValueInt(root) {
			successor = root
			root = tree.GetLeftNode(root)
		} else if tree.GetValueInt(node) > tree.GetValueInt(root) {
			root = tree.GetRightNode(root)
		} else {
			break
		}
	}

	return tree.GetValueInt(successor)
}

func TestInorderSuccessor(t *testing.T) {
	data := []int{20, 8, 22, 4, 12, 10, 14}

	bst := tree.New()
	for _, i := range data {
		bst.InsertBST(i)
	}

	fmt.Println("\n\n\t\t***** Inorder Successor *****")
	fmt.Println("Tree data: ", data)

	_, givenNode := bst.SearchBST(8)
	fmt.Println("Inorder Successor of  8 is (10): ", inorderSuccessor(bst.GetRoot(), givenNode))

	_, givenNode = bst.SearchBST(10)
	fmt.Println("Inorder Successor of  10 is (12): ", inorderSuccessor(bst.GetRoot(), givenNode))

	_, givenNode = bst.SearchBST(14)
	fmt.Println("Inorder Successor of  14 is (20): ", inorderSuccessor(bst.GetRoot(), givenNode))
}
