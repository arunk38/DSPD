package tripletsumtozero_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Given an array of unsorted numbers, find all unique triplets in it that add up to zero.
 *
 *	Example 1:
 *	  Input: [-3, 0, 1, 2, -1, 1, -2]
 *	  Output: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
 *	  Explanation: There are four unique triplets whose sum is equal to zero.
 *
 *	Example 2:
 *	  Input: [-5, 2, -1, -2, 3]
 *	  Output: [[-5, 2, 3], [-2, -1, 3]]
 *	  Explanation: There are two unique triplets whose sum is equal to zero.
 */

func searchPair(input []int, target int, left int) [][]int {
	right := len(input) - 1
	tripletArr := [][]int{}

	for left < right {
		currSum := input[left] + input[right]

		if currSum == target { // found the triplet
			tripletArr = append(tripletArr, []int{-target, input[left], input[right]})
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
	return tripletArr
}

func tripletSumZero(input []int) [][]int {

	tripletArray := [][]int{}
	sort.Ints(input)

	for i := 0; i < len(input)-2; i++ {
		if i > 0 && input[i] == input[i-1] { // skip same elements to avoid duplicates
			continue
		}

		tripletArray = append(tripletArray, searchPair(input, -input[i], i+1)...)
	}

	return tripletArray
}

func TestTripletSumZero(t *testing.T) {

	type args struct {
		arg1 []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Triplet Sum [-3, 0, 1, 2, -1, 1, -2]",
			args: args{[]int{-3, 0, 1, 2, -1, 1, -2}},
			want: [][]int{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}},
		},
		{
			name: "Triplet Sum [-5, 2, -1, -2, 3]",
			args: args{[]int{-5, 2, -1, -2, 3}},
			want: [][]int{{-5, 2, 3}, {-2, -1, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tripletSumZero(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
