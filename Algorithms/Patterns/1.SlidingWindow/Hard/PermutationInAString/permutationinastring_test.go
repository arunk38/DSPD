package permutationinastring_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Given a string and a pattern, find out if the string contains any permutation of the pattern.
 *
 * Example 1:
 *	Input: String="oidbcaf", Pattern="abc"
 *	Output: true
 *	Explanation: The string contains "bca" which is a permutation of the given pattern.
 *
 *	Example 2:
 *	 Input: String="odicf", Pattern="dc"
 *	 Output: false
 *	 Explanation: No permutation of the pattern is present in the given string as a substring.
 *
 *	Example 3:
 *	 Input: String="bcdxabcdy", Pattern="bcdyabcdx"
 *	 Output: true
 *	 Explanation: Both the string and the pattern are a permutation of each other.
 *
 *	Example 4:
 *	 Input: String="aaacb", Pattern="abc"
 *	 Output: true
 *	 Explanation: The string contains "acb" which is a permutation of the given pattern.
 *
 */

func findPermutation(str string, pattern string) bool {

	freqencyMap := make(map[rune]int)
	windowStart := 0
	matched := 0

	// calculate frequqency of all chars in the pattern
	for _, val := range pattern {
		if _, ok := freqencyMap[val]; !ok {
			freqencyMap[val] = 0
		}
		freqencyMap[val] += 1
	}

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]

		// decrement the frequency of matched char
		if _, ok := freqencyMap[rune(rightChar)]; ok {
			freqencyMap[rune(rightChar)] -= 1

			if freqencyMap[rune(rightChar)] == 0 {
				matched += 1 // current char is completly matched
			}
		}

		// If at any time, the number of characters matched is equal to the number of
		// distinct characters in the pattern (i.e., total characters in the HashMap),
		// we have gotten our required permutation.
		if matched == len(freqencyMap) {
			return true
		}

		// shrink the window if going out of pattern size
		if windowEnd >= len(pattern)-1 {
			leftChar := str[windowStart]

			if _, ok := freqencyMap[rune(leftChar)]; ok {
				// Put the going out char back into frequency map
				// we decremented its frequeceny when we entered this window
				if freqencyMap[rune(leftChar)] == 0 {
					matched -= 1 // decrement the matched count as this char is going out
				}
				// put the char back for matching
				freqencyMap[rune(leftChar)] += 1
			}
			windowStart++
		}
	}

	return false
}

func TestSmallestSubArray(t *testing.T) {

	type args struct {
		arg1 string
		arg2 string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "String=oidbcaf, Pattern=abc",
			args: args{"oidbcaf", "abc"},
			want: true,
		},
		{
			name: "String=odicf, Pattern=dc",
			args: args{"odicf", "dc"},
			want: false,
		},
		{
			name: "String=bcdxabcdy, Pattern=bcdyabcdx",
			args: args{"bcdxabcdy", "bcdyabcdx"},
			want: true,
		},
		{
			name: "String=aaacb, Pattern=abc",
			args: args{"aaacb", "abc"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPermutation(tt.args.arg1, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
