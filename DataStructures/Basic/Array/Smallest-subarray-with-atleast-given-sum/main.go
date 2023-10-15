package main

/*
 * Given an array of integers and a number x, find the smallest subarray with a sum greater than the given value. 
 * Example:
 *	Input:
 *		arr[] = {1, 4, 45, 6, 0, 19}
 *		x  =  51
 *	Output: 
 *		sub array length: 3
 *		Minimum length subarray is {4, 45, 6}
 *
 * Approach:
 *		sliding window technique
 *			- Expand the window by adding the current value to sum till now
 *			- shrink the window by removing the element at start index of current window
 *				- dunring shirnk, update the min as per condition
 *				- remove the element at start index
 *			Complexity:
 *				- Time: O(n)
 *		
 */

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	min := len(nums) + 1
	sum := 0
	start := 0

	for end := 0; end < len(nums); end++ {
		sum += nums[end]

		// slide the start -> end window by removing starting elements from index
		for sum > target {

			// if the window is smaller, update min to window size
			if (end - start + 1) < min {
				min = (end - start + 1)
			}

			// repeat the same as long as the sum is more than target
			sum -= nums[start]
			start++
		}
	}

	// min value is not changed, finding sub array is not possible
	if min == len(nums) + 1{
		return 0
	}

	return min
}

func main() {

	input_arr := []int{1, 4, 45, 6, 10, 19}
	X := 51
	fmt.Println("Array: ", input_arr, "Sum value: ", X)
	res := minSubArrayLen(X, input_arr)
	if res == 0 {
		fmt.Println("Not possible")
	} else {
		fmt.Println("Sub array length: ", res)
	}

	input_arr = []int{ 1, 10, 5, 2, 7 }
	X = 9
	fmt.Println("\nArray: ", input_arr, "Sum value: ", X)
	res = minSubArrayLen(X, input_arr)
	if res == 0 {
		fmt.Println("Not possible")
	} else {
		fmt.Println("Sub array length: ", res)
	}

	input_arr = []int{1, 2, 4}
	X = 8
	fmt.Println("\nArray: ", input_arr, "Sum value: ", X)
	res = minSubArrayLen(X, input_arr)
	if res == 0 {
		fmt.Println("Not possible")
	} else {
		fmt.Println("Sub array length: ", res)
	}

	input_arr = []int{1, 11, 100, 1, 0, 200, 3, 2, 1, 250}
	X = 280
	fmt.Println("\nArray: ", input_arr, "Sum value: ", X)
	res = minSubArrayLen(X, input_arr)
	if res == 0 {
		fmt.Println("Not possible")
	} else {
		fmt.Println("Sub array length: ", res)
	}
}