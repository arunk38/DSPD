package main

import "fmt"

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maximumSumSubarraySizeK(arr []int, k int) int {

	maxSum := 0
	windowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd] // add the next element

		// slide the window once we reach the end of first window
		if windowEnd >= (k - 1) {
			maxSum = getMax(windowSum, maxSum) // calculate the max sum
			windowSum -= arr[windowStart]      // substract the element going out
			windowStart++                      // slide the windoe start
		}
	}

	return maxSum
}

func main() {
	fmt.Println("Max Sum in [2, 1, 5, 1, 3, 2] of subarrays of size 3:", maximumSumSubarraySizeK([]int{2, 1, 5, 1, 3, 2}, 3))
	fmt.Println("Max Sum in [2, 3, 4, 1, 5] of subarrays of size 2:", maximumSumSubarraySizeK([]int{2, 3, 4, 1, 5}, 2))
}
