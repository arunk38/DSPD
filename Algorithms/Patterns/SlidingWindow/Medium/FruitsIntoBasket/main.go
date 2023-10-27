package main

/*
 * Given an array of characters where each character represents a fruit tree,
 * you are given two baskets and your goal is to put maximum number of fruits in each basket.
 * The only restriction is that each basket can have only one type of fruit.
 *
 * Start with any tree, but once you have started you can’t skip a tree.
 * You will pick one fruit from each tree until you cannot, i.e., you will stop when you have to pick from a third fruit type.
 *
 * Example 1:
 *	Input: Fruit=['A', 'B', 'C', 'A', 'C']
 *	Output: 3
 *	Explanation: We can put 2 'C' in one basket and one 'A' in the other from the subarray ['C', 'A', 'C']
 */

import (
	"fmt"

	algos "algos.example.com/api/v1"
)

func getMaxFruitCount(input []rune) int {
	hashMap := make(map[rune]int)
	windowStart := 0
	maxLength := 0

	// extend the loop [windowStart, windowEnd]
	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		// initialize the key in hash map if not present
		if _, ok := hashMap[input[windowEnd]]; !ok {
			hashMap[input[windowEnd]] = 0
		}
		// incremenet key count
		hashMap[input[windowEnd]]++

		// shrink the window until only 2 distinct keys are present
		for len(hashMap) > 2 {
			leftChar := input[windowStart] // get the left most fruit from window
			hashMap[leftChar] -= 1         // remove that from window and check value and remove key if 0
			if hashMap[leftChar] == 0 {
				delete(hashMap, leftChar)
			}
			windowStart++ // shrink the window
		}
		maxLength = algos.GetMax(maxLength, windowEnd-windowStart+1)
	}
	return maxLength
}

func main() {
	fmt.Println("Maximum number of fruits: ", getMaxFruitCount([]rune{'A', 'B', 'C', 'A', 'C'}))
	fmt.Println("Maximum number of fruits: ", getMaxFruitCount([]rune{'A', 'B', 'C', 'B', 'B', 'C'}))
}
