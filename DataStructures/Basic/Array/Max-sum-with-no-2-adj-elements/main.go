package main

/*
 * Given an array of positive numbers, find the maximum sum of a subsequence with the
 * constraint that no 2 numbers in the sequence should be adjacent in the array.
 * Example:
 *	Input: arr[] = {5, 5, 10, 100, 10, 5}
 *	Output: 110
 *
 * Approach:
 *		Each element has two choices. If one element is picked then its neighbours cannot
 *		be picked. Otherwise, its neighbours may be picked or may not be.
 *		So the maximum sum till ith index has two possibilities: the subsequence includes arr[i] or it does not include arr[i].
 *		If arr[i] is included then the maximum sum depends on the maximum subsequence sum till (i-1)th element excluding arr[i-1].
 *
 *			- Start with two sums `excl` and `incl`.
 *			- `excl` stores the value of the maximum subsequence sum till i-1 when arr[i-1] is excluded.
 *			- `incl` stores the value of the maximum subsequence sum till i-1 when arr[i-1] is included.
 *			- The value of `excl` for the current state(say excl_new) will be max(excl ,incl). And the value of `incl` will be updated to excl + arr[i].
 *			- At the end of the loop max sum will be max(incl, excl)
 */

import (
	"fmt"
	"math"
)

func MaxSumWithNoAdjacents(input_arr []int) {
	excl := 0.0
	incl := 0.0

	fmt.Println("\nInput array: ", input_arr)
	for _, i := range input_arr {
		excl_new := math.Max(incl, excl)
		incl = excl + float64(i)
		excl = excl_new
	}
	fmt.Println("Max sum: ", int(math.Max(incl, excl)))
}

func main() {

	input_arr1 := []int{5, 5, 10, 100, 10, 5}
	input_arr2 := []int{1, 2, 3}
	input_arr3 := []int{1, -20, 3}

	MaxSumWithNoAdjacents(input_arr1)
	MaxSumWithNoAdjacents(input_arr2)
	MaxSumWithNoAdjacents(input_arr3)
}
