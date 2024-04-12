package removeDuplicates_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Given an sorted array, remove duplicate elements and return length of new array
 * Example:
 * 	Input: [3, 2, 3, 6, 3, 10, 9, 3], Key=3
 *	Output: 4
 *
 * Smiliar Questions:
 * Given an unsorted array of numbers and a target ‘key’,
 * remove all instances of ‘key’ in-place and return the new length of the array.
 */

func removeDuplicates(input []int, k int) int {
	nextNonDuplicate := 1
	start := 1
	if k != -1 {
		nextNonDuplicate = 0
		start = 0
	}

	for i := start; i < len(input); i++ {
		compareKey := input[nextNonDuplicate]
		// if key is sepecified, we are dealing with unsorted array to remove kwy from input
		if k != -1 {
			compareKey = k
		}

		// update nextNonDuplicate index only if current and prev are not same
		// input[:nextNonDuplicate] will be array after removing duplicates
		if input[i] != compareKey {
			input[nextNonDuplicate] = input[i]
			nextNonDuplicate++
		}
	}

	fmt.Println("Array after rmeoving duplicates: ", input[:nextNonDuplicate])
	return nextNonDuplicate
}

func TestLongestSubstringKDistinct(t *testing.T) {

	type args struct {
		arg1 []int
		arg2 int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "remove duplicates [2, 3, 3, 3, 6, 9, 9]",
			args: args{[]int{2, 3, 3, 3, 6, 9, 9}, -1},
			want: 4,
		},
		{
			name: "remove duplicates [2, 2, 2, 11]",
			args: args{[]int{2, 2, 2, 11}, -1},
			want: 2,
		},
		{
			name: "remove key [3, 2, 3, 6, 3, 10, 9, 3], Key=3",
			args: args{[]int{3, 2, 3, 6, 3, 10, 9, 3}, 3},
			want: 4,
		},
		{
			name: "remove key [2, 11, 2, 2, 1] K=2",
			args: args{[]int{2, 11, 2, 2, 1}, 2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.arg1, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
