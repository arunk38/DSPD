package main

import (
	"fmt"
	tree "trees.example.com/api/v1"
)

func main() {
	fmt.Println("\t\t***** Binary Search Tree *****")

	bst := tree.New()

	data := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	fmt.Println("Tree data: ", data)
	for _, i := range data {
		bst.InsertBST(i)
	}
	bst.PrintRootValue()

	bst.PrintInOrder() 

	fmt.Println("Searching if 5 is present ? : ", bst.SearchBST(5))
	fmt.Println("Searching if 6 is present ? : ", bst.SearchBST(6))

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
}