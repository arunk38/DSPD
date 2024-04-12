package main

import (
	"testing"

	alogs "algos.example.com/api/v1"
	"github.com/stretchr/testify/assert"
)

/*
 * Given a string, find the length of the longest substring in it with no more than K distinct characters.
 */

func longestSubstringKDistinct(input string, k int) int {

	hashMap := make(map[byte]int)
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

		// shrink the window until only k distinct keys are present
		for len(hashMap) > k {
			leftChar := input[windowStart] // get the left most fruit from window
			hashMap[leftChar] -= 1         // remove that from window and check value and remove key if 0
			if hashMap[leftChar] == 0 {
				delete(hashMap, leftChar)
			}
			windowStart++ // shrink the window
		}
		maxLength = alogs.GetMax(maxLength, windowEnd-windowStart+1)
	}
	return maxLength
}

func TestLongestSubstringKDistinct(t *testing.T) {

	type args struct {
		arg1 string
		arg2 int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "String=araaci K=2",
			args: args{"araaci", 2},
			want: 4,
		},
		{
			name: "String=araaci K=1",
			args: args{"araaci", 1},
			want: 2,
		},
		{
			name: "String=cbbebi K=3",
			args: args{"cbbebi", 3},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubstringKDistinct(tt.args.arg1, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
