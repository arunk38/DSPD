package main

/*
 * Example:- Binary Tree
 *	Inorder Traversal:
 *		9 12 14 17 19 23 67 50 54 72 76
 *	Preorder Traversal:
 *		50 17 12 9 14 23 19 67 72 54 76
 *	Postorder Traversal:
 *		9 14 12 19 67 23 17 54 76 72 50
 *	BFS:
 *		50 17 72 12 23 54 76 9 14 19 67
 */

import (
	"fmt"

	tree "trees.example.com/api/v1"
)

func main() {
	bt := tree.New()

	bt.Insert(50)
	bt.Insert(17)
	bt.Insert(72)
	bt.Insert(12)
	bt.Insert(23)
	bt.Insert(54)
	bt.Insert(76)
	bt.Insert(9)
	bt.Insert(14)
	bt.Insert(19)
	bt.Insert(67)

	bt.PrintInOrder()
	bt.PrintPreOrder()
	bt.PrintPostOrder()

	bt.PrintBFS()

	fmt.Println("Number of nodes in tree: ", bt.Size())
	fmt.Println("Height of the tree: ", bt.Height())
	fmt.Println("Maximum element of the tree: ", bt.Maximum())
}
