package main

import (
	"fmt"

	alogs "algos.example.com/api/v1"
)

/*
 * Given a string with lowercase letters only, you are allowed to replace no more than ‘k’ letters with any letter,
 * find the length of the longest substring having the same letters after replacement.
 *
 * Example:
 *	Input: String="aabccbb", k=2
 *	Output: 5
 *	Explanation: Replace the two 'c' with 'b' to have a longest repeating substring "bbbbb".
 */

func findLength(input []int, k int) {
	windowStart := 0
	maxLength := 0
	maxOnesCount := 0

	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		if input[windowEnd] == 1 {
			maxOnesCount += 1
		}

		// Current window size is from windowStart to windowEnd, overall we have a maximum of 1's repeating
		// 'maxOnesCount' times, means we have window which has  'maxOnesCount' times 1's and remaining 0's we should/could replace.
		// If the remaining 0's are more than k, need to shrink the window
		if (windowEnd - windowStart + 1 - maxOnesCount) > k {
			if input[windowStart] == 1 {
				maxOnesCount -= 1
			}
			windowStart++
		}
		maxLength = alogs.GetMax(maxLength, windowEnd-windowStart+1)
	}

	fmt.Println("Longest sub-array after replacement : ", maxLength)
}

func main() {
	findLength([]int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 2)
	findLength([]int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 3)
}
