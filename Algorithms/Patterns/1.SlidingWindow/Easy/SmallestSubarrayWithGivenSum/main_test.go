package main

/*
 * Given an array of positive numbers and a positive number ‘S’,
 * find the length of the smallest contiguous subarray whose sum is greater than or equal to ‘S’.
 * Return 0, if no such subarray exists.
 */

import (
	"fmt"
	"math"
	"testing"

	algos "algos.example.com/api/v1"
	"github.com/stretchr/testify/assert"
)

func SmallestSubArray(arr []int, S int) int {

	windowSum := 0
	windowStart := 0
	arrayLength := math.MaxInt

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd] // add the next element

		// shrink the window as small as possible until is windowSum in greater
		for windowSum >= S { // current window sum is greater than given value
			arrayLength = algos.GetMin(arrayLength, windowEnd-windowStart+1) // get the minimum sub array
			windowSum -= arr[windowStart]                                    // substract the element going out
			windowStart++                                                    // slide the windoe start
		}
	}

	if arrayLength == math.MaxInt {
		arrayLength = 0
	}
	fmt.Printf("The length of smallest subarray with a sum greater than or equal to '%d' is %d\n", S, arrayLength)
	return arrayLength
}

func TestSmallestSubArray(t *testing.T) {

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
			name: "Array {2, 1, 5, 2, 3, 2} Sum 7",
			args: args{[]int{2, 1, 5, 2, 3, 2}, 7},
			want: 2,
		},
		{
			name: "Array {2, 1, 5, 2, 8}, 7 Sum 7",
			args: args{[]int{2, 1, 5, 2, 8}, 7},
			want: 1,
		},
		{
			name: "Array {3, 4, 1, 1, 6} Sum 8",
			args: args{[]int{3, 4, 1, 1, 6}, 8},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SmallestSubArray(tt.args.arg1, tt.args.arg2)
			assert.Equal(t, tt.want, got)
		})
	}
}
