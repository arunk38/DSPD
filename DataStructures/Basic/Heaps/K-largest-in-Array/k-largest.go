package main

/*
 * Given an array, print the k largest elements
 */

import (
	"fmt"

	heap "heaps.example.com/api/v1"
)

func kLargest(input []int, k int) []int {
	if k == 0 || k > len(input) {
		return []int{}
	}
	h := heap.NewHeap(func(a, b int) bool { return a < b }).BuildHeap(input)

	var returnArray []int

	for i := 0; i < k; i++ {
		returnArray = append(returnArray, h.Peek())
		h.Pop()
	}

	return returnArray
}

func main() {

	fmt.Println(kLargest([]int{1, 23, 12, 9, 30, 2, 50}, 3))
}
