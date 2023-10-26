package main

/*
 * Given an array, find the average of all contiguous subarrays of size ‘K’ in it.
 */

import "fmt"

func averageOfSubarraySizeK(arr []int, k int) []float64 {
	var output []float64

	windowStart := 0
	windowSum := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd] // add the next element

		// slide the window once we reach the end of first window
		if windowEnd >= (k - 1) {
			output = append(output, float64(windowSum)/float64(k)) // calculate the average
			windowSum -= arr[windowStart]                          // substract the element going out
			windowStart++                                          // slide the windoe start
		}
	}

	return output
}

func main() {
	fmt.Println("Averages of subarrays of size K:", averageOfSubarraySizeK([]int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5))
}
