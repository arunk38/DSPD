package stringanagrams_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Given a string and a pattern, find all anagrams of the pattern in the given string.
 *
 * Example 1:
 *	Input: String="ppqp", Pattern="pq"
 *	Output: [1, 2]
 *	Explanation: The two anagrams of the pattern in the given string are "pq" and "qp".
 *
 *	Example 2:
 *	 Input: String="abbcabc", Pattern="abc"
 *	 Output: [2, 3, 4]
 *	 Explanation: The three anagrams of the pattern in the given string are "bca", "cab", and "abc".
 *
 */

func findStringAnagrams(str string, pattern string) []int {

	freqencyMap := make(map[rune]int)
	windowStart := 0
	matched := 0
	var anagramlist []int

	// calculate frequqency of all chars in the pattern
	for _, val := range pattern {
		if _, ok := freqencyMap[val]; !ok {
			freqencyMap[val] = 0
		}
		freqencyMap[val] += 1
	}

	// match all the chars from the map wuth the current window
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]

		// decrement the frequency of matched char
		if _, ok := freqencyMap[rune(rightChar)]; ok {
			freqencyMap[rune(rightChar)] -= 1

			if freqencyMap[rune(rightChar)] == 0 {
				matched += 1 // current char is completly matched, add to list
			}
		}

		// we found an anagram
		if matched == len(freqencyMap) {
			anagramlist = append(anagramlist, windowStart)
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

	return anagramlist
}

func TestSmallestSubArray(t *testing.T) {

	type args struct {
		arg1 string
		arg2 string
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "String=ppqp, Pattern=pq",
			args: args{"ppqp", "pq"},
			want: []int{1, 2},
		},
		{
			name: "String=abbcabc, Pattern=abc",
			args: args{"abbcabc", "abc"},
			want: []int{2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findStringAnagrams(tt.args.arg1, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
