package main

/*
 * Given an array, print the k largest elements
 */

import (
	"fmt"

	heap "heaps.example.com/api/v1"
)

func smallerThanUtil(input []int, k int, pos int) []int {
	var returnArray []int

	if pos >= len(input) || input[pos] >= k {
		return []int{}
	}

	returnArray = append(returnArray, input[pos])
	subArray := append(smallerThanUtil(input, k, (2*pos+1)), smallerThanUtil(input, k, (2*pos+2))...)
	return append(returnArray, subArray...)
}

func smallerThan(input []int, k int) []int {

	// build min heap
	h := heap.NewHeap(func(a, b int) bool { return a > b }).BuildHeap(input)
	return smallerThanUtil(h.GetHeapData(), k, 0)
}

func main() {

	fmt.Println(smallerThan([]int{1, 23, 12, 9, 30, 2, 50}, 3))
}
