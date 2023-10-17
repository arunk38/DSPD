package main

/*
 *	Find the length of the longest contiguous subarray such that every element in the
 *	subarray is strictly greater than its previous element in the same subarray.
 *	Example:
 *		input: arr[] = {5, 6, 3, 5, 7, 8, 9, 1, 2}
 *		Output: 5   Subarray: {3, 5, 7, 8, 9}
 *
 *	Approach:
 *		If current element is greater than previous element increment the curr_length.
 *		Else, as current element is smaller, check if curr_length is greater than max_length and set max_length as curr_length, end=i and curr_length=1.
 *
 *		Complexity:
 *			- Time: O(n)
 *			- Aux. Space: O(1)
 */

import "fmt"

func longestIncreasingSubarray(arr []int) {
	fmt.Println("\nInput array: ", arr)

	max := 1
	curen_len := 1
	end := 0
	n := len(arr)

	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			curen_len += 1
		} else {
			if curen_len > max {
				max = curen_len
				curen_len = 1
				end = i
			}
		}
	}

	if curen_len > max {
		max = curen_len
		end = n
	}

	fmt.Println("Longest Increasing Subarray: ", arr[end-max:end])
}

func main() {

	input_arr := []int{5, 6, 3, 5, 7, 8, 9, 1, 2}
	longestIncreasingSubarray(input_arr)

	input_arr = []int{12, 13, 1, 5, 4, 7, 18, 10, 10, 11}
	longestIncreasingSubarray(input_arr)
}
