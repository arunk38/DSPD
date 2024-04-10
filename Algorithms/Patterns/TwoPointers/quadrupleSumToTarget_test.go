package twopointers

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Given an array containing 0s, 1s and 2s, sort the array in-place.
 *
 * Time complexity:
 *		Sorting - O(N * logN)
 *		Overall -  O(N * logN + N^3) ---> O(N^3)
 *
 * Space complexity:
 *		O(N) - required for sorting
 */

func searchPair(input []int, target int, first int, second int) [][]int {
	right := len(input) - 1
	left := second + 1
	quadruplets := [][]int{}

	for left < right {
		currSum := input[first] + input[second] + input[left] + input[right]

		if currSum == target { // found the triplet
			quadruplets = append(quadruplets, []int{input[first], input[second], input[left], input[right]})
			left++
			right--

			// skip same elements to avoid duplicate triplets
			for left < right && input[left] == input[left-1] {
				left++
			}
			for left < right && input[right] == input[right+1] {
				right--
			}
		} else if target > currSum {
			left++ // we need pair wtih bigger sum
		} else {
			right-- // we need pair with smaller sum
		}
	}
	return quadruplets
}

func quadrupleTargetSum(input []int, target int) [][]int {
	// sort the given array
	sort.Ints(input)

	quadruplets := [][]int{}

	for i := 0; i < len(input)-3; i++ {
		if i > 0 && input[i] == input[i-1] { // skip same elements to avoid cuplicates
			continue
		}
		for j := i + 1; j < len(input)-2; j++ {
			if j > i+1 && input[j] == input[j-1] { // skip same elements to avoid cuplicates
				continue
			}
			quadruplets = append(quadruplets, searchPair(input, target, i, j)...)
		}
	}
	return quadruplets
}

func TestQuadrupleTargetSum(t *testing.T) {

	type args struct {
		arg1 []int
		arg2 int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[4, 1, 2, -1, 1, -3], target=1",
			args: args{[]int{4, 1, 2, -1, 1, -3}, 1},
			want: [][]int{{-3, -1, 1, 4}, {-3, 1, 1, 2}},
		},
		{
			name: "[2, 0, -1, 1, -2, 2], target=2",
			args: args{[]int{2, 0, -1, 1, -2, 2}, 2},
			want: [][]int{{-2, 0, 2, 2}, {-1, 0, 1, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := quadrupleTargetSum(tt.args.arg1, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
