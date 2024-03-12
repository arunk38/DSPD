package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Given an array containing 0s, 1s and 2s, sort the array in-place.
 */

func dutchFlag(input []int) []int {
	low, high := 0, len(input)-1
	i := 0

	// all the elemets < low are o and all > high are 2
	// all elements from >= low < i are 1
	for i <= high {
		if input[i] == 0 {
			input[i], input[low] = input[low], input[i]
			i++
			low++
		} else if input[i] == 1 {
			i++
		} else {
			input[i], input[high] = input[high], input[i]
			high--
		}
	}
	return input
}

func TestLongestSubstringKDistinct(t *testing.T) {

	type args struct {
		arg1 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1, 0, 2, 1, 0]",
			args: args{[]int{1, 0, 2, 1, 0}},
			want: []int{0, 0, 1, 1, 2},
		},
		{
			name: "[2, 2, 0, 1, 2, 0]",
			args: args{[]int{2, 2, 0, 1, 2, 0}},
			want: []int{0, 0, 1, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dutchFlag(tt.args.arg1)
			assert.Equal(t, tt.want, got)
		})
	}
}
