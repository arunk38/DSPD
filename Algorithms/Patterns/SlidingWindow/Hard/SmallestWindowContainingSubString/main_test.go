package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Given a string and a pattern, find the smallest substring in the given string which has all the characters of the given pattern.
 *
 *	Example 1:
 *		Input: String="aabdec", Pattern="abc"
 *		Output: "abdec"
 *		Explanation: The smallest substring having all characters of the pattern is "abdec"
 *
 *	Example 2:
 *		Input: String="abdabca", Pattern="abc"
 *		Output: "abc"
 *		Explanation: The smallest substring having all characters of the pattern is "abc".
 *
 *	Example 3:
 *		Input: String="adcad", Pattern="abc"
 *		Output: ""
 *		Explanation: No substring in the given string has all characters of the pattern.
 *
 */

func finSubString(str string, pattern string) string {

	freqencyMap := make(map[rune]int)
	windowStart := 0
	matched := 0
	minLength := len(str) + 1
	subStrStart := 0

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

			if freqencyMap[rune(rightChar)] >= 0 {
				matched += 1 // count every matching of a char
			}
		}

		// Whenever we have matched all the characters, we will try to shrink the window
		// from the beginning, keeping track of the smallest substring that has all the matching characters.
		for matched == len(pattern) {
			if minLength > windowEnd-windowStart+1 {
				minLength = windowEnd - windowStart + 1
				subStrStart = windowStart
			}
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

	if minLength > len(str) {
		return ""
	} else {
		return str[subStrStart : subStrStart+minLength]
	}
}

func TestSmallestSubArray(t *testing.T) {

	type args struct {
		arg1 string
		arg2 string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String=aabdec, Pattern=abc",
			args: args{"aabdec", "abc"},
			want: "abdec",
		},
		{
			name: "String=abdabca, Pattern=abc",
			args: args{"abdabca", "abc"},
			want: "abc",
		},
		{
			name: "String=adcad, Pattern=abc",
			args: args{"adcad", "abc"},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := finSubString(tt.args.arg1, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
