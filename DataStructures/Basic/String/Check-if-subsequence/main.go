package main

/*
 * Given two strings str1 and str2, find if str1 is a subsequence of str2.
 * A subsequence is a sequence that can be derived from another sequence by
 * deleting some elements without changing the order of the remaining elements.
 * Example:
 *		Input: str1 = “AXY”, str2 = “ADXCPY”
 *		Output: True
 *		Input: str1 = “AXY”, str2 = “YADXCP”
 *		Output: False
 *		Input: str1 = “mtsdet”, str2 = “meetsandmeets”
 *		Output: True
 *
 * Approach:
 *		Traverse both strings from one side to other side.
 *		If a matching character found, move ahead in both strings, otherwise move ahead only in str2.
 *
 *		Complexity:
 *			- Time: O(n)
 *			- Aux. Space: O(1)
 *
 */

import "fmt"

func checkSubsequence(str1 string, str2 string) {
	fmt.Println("\nInput strings -", "string1 :", str1, "string2 :", str2)

	n1 := len(str1)
	n2 := len(str2)
	var i, j = 0, 0

	for i < n1 && j < n2 {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}

	fmt.Println(i == n1)
}

func main() {
	checkSubsequence("AXY", "ADXCPY")
	checkSubsequence("AXY", "YADXCP")
	checkSubsequence("mtsdet", "meetsandmeets")
}
