package main

import (
	"fmt"

	alogs "algos.example.com/api/v1"
)

/*
 * Given a string, find the length of the longest substring which has no repeating characters.
 */

func noRepeatSubString(input string) {
	// hash map to store char index
	hashMap := make(map[byte]int)
	windowStart := 0
	maxLength := 0

	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		rightChar := input[windowEnd]

		// this is tricky; in the current window, we will not have any 'rightChar' after its previous index
		// and if 'windowStart' is already ahead of the last index of 'rightChar', we'll keep 'windowStart'
		if _, ok := hashMap[rightChar]; ok {
			// keep moving forward with farthest index
			windowStart = alogs.GetMax(windowStart, hashMap[rightChar]+1)
		}
		// insert right char into map
		hashMap[rightChar] = windowEnd
		maxLength = alogs.GetMax(maxLength, windowEnd-windowStart+1)
	}

	fmt.Println("Length of the longest substring: ", maxLength)
}

func main() {
	noRepeatSubString("aabccbb")
	noRepeatSubString("abbbb")
	noRepeatSubString("abccde")
}
