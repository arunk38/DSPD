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

func findLength(input string, k int) {
	// hash map to store char index
	hashMap := make(map[byte]int)
	windowStart := 0
	maxLength := 0
	maxRepeatLetterCount := 0

	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		rightChar := input[windowEnd]
		// add the letter to hashmap if not present
		if _, ok := hashMap[rightChar]; !ok {
			hashMap[rightChar] = 0
		}
		// increment the letter frequency in hash map
		hashMap[rightChar] += 1

		// tracks the count of letter which is repeating max times in current times
		maxRepeatLetterCount = alogs.GetMax(maxRepeatLetterCount, hashMap[rightChar])

		// Current window size is from windowStart to windowEnd, overall we have a letter which is repeating
		// 'maxRepeatLetterCount' times, means we have window which has one letter repeating 'maxRepeatLetterCount' times
		// and remaining letters we should/could replace.
		// If the remaining letters are more than k, need to shrink the window
		if (windowEnd - windowStart + 1 - maxRepeatLetterCount) > k {
			leftChar := input[windowStart]
			// remove letter going out of window from hash map
			hashMap[leftChar] -= 1
			windowStart++
		}
		maxLength = alogs.GetMax(maxLength, windowEnd-windowStart+1)
	}

	fmt.Println("Longest sub-string after replacement : ", maxLength)
}

func main() {
	findLength("aabccbb", 2)
	findLength("abbcb", 1)
	findLength("abccde", 1)
}
